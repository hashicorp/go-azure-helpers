package authentication

import (
	"reflect"
	"testing"
)

func TestOIDC_builder(t *testing.T) {
	builder := Builder{
		Environment:         "some-environment",
		TenantID:            "some-tenant-id",
		AuxiliaryTenantIDs:  []string{"a-tenant-id", "b-tenant-id"},
		ClientID:            "some-client-id",
		IDTokenRequestURL:   "https://token.endpoint",
		IDTokenRequestToken: "fedcba9876543210",
	}

	method, err := oidcAuth{}.build(builder)
	if err != nil {
		t.Fatalf("Error building OIDC Auth: %+v", err)
	}

	authMethod := method.(oidcAuth)
	if builder.Environment != authMethod.environment {
		t.Fatalf("Expected Environment to be %q but got %q", builder.Environment, authMethod.environment)
	}
	if builder.TenantID != authMethod.tenantId {
		t.Fatalf("Expected Tenant ID to be %q but got %q", builder.TenantID, authMethod.tenantId)
	}
	if !reflect.DeepEqual(builder.AuxiliaryTenantIDs, authMethod.auxiliaryTenantIds) {
		t.Fatalf("Expected Aux Tenant IDs to be %#v but got %#v", builder.AuxiliaryTenantIDs, authMethod.auxiliaryTenantIds)
	}
	if builder.ClientID != authMethod.clientId {
		t.Fatalf("Expected Client ID to be %q but got %q", builder.ClientID, authMethod.clientId)
	}
	if builder.IDTokenRequestURL != authMethod.idTokenRequestUrl {
		t.Fatalf("Expected ID Token Request URL to be %q but got %q", builder.IDTokenRequestURL, authMethod.idTokenRequestUrl)
	}
	if builder.IDTokenRequestToken != authMethod.idTokenRequestToken {
		t.Fatalf("Expected ID Token Request Token to be %q but got %q", builder.IDTokenRequestToken, authMethod.idTokenRequestToken)
	}
}

func TestOIDC_isApplicable(t *testing.T) {
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
				SupportsOIDCAuth: false,
			},
			Valid: false,
		},
		{
			Description: "Feature Toggled on",
			Builder: Builder{
				SupportsOIDCAuth:    true,
				IDTokenRequestURL:   "https://token.endpoint",
				IDTokenRequestToken: "abcdeftoken",
				UseMicrosoftGraph:   true,
			},
			Valid: true,
		},
	}

	for _, v := range cases {
		applicable := oidcAuth{}.isApplicable(v.Builder)
		if v.Valid != applicable {
			t.Fatalf("Expected %q to be %t but got %t", v.Description, v.Valid, applicable)
		}
	}
}

func TestOIDC_populateConfig(t *testing.T) {
	config := &Config{}
	err := oidcAuth{}.populateConfig(config)
	if err != nil {
		t.Fatalf("Error populating config: %s", err)
	}
	if !config.AuthenticatedAsAServicePrincipal {
		t.Fatalf("Expected `AuthenticatedAsAServicePrincipal` to be true but it wasn't")
	}
}

func TestOIDC_validate(t *testing.T) {
	cases := []struct {
		Description string
		Config      oidcAuth
		ExpectError bool
	}{
		{
			Description: "Empty Configuration",
			Config:      oidcAuth{},
			ExpectError: true,
		},
		{
			Description: "Valid Configuration",
			Config: oidcAuth{
				auxiliaryTenantIds:  []string{"a-tenant-id", "b-tenant-id"},
				clientId:            "client-id",
				environment:         "environment",
				idTokenRequestUrl:   "https://token.endpoint",
				idTokenRequestToken: "abcdeftoken",
				tenantId:            "tenant-id",
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
