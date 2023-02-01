// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authentication

import "testing"

func TestAzureBearToken_builder(t *testing.T) {
	builder := Builder{
		BearToken: "token",
	}

	method, err := azureBearTokenAuth{}.build(builder)
	if err != nil {
		t.Fatalf("Error building Azure Bear Token Auth: %+v", err)
	}

	authMethod := method.(azureBearTokenAuth)
	if builder.BearToken != authMethod.bearToken {
		t.Fatalf("Expected bear token to be %q but got %q", builder.BearToken, authMethod.bearToken)
	}
}

func TestAzureBearToken_isApplicable(t *testing.T) {
	cases := []struct {
		Description string
		Builder     Builder
		Valid       bool
	}{
		{
			Description: "Empty Bear Token",
			Builder:     Builder{},
			Valid:       false,
		},
		{
			Description: "Valid Bear Token",
			Builder: Builder{
				BearToken: "token",
			},
			Valid: true,
		},
	}

	for _, v := range cases {
		applicable := azureBearTokenAuth{}.isApplicable(v.Builder)
		if v.Valid != applicable {
			t.Fatalf("Expected %q to be %t but got %t", v.Description, v.Valid, applicable)
		}
	}
}

func TestAzureBearToken_populateConfig(t *testing.T) {
	config := &Config{}
	err := azureBearTokenAuth{}.populateConfig(config)
	if err != nil {
		t.Fatalf("Error populating config: %s", err)
	}

	// nothing to check since it's not doing anything
}

func TestAzureBearToken_validate(t *testing.T) {
	cases := []struct {
		Description string
		Config      azureBearTokenAuth
		ExpectError bool
	}{
		{
			Description: "Empty Bear Token",
			Config:      azureBearTokenAuth{},
			ExpectError: true,
		},
		{
			Description: "Valid Bear Token",
			Config: azureBearTokenAuth{
				bearToken: "token",
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
