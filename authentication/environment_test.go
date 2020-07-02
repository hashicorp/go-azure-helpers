package authentication

import (
	"context"
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
	_, err := AzureEnvironmentByNameFromEndpoint(context.TODO(), "management.azure.com", "public")
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
	_, err = AzureEnvironmentByNameFromEndpoint(context.TODO(), "management.azure.com", "AzureCloud")
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
	_, err = AzureEnvironmentByNameFromEndpoint(context.TODO(), "badurl", "AzureCloud")
	if err == nil {
		t.Fatal("Expected error from bad endpoint")
	}
	_, err = AzureEnvironmentByNameFromEndpoint(context.TODO(), "management.azure.com", "badEnvironment")
	if err == nil {
		t.Fatal("Expected error from bad environment")
	}
}

func TestAccIsEnvironmentAzureStack(t *testing.T) {
	ok, err := IsEnvironmentAzureStack(context.TODO(),"management.azure.com", "public")
	if err != nil {
		t.Fatalf("Error getting Endpoint: %s", err)
	}
	if ok {
		t.Fatal("Expected `public` environment to not be Azure Stack")
	}
}
