// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ChaosStudioCapabilityId{}

func TestNewChaosStudioCapabilityId(t *testing.T) {
	id := NewChaosStudioCapabilityID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "targetName", "capabilityName")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.TargetName != "targetName" {
		t.Fatalf("Expected %q but got %q for Segment 'Target Name'", id.TargetName, "targetName")
	}

	if id.CapabilityName != "capabilityName" {
		t.Fatalf("Expected %q but got %q for Segment 'Capability Name'", id.CapabilityName, "capabilityName")
	}
}

func TestFormatChaosStudioCapabilityId(t *testing.T) {
	actual := NewChaosStudioCapabilityID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "targetName", "capabilityName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Chaos/targets/targetName/capabilities/capabilityName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseChaosStudioCapabilityId(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ChaosStudioCapabilityId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Chaos/targets/targetName/capabilities/capabilityName",
			Expected: &ChaosStudioCapabilityId{
				Scope:          "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				TargetName:     "targetName",
				CapabilityName: "capabilityName",
			},
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseChaosStudioCapabilityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.TargetName != v.Expected.TargetName {
			t.Fatalf("Expected %q but got %q for Target Name", v.Expected.TargetName, actual.TargetName)
		}

		if actual.CapabilityName != v.Expected.CapabilityName {
			t.Fatalf("Expected %q but got %q for Capability Name", v.Expected.CapabilityName, actual.CapabilityName)
		}

	}
}

func TestParseChaosStudioCapabilityIdInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ChaosStudioCapabilityId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Chaos/targets/targetName/capabilities/capabilityName",
			Expected: &ChaosStudioCapabilityId{
				Scope:          "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				TargetName:     "targetName",
				CapabilityName: "capabilityName",
			},
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/providers/Microsoft.Chaos/TaRgEtS/targetName/CaPaBiLitIeS/capabilityName",
			Expected: &ChaosStudioCapabilityId{
				Scope:          "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				TargetName:     "targetName",
				CapabilityName: "capabilityName",
			},
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseChaosStudioCapabilityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.TargetName != v.Expected.TargetName {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.TargetName, actual.TargetName)
		}

		if actual.CapabilityName != v.Expected.CapabilityName {
			t.Fatalf("Expected %q but got %q for Capability Name", v.Expected.CapabilityName, actual.CapabilityName)
		}

	}
}

func TestSegmentsForChaosStudioCapabilityId(t *testing.T) {
	segments := ChaosStudioCapabilityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ChaosStudioTargetId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
