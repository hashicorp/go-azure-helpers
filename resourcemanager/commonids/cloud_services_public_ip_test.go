// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = &CloudServicesPublicIPAddressId{}

func TestNewCloudServicesPublicIPAddressID(t *testing.T) {
	id := NewCloudServicesPublicIPAddressID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServiceValue", "roleInstanceValue", "networkInterfaceValue", "ipConfigurationValue", "publicIPAddressValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
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

	if id.PublicIPAddressName != "publicIPAddressValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PublicIPAddressName'", id.PublicIPAddressName, "publicIPAddressValue")
	}
}

func TestFormatCloudServicesPublicIPAddressID(t *testing.T) {
	actual := NewCloudServicesPublicIPAddressID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServiceValue", "roleInstanceValue", "networkInterfaceValue", "ipConfigurationValue", "publicIPAddressValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses/publicIPAddressValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCloudServicesPublicIPAddressID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CloudServicesPublicIPAddressId
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
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses/publicIPAddressValue",
			Expected: &CloudServicesPublicIPAddressId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "example-resource-group",
				CloudServiceName:     "cloudServiceValue",
				RoleInstanceName:     "roleInstanceValue",
				NetworkInterfaceName: "networkInterfaceValue",
				IpConfigurationName:  "ipConfigurationValue",
				PublicIPAddressName:  "publicIPAddressValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses/publicIPAddressValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCloudServicesPublicIPAddressID(v.Input)
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

		if actual.PublicIPAddressName != v.Expected.PublicIPAddressName {
			t.Fatalf("Expected %q but got %q for PublicIPAddressName", v.Expected.PublicIPAddressName, actual.PublicIPAddressName)
		}

	}
}

func TestParseCloudServicesPublicIPAddressIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CloudServicesPublicIPAddressId
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
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/rOlEiNsTaNcEs/rOlEiNsTaNcEvAlUe/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS/iPcOnFiGuRaTiOnVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/rOlEiNsTaNcEs/rOlEiNsTaNcEvAlUe/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS/iPcOnFiGuRaTiOnVaLuE/pUbLiCiPaDdReSsEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses/publicIPAddressValue",
			Expected: &CloudServicesPublicIPAddressId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "example-resource-group",
				CloudServiceName:     "cloudServiceValue",
				RoleInstanceName:     "roleInstanceValue",
				NetworkInterfaceName: "networkInterfaceValue",
				IpConfigurationName:  "ipConfigurationValue",
				PublicIPAddressName:  "publicIPAddressValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/cloudServices/cloudServiceValue/roleInstances/roleInstanceValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses/publicIPAddressValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/rOlEiNsTaNcEs/rOlEiNsTaNcEvAlUe/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS/iPcOnFiGuRaTiOnVaLuE/pUbLiCiPaDdReSsEs/pUbLiCiPaDdReSsVaLuE",
			Expected: &CloudServicesPublicIPAddressId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "eXaMpLe-rEsOuRcE-GrOuP",
				CloudServiceName:     "cLoUdSeRvIcEvAlUe",
				RoleInstanceName:     "rOlEiNsTaNcEvAlUe",
				NetworkInterfaceName: "nEtWoRkInTeRfAcEvAlUe",
				IpConfigurationName:  "iPcOnFiGuRaTiOnVaLuE",
				PublicIPAddressName:  "pUbLiCiPaDdReSsVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cLoUdSeRvIcEs/cLoUdSeRvIcEvAlUe/rOlEiNsTaNcEs/rOlEiNsTaNcEvAlUe/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS/iPcOnFiGuRaTiOnVaLuE/pUbLiCiPaDdReSsEs/pUbLiCiPaDdReSsVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCloudServicesPublicIPAddressIDInsensitively(v.Input)
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

		if actual.PublicIPAddressName != v.Expected.PublicIPAddressName {
			t.Fatalf("Expected %q but got %q for PublicIPAddressName", v.Expected.PublicIPAddressName, actual.PublicIPAddressName)
		}

	}
}

func TestSegmentsForCloudServicesPublicIPAddressId(t *testing.T) {
	segments := CloudServicesPublicIPAddressId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CloudServicesPublicIPAddressId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
