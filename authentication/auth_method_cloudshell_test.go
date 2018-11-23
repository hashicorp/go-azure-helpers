package authentication

import (
	"os"
	"testing"

	"github.com/Azure/go-autorest/autorest/adal"
)

func TestCloudShellAuth_isApplicable(t *testing.T) {
	cases := []struct {
		Description string
		Builder     Builder
		SetEnvVar   bool
		Valid       bool
	}{
		{
			Description: "Empty Configuration",
			Builder:     Builder{},
			SetEnvVar:   false,
			Valid:       false,
		},
		{
			Description: "Feature Toggled off",
			Builder: Builder{
				SupportsCloudShell: false,
			},
			SetEnvVar: false,
			Valid:     false,
		},
		{
			Description: "Feature Toggled off with environment variable",
			Builder: Builder{
				SupportsCloudShell: false,
			},
			SetEnvVar: true,
			Valid:     false,
		},
		{
			Description: "Feature Toggled on but no environment variable",
			Builder: Builder{
				SupportsCloudShell: true,
			},
			SetEnvVar: false,
			Valid:     false,
		},
		{
			Description: "Feature Toggled on with environment variable",
			Builder: Builder{
				SupportsCloudShell: true,
			},
			SetEnvVar: true,
			Valid:     true,
		},
	}

	for _, v := range cases {
		// simulate running within CloudShell
		if v.SetEnvVar {
			os.Setenv("ACC_CLOUD", "PROD")
		} else {
			os.Unsetenv("ACC_CLOUD")
		}

		applicable := cloudShellAuth{}.isApplicable(v.Builder)
		if v.Valid != applicable {
			t.Fatalf("Expected %q to be %t but got %t", v.Description, v.Valid, applicable)
		}
	}
}

func TestCloudShellAuth_populateConfig(t *testing.T) {
	config := &Config{}
	auth := cloudShellAuth{
		profile: &azureCLIProfile{
			clientId:       "some-subscription-id",
			environment:    "dimension-c137",
			subscriptionId: "some-subscription-id",
			tenantId:       "some-tenant-id",
		},
	}

	err := auth.populateConfig(config)
	if err != nil {
		t.Fatalf("Error populating config: %s", err)
	}

	if auth.profile.clientId != config.ClientID {
		t.Fatalf("Expected Client ID to be %q but got %q", auth.profile.tenantId, config.TenantID)
	}

	if auth.profile.environment != config.Environment {
		t.Fatalf("Expected Environment to be %q but got %q", auth.profile.tenantId, config.TenantID)
	}

	if auth.profile.subscriptionId != config.SubscriptionID {
		t.Fatalf("Expected Subscription ID to be %q but got %q", auth.profile.tenantId, config.TenantID)
	}

	if auth.profile.tenantId != config.TenantID {
		t.Fatalf("Expected Tenant ID to be %q but got %q", auth.profile.tenantId, config.TenantID)
	}
}

func TestCloudShellAuth_validate(t *testing.T) {
	cases := []struct {
		Description string
		Config      cloudShellAuth
		ExpectError bool
	}{
		{
			Description: "Empty Configuration",
			Config:      cloudShellAuth{},
			ExpectError: true,
		},
		{
			Description: "Missing Access Token",
			Config: cloudShellAuth{
				profile: &azureCLIProfile{
					clientId:       "62e73395-5017-43b6-8ebf-d6c30a514cf1",
					subscriptionId: "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
					tenantId:       "9834f8d0-24b3-41b7-8b8d-c611c461a129",
				},
			},
			ExpectError: true,
		},
		{
			Description: "Missing Client ID",
			Config: cloudShellAuth{
				profile: &azureCLIProfile{
					accessToken:    &adal.Token{},
					subscriptionId: "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
					tenantId:       "9834f8d0-24b3-41b7-8b8d-c611c461a129",
				},
			},
			ExpectError: true,
		},
		{
			Description: "Missing Subscription ID",
			Config: cloudShellAuth{
				profile: &azureCLIProfile{
					accessToken: &adal.Token{},
					clientId:    "62e73395-5017-43b6-8ebf-d6c30a514cf1",
					tenantId:    "9834f8d0-24b3-41b7-8b8d-c611c461a129",
				},
			},
			ExpectError: true,
		},
		{
			Description: "Missing Tenant ID",
			Config: cloudShellAuth{
				profile: &azureCLIProfile{
					accessToken:    &adal.Token{},
					clientId:       "62e73395-5017-43b6-8ebf-d6c30a514cf1",
					subscriptionId: "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
				},
			},
			ExpectError: true,
		},
		{
			Description: "Valid Configuration",
			Config: cloudShellAuth{
				profile: &azureCLIProfile{
					accessToken:    &adal.Token{},
					clientId:       "62e73395-5017-43b6-8ebf-d6c30a514cf1",
					subscriptionId: "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
					tenantId:       "9834f8d0-24b3-41b7-8b8d-c611c461a129",
				},
			},
			ExpectError: false,
		},
	}

	for _, v := range cases {
		err := v.Config.validate()

		if v.ExpectError && err == nil {
			t.Fatalf("Expected an error for %q: didn't get one", v.Description)
		}

		if !v.ExpectError && err != nil {
			t.Fatalf("Expected there to be no error for %q - but got: %v", v.Description, err)
		}
	}
}
