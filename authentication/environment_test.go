package authentication

import (
	"context"
	"github.com/hashicorp/go-azure-helpers/testhelpers"
	"testing"
)

func TestAzureEnvironmentNames(t *testing.T) {
	testData := map[string]string{
		"":                       "public",
		"AzureChinaCloud":        "china",
		"AzureCloud":             "public",
		"AzureGermanCloud":       "german",
		"AZUREUSGOVERNMENTCLOUD": "usgovernment",
		"AzurePublicCloud":       "public",
	}

	for input, expected := range testData {
		actual := normalizeEnvironmentName(input)
		if actual != expected {
			t.Fatalf("Expected %q for input %q: got %q!", expected, input, actual)
		}
	}
}

func TestAccAzureEnvironmentByNameFromEndpoint(t *testing.T) {
	c, err := testhelpers.BuildAuthClient()
	if err != nil {
		t.Fatalf("Error building client: %s", err)
	}

	_, err = AzureEnvironmentByNameFromEndpoint(context.TODO(), c.MetadataURL, c.Environment)
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
}

func TestAccIsEnvironmentAzureStack(t *testing.T) {
	c, err := testhelpers.BuildAuthClient()
	if err != nil {
		t.Fatalf("Error building client: %s", err)
	}

	_, err = IsEnvironmentAzureStack(context.TODO(), c.MetadataURL, c.Environment)
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
}