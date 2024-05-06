// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourceids

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/features"
)

func TestMatch(t *testing.T) {
	// force a case-sensitive comparison for this test
	features.TreatUserSpecifiedSegmentsAsCaseInsensitive = false

	testData := []struct {
		first    ResourceId
		second   ResourceId
		expected bool
	}{
		{
			// two instances of the same Resource ID with the same value should match
			first:    newPlanetID("mars"),
			second:   newPlanetID("mars"),
			expected: true,
		},
		{
			// two instances of the same Resource ID with the same value in differing casing
			// shouldn't match
			first:    newPlanetID("mars"),
			second:   newPlanetID("mArs"),
			expected: false,
		},
		{
			// two instances of the same Resource ID with differing values shouldn't match
			first:    newPlanetID("earth"),
			second:   newPlanetID("mars"),
			expected: false,
		},
		{
			// two different Resource ID types shouldn't match - same URI
			first:    newPlanetID("earth"),
			second:   newOtherPlanetID("earth"),
			expected: false,
		},
		{
			// two different Resource ID types shouldn't match - different URIs
			first:    newPlanetID("earth"),
			second:   newSolarSystemPlanetID("milkyWay", "earth"),
			expected: false,
		},
	}
	for i, data := range testData {
		t.Logf("Iteration %d", i)
		actual := Match(data.first, data.second)
		if actual != data.expected {
			t.Fatalf("expected Match to return %t but got %t", data.expected, actual)
		}
	}
}

func TestMatch_Insensitive(t *testing.T) {
	// force a case-sensitive comparison for this test
	features.TreatUserSpecifiedSegmentsAsCaseInsensitive = true

	testData := []struct {
		first    ResourceId
		second   ResourceId
		expected bool
	}{
		{
			// two instances of the same Resource ID with the same value should match
			first:    newPlanetID("mars"),
			second:   newPlanetID("mars"),
			expected: true,
		},
		{
			// two instances of the same Resource ID with the same value in differing casing
			// should match since the feature-flag is enabled
			first:    newPlanetID("mars"),
			second:   newPlanetID("mArs"),
			expected: true,
		},
		{
			// two instances of the same Resource ID with differing values shouldn't match
			first:    newPlanetID("earth"),
			second:   newPlanetID("mars"),
			expected: false,
		},
		{
			// two different Resource ID types shouldn't match - same URI
			first:    newPlanetID("earth"),
			second:   newOtherPlanetID("earth"),
			expected: false,
		},
		{
			// two different Resource ID types shouldn't match - different URIs
			first:    newPlanetID("earth"),
			second:   newSolarSystemPlanetID("milkyWay", "earth"),
			expected: false,
		},
	}
	for i, data := range testData {
		t.Logf("Iteration %d", i)
		actual := Match(data.first, data.second)
		if actual != data.expected {
			t.Fatalf("expected Match to return %t but got %t", data.expected, actual)
		}
	}
}

// NOTE: the ResourceId implementations below are purely for test purposes and so intentionally
// incomplete implementations/panicking for unexpected/unused methods

var _ ResourceId = &planetResourceId{}

type planetResourceId struct {
	planetName string
}

func newPlanetID(planetName string) *planetResourceId {
	return &planetResourceId{
		planetName: planetName,
	}
}

func (f *planetResourceId) FromParseResult(_ ParseResult) error {
	panic("not implemented since this codepath should not be used")
}

func (f *planetResourceId) ID() string {
	return fmt.Sprintf("/planets/%s", f.planetName)
}

func (f *planetResourceId) String() string {
	panic("not implemented since this codepath should not be used")
}

func (f *planetResourceId) Segments() []Segment {
	return []Segment{
		StaticSegment("planets", "planets", "planets"),
		UserSpecifiedSegment("planetName", "earth"),
	}
}

var _ ResourceId = &otherPlanetResourceId{}

type otherPlanetResourceId struct {
	planetName string
}

func newOtherPlanetID(planetName string) *otherPlanetResourceId {
	return &otherPlanetResourceId{
		planetName: planetName,
	}
}

func (f *otherPlanetResourceId) FromParseResult(_ ParseResult) error {
	panic("not implemented since this codepath should not be used")
}

func (f *otherPlanetResourceId) ID() string {
	return fmt.Sprintf("/planets/%s", f.planetName)
}

func (f *otherPlanetResourceId) String() string {
	panic("not implemented since this codepath should not be used")
}

func (f *otherPlanetResourceId) Segments() []Segment {
	return []Segment{
		StaticSegment("planets", "planets", "planets"),
		UserSpecifiedSegment("planetName", "earth"),
	}
}

var _ ResourceId = &solarSystemPlanetResourceId{}

type solarSystemPlanetResourceId struct {
	planetName      string
	solarSystemName string
}

func newSolarSystemPlanetID(solarSystemName, planetName string) *solarSystemPlanetResourceId {
	return &solarSystemPlanetResourceId{
		planetName:      planetName,
		solarSystemName: solarSystemName,
	}
}

func (f *solarSystemPlanetResourceId) FromParseResult(_ ParseResult) error {
	panic("not implemented since this codepath should not be used")
}

func (f *solarSystemPlanetResourceId) ID() string {
	return fmt.Sprintf("/solarSystems/%s/planets/%s", f.solarSystemName, f.planetName)
}

func (f *solarSystemPlanetResourceId) String() string {
	panic("not implemented since this codepath should not be used")
}

func (f *solarSystemPlanetResourceId) Segments() []Segment {
	return []Segment{
		StaticSegment("solarSystems", "solarSystems", "solarSystems"),
		UserSpecifiedSegment("solarSystemName", "milkyWay"),
		StaticSegment("planets", "planets", "planets"),
		UserSpecifiedSegment("planetName", "earth"),
	}
}
