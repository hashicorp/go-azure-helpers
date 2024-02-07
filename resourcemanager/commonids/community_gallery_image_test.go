// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import "testing"

func TestNewCommunityGalleryImageID(t *testing.T) {
	id := NewCommunityGalleryImageID("communityGalleryValue", "imageValue")

	if id.CommunityGalleryName != "communityGalleryValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CommunityGalleryName'", "communityGalleryValue", id.CommunityGalleryName)
	}

	if id.ImageName != "imageValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ImageName'", "imageValue", id.ImageName)
	}
}

func TestFormatCommunityGalleryImageID(t *testing.T) {
	actual := NewCommunityGalleryImageID("communityGalleryValue", "imageValue").ID()
	expected := "/communityGalleries/communityGalleryValue/images/imageValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCommunityGalleryImageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CommunityGalleryImageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/communityGalleries",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/communityGalleries/communityGalleryValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/communityGalleries/communityGalleryValue/images/",
			Error: true,
		},
		{
			// Valid URI
			Input: "/communityGalleries/communityGalleryValue/images/imageValue",
			Expected: &CommunityGalleryImageId{
				CommunityGalleryName: "communityGalleryValue",
				ImageName:            "imageValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/communityGalleries/communityGalleryValue/images/imageValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCommunityGalleryImageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CommunityGalleryName != v.Expected.CommunityGalleryName {
			t.Fatalf("Expected %q but got %q for CommunityGalleryName", v.Expected.CommunityGalleryName, actual.CommunityGalleryName)
		}

		if actual.ImageName != v.Expected.ImageName {
			t.Fatalf("Expected %q but got %q for CommunityGalleryImageName", v.Expected.ImageName, actual.ImageName)
		}
	}
}

func TestParseCommunityGalleryImageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CommunityGalleryImageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/communityGalleries",
			Error: true,
		},
		{
			// Incomplete URI (Insensitively)
			Input: "/CoMmunItYGaLleRiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/communityGalleries/communityGalleryValue",
			Error: true,
		},
		{
			// Incomplete URI (Insensitively)
			Input: "/CoMmunItYGaLleRiEs/CoMmunItYGalLERYVaLue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/communityGalleries/communityGalleryValue/images/",
			Error: true,
		},
		{
			// Incomplete URI (Insensitively)
			Input: "/CoMmunItYGaLleRiEs/CoMmunItYGalLERYVaLue/imAGes/",
			Error: true,
		},
		{
			// Valid URI
			Input: "/communityGalleries/communityGalleryValue/images/imageValue",
			Expected: &CommunityGalleryImageId{
				CommunityGalleryName: "communityGalleryValue",
				ImageName:            "imageValue",
			},
		},
		{
			// Valid URI (Insensitively)
			Input: "/CoMmunItYGaLleRiEs/CoMmunItYGalLERYVaLue/imAGes/iMaGeVaLue",
			Expected: &CommunityGalleryImageId{
				CommunityGalleryName: "CoMmunItYGalLERYVaLue",
				ImageName:            "iMaGeVaLue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/communityGalleries/communityGalleryValue/images/imageValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCommunityGalleryImageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CommunityGalleryName != v.Expected.CommunityGalleryName {
			t.Fatalf("Expected %q but got %q for CommunityGalleryName", v.Expected.CommunityGalleryName, actual.CommunityGalleryName)
		}

		if actual.ImageName != v.Expected.ImageName {
			t.Fatalf("Expected %q but got %q for CommunityGalleryImageName", v.Expected.ImageName, actual.ImageName)
		}
	}
}

func TestSegmentsForCommunityGalleryImageId(t *testing.T) {
	segments := CommunityGalleryImageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CommunityGalleryImageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
