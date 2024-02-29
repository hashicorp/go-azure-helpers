// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import "testing"

func TestNewBillingEnrollmentAccountID(t *testing.T) {
	id := NewBillingEnrollmentAccountID("enrollmentAccountNameValue")

	if id.EnrollmentAccountName != "enrollmentAccountNameValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentAccountName'", id.EnrollmentAccountName, "enrollmentAccountNameValue")
	}
}

func TestFormatBillingEnrollmentAccountID(t *testing.T) {
	actual := NewBillingEnrollmentAccountID("enrollmentAccountNameValue").ID()
	expected := "/providers/Microsoft.Billing/enrollmentAccounts/enrollmentAccountNameValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingEnrollmentAccountID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingEnrollmentAccountId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/enrollmentAccounts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/enrollmentAccounts/enrollmentAccountNameValue",
			Expected: &BillingEnrollmentAccountId{
				EnrollmentAccountName: "enrollmentAccountNameValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/enrollmentAccounts/enrollmentAccountNameValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingEnrollmentAccountID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EnrollmentAccountName != v.Expected.EnrollmentAccountName {
			t.Fatalf("Expected %q but got %q for EnrollmentAccountName", v.Expected.EnrollmentAccountName, actual.EnrollmentAccountName)
		}
	}
}

func TestParseBillingEnrollmentAccountIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingEnrollmentAccountId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (Insensitively)
			Input: "/pRoVideRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing",
			Error: true,
		},
		{
			// Incomplete URI (Insensitively)
			Input: "/ProvIders/MicroSOFT.biLLing",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/enrollmentAccounts",
			Error: true,
		},
		{
			// Incomplete URI  (Insensitively)
			Input: "/pRoVideRs/MiCrOsOfT.biLLing/EnRoLlMeNtACcOuNts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/enrollmentAccounts/enrollmentAccountNameValue",
			Expected: &BillingEnrollmentAccountId{
				EnrollmentAccountName: "enrollmentAccountNameValue",
			},
		},
		{
			// Valid URI (Insensitively)
			Input: "/pRoVideRs/MiCrOsOfT.biLLing/EnRoLlMeNtACcOuNts/enRolLmEntAcCouNtNaMeVaLue",
			Expected: &BillingEnrollmentAccountId{
				EnrollmentAccountName: "enRolLmEntAcCouNtNaMeVaLue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/enrollmentAccounts/enrollmentAccountNameValue/extra",
			Error: true,
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/pRoVideRs/MiCrOsOfT.biLLing/EnRoLlMeNtACcOuNts/enRolLmEntAcCouNtNaMeVaLue/exTra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingEnrollmentAccountIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EnrollmentAccountName != v.Expected.EnrollmentAccountName {
			t.Fatalf("Expected %q but got %q for EnrollmentAccountName", v.Expected.EnrollmentAccountName, actual.EnrollmentAccountName)
		}

	}
}

func TestSegmentsForBillingEnrollmentAccountId(t *testing.T) {
	segments := BillingEnrollmentAccountId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingEnrollmentAccountId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
