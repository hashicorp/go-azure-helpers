package authentication

import "testing"

func TestServicePrincipalClientSecretMultiTenantAuth_builder(t *testing.T) {
	builder := Builder{
		ClientID:           "some-client-id",
		ClientSecret:       "some-client-secret",
		SubscriptionID:     "some-subscription-id",
		TenantID:           "some-tenant-id",
		AuxiliaryTenantIDs: []string{"aux-tenant-id1", "aux-tenant-id2"},
	}
	config, err := servicePrincipalClientSecretMultiTenantAuth{}.build(builder)
	if err != nil {
		t.Fatalf("Error building client secret auth: %s", err)
	}
	servicePrincipal := config.(servicePrincipalClientSecretMultiTenantAuth)

	if builder.ClientID != servicePrincipal.clientId {
		t.Fatalf("Expected Client ID to be %q but got %q", builder.ClientID, servicePrincipal.clientId)
	}

	if builder.ClientSecret != servicePrincipal.clientSecret {
		t.Fatalf("Expected Client Secret to be %q but got %q", builder.ClientSecret, servicePrincipal.clientSecret)
	}

	if builder.SubscriptionID != servicePrincipal.subscriptionId {
		t.Fatalf("Expected Subscription ID to be %q but got %q", builder.SubscriptionID, servicePrincipal.subscriptionId)
	}

	if builder.TenantID != servicePrincipal.tenantId {
		t.Fatalf("Expected Tenant ID to be %q but got %q", builder.TenantID, servicePrincipal.tenantId)
	}

	if builder.AuxiliaryTenantIDs[0] != servicePrincipal.auxiliaryTenantIDs[0] {
		t.Fatalf("Expected Auxiliary Tenant ID 1 to be %q but got %q", builder.TenantID[0], servicePrincipal.tenantId[0])
	}

	if builder.AuxiliaryTenantIDs[1] != servicePrincipal.auxiliaryTenantIDs[1] {
		t.Fatalf("Expected Auxiliary Tenant ID 2 to be %q but got %q", builder.TenantID[1], servicePrincipal.tenantId[1])
	}

	if len(builder.AuxiliaryTenantIDs) != len(servicePrincipal.auxiliaryTenantIDs) {
		t.Fatalf("Expected len(Auxiliary Tenant ID) to be %q but got %q", len(builder.TenantID), len(servicePrincipal.tenantId))
	}
}

func TestServicePrincipalClientSecretMultiTenantAuth_isApplicable(t *testing.T) {
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
				SupportsClientSecretAuth: false,
			},
			Valid: false,
		},
		{
			Description: "Feature Toggled on but no secret specified",
			Builder: Builder{
				SupportsClientSecretAuth: true,
			},
			Valid: false,
		},
		{
			Description: "Secret specified but feature toggled off",
			Builder: Builder{
				ClientSecret: "I turned myself into a pickle morty!",
			},
			Valid: false,
		},
		{
			Description: "Multi Tenant not enabled",
			Builder: Builder{
				SupportsClientSecretAuth: true,
				ClientSecret:             "I turned myself into a pickle morty!",
			},
			Valid: false,
		},
		{
			Description: "Missing Auxiliary Tenants",
			Builder: Builder{
				SupportsClientSecretAuth: true,
				SupportsAuxiliaryTenants: true,
				ClientSecret:             "I turned myself into a pickle morty!",
			},
			Valid: false,
		},
		{
			Description: "Valid configuration",
			Builder: Builder{
				SupportsClientSecretAuth: true,
				SupportsAuxiliaryTenants: true,
				AuxiliaryTenantIDs:       []string{"aux-tenant-id1", "aux-tenant-id2"},
				ClientSecret:             "I turned myself into a pickle morty!",
			},
			Valid: true,
		},
	}

	for _, v := range cases {
		applicable := servicePrincipalClientSecretMultiTenantAuth{}.isApplicable(v.Builder)
		if v.Valid != applicable {
			t.Fatalf("Expected %q to be %t but got %t", v.Description, v.Valid, applicable)
		}
	}
}

func TestServicePrincipalClientSecretMultiTenantAuth_populateConfig(t *testing.T) {
	config := &Config{}
	err := servicePrincipalClientSecretMultiTenantAuth{}.populateConfig(config)
	if err != nil {
		t.Fatalf("Error populating config: %s", err)
	}

	if !config.AuthenticatedAsAServicePrincipal {
		t.Fatalf("Expected `AuthenticatedAsAServicePrincipal` to be true but it wasn't")
	}
}

func TestServicePrincipalClientSecretMultiTenantAuth_validate(t *testing.T) {
	cases := []struct {
		Description string
		Config      servicePrincipalClientSecretMultiTenantAuth
		ExpectError bool
	}{
		{
			Description: "Empty Configuration",
			Config:      servicePrincipalClientSecretMultiTenantAuth{},
			ExpectError: true,
		},
		{
			Description: "Missing Client ID",
			Config: servicePrincipalClientSecretMultiTenantAuth{
				subscriptionId: "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
				clientSecret:   "Does Hammer Time have Daylight Savings Time?",
				tenantId:       "9834f8d0-24b3-41b7-8b8d-c611c461a129",
			},
			ExpectError: true,
		},
		{
			Description: "Missing Subscription ID",
			Config: servicePrincipalClientSecretMultiTenantAuth{
				clientId:     "62e73395-5017-43b6-8ebf-d6c30a514cf1",
				clientSecret: "Does Hammer Time have Daylight Savings Time?",
				tenantId:     "9834f8d0-24b3-41b7-8b8d-c611c461a129",
			},
			ExpectError: true,
		},
		{
			Description: "Missing Client Secret",
			Config: servicePrincipalClientSecretMultiTenantAuth{
				clientId:       "62e73395-5017-43b6-8ebf-d6c30a514cf1",
				subscriptionId: "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
				tenantId:       "9834f8d0-24b3-41b7-8b8d-c611c461a129",
			},
			ExpectError: true,
		},
		{
			Description: "Missing Tenant ID",
			Config: servicePrincipalClientSecretMultiTenantAuth{
				clientId:       "62e73395-5017-43b6-8ebf-d6c30a514cf1",
				subscriptionId: "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
				clientSecret:   "Does Hammer Time have Daylight Savings Time?",
			},
			ExpectError: true,
		},
		{
			Description: "Missing Auxiliary Tenants ID",
			Config: servicePrincipalClientSecretMultiTenantAuth{
				clientId:       "62e73395-5017-43b6-8ebf-d6c30a514cf1",
				subscriptionId: "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
				clientSecret:   "Does Hammer Time have Daylight Savings Time?",
			},
			ExpectError: true,
		},
		{
			Description: "Valid Configuration",
			Config: servicePrincipalClientSecretMultiTenantAuth{
				clientId:           "62e73395-5017-43b6-8ebf-d6c30a514cf1",
				subscriptionId:     "8e8b5e02-5c13-4822-b7dc-4232afb7e8c2",
				clientSecret:       "Does Hammer Time have Daylight Savings Time?",
				tenantId:           "9834f8d0-24b3-41b7-8b8d-c611c461a129",
				auxiliaryTenantIDs: []string{"9834f8d0-0707-1984-bd35-c611c461a129", "9834f8d0-1984-0707-bd35-c611c461a129"},
			},
			ExpectError: false,
		},
		{
			Description: "Invalid TenantOnly Configuration",
			Config: servicePrincipalClientSecretMultiTenantAuth{
				clientId:           "62e73395-5017-43b6-8ebf-d6c30a514cf1",
				clientSecret:       "Does Hammer Time have Daylight Savings Time?",
				tenantOnly:         true,
				auxiliaryTenantIDs: []string{"9834f8d0-0707-1984-bd35-c611c461a129", "9834f8d0-1984-0707-bd35-c611c461a129"},
			},
			ExpectError: true,
		},
		{
			Description: "Valid TenantOnly Configuration",
			Config: servicePrincipalClientSecretMultiTenantAuth{
				clientId:           "62e73395-5017-43b6-8ebf-d6c30a514cf1",
				clientSecret:       "Does Hammer Time have Daylight Savings Time?",
				tenantId:           "9834f8d0-24b3-41b7-8b8d-c611c461a129",
				tenantOnly:         true,
				auxiliaryTenantIDs: []string{"9834f8d0-0707-1984-bd35-c611c461a129", "9834f8d0-1984-0707-bd35-c611c461a129"},
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
