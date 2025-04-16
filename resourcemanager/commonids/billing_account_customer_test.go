// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import "testing"

func TestNewBillingAccountCustomerID(t *testing.T) {
	id := NewBillingAccountCustomerID("billingAccountNameValue", "customerNameValue")

	if id.BillingAccountName != "billingAccountNameValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", "billingAccountNameValue", id.BillingAccountName)
	}

	if id.CustomerName != "customerNameValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomerName'", "customerNameValue", id.CustomerName)
	}
}

func TestFormatBillingAccountCustomerID(t *testing.T) {
	actual := NewBillingAccountCustomerID("billingAccountNameValue", "customerNameValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/customers/customerNameValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingAccountCustomerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingAccountCustomerId
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
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/customers/customerNameValue",
			Expected: &BillingAccountCustomerId{
				BillingAccountName: "billingAccountNameValue",
				CustomerName:       "customerNameValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/customers/customerNameValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingAccountCustomerID(v.Input)
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

		if actual.CustomerName != v.Expected.CustomerName {
			t.Fatalf("Expected %q but got %q for CustomerName", v.Expected.CustomerName, actual.CustomerName)
		}
	}
}

func TestParseBillingAccountCustomerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingAccountCustomerId
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
			Input: "/pRoVideRs/MiCrOsOfT.biLLing/BillInGAccouNts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue",
			Error: true,
		},
		{
			// Incomplete URI (Insensitively)
			Input: "/pRoVideRs/MiCrOsOfT.biLLing/BillInGAccouNts/BillingAcCOUNTNameValue",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/customers/customerNameValue",
			Expected: &BillingAccountCustomerId{
				BillingAccountName: "billingAccountNameValue",
				CustomerName:       "customerNameValue",
			},
		},
		{
			// Valid URI (Insensitively)
			Input: "/ProvIdErs/MIcroSoft.Billing/billingACcOunts/billingAccOuntNameVaLue/Customers/CustomerNaMevALue",
			Expected: &BillingAccountCustomerId{
				BillingAccountName: "billingAccOuntNameVaLue",
				CustomerName:       "CustomerNaMevALue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountNameValue/customers/customerNameValue/extra",
			Error: true,
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/ProvIdErs/MIcroSoft.Billing/billingACcOunts/billingAccOuntNameVaLue/Customers/CustomerNaMevAlue/exTra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingAccountCustomerIDInsensitively(v.Input)
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

		if actual.CustomerName != v.Expected.CustomerName {
			t.Fatalf("Expected %q but got %q for CustomerName", v.Expected.CustomerName, actual.CustomerName)
		}
	}
}

func TestSegmentsForBillingAccountCustomerId(t *testing.T) {
	segments := BillingAccountCustomerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingAccountCustomerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
