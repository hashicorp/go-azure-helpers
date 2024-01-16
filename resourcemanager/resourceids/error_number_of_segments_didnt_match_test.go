// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourceids_test

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

func TestNumberOfSegmentsDidntMatchError_CommonIdAllMissing(t *testing.T) {
	id := commonids.ResourceGroupId{}
	parseResult := resourceids.ParseResult{
		Parsed:   map[string]string{},
		RawInput: "/some-value",
	}
	expected := `
parsing the ResourceGroup ID: the number of segments didn't match

Expected a ResourceGroup ID that matched (containing 4 segments):

> /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group

However this value was provided (which was parsed into 0 segments):

> /some-value

The following Segments are expected:

* Segment 0 - this should be the literal value "subscriptions"
* Segment 1 - this should be the UUID of the Azure Subscription
* Segment 2 - this should be the literal value "resourceGroups"
* Segment 3 - this should be the name of the Resource Group

The following Segments were parsed:

* Segment 0 - not found
* Segment 1 - not found
* Segment 2 - not found
* Segment 3 - not found
`
	actual := resourceids.NewNumberOfSegmentsDidntMatchError(&id, parseResult).Error()
	assertTemplatedCodeMatches(t, expected, actual)
}

func TestNumberOfSegmentsDidntMatchError_CommonIdMissingSegment(t *testing.T) {
	id := commonids.ResourceGroupId{}
	parseResult := resourceids.ParseResult{
		Parsed: map[string]string{
			"subscriptions":  "subscriptions",
			"subscriptionId": "1234",
		},
		RawInput: "/subscriptions/1234/resourcegroups",
	}
	expected := `
parsing the ResourceGroup ID: the number of segments didn't match

Expected a ResourceGroup ID that matched (containing 4 segments):

> /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group

However this value was provided (which was parsed into 2 segments):

> /subscriptions/1234/resourcegroups

The following Segments are expected:

* Segment 0 - this should be the literal value "subscriptions"
* Segment 1 - this should be the UUID of the Azure Subscription
* Segment 2 - this should be the literal value "resourceGroups"
* Segment 3 - this should be the name of the Resource Group

The following Segments were parsed:

* Segment 0 - parsed as "subscriptions"
* Segment 1 - parsed as "1234"
* Segment 2 - not found
* Segment 3 - not found
`
	actual := resourceids.NewNumberOfSegmentsDidntMatchError(&id, parseResult).Error()
	assertTemplatedCodeMatches(t, expected, actual)
}
