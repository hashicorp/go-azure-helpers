package keyvault

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
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
		id, err := NewNestedItemID(tc.KeyVaultBaseURL, NestedItemTypeKey, "test", pointer.To("testVersionString"))
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

// TODO: additional tests
