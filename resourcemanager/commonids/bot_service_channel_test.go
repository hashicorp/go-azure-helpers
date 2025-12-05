// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package commonids

import "testing"

func TestNewBotServiceChannelID(t *testing.T) {
	id := NewBotServiceChannelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceValue", EmailBotServiceChannelType)

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.BotServiceName != "botServiceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BotServiceName'", id.BotServiceName, "botServiceValue")
	}

	if id.ChannelType != EmailBotServiceChannelType {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelType'", id.ChannelType, EmailBotServiceChannelType)
	}
}

func TestFormatBotServiceChannelID(t *testing.T) {
	actual := NewBotServiceChannelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "botServiceValue", EmailBotServiceChannelType).ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue/channels/EmailChannel"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBotServiceChannelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BotServiceChannelId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotServices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue/channels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue/channels/EmailChannel",
			Expected: &BotServiceChannelId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				BotServiceName:    "botServiceValue",
				ChannelType:       EmailBotServiceChannelType,
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue/channels/EmailChannel/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBotServiceChannelID(v.Input)
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

		if actual.BotServiceName != v.Expected.BotServiceName {
			t.Fatalf("Expected %q but got %q for BotServiceName", v.Expected.BotServiceName, actual.BotServiceName)
		}
	}
}

func TestParseBotServiceChannelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BotServiceChannelId
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
			// Incomplete URI (Insensitively)
			Input: "/sUbScRiPtIoNs",
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
			// Incomplete URI (Insensitively)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-ReSoUrCe-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (Insensitively)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-ReSoUrCe-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotServices",
			Error: true,
		},
		{
			// Incomplete URI (Insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-ReSoUrCe-GrOuP/pRoViDeRs/MiCrOsOfT.BoTsErViCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices",
			Error: true,
		},
		{
			// Incomplete URI (Insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-ReSoUrCe-GrOuP/pRoViDeRs/MiCrOsOfT.BoTsErViCe/BoTsErViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue",
			Error: true,
		},
		{
			// Incomplete URI (Insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-ReSoUrCe-GrOuP/pRoViDeRs/MiCrOsOfT.BoTsErViCe/BoTsErViCeS/bOtSeRvIcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue/channels",
			Error: true,
		},
		{
			// Incomplete URI (Insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-ReSoUrCe-GrOuP/pRoViDeRs/MiCrOsOfT.BoTsErViCe/BoTsErViCeS/bOtSeRvIcEvAlUe/ChAnNeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue/channels/EmailChannel",
			Expected: &BotServiceChannelId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				BotServiceName:    "botServiceValue",
				ChannelType:       EmailBotServiceChannelType,
			},
		},
		{
			// Valid URI Insensitively
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-ReSoUrCe-GrOuP/pRoViDeRs/MiCrOsOfT.BoTsErViCe/BoTsErViCeS/botServiceValue/ChAnNeLs/EmAiLChAnNeL",
			Expected: &BotServiceChannelId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "eXaMpLe-ReSoUrCe-GrOuP",
				BotServiceName:    "botServiceValue",
				ChannelType:       EmailBotServiceChannelType,
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue/channels/EmailChannel/extra",
			Error: true,
		},
		{
			// Invalid (Valid Uri with Extra segment) insensitively
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-ReSoUrCe-GrOuP/pRoViDeRs/MiCrOsOfT.BoTsErViCe/BoTsErViCeS/botServiceValue/cHaNnElS/EmAiLChAnNeL/ExTra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBotServiceChannelIDInsensitively(v.Input)
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

		if actual.BotServiceName != v.Expected.BotServiceName {
			t.Fatalf("Expected %q but got %q for BotServiceName", v.Expected.BotServiceName, actual.BotServiceName)
		}

		if actual.ChannelType != v.Expected.ChannelType {
			t.Fatalf("Expected %q but got %q for ChannelType", v.Expected.ChannelType, actual.ChannelType)
		}
	}
}

func TestSegmentsForBotServiceChannelId(t *testing.T) {
	segments := BotServiceChannelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BotServiceChannelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
