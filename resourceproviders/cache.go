package resourceproviders

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/2017-03-09/resources/mgmt/resources"
)

// cachedResourceProviders can be (validly) nil - as such this shouldn't be relied on
var cachedResourceProviders *[]resources.Provider

// CacheSupportedProviders attempts to retrieve the supported Resource Providers from the Resource Manager API
// and caches them, for used in enhanced validation
func CacheSupportedProviders(ctx context.Context, client *resources.ProvidersClient) error {
	if cachedResourceProviders != nil {
		return nil
	}
	providers, err := availableResourceProviders(ctx, client)
	if err != nil {
		return err
	}
	cached := make([]resources.Provider, 0)
	for _, provider := range providers {
		cached = append(cached, resources.Provider{
			Namespace:         provider.Namespace,
			RegistrationState: provider.RegistrationState,
		})
	}
	cachedResourceProviders = &cached
	return nil
}
