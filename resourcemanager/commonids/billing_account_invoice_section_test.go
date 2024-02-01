// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import "testing"

func TestNewBillingAccountInvoiceSectionID(t *testing.T) {
	id := NewBillingAccountInvoiceSectionID("billingAccountNameValue", "billingProfileNameValue", "invoiceSectionNameValue")

	if id.BillingAccountName != "billingAccountNameValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountNameValue")
	}

	if id.BillingProfileName != "billingProfileNameValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingProfileName'", id.BillingProfileName, "billingProfileNameValue")
	}

	if id.InvoiceSectionName != "invoiceSectionNameValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InvoiceSectionName'", id.InvoiceSectionName, "invoiceSectionNameValue")
	}

}

func TestFormatBillingAccountInvoiceSectionID(t *testing.T) {
	actual := NewBillingAccountInvoiceSectionID("billingAccountNameValue", "billingProfileNameValue", "invoiceSectionNameValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/billingProfiles/billingProfileNameValue/invoiceSections/invoiceSectionNameValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingAccountInvoiceSectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingAccountInvoiceSectionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/invoiceSections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/billingProfiles/billingProfileNameValue/invoiceSections/invoiceSectionNameValue",
			Expected: &BillingAccountInvoiceSectionId{
				BillingAccountName: "billingAccountNameValue",
				BillingProfileName: "billingProfileNameValue",
				InvoiceSectionName: "invoiceSectionNameValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/billingProfiles/billingProfileNameValue/invoiceSections/invoiceSectionNameValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingAccountInvoiceSectionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingAccountName != v.Expected.BillingAccountName {
			t.Fatalf("Expected %q but got %q for BillingAccountName", v.Expected.BillingAccountName, actual.BillingAccountName)
		}

		if actual.BillingProfileName != v.Expected.BillingProfileName {
			t.Fatalf("Expected %q but got %q for BillingProfileName", v.Expected.BillingProfileName, actual.BillingProfileName)
		}

		if actual.InvoiceSectionName != v.Expected.InvoiceSectionName {
			t.Fatalf("Expected %q but got %q for InvoiceSectionName", v.Expected.InvoiceSectionName, actual.InvoiceSectionName)
		}
	}
}

func TestParseBillingAccountInvoiceSectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingAccountInvoiceSectionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts",
			Error: true,
		},
		{
			// Incomplete URI  (Insensitively)
			Input: "/pRoViVeRs/MiCrOsOfT.biLLing/BillInGAccouNts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue",
			Error: true,
		},
		{
			// Incomplete URI (Insensitively)
			Input: "/pRoViVeRs/MiCrOsOfT.biLLing/BillInGAccouNts/BillingAcCOUNTNameValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/invoiceSections",
			Error: true,
		},
		{
			// Incomplete URI (Insensitively)
			Input: "/PRoVideRs/MiCrOsOfT.biLLing/BillInGAccouNts/BillingAcCOUNTNameValue/iNvoICESections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/billingProfiles/billingProfileNameValue/invoiceSections/invoiceSectionNameValue",
			Expected: &BillingAccountInvoiceSectionId{
				BillingAccountName: "billingAccountNameValue",
				BillingProfileName: "billingProfileNameValue",
				InvoiceSectionName: "invoiceSectionNameValue",
			},
		},
		{
			// Valid URI (Insensitively)
			Input: "/ProvIdErs/MIcroSoft.Billing/billingACcOunts/billingAccOuntNameVaLue/biLlingPrOfiLes/bIllingProfileNaMevALue/InVoiceSEctions/invoiceSectionNameValue",
			Expected: &BillingAccountInvoiceSectionId{
				BillingAccountName: "billingAccOuntNameVaLue",
				BillingProfileName: "bIllingProfileNaMevALue",
				InvoiceSectionName: "invoiceSectionNameValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/billingProfiles/billingProfileNameValue/invoiceSections/invoiceSectionNameValue/extra",
			Error: true,
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/ProvIdErs/MIcroSoft.Billing/billingACcOunts/billingAccOuntNameVaLue/biLlingPrOfiLes/bIllingProfileNaMevAlue/invoiCeSections/invoiceSectionNAmeValue/exTra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingAccountInvoiceSectionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingAccountName != v.Expected.BillingAccountName {
			t.Fatalf("Expected %q but got %q for BillingAccountName", v.Expected.BillingAccountName, actual.BillingAccountName)
		}

		if actual.BillingProfileName != v.Expected.BillingProfileName {
			t.Fatalf("Expected %q but got %q for BillingProfileName", v.Expected.BillingProfileName, actual.BillingProfileName)
		}

		if actual.InvoiceSectionName != v.Expected.InvoiceSectionName {
			t.Fatalf("Expected %q but got %q for InvoiceSectionName", v.Expected.InvoiceSectionName, actual.InvoiceSectionName)
		}
	}
}

func TestSegmentsForBillingAccountInvoiceSectionId(t *testing.T) {
	segments := BillingAccountInvoiceSectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingAccountInvoiceSectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
