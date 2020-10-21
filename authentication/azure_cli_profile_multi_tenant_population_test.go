package authentication

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/azure/cli"
)

func TestAzureCliProfileMultiTenant_populateSubscriptionIdMissing(t *testing.T) {
	cliProfile := azureCLIProfileMultiTenant{
		profile: cli.Profile{
			Subscriptions: []cli.Subscription{},
		},
	}

	err := cliProfile.populateSubscriptionID()
	if err == nil {
		t.Fatalf("Expected an error to be returned - but didn't get one")
	}
}

func TestAzureCliProfileMultiTenant_populateSubscriptionIdNoDefault(t *testing.T) {
	cliProfile := azureCLIProfileMultiTenant{
		profile: cli.Profile{
			Subscriptions: []cli.Subscription{
				{
					IsDefault: false,
					ID:        "abc123",
				},
			},
		},
	}

	err := cliProfile.populateSubscriptionID()
	if err == nil {
		t.Fatalf("Expected an error to be returned - but didn't get one")
	}
}

func TestAzureCliProfileMultiTenant_populateSubscriptionIdValid(t *testing.T) {
	subscriptionId := "abc123"
	cliProfile := azureCLIProfileMultiTenant{
		profile: cli.Profile{
			Subscriptions: []cli.Subscription{
				{
					IsDefault: true,
					ID:        subscriptionId,
				},
			},
		},
	}

	err := cliProfile.populateSubscriptionID()
	if err != nil {
		t.Fatalf("Expected no error to be returned - but got: %+v", err)
	}

	if cliProfile.subscriptionId != subscriptionId {
		t.Fatalf("Expected the Subscription ID to be %q but got %q", subscriptionId, cliProfile.subscriptionId)
	}
}

func TestAzureCliProfileMultiTenant_populateTenantIdEmpty(t *testing.T) {
	cliProfile := azureCLIProfileMultiTenant{
		profile: cli.Profile{
			Subscriptions: []cli.Subscription{},
		},
	}

	err := cliProfile.populateEnvironment()
	if err == nil {
		t.Fatalf("Expected an error to be returned - but didn't get one")
	}
}

func TestAzureCliProfileMultiTenant_populateTenantIdMissingSubscription(t *testing.T) {
	cliProfile := azureCLIProfileMultiTenant{
		subscriptionId: "bcd234",
		profile: cli.Profile{
			Subscriptions: []cli.Subscription{
				{
					IsDefault: false,
					ID:        "abc123",
				},
			},
		},
	}

	err := cliProfile.populateTenantID()
	if err == nil {
		t.Fatalf("Expected an error to be returned - but didn't get one")
	}
}

func TestAzureCliProfileMultiTenant_populateTenantIdValid(t *testing.T) {
	cliProfile := azureCLIProfileMultiTenant{
		subscriptionId: "abc123",
		profile: cli.Profile{
			Subscriptions: []cli.Subscription{
				{
					IsDefault: false,
					ID:        "abc123",
					TenantID:  "bcd234",
				},
			},
		},
	}

	err := cliProfile.populateTenantID()
	if err != nil {
		t.Fatalf("Expected no error to be returned - but got: %+v", err)
	}

	if cliProfile.subscriptionId != "abc123" {
		t.Fatalf("Expected Subscription ID to be 'abc123' - got %q", cliProfile.subscriptionId)
	}

	if cliProfile.tenantId != "bcd234" {
		t.Fatalf("Expected Tenant ID to be 'bcd234' - got %q", cliProfile.tenantId)
	}
}
