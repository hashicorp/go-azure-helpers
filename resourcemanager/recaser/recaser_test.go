package recaser_test

import (
	"log"
	"testing"

	ids "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
)

func TestRecaserWithIncorrectCasing(t *testing.T) {
	// init package
	resourceGroupId := ids.NewResourceGroupID("0000", "hello")
	log.Printf(resourceGroupId.ID())
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO"
	actual := recaser.ReCase("/subscriptions/11111/resourcegroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}

}

func TestRecaserWithCorrectCasing(t *testing.T) {
	// init package
	resourceGroupId := ids.NewResourceGroupID("0000", "hello")
	log.Printf(resourceGroupId.ID())
	expected := "/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO"
	actual := recaser.ReCase("/subscriptions/11111/resourceGroups/bobby/providers/Microsoft.Compute/availabilitySets/HeYO")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}

}

func TestRecaserWithCorrectCasingResourceGroupId(t *testing.T) {
	// init package
	resourceGroupId := ids.NewResourceGroupID("0000", "hello")
	log.Printf(resourceGroupId.ID())
	expected := "/subscriptions/11111/resourceGroups/bobby"
	actual := recaser.ReCase("/subscriptions/11111/resourceGroups/bobby")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithIncorrectCasingResourceGroupId(t *testing.T) {
	// init package
	resourceGroupId := ids.NewResourceGroupID("0000", "hello")
	log.Printf(resourceGroupId.ID())
	expected := "/subscriptions/11111/resourceGroups/bobby"
	actual := recaser.ReCase("/Subscriptions/11111/resourcegroups/bobby")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithUnknownId(t *testing.T) {
	// init package
	resourceGroupId := ids.NewResourceGroupID("0000", "hello")
	log.Printf(resourceGroupId.ID())

	// should return string without recasing
	expected := "/blah/11111/Blah"
	actual := recaser.ReCase("/blah/11111/Blah")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithUnkownIdContainingSubscriptions(t *testing.T) {
	// init package
	resourceGroupId := ids.NewResourceGroupID("0000", "hello")
	log.Printf(resourceGroupId.ID())

	expected := "/subscriptions/11111/Blah"
	actual := recaser.ReCase("/suBsCrIpTiOnS/11111/Blah")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithUnkownIdContainingSubscriptionsAndResourceGroups(t *testing.T) {
	// init package
	resourceGroupId := ids.NewResourceGroupID("0000", "hello")
	log.Printf(resourceGroupId.ID())

	expected := "/subscriptions/11111/resourceGroups/group1/blah/"
	actual := recaser.ReCase("/suBscriptions/11111/ReSoUrCeGRoUps/group1/blah/")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestRecaserWithEmptyString(t *testing.T) {
	// init package
	resourceGroupId := ids.NewResourceGroupID("0000", "hello")
	log.Printf(resourceGroupId.ID())

	expected := ""
	actual := recaser.ReCase("")

	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}
