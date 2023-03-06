package resourceproviders

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/2017-03-09/resources/mgmt/resources"
)

func TestEnhancedValidationEnabled(t *testing.T) {
	testCases := []struct {
		input string
		valid bool
	}{
		{
			input: "",
			valid: false,
		},
		{
			input: "micr0soft",
			valid: false,
		},
		{
			input: "microsoft.compute",
			valid: false,
		},
		{
			input: "Microsoft.Compute",
			valid: true,
		},
	}
	namespace := "Microsoft.Compute"
	cachedResourceProviders = &[]resources.Provider{{Namespace: &namespace}}
	defer func() {
		cachedResourceProviders = nil
	}()

	for _, testCase := range testCases {
		t.Logf("Testing %q..", testCase.input)

		warnings, errors := EnhancedValidate(testCase.input, "name")
		valid := len(warnings) == 0 && len(errors) == 0
		if testCase.valid != valid {
			t.Errorf("Expected %t but got %t", testCase.valid, valid)
		}
	}
}
