package authentication

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Azure/go-autorest/autorest/azure"
)

type Environment struct {
	Portal                  string         `json:"portal"`
	Authentication          Authentication `json:"authentication"`
	Media                   string         `json:"media"`
	GraphAudience           string         `json:"graphAudience"`
	Graph                   string         `json:"graph"`
	Name                    string         `json:"name"`
	Suffixes                Suffixes       `json:"suffixes"`
	Batch                   string         `json:"batch"`
	ResourceManager         string         `json:"resourceManager"`
	VmImageAliasDoc         string         `json:"vmImageAliasDoc"`
	ActiveDirectoryDataLake string         `json:"activeDirectoryDataLake"`
	SqlManagement           string         `json:"sqlManagement"`
	Gallery                 string         `json:"gallery"`
}

type Authentication struct {
	LoginEndpoint    string   `json:"loginEndpoint"`
	Audiences        []string `json:"audiences"`
	Tenant           string   `json:"tenant"`
	IdentityProvider string   `json:"identityProvider"`
}

type Suffixes struct {
	AzureDataLakeStoreFileSystem        string `json:"azureDataLakeStoreFileSystem"`
	AcrLoginServer                      string `json:"acrLoginServer"`
	SqlServerHostname                   string `json:"sqlServerHostname"`
	AzureDataLakeAnalyticsCatalogAndJob string `json:"azureDataLakeAnalyticsCatalogAndJob"`
	KeyVaultDns                         string `json:"keyVaultDns"`
	Storage                             string `json:"storage"`
	AzureFrontDoorEndpointSuffix        string `json:"azureFrontDoorEndpointSuffix"`
}


// DetermineEnvironment determines what the Environment name is within
// the Azure SDK for Go and then returns the association environment, if it exists.
func DetermineEnvironment(name string) (*azure.Environment, error) {
	// detect cloud from environment
	env, envErr := azure.EnvironmentFromName(name)

	if envErr != nil {
		// try again with wrapped value to support readable values like german instead of AZUREGERMANCLOUD
		wrapped := fmt.Sprintf("AZURE%sCLOUD", name)
		env, envErr = azure.EnvironmentFromName(wrapped)
		if envErr != nil {
			return nil, fmt.Errorf("An Azure Environment with name %q was not found: %+v", name, envErr)
		}
	}

	return &env, nil
}

// LoadEnvironmentFromUrl attempts to load the specified environment from the endpoint.
// if the endpoint is an empty string, or an environment can't be
// found at the endpoint url then an error is returned
func LoadEnvironmentFromUrl(endpoint string) (*azure.Environment, error) {
	if endpoint == "" {
		return nil, fmt.Errorf("Endpoint was not set!")
	}

	env, err := azure.EnvironmentFromURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving Environment from Endpoint %q: %+v", endpoint, err)
	}

	return &env, nil
}

func normalizeEnvironmentName(input string) string {
	// Environment is stored as `Azure{Environment}Cloud`
	output := strings.ToLower(input)
	output = strings.TrimPrefix(output, "azure")
	output = strings.TrimSuffix(output, "cloud")

	// however Azure Public is `AzureCloud` in the CLI Profile and not `AzurePublicCloud`.
	if output == "" {
		return "public"
	}
	return output
}

// AzureEnvironmentByName returns a specific Azure Environment from the specified endpoint
func AzureEnvironmentByNameFromEndpoint(ctx context.Context, endpoint string, environmentName string) (*azure.Environment, error) {
	var environmentTranslationMap = map[string]string{
		"public": "AzureCloud",
		"usgovernment": "AzureUSGovernment",
		"german": "AzureGermanCloud",
		"china": "AzureChinaCloud",
	}

	uri := fmt.Sprintf("https://%s//metadata/endpoints?api-version=2019-05-01", endpoint)
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrieving environments from Azure MetaData service: %+v", err)
	}

	dec := json.NewDecoder(resp.Body)
	if _, err = dec.Token(); err != nil {
		return nil, err
	}

	// while the array contains values
	for dec.More() {
		var env Environment
		// decode an array value (Message)
		err := dec.Decode(&env)
		if err != nil {
			return nil, fmt.Errorf("unable to decode environment from %q response: %+v", uri, err)
		}
		if strings.EqualFold(env.Name, environmentName) || strings.EqualFold(env.Name, environmentTranslationMap[environmentName]) {
			// This check is to prevent azure stack users from getting their environment this way
			if env.Authentication.Tenant != "common" && env.Authentication.IdentityProvider != "AAD" {
				return nil, fmt.Errorf("environment %q from metadata_url %q is not supported", environmentName, endpoint)
			}

			aEnv := &azure.Environment{
				ResourceManagerEndpoint: env.ResourceManager,
				StorageEndpointSuffix:   env.Suffixes.Storage,
				ActiveDirectoryEndpoint: env.Authentication.LoginEndpoint,
				GraphEndpoint:           env.Graph,
				ResourceIdentifiers: azure.ResourceIdentifier{
					// This isn't returned from the metadata url and is universal across all environments
					Storage: "https://storage.azure.com/",
				},
			}

			if len(env.Authentication.Audiences) > 0 {
				aEnv.TokenAudience = env.Authentication.Audiences[0]
			} else {
				return nil, fmt.Errorf("unable to find token audience for environment %q from endpoint %q", environmentName, endpoint)
			}

			return aEnv, nil
		}
	}

	return nil, fmt.Errorf("unable to find environment %q from endpoint %q", environmentName, endpoint)
}
