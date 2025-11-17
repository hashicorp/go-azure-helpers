package keyvault

import (
	"testing"
)

func TestNewNestedItemID(t *testing.T) {
	cases := []struct {
		Scenario        string
		KeyVaultBaseURL string
		Expected        string
		ExpectError     bool
	}{
		{
			Scenario:        "empty value",
			KeyVaultBaseURL: "",
			Expected:        "",
			ExpectError:     true,
		},
		{
			Scenario:        "valid, no port",
			KeyVaultBaseURL: "https://test.vault.azure.net",
			Expected:        "https://test.vault.azure.net/keys/test/testVersionString",
			ExpectError:     false,
		},
		{
			Scenario:        "valid, with port",
			KeyVaultBaseURL: "https://test.vault.azure.net:443",
			Expected:        "https://test.vault.azure.net/keys/test/testVersionString",
			ExpectError:     false,
		},
		{
			Scenario:        "managed hsm valid, no port",
			KeyVaultBaseURL: "https://test.managedhsm.azure.net",
			Expected:        "https://test.managedhsm.azure.net/keys/test/testVersionString",
			ExpectError:     false,
		},
		{
			Scenario:        "managed hsm valid, with port",
			KeyVaultBaseURL: "https://test.managedhsm.azure.net:443",
			Expected:        "https://test.managedhsm.azure.net/keys/test/testVersionString",
			ExpectError:     false,
		},
	}

	for _, tc := range cases {
		id, err := NewNestedItemID(tc.KeyVaultBaseURL, NestedItemTypeKey, "test", "testVersionString")
		if err != nil {
			if !tc.ExpectError {
				t.Fatalf("%s: unexpected error: %+v", tc.Scenario, err)
			}

			continue
		}

		if id.ID() != tc.Expected {
			t.Fatalf("%s: expected `%s`, but got `%s`", tc.Scenario, tc.Expected, id.ID())
		}
	}
}

func TestParseNestedItemID(t *testing.T) {
	cases := []struct {
		Input       string
		Expected    NestedItemID
		ExpectError bool
	}{
		{
			Input:       "",
			ExpectError: true,
		},
		{
			Input:       "https",
			ExpectError: true,
		},
		{
			Input:       "https://",
			ExpectError: true,
		},
		{
			Input:       "https://my-keyvault.vault.azure.net",
			ExpectError: true,
		},
		{
			Input:       "https://my-keyvault.vault.azure.net/",
			ExpectError: true,
		},
		{
			Input:       "https://my-keyvault.vault.azure.net/secrets",
			ExpectError: true,
		},
		{
			Input:       "https://my-keyvault.vault.azure.net/invalidNestedItemObjectType/bird/fdf067c93bbb4b22bff4d8b7a9a56217",
			ExpectError: true,
		},
		{
			Input:       "https://my-keyvault.vault.azure.net/secrets/bird/fdf067c93bbb4b22bff4d8b7a9a56217/XXX",
			ExpectError: true,
		},
		{
			Input:       "https://my-keyvault.vault.azure.net/secrets/bird/fdf067c93bbb4b22bff4d8b7a9a56217",
			ExpectError: false,
			Expected: NestedItemID{
				Name:            "bird",
				NestedItemType:  NestedItemTypeSecret,
				KeyVaultBaseURL: "https://my-keyvault.vault.azure.net",
				Version:         "fdf067c93bbb4b22bff4d8b7a9a56217",
			},
		},
		{
			Input:       "https://my-keyvault.vault.azure.net/secrets/bird",
			ExpectError: false,
			Expected: NestedItemID{
				Name:            "bird",
				NestedItemType:  NestedItemTypeSecret,
				KeyVaultBaseURL: "https://my-keyvault.vault.azure.net",
			},
		},
		{
			Input:       "https://my-keyvault.vault.azure.net/certificates/hello/world",
			ExpectError: false,
			Expected: NestedItemID{
				Name:            "hello",
				NestedItemType:  NestedItemTypeCertificate,
				KeyVaultBaseURL: "https://my-keyvault.vault.azure.net",
				Version:         "world",
			},
		},
		{
			Input:       "https://my-keyvault.vault.azure.net/keys/castle/1492",
			ExpectError: false,
			Expected: NestedItemID{
				Name:            "castle",
				NestedItemType:  NestedItemTypeKey,
				KeyVaultBaseURL: "https://my-keyvault.vault.azure.net",
				Version:         "1492",
			},
		},
		{
			Input:       "https://my-keyvault.managedhsm.azure.net/keys/castle/1492",
			ExpectError: false,
			Expected: NestedItemID{
				Name:            "castle",
				NestedItemType:  NestedItemTypeKey,
				KeyVaultBaseURL: "https://my-keyvault.managedhsm.azure.net",
				Version:         "1492",
			},
		},
	}

	for _, tc := range cases {
		id, err := ParseNestedItemID(tc.Input, VersionTypeAny, NestedItemTypeAny)
		if err != nil {
			if !tc.ExpectError {
				t.Fatalf("%s: unexpected error: %+v", tc.Input, err)
			}

			continue
		}

		if id == nil {
			t.Fatalf("expected a valid NestedItemID for `%s`, got nil.", tc.Input)
		}

		if tc.Expected.KeyVaultBaseURL != id.KeyVaultBaseURL {
			t.Fatalf("expected `KeyVaultBaseURL` to be `%s`, got `%s` for ID `%s`", tc.Expected.KeyVaultBaseURL, id.KeyVaultBaseURL, tc.Input)
		}

		if tc.Expected.NestedItemType != id.NestedItemType {
			t.Fatalf("expected `NestedItemType` to be `%s`, got `%s` for ID `%s`", tc.Expected.NestedItemType, id.NestedItemType, tc.Input)
		}

		if tc.Expected.Name != id.Name {
			t.Fatalf("expected `Name` to be `%s`, got `%s` for ID `%s`", tc.Expected.Name, id.Name, tc.Input)
		}

		if tc.Expected.Version != id.Version {
			t.Fatalf("expected `Version` to be `%s`, got `%s` for ID `%s`", tc.Expected.Version, id.Version, tc.Input)
		}

		if tc.Input != id.ID() {
			t.Fatalf("expected `ID()` to be `%s`, got `%s`", tc.Input, id.ID())
		}
	}
}
