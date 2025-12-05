// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package resourceids_test

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// NOTE: most of the below tests reference the Id `planetId`, this is intentionally private for testing
// purposes and isn't being TitleCased because this shouldn't happen outside of testing

func TestSegmentNotSpecifiedError_CommonId(t *testing.T) {
	subId := commonids.ResourceGroupId{}
	result := resourceids.ParseResult{
		RawInput: "/example",
	}
	actual := resourceids.NewSegmentNotSpecifiedError(&subId, "resourceGroupName", result)
	expected := `parsing the ResourceGroup ID: the segment at position 3 didn't match

Expected a ResourceGroup ID that matched:

> /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group

However this value was provided:

> /example

The parsed Resource ID was missing a value for the segment at position 3
(which should be the name of the Resource Group).
`
	assertTemplatedCodeMatches(t, expected, actual.Error())
}

func TestSegmentNotSpecifiedError_MissingSegment(t *testing.T) {
	subId := commonids.ResourceGroupId{}
	result := resourceids.ParseResult{
		RawInput: "/example",
	}
	actual := resourceids.NewSegmentNotSpecifiedError(&subId, "somethingRandom", result)
	expected := `internal-error: couldn't determine the position for segment "somethingRandom"`
	assertTemplatedCodeMatches(t, expected, actual.Error())
}

func TestSegmentNotSpecifiedError_Constant(t *testing.T) {
	id := planetId{
		segments: []resourceids.Segment{
			resourceids.StaticSegment("planets", "planets", "planets"),
			resourceids.ConstantSegment("planetName", []string{"earth", "mars"}, "earth"),
		},
	}
	result := resourceids.ParseResult{
		RawInput: "/planets",
	}
	actual := resourceids.NewSegmentNotSpecifiedError(&id, "planetName", result)
	expected := `parsing the planet ID: the segment at position 1 didn't match

Expected a planet ID that matched:

> /planets/earth

However this value was provided:

> /planets

The parsed Resource ID was missing a value for the segment at position 1
(which should be a Constant with one of the following values ["earth", "mars"]).
`
	assertTemplatedCodeMatches(t, expected, actual.Error())
}

func TestSegmentNotSpecifiedError_ResourceGroup(t *testing.T) {
	id := planetId{
		segments: []resourceids.Segment{
			resourceids.StaticSegment("planets", "planets", "planets"),
			resourceids.ResourceGroupSegment("resourceGroupName", "example-resources"),
		},
	}
	result := resourceids.ParseResult{
		RawInput: "/planets",
	}
	actual := resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", result)
	expected := `parsing the planet ID: the segment at position 1 didn't match

Expected a planet ID that matched:

> /planets/example-resources

However this value was provided:

> /planets

The parsed Resource ID was missing a value for the segment at position 1
(which should be the name of the Resource Group).
`
	assertTemplatedCodeMatches(t, expected, actual.Error())
}

func TestSegmentNotSpecifiedError_ResourceProvider(t *testing.T) {
	id := planetId{
		segments: []resourceids.Segment{
			resourceids.StaticSegment("planets", "planets", "planets"),
			resourceids.ResourceProviderSegment("resourceProvider", "Example.ResourceProvider", "Example.ResourceProvider"),
		},
	}
	result := resourceids.ParseResult{
		RawInput: "/planets",
	}
	actual := resourceids.NewSegmentNotSpecifiedError(id, "resourceProvider", result)
	expected := `parsing the planet ID: the segment at position 1 didn't match

Expected a planet ID that matched:

> /planets/Example.ResourceProvider

However this value was provided:

> /planets

The parsed Resource ID was missing a value for the segment at position 1
(which should be the name of the Resource Provider [for example 'Example.ResourceProvider']).
`
	assertTemplatedCodeMatches(t, expected, actual.Error())
}

func TestSegmentNotSpecifiedError_Scope(t *testing.T) {
	id := planetId{
		segments: []resourceids.Segment{
			resourceids.StaticSegment("planets", "planets", "planets"),
			resourceids.ScopeSegment("scopeSegment", "/some/scope/value"),
		},
	}
	result := resourceids.ParseResult{
		RawInput: "/planets",
	}
	actual := resourceids.NewSegmentNotSpecifiedError(id, "scopeSegment", result)
	expected := `parsing the planet ID: the segment at position 1 didn't match

Expected a planet ID that matched:

> /planets/some/scope/value

However this value was provided:

> /planets

The parsed Resource ID was missing a value for the segment at position 1
(which specifies the Resource ID that should be used as a Scope [for example '/some/scope/value']).
`
	assertTemplatedCodeMatches(t, expected, actual.Error())
}

func TestSegmentNotSpecifiedError_Static(t *testing.T) {
	id := planetId{
		segments: []resourceids.Segment{
			resourceids.StaticSegment("resourceGroups", "resourceGroups", "resourceGroups"),
			resourceids.ResourceGroupSegment("resourceGroupName", "example-resources"),
		},
	}
	result := resourceids.ParseResult{
		RawInput: "/resourcegroups/example",
	}
	actual := resourceids.NewSegmentNotSpecifiedError(id, "resourceGroups", result)
	expected := `parsing the planet ID: the segment at position 0 didn't match

Expected a planet ID that matched:

> /resourceGroups/example-resources

However this value was provided:

> /resourcegroups/example

The parsed Resource ID was missing a value for the segment at position 0
(which should be the literal value "resourceGroups").
`
	assertTemplatedCodeMatches(t, expected, actual.Error())
}

func TestSegmentNotSpecifiedError_SubscriptionId(t *testing.T) {
	id := planetId{
		segments: []resourceids.Segment{
			resourceids.StaticSegment("planets", "planets", "planets"),
			resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-1234-1234-1234567890123"),
		},
	}
	result := resourceids.ParseResult{
		RawInput: "/planets",
	}
	actual := resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", result)
	expected := `parsing the planet ID: the segment at position 1 didn't match

Expected a planet ID that matched:

> /planets/12345678-1234-1234-1234-1234567890123

However this value was provided:

> /planets

The parsed Resource ID was missing a value for the segment at position 1
(which should be the UUID of the Azure Subscription).
`
	assertTemplatedCodeMatches(t, expected, actual.Error())
}

func TestSegmentNotSpecifiedError_UserSpecified(t *testing.T) {
	id := planetId{
		segments: []resourceids.Segment{
			resourceids.StaticSegment("planets", "planets", "planets"),
			resourceids.UserSpecifiedSegment("planetName", "example-planet"),
		},
	}
	result := resourceids.ParseResult{
		RawInput: "/planets",
	}
	actual := resourceids.NewSegmentNotSpecifiedError(id, "planetName", result)
	expected := `parsing the planet ID: the segment at position 1 didn't match

Expected a planet ID that matched:

> /planets/example-planet

However this value was provided:

> /planets

The parsed Resource ID was missing a value for the segment at position 1
(which should be the user specified value for this planet [for example "example-planet"]).
`
	assertTemplatedCodeMatches(t, expected, actual.Error())
}

func TestSegmentNotSpecifiedError_UnknownType(t *testing.T) {
	id := planetId{
		segments: []resourceids.Segment{
			{
				Name: "nah",
				Type: "does-not-matter",
			},
		},
	}
	result := resourceids.ParseResult{
		RawInput: "/planets",
	}
	actual := resourceids.NewSegmentNotSpecifiedError(id, "nah", result)
	expected := `internal-error: building description for segment: internal-error: the Segment Type "does-not-matter" was not implemented for Segment "nah"`
	assertTemplatedCodeMatches(t, expected, actual.Error())
}

var _ resourceids.ResourceId = planetId{}

type planetId struct {
	segments []resourceids.Segment
}

func (p planetId) ID() string {
	panic("should not be called in test")
}

func (p planetId) String() string {
	panic("should not be called in test")
}

func (p planetId) Segments() []resourceids.Segment {
	return p.segments
}

func (p planetId) FromParseResult(resourceids.ParseResult) error {
	panic("should not be called in test")
}
