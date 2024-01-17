// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = &VirtualMachineScaleSetPublicIPAddressId{}

func TestNewVirtualMachineScaleSetPublicIPAddressID(t *testing.T) {
	id := NewVirtualMachineScaleSetPublicIPAddressID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetValue", "virtualMachineIndexValue", "networkInterfaceValue", "ipConfigurationValue", "publicIpAddressValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.VirtualMachineScaleSetName != "virtualMachineScaleSetValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VirtualMachineScaleSetName'", id.VirtualMachineScaleSetName, "virtualMachineScaleSetValue")
	}

	if id.VirtualMachineIndex != "virtualMachineIndexValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VirtualMachineIndex'", id.VirtualMachineIndex, "virtualMachineIndexValue")
	}

	if id.NetworkInterfaceName != "networkInterfaceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'NetworkInterfaceName'", id.NetworkInterfaceName, "networkInterfaceValue")
	}

	if id.IpConfigurationName != "ipConfigurationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IpConfigurationName'", id.IpConfigurationName, "ipConfigurationValue")
	}

	if id.PublicIpAddressName != "publicIpAddressValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PublicIpAddressName'", id.PublicIpAddressName, "publicIpAddressValue")
	}
}

func TestFormatVirtualMachineScaleSetPublicIPAddressID(t *testing.T) {
	actual := NewVirtualMachineScaleSetPublicIPAddressID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetValue", "virtualMachineIndexValue", "networkInterfaceValue", "ipConfigurationValue", "publicIpAddressValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses/publicIpAddressValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVirtualMachineScaleSetPublicIPAddressID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualMachineScaleSetPublicIPAddressId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue/ipConfigurations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses/publicIpAddressValue",
			Expected: &VirtualMachineScaleSetPublicIPAddressId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				VirtualMachineScaleSetName: "virtualMachineScaleSetValue",
				VirtualMachineIndex:        "virtualMachineIndexValue",
				NetworkInterfaceName:       "networkInterfaceValue",
				IpConfigurationName:        "ipConfigurationValue",
				PublicIpAddressName:        "publicIpAddressValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses/publicIpAddressValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualMachineScaleSetPublicIPAddressID(v.Input)
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

		if actual.VirtualMachineScaleSetName != v.Expected.VirtualMachineScaleSetName {
			t.Fatalf("Expected %q but got %q for VirtualMachineScaleSetName", v.Expected.VirtualMachineScaleSetName, actual.VirtualMachineScaleSetName)
		}

		if actual.VirtualMachineIndex != v.Expected.VirtualMachineIndex {
			t.Fatalf("Expected %q but got %q for VirtualMachineIndex", v.Expected.VirtualMachineIndex, actual.VirtualMachineIndex)
		}

		if actual.NetworkInterfaceName != v.Expected.NetworkInterfaceName {
			t.Fatalf("Expected %q but got %q for NetworkInterfaceName", v.Expected.NetworkInterfaceName, actual.NetworkInterfaceName)
		}

		if actual.IpConfigurationName != v.Expected.IpConfigurationName {
			t.Fatalf("Expected %q but got %q for IpConfigurationName", v.Expected.IpConfigurationName, actual.IpConfigurationName)
		}

		if actual.PublicIpAddressName != v.Expected.PublicIpAddressName {
			t.Fatalf("Expected %q but got %q for PublicIpAddressName", v.Expected.PublicIpAddressName, actual.PublicIpAddressName)
		}

	}
}

func TestParseVirtualMachineScaleSetPublicIPAddressIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualMachineScaleSetPublicIPAddressId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs/vIrTuAlMaChInEiNdExVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs/vIrTuAlMaChInEiNdExVaLuE/nEtWoRkInTeRfAcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs/vIrTuAlMaChInEiNdExVaLuE/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue/ipConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs/vIrTuAlMaChInEiNdExVaLuE/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs/vIrTuAlMaChInEiNdExVaLuE/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS/iPcOnFiGuRaTiOnVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs/vIrTuAlMaChInEiNdExVaLuE/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS/iPcOnFiGuRaTiOnVaLuE/pUbLiCiPaDdReSsEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses/publicIpAddressValue",
			Expected: &VirtualMachineScaleSetPublicIPAddressId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				VirtualMachineScaleSetName: "virtualMachineScaleSetValue",
				VirtualMachineIndex:        "virtualMachineIndexValue",
				NetworkInterfaceName:       "networkInterfaceValue",
				IpConfigurationName:        "ipConfigurationValue",
				PublicIpAddressName:        "publicIpAddressValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/virtualMachineIndexValue/networkInterfaces/networkInterfaceValue/ipConfigurations/ipConfigurationValue/publicIPAddresses/publicIpAddressValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs/vIrTuAlMaChInEiNdExVaLuE/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS/iPcOnFiGuRaTiOnVaLuE/pUbLiCiPaDdReSsEs/pUbLiCiPaDdReSsVaLuE",
			Expected: &VirtualMachineScaleSetPublicIPAddressId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "eXaMpLe-rEsOuRcE-GrOuP",
				VirtualMachineScaleSetName: "vIrTuAlMaChInEsCaLeSeTvAlUe",
				VirtualMachineIndex:        "vIrTuAlMaChInEiNdExVaLuE",
				NetworkInterfaceName:       "nEtWoRkInTeRfAcEvAlUe",
				IpConfigurationName:        "iPcOnFiGuRaTiOnVaLuE",
				PublicIpAddressName:        "pUbLiCiPaDdReSsVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs/vIrTuAlMaChInEiNdExVaLuE/nEtWoRkInTeRfAcEs/nEtWoRkInTeRfAcEvAlUe/iPcOnFiGuRaTiOnS/iPcOnFiGuRaTiOnVaLuE/pUbLiCiPaDdReSsEs/pUbLiCiPaDdReSsVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualMachineScaleSetPublicIPAddressIDInsensitively(v.Input)
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

		if actual.VirtualMachineScaleSetName != v.Expected.VirtualMachineScaleSetName {
			t.Fatalf("Expected %q but got %q for VirtualMachineScaleSetName", v.Expected.VirtualMachineScaleSetName, actual.VirtualMachineScaleSetName)
		}

		if actual.VirtualMachineIndex != v.Expected.VirtualMachineIndex {
			t.Fatalf("Expected %q but got %q for VirtualMachineIndex", v.Expected.VirtualMachineIndex, actual.VirtualMachineIndex)
		}

		if actual.NetworkInterfaceName != v.Expected.NetworkInterfaceName {
			t.Fatalf("Expected %q but got %q for NetworkInterfaceName", v.Expected.NetworkInterfaceName, actual.NetworkInterfaceName)
		}

		if actual.IpConfigurationName != v.Expected.IpConfigurationName {
			t.Fatalf("Expected %q but got %q for IpConfigurationName", v.Expected.IpConfigurationName, actual.IpConfigurationName)
		}

		if actual.PublicIpAddressName != v.Expected.PublicIpAddressName {
			t.Fatalf("Expected %q but got %q for PublicIpAddressName", v.Expected.PublicIpAddressName, actual.PublicIpAddressName)
		}

	}
}

func TestSegmentsForVirtualMachineScaleSetPublicIPAddressId(t *testing.T) {
	segments := VirtualMachineScaleSetPublicIPAddressId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VirtualMachineScaleSetPublicIPAddressId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
