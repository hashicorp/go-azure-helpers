package authentication

import (
	"context"
	"strings"
	"testing"
)

func TestAzureEnvironmentNames(t *testing.T) {
	testData := map[string]string{
		"":                       "public",
		"AzureChinaCloud":        "china",
		"AzureCloud":             "public",
		"AZUREUSGOVERNMENTCLOUD": "usgovernment",
		"AzurePublicCloud":       "public",
	}

	for input, expected := range testData {
		actual := normalizeEnvironmentName(input)
		if actual != expected {
			t.Fatalf("Expected %q for input %q: got %q!", expected, input, actual)
		}
	}
}

func TestAccAzureEnvironmentByName(t *testing.T) {
	env, err := AzureEnvironmentByNameFromEndpoint(context.TODO(), "management.azure.com", "public")
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
	if !strings.EqualFold(env.Name, "AzurePublicCloud") {
		t.Fatalf("Incorrect environment name returned. Expected: %q. Received: %q", "AzurePublicCloud", env.Name)
	}
	env, err = AzureEnvironmentByNameFromEndpoint(context.TODO(), "management.azure.com", "usgovernment")
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
	if !strings.EqualFold(env.Name, "AzureUSGovernmentCloud") {
		t.Fatalf("Incorrect environment name returned. Expected: %q. Received: %q", "AzureUSGovernmentCloud", env.Name)
	}
	env, err = AzureEnvironmentByNameFromEndpoint(context.TODO(), "management.azure.com", "china")
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
	if !strings.EqualFold(env.Name, "AzureChinaCloud") {
		t.Fatalf("Incorrect environment name returned. Expected: %q. Received: %q", "AzureChinaCloud", env.Name)
	}

}

func TestAccAzureEnvironmentByNameFromEndpoint(t *testing.T) {
	env, err := AzureEnvironmentByNameFromEndpoint(context.TODO(), "management.azure.com", "AzureCloud")
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
	if !strings.EqualFold(env.Name, "AzureCloud") {
		t.Fatalf("Incorrect environment name returned. Expected: %q. Received: %q", "AzureCloud", env.Name)
	}
	env, err = AzureEnvironmentByNameFromEndpoint(context.TODO(), "management.azure.com", "AzureChinaCloud")
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
	if !strings.EqualFold(env.Name, "AzureChinaCloud") {
		t.Fatalf("Incorrect environment name returned. Expected: %q. Received: %q", "AzureChinaCloud", env.Name)
	}
	env, err = AzureEnvironmentByNameFromEndpoint(context.TODO(), "management.azure.com", "AzureUSGovernment")
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
	if !strings.EqualFold(env.Name, "AzureUSGovernment") {
		t.Fatalf("Incorrect environment name returned. Expected: %q. Received: %q", "AzureUSGovernment", env.Name)
	}
	_, err = AzureEnvironmentByNameFromEndpoint(context.TODO(), "badurl", "AzureChinaCloud")
	if err == nil {
		t.Fatal("Expected error from bad endpoint")
	}
	_, err = AzureEnvironmentByNameFromEndpoint(context.TODO(), "management.azure.com", "badEnvironment")
	if err == nil {
		t.Fatal("Expected error from bad environment")
	}
}

func TestAccIsEnvironmentAzureStack(t *testing.T) {
	ok, err := IsEnvironmentAzureStack(context.TODO(), "management.azure.com", "public")
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
	if ok {
		t.Fatal("Expected `public` environment to not be Azure Stack")
	}
}
