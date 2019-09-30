package authentication

import "testing"

func TestManagedServiceIdentity_builder(t *testing.T) {
	builder := Builder{
		MsiEndpoint: "https://hello-world",
		ClientID:    "some-client-id",
	}

	method, err := managedServiceIdentityAuth{}.build(builder)
	if err != nil {
		t.Fatalf("Error building MSI Identity Auth: %+v", err)
	}

	authMethod := method.(managedServiceIdentityAuth)
	if builder.MsiEndpoint != authMethod.msiEndpoint {
		t.Fatalf("Expected MSI Endpoint to be %q but got %q", builder.MsiEndpoint, authMethod.msiEndpoint)
	}
	if builder.ClientID != authMethod.clientID {
		t.Fatalf("Expected MSI Client ID to be %q but got %q", builder.ClientID, authMethod.clientID)
	}
}

func TestManagedServiceIdentity_isApplicable(t *testing.T) {
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
				SupportsManagedServiceIdentity: false,
			},
			Valid: false,
		},
		{
			Description: "Feature Toggled on",
			Builder: Builder{
				SupportsManagedServiceIdentity: true,
			},
			Valid: false,
		},
	}

	for _, v := range cases {
		applicable := servicePrincipalClientSecretAuth{}.isApplicable(v.Builder)
		if v.Valid != applicable {
			t.Fatalf("Expected %q to be %t but got %t", v.Description, v.Valid, applicable)
		}
	}
}

func TestManagedServiceIdentity_populateConfig(t *testing.T) {
	config := &Config{}
	err := servicePrincipalClientSecretAuth{}.populateConfig(config)
	if err != nil {
		t.Fatalf("Error populating config: %s", err)
	}

	// nothing to check since it's not doing anything
}

func TestManagedServiceIdentity_validate(t *testing.T) {
	cases := []struct {
		Description string
		Config      managedServiceIdentityAuth
		ExpectError bool
	}{
		{
			Description: "Empty Configuration",
			Config:      managedServiceIdentityAuth{},
			ExpectError: true,
		},
		{
			Description: "Valid Configuration",
			Config: managedServiceIdentityAuth{
				msiEndpoint: "https://some-location",
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
