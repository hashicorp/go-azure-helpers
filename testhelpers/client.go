package testhelpers

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-azure-helpers/authentication"
)

func BuildAuthClient() (*authentication.Config, error) {
	builder := &authentication.Builder{
		SubscriptionID: os.Getenv("ARM_SUBSCRIPTION_ID"),
		ClientID:       os.Getenv("ARM_CLIENT_ID"),
		ClientSecret:   os.Getenv("ARM_CLIENT_SECRET"),
		TenantID:       os.Getenv("ARM_TENANT_ID"),
		Environment:    os.Getenv("ARM_ENVIRONMENT"),

		// Feature Toggles
		SupportsClientSecretAuth: true,
	}

	c, err := builder.Build()
	if err != nil {
		return nil, fmt.Errorf("Error building AzureRM Client: %s", err)
	}

	return c, nil
}
