package authentication

import (
	"testing"
)

func TestAzureCLITokenMultiTenantAuth_isApplicable(t *testing.T) {
	cases := []struct {
		Description string
		Builder     Builder
		Valid       bool
	}{
		{
			Description: "Empty Configuration",
			Builder:     Builder{},
			Valid:       false,
		},
		{
			Description: "Feature Toggled off",
			Builder: Builder{
				SupportsAzureCliToken:    false,
				SupportsAuxiliaryTenants: true,
				AuxiliaryTenantIDs:       []string{"test"},
			},
			Valid: false,
		},
		{
			Description: "Aux Tenant Feature Toggled off",
			Builder: Builder{
				SupportsAzureCliToken:    true,
				SupportsAuxiliaryTenants: false,
				AuxiliaryTenantIDs:       []string{"test"},
			},
			Valid: false,
		},
		{
			Description: "Empty Aux Tenants",
			Builder: Builder{
				SupportsAzureCliToken:    false,
				SupportsAuxiliaryTenants: true,
				AuxiliaryTenantIDs:       []string{},
			},
			Valid: false,
		},
		{
			Description: "Feature Toggled on",
			Builder: Builder{
				SupportsAzureCliToken:    true,
				SupportsAuxiliaryTenants: true,
				AuxiliaryTenantIDs:       []string{"test"},
			},
			Valid: true,
		},
	}

	for _, v := range cases {
		applicable := azureCliTokenMultiTenantAuth{}.isApplicable(v.Builder)
		if v.Valid != applicable {
			t.Fatalf("Expected %q to be %t but got %t", v.Description, v.Valid, applicable)
		}
	}
}

func TestAzureCLITokenMultiTenantAuth_populateConfig(t *testing.T) {
	config := &Config{}
	auth := azureCliTokenMultiTenantAuth{
		clientId: "some-subscription-id",
		profile: &azureCLIProfileMultiTenant{
			environment:        "dimension-c137",
			subscriptionId:     "some-subscription-id",
			tenantId:           "some-tenant-id",
			auxiliaryTenantIDs: []string{"aux-tenant-id"},
		},
	}

	err := auth.populateConfig(config)
	if err != nil {
		t.Fatalf("Error populating config: %s", err)
	}

	if auth.clientId != config.ClientID {
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

func TestAzureCLITokenMultiTenantAuth_validate(t *testing.T) {
	cases := []struct {
		Description string
		Config      azureCliTokenMultiTenantAuth
		ExpectError bool
	}{
		{
			Description: "Empty Configuration",
			Config:      azureCliTokenMultiTenantAuth{},
			ExpectError: true,
		},
		{
			Description: "Missing Subscription ID",
			Config: azureCliTokenMultiTenantAuth{
				clientId: "62e73395-5017-43b6-8ebf-d6c30a514cf1",
				profile: &azureCLIProfileMultiTenant{
					tenantId:           "9834f8d0-24b3-41b7-8b8d-c611c461a129",
					auxiliaryTenantIDs: []string{"9834f8d0-24b3-41b7-8b8d-000000000000"},
				},
			},
			ExpectError: true,
		},
		{
			Description: "Missing Tenant ID",
			Config: azureCliTokenMultiTenantAuth{
				clientId: "62e73395-5017-43b6-8ebf-d6c30a514cf1",
				profile: &azureCLIProfileMultiTenant{
					subscriptionId:     "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
					auxiliaryTenantIDs: []string{"9834f8d0-24b3-41b7-8b8d-000000000000"},
				},
			},
			ExpectError: true,
		},
		{
			Description: "Missing aux tenant IDs",
			Config: azureCliTokenMultiTenantAuth{
				clientId: "62e73395-5017-43b6-8ebf-d6c30a514cf1",
				profile: &azureCLIProfileMultiTenant{
					subscriptionId: "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
					tenantId:       "9834f8d0-24b3-41b7-8b8d-c611c461a129",
				},
			},
			ExpectError: true,
		},
		{
			Description: "Valid Configuration",
			Config: azureCliTokenMultiTenantAuth{
				clientId: "62e73395-5017-43b6-8ebf-d6c30a514cf1",
				profile: &azureCLIProfileMultiTenant{
					subscriptionId:     "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
					tenantId:           "9834f8d0-24b3-41b7-8b8d-c611c461a129",
					auxiliaryTenantIDs: []string{"9834f8d0-24b3-41b7-8b8d-000000000000"},
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
