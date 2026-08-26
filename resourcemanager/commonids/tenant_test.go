// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = &TenantId{}

func TestNewTenantID(t *testing.T) {
	id := NewTenantID("12345678-1234-9876-4563-123456789012")

	if id.TenantId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'TenantId'", id.TenantId, "12345678-1234-9876-4563-123456789012")
	}
}

func TestFormatTenantID(t *testing.T) {
	actual := NewTenantID("12345678-1234-9876-4563-123456789012").ID()
	expected := "/tenants/12345678-1234-9876-4563-123456789012"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseTenantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TenantId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/tenants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/tenants/12345678-1234-9876-4563-123456789012",
			Expected: &TenantId{
				TenantId: "12345678-1234-9876-4563-123456789012",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/tenants/12345678-1234-9876-4563-123456789012/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTenantID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TenantId != v.Expected.TenantId {
			t.Fatalf("Expected %q but got %q for TenantId", v.Expected.TenantId, actual.TenantId)
		}

	}
}

func TestParseTenantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TenantId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/tenants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/TeNaNtS/",
			Error: true,
		},
		{
			// Valid URI
			Input: "/tenants/12345678-1234-9876-4563-123456789012",
			Expected: &TenantId{
				TenantId: "12345678-1234-9876-4563-123456789012",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/tenants/12345678-1234-9876-4563-123456789012/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/TeNaNtS/12345678-1234-9876-4563-123456789012",
			Expected: &TenantId{
				TenantId: "12345678-1234-9876-4563-123456789012",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/TeNaNtS/12345678-1234-9876-4563-123456789012/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTenantIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TenantId != v.Expected.TenantId {
			t.Fatalf("Expected %q but got %q for TenantId", v.Expected.TenantId, actual.TenantId)
		}

	}
}

func TestSegmentsForTenantId(t *testing.T) {
	segments := TenantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("TenantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
