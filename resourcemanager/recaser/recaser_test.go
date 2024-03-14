package recaser_test

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
)

func TestRecaserWithIncorrectCasing(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO"
	actual := recaser.ReCase("/Subscriptions/11111/resourcegroups/bobby/Providers/Microsoft.Compute/AvailabilitySets/HeYO")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}

}

func TestRecaserWithCorrectCasing(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO"
	actual := recaser.ReCase("/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}

}

func TestRecaserWithCorrectCasingResourceGroupId(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby"
	actual := recaser.ReCase("/subscriptions/11111/resourceGroups/bobby")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithIncorrectCasingResourceGroupId(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby"
	actual := recaser.ReCase("/Subscriptions/11111/resourcegroups/bobby")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithUnknownId(t *testing.T) {
	// should return string without recasing
	expected := "/blah/11111/Blah"
	actual := recaser.ReCase("/blah/11111/Blah")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithUnkownIdContainingSubscriptions(t *testing.T) {

	expected := "/subscriptions/11111/Blah"
	actual := recaser.ReCase("/suBsCrIpTiOnS/11111/Blah")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithUnkownIdContainingSubscriptionsAndResourceGroups(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/group1/blah/"
	actual := recaser.ReCase("/suBscriptions/11111/ReSoUrCeGRoUps/group1/blah/")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithEmptyString(t *testing.T) {
	expected := ""
	actual := recaser.ReCase("")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithMultipleProviderSegmentsAndCorrectCasing(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO/providers/Microsoft.Compute"
	actual := recaser.ReCase("/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO/providers/Microsoft.Compute")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithMultipleProviderSegmentsAndIncorrectCasing(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO/providers/Microsoft.Compute"
	actual := recaser.ReCase("/Subscriptions/11111/resourcegroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO/providers/Microsoft.Compute")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithIncompleteProviderSegments(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/"
	actual := recaser.ReCase("/Subscriptions/11111/resourcegroups/bobby/providers/")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithOddNumberOfSegmentsAndCorrectCasing(t *testing.T) {
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/"
	actual := recaser.ReCase("/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithOddNumberOfSegmentsAndIncorrectCasing(t *testing.T) {
	// expect /subscriptions/ and /resourceGroups/ to be recased but not /AvaiLabilitySets/
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/AvaiLabilitySets/"
	actual := recaser.ReCase("/SubsCriptions/11111/ResourceGroups/bobby/providers/Microsoft.Compute/AvaiLabilitySets/")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}
