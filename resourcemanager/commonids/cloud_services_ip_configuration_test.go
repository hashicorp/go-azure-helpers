package commonids

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CloudServicesIPConfigurationId{}

func TestNewCloudServicesIPConfigurationID(t *testing.T) {
	id := NewCloudServicesIPConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServiceValue", "roleInstanceValue", "networkInterfaceValue", "ipConfigurationValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroup'", id.ResourceGroupName, "example-resource-group")
	}

	if id.CloudServiceName != "cloudServiceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudServiceName'", id.CloudServiceName, "cloudServiceValue")
	}

	if id.RoleInstanceName != "roleInstanceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleInstanceName'", id.RoleInstanceName, "roleInstanceValue")
	}

	if id.NetworkInterfaceName != "networkInterfaceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'NetworkInterfaceName'", id.NetworkInterfaceName, "networkInterfaceValue")
	}

	if id.IpConfigurationName != "ipConfigurationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IpConfigurationName'", id.IpConfigurationName, "ipConfigurationValue")
	}
}

func TestFormatCloudServicesIPConfigurationID(t *testing.T) {
	actual := NewCloudServicesIPConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServiceValue", "roleInstanceValue", "networkInterfaceValue", "ipConfigurationValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCloudServicesIPConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CloudServicesIPConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue",
			Expected: &CloudServicesIPConfigurationId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "example-resource-group",
				CloudServiceName:     "cloudServiceValue",
				RoleInstanceName:     "roleInstanceValue",
				NetworkInterfaceName: "networkInterfaceValue",
				IpConfigurationName:  "ipConfigurationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCloudServicesIPConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.CloudServiceName != v.Expected.CloudServiceName {
			t.Fatalf("Expected %q but got %q for CloudServiceName", v.Expected.CloudServiceName, actual.CloudServiceName)
		}

		if actual.RoleInstanceName != v.Expected.RoleInstanceName {
			t.Fatalf("Expected %q but got %q for RoleInstanceName", v.Expected.RoleInstanceName, actual.RoleInstanceName)
		}

		if actual.NetworkInterfaceName != v.Expected.NetworkInterfaceName {
			t.Fatalf("Expected %q but got %q for NetworkInterfaceName", v.Expected.NetworkInterfaceName, actual.NetworkInterfaceName)
		}

		if actual.IpConfigurationName != v.Expected.IpConfigurationName {
			t.Fatalf("Expected %q but got %q for IpConfigurationName", v.Expected.IpConfigurationName, actual.IpConfigurationName)
		}

	}
}

func TestParseCloudServicesIPConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CloudServicesIPConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/rOlEiNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/rOlEiNsTaNcEs/rOlEiNsTaNcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/rOlEiNsTaNcEs/rOlEiNsTaNcEvAlUe/nEtWoRkInTeRfAcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/rOlEiNsTaNcEs/rOlEiNsTaNcEvAlUe/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/rOlEiNsTaNcEs/rOlEiNsTaNcEvAlUe/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue",
			Expected: &CloudServicesIPConfigurationId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "example-resource-group",
				CloudServiceName:     "cloudServiceValue",
				RoleInstanceName:     "roleInstanceValue",
				NetworkInterfaceName: "networkInterfaceValue",
				IpConfigurationName:  "ipConfigurationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/rOlEiNsTaNcEs/rOlEiNsTaNcEvAlUe/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS/iPcOnFiGuRaTiOnVaLuE",
			Expected: &CloudServicesIPConfigurationId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "eXaMpLe-rEsOuRcE-GrOuP",
				CloudServiceName:     "cLoUdSeRvIcEvAlUe",
				RoleInstanceName:     "rOlEiNsTaNcEvAlUe",
				NetworkInterfaceName: "nEtWoRkInTeRfAcEvAlUe",
				IpConfigurationName:  "iPcOnFiGuRaTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/rOlEiNsTaNcEs/rOlEiNsTaNcEvAlUe/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS/iPcOnFiGuRaTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCloudServicesIPConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.CloudServiceName != v.Expected.CloudServiceName {
			t.Fatalf("Expected %q but got %q for CloudServiceName", v.Expected.CloudServiceName, actual.CloudServiceName)
		}

		if actual.RoleInstanceName != v.Expected.RoleInstanceName {
			t.Fatalf("Expected %q but got %q for RoleInstanceName", v.Expected.RoleInstanceName, actual.RoleInstanceName)
		}

		if actual.NetworkInterfaceName != v.Expected.NetworkInterfaceName {
			t.Fatalf("Expected %q but got %q for NetworkInterfaceName", v.Expected.NetworkInterfaceName, actual.NetworkInterfaceName)
		}

		if actual.IpConfigurationName != v.Expected.IpConfigurationName {
			t.Fatalf("Expected %q but got %q for IpConfigurationName", v.Expected.IpConfigurationName, actual.IpConfigurationName)
		}

	}
}

func TestSegmentsForCloudServicesIPConfigurationId(t *testing.T) {
	segments := CloudServicesIPConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CloudServicesIPConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
