// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import (
	"fmt"
	"testing"
)

func TestNewCompositeResourceID(t *testing.T) {
	botIdString := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue"
	appIdString := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue"

	botId, err := ParseBotServiceID(botIdString)
	if err != nil {
		t.Fatalf("parsing resource ID for composite resource ID test: %+v", err)
	}

	appId, err := ParseAppServiceID(appIdString)
	if err != nil {
		t.Fatalf("parsing resource ID for composite resource ID test: %+v", err)
	}

	id := NewCompositeResourceID(botId, appId)

	if id.First.ID() != botIdString {
		t.Fatalf("expected First ID string to be %q but got %q", botIdString, id.First.ID())
	}

	if id.Second.ID() != appIdString {
		t.Fatalf("expected Second ID string to be %q but got %q", appIdString, id.Second.ID())
	}

	expectedIdString := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue|/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue"

	if id.ID() != expectedIdString {
		t.Fatalf("Expected ID string to be %q but got %q", expectedIdString, id.ID())
	}
}

func TestCompositeResourceID(t *testing.T) {
	idString := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue|/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue"

	id1 := BotServiceId{}
	id2 := AppServiceId{}
	id, err := ParseCompositeResourceID(idString, &id1, &id2)
	if err != nil {
		t.Fatalf("Expected CompositeResourceID to parse successfully but got Error: %q", err)
	}

	expectedId1 := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue"
	actualId1 := id.First.ID()
	if expectedId1 != actualId1 {
		t.Fatalf("Expected the First ID to be %q but got %q", expectedId1, actualId1)
	}

	expectedId2 := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue"
	actualId2 := id.Second.ID()
	if expectedId2 != actualId2 {
		t.Fatalf("Expected the Second ID to be %q but got %q", expectedId2, actualId2)
	}

	if idString != id.ID() {
		t.Fatalf("Expected the Composite ID to be %q but got %q", idString, id.ID())
	}

	expectedIdString1 := `Bot Service (Subscription: "12345678-1234-9876-4563-123456789012"
Resource Group Name: "example-resource-group"
Bot Service Name: "botServiceValue")`
	actualIdString1 := id.First.String()
	if expectedIdString1 != actualIdString1 {
		t.Fatalf("Expected the First ID String to be %q but got %q", expectedIdString1, actualIdString1)
	}

	expectedIdString2 := `App Service (Subscription: "12345678-1234-9876-4563-123456789012"
Resource Group Name: "example-resource-group"
Site Name: "siteValue")`
	actualIdString2 := id.Second.String()
	if expectedIdString2 != actualIdString2 {
		t.Fatalf("Expected the Second ID String to be %q but got %q", expectedIdString2, actualIdString2)
	}

	expectedCompositeString := fmt.Sprintf("Composite Resource ID (%s | %s)", expectedIdString1, expectedIdString2)
	actualCompositeString := id.String()
	if expectedCompositeString != actualCompositeString {
		t.Fatalf("Expected the Composite ID String to be %q but got %q", expectedCompositeString, actualCompositeString)
	}
}

func TestParseCompositeResourceIDInsensitively(t *testing.T) {
	idIncorrectCasing := "/subscriptions/12345678-1234-9876-4563-123456789012/ReSoUrCeGrOuPs/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue|/subscriptions/12345678-1234-9876-4563-123456789012/ReSoUrCeGrOuPs/example-resource-group/providers/Microsoft.Web/sites/siteValue"
	idCorrectCasing := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue|/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue"

	id1 := BotServiceId{}
	id2 := AppServiceId{}
	id, err := ParseCompositeResourceIDInsensitively(idIncorrectCasing, &id1, &id2)
	if err != nil {
		t.Fatalf("Expected CompositeResourceID to parse successfully but got Error: %q", err)
	}

	expectedId1 := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue"
	actualId1 := id.First.ID()
	if expectedId1 != actualId1 {
		t.Fatalf("Expected the First ID to be %q but got %q", expectedId1, actualId1)
	}

	expectedId2 := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue"
	actualId2 := id.Second.ID()
	if expectedId2 != actualId2 {
		t.Fatalf("Expected the Second ID to be %q but got %q", expectedId2, actualId2)
	}

	if idCorrectCasing != id.ID() {
		t.Fatalf("Expected the Composite ID to be %q but got %q", idIncorrectCasing, id.ID())
	}

	expectedIdString1 := `Bot Service (Subscription: "12345678-1234-9876-4563-123456789012"
Resource Group Name: "example-resource-group"
Bot Service Name: "botServiceValue")`
	actualIdString1 := id.First.String()
	if expectedIdString1 != actualIdString1 {
		t.Fatalf("Expected the First ID String to be %q but got %q", expectedIdString1, actualIdString1)
	}

	expectedIdString2 := `App Service (Subscription: "12345678-1234-9876-4563-123456789012"
Resource Group Name: "example-resource-group"
Site Name: "siteValue")`
	actualIdString2 := id.Second.String()
	if expectedIdString2 != actualIdString2 {
		t.Fatalf("Expected the Second ID String to be %q but got %q", expectedIdString2, actualIdString2)
	}

	expectedCompositeString := fmt.Sprintf("Composite Resource ID (%s | %s)", expectedIdString1, expectedIdString2)
	actualCompositeString := id.String()
	if expectedCompositeString != actualCompositeString {
		t.Fatalf("Expected the Composite ID String to be %q but got %q", expectedCompositeString, actualCompositeString)
	}
}

func TestCompositeResourceIDNumberOfIDsErrors(t *testing.T) {
	id1 := BotServiceId{}
	id2 := AppServiceId{}
	testData := []struct {
		Input         string
		ExpectedError string
	}{
		{
			// 1 ID
			Input:         "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue",
			ExpectedError: "expected 2 resourceids but got 1",
		},
		{
			// 3 IDs
			Input:         "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue|/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue|/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue",
			ExpectedError: "expected 2 resourceids but got 3",
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		_, err := ParseCompositeResourceID(v.Input, &id1, &id2)
		if err == nil {
			t.Fatalf("Expected an error but didn't get one")
		}
		if err.Error() != v.ExpectedError {
			t.Fatalf("Expected error %q but got %q", v.ExpectedError, err.Error())
		}
	}
}

func TestCompositeResourceIDInvalidIDs(t *testing.T) {

	id1 := BotServiceId{}
	id2 := AppServiceId{}

	idStringFirstInvalid := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService|/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/sites/siteValue"
	_, err := ParseCompositeResourceID(idStringFirstInvalid, &id1, &id2)
	if err == nil {
		t.Fatalf("Expected error but didn't get one")
	}

	idStringSecondInvalid := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.BotService/botServices/botServiceValue|/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web"
	_, err = ParseCompositeResourceID(idStringSecondInvalid, &id1, &id2)
	if err == nil {
		t.Fatalf("Expected error but didn't get one")
	}
}
