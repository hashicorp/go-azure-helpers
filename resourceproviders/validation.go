package resourceproviders

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// EnhancedValidate returns a validation function which attempts to validate the Resource Provider
// against the list of Resource Provider supported by this Azure Environment.
//
// NOTE: this is best-effort - if the users offline, or the API doesn't return it we'll
// fall back to the original approach
func EnhancedValidate(i interface{}, k string) ([]string, []error) {
	if cachedResourceProviders == nil {
		return validation.StringIsNotEmpty(i, k)
	}

	return enhancedValidation(i, k)
}

func enhancedValidation(i interface{}, k string) ([]string, []error) {
	v, ok := i.(string)
	if !ok {
		return nil, []error{fmt.Errorf("expected type of %q to be string", k)}
	}

	if v == "" {
		return nil, []error{fmt.Errorf("%q must not be empty", k)}
	}

	// enhanced validation is unavailable, but we're in this method..
	if cachedResourceProviders == nil {
		return nil, nil
	}

	found := false
	for _, provider := range *cachedResourceProviders {
		if provider.Namespace != nil && *provider.Namespace == v {
			found = true
		}
	}

	if !found {
		cachedResourceProvidersNames := make([]string, 0)
		for _, provider := range *cachedResourceProviders {
			if provider.Namespace != nil {
				cachedResourceProvidersNames = append(cachedResourceProvidersNames, *provider.Namespace)
			}
		}
		providersJoined := strings.Join(cachedResourceProvidersNames, ", ")
		return nil, []error{
			fmt.Errorf("%q was not found in the list of supported Resource Providers: %q", v, providersJoined),
		}
	}

	return nil, nil
}
