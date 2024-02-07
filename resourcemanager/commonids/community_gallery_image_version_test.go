// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import "testing"

func TestNewCommunityGalleryImageVersionID(t *testing.T) {
	id := NewCommunityGalleryImageVersionID("communityGalleryValue", "imageValue", "versionValue")

	if id.CommunityGalleryName != "communityGalleryValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CommunityGalleryName'", "communityGalleryValue", id.CommunityGalleryName)
	}

	if id.ImageName != "imageValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ImageName'", "imageValue", id.ImageName)
	}

	if id.VersionName != "versionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", "versionValue", id.VersionName)
	}
}

func TestFormatCommunityGalleryImageVersionID(t *testing.T) {
	actual := NewCommunityGalleryImageVersionID("communityGalleryValue", "imageValue", "versionValue").ID()
	expected := "/communityGalleries/communityGalleryValue/images/imageValue/versions/versionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCommunityGalleryImageVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CommunityGalleryImageVersionId
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
			// Incomplete URI
			Input: "/communityGalleries/communityGalleryValue/images/imageValue/",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/communityGalleries/communityGalleryValue/images/imageValue/versions/",
			Error: true,
		},
		{
			// Valid URI
			Input: "/communityGalleries/communityGalleryValue/images/imageValue/versions/versionValue",
			Expected: &CommunityGalleryImageVersionId{
				CommunityGalleryName: "communityGalleryValue",
				ImageName:            "imageValue",
				VersionName:          "versionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/communityGalleries/communityGalleryValue/images/imageValue/versions/versionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCommunityGalleryImageVersionID(v.Input)
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
			t.Fatalf("Expected %q but got %q for ImageName", v.Expected.ImageName, actual.ImageName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}
	}
}

func TestParseCommunityGalleryImageVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CommunityGalleryImageVersionId
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
			// Incomplete URI
			Input: "/CoMmunItYGaLleRiEs/CoMmunItYGalLERYVaLue/imAGes/imageValue/",
			Error: true,
		},
		{
			// Incomplete URI (Insensitively)
			Input: "/CoMmunItYGaLleRiEs/CoMmunItYGalLERYVaLue/imAGes/iMAgeVaLue/",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/communityGalleries/communityGalleryValue/images/imageValue/versions/",
			Error: true,
		},
		{
			// Incomplete URI (Insensitively)
			Input: "/CoMmunItYGaLleRiEs/CoMmunItYGalLERYVaLue/imAGes/iMAgeVaLue/veRSiOns/",
			Error: true,
		},
		{
			// Valid URI
			Input: "/communityGalleries/communityGalleryValue/images/imageValue/versions/versionValue",
			Expected: &CommunityGalleryImageVersionId{
				CommunityGalleryName: "communityGalleryValue",
				ImageName:            "imageValue",
				VersionName:          "versionValue",
			},
		},
		{
			// Valid URI (Insensitively)
			Input: "/CoMmunItYGaLleRiEs/CoMmunItYGalLERYVaLue/imAGes/iMaGeVaLue/vErSiOnS/VeRsIoNVaLuE",
			Expected: &CommunityGalleryImageVersionId{
				CommunityGalleryName: "CoMmunItYGalLERYVaLue",
				ImageName:            "iMaGeVaLue",
				VersionName:          "VeRsIoNVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/communityGalleries/communityGalleryValue/images/imageValue/versions/versionValue/extra",
			Error: true,
		},
		{
			// Invalid (Valid Uri with Extra segment) insensitively
			Input: "/CoMmunItYGaLleRiEs/CoMmunItYGalLERYVaLue/imAGes/iMaGeVaLue/vErSiOnS/VeRsIoNVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCommunityGalleryImageVersionIDInsensitively(v.Input)
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
			t.Fatalf("Expected %q but got %q for ImageName", v.Expected.ImageName, actual.ImageName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}
	}
}

func TestSegmentsForCommunityGalleryImageVersionId(t *testing.T) {
	segments := CommunityGalleryImageVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CommunityGalleryImageVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
