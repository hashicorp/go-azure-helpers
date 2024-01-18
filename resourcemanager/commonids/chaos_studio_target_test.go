// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = &ChaosStudioTargetId{}

func TestNewChaosStudioTargetId(t *testing.T) {
	id := NewChaosStudioTargetID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "targetName")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.TargetName != "targetName" {
		t.Fatalf("Expected %q but got %q for Segment 'Target Name'", id.TargetName, "targetName")
	}
}

func TestFormatChaosStudioTargetId(t *testing.T) {
	actual := NewChaosStudioTargetID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "targetName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Chaos/targets/targetName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseChaosStudioTargetId(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ChaosStudioTargetId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Chaos/targets/targetName",
			Expected: &ChaosStudioTargetId{
				Scope:      "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				TargetName: "targetName",
			},
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseChaosStudioTargetID(v.Input)
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

	}
}

func TestParseChaosStudioTargetIdInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ChaosStudioTargetId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Chaos/targets/targetName",
			Expected: &ChaosStudioTargetId{
				Scope:      "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				TargetName: "targetName",
			},
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/providers/Microsoft.Chaos/TaRgEtS/targetName",
			Expected: &ChaosStudioTargetId{
				Scope:      "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				TargetName: "targetName",
			},
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseChaosStudioTargetIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForChaosStudioTargetId(t *testing.T) {
	segments := ChaosStudioTargetId{}.Segments()
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
