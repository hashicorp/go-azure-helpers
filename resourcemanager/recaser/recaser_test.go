package recaser

import (
	"strings"
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

func TestReCaserWithIncorrectCasing(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO"

	actual := reCaseWithIds("/Subscriptions/11111/resourcegroups/bobby/Providers/Microsoft.Compute/AvailabilitySets/HeYO", getTestIds())
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithCorrectCasing(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO"
	actual := reCaseWithIds("/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}

}

func TestReCaserWithCorrectCasingResourceGroupId(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby"
	actual := reCaseWithIds("/subscriptions/11111/resourceGroups/bobby", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithIncorrectCasingResourceGroupId(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby"
	actual := reCaseWithIds("/Subscriptions/11111/resourcegroups/bobby", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithUnknownId(t *testing.T) {
	// should return string without recasing
	expected := "/blah/11111/Blah"
	actual := reCaseWithIds("/blah/11111/Blah", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithUnkownIdContainingSubscriptions(t *testing.T) {

	expected := "/subscriptions/11111/Blah"
	actual := reCaseWithIds("/suBsCrIpTiOnS/11111/Blah", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithUnkownIdContainingSubscriptionsAndResourceGroups(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/group1/blah/"
	actual := reCaseWithIds("/suBscriptions/11111/ReSoUrCeGRoUps/group1/blah/", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithEmptyString(t *testing.T) {
	expected := ""
	actual := reCaseWithIds("", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithMultipleProviderSegmentsAndCorrectCasing(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO/providers/Microsoft.Compute"
	actual := reCaseWithIds("/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO/providers/Microsoft.Compute", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithMultipleProviderSegmentsAndIncorrectCasing(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO/providers/Microsoft.Compute"
	actual := reCaseWithIds("/Subscriptions/11111/resourcegroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO/providers/Microsoft.Compute", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithIncompleteProviderSegments(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/"
	actual := reCaseWithIds("/Subscriptions/11111/resourcegroups/bobby/providers/", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithOddNumberOfSegmentsAndCorrectCasing(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/"
	actual := reCaseWithIds("/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithOddNumberOfSegmentsAndIncorrectCasing(t *testing.T) {
	// expect /subscriptions/ and /resourceGroups/ to be recased but not /AvaiLabilitySets/
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/AvaiLabilitySets/"
	actual := reCaseWithIds("/SubsCriptions/11111/ResourceGroups/bobby/providers/Microsoft.Compute/AvaiLabilitySets/", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithURIAndCorrectCasing(t *testing.T) {
	expected := "https://management.azure.com:80/subscriptions/12345"
	actual := reCaseWithIds("https://management.azure.com:80/subscriptions/12345", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithURIAndIncorrectCasing(t *testing.T) {
	expected := "https://management.azure.com:80/subscriptions/12345"
	actual := reCaseWithIds("https://management.azure.com:80/SuBsCriPTions/12345", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReCaserWithDataPlaneURI(t *testing.T) {
	expected := "https://example.blob.storage.azure.com/container1"
	actual := reCaseWithIds("https://example.blob.storage.azure.com/container1", getTestIds())

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func getTestIds() map[string]resourceids.ResourceId {
	return map[string]resourceids.ResourceId{
		strings.ToLower(commonids.AppServiceId{}.ID()):      &commonids.AppServiceId{},
		strings.ToLower(commonids.AvailabilitySetId{}.ID()): &commonids.AvailabilitySetId{},
		strings.ToLower(commonids.BotServiceId{}.ID()):      &commonids.BotServiceId{},
	}
}
