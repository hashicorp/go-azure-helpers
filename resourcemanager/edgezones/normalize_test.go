// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgezones

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
)

func TestNormalizeLocation(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "MicrosoftLosAngeles1",
			expected: "microsoftlosangeles1",
		},
		{
			input:    "Microsoft Los Angeles 1",
			expected: "microsoftlosangeles1",
		},
	}

	for _, v := range cases {
		actual := Normalize(v.input)
		if v.expected != actual {
			t.Fatalf("Expected %q but got %q", v.expected, actual)
		}
	}
}

func TestNormalizeNilableLocation(t *testing.T) {
	cases := []struct {
		input    *string
		expected string
	}{
		{
			input:    pointer.FromString("MicrosoftLosAngeles1"),
			expected: "microsoftlosangeles1",
		},
		{
			input:    pointer.FromString("Microsoft Los Angeles 1"),
			expected: "microsoftlosangeles1",
		},
		{
			input:    nil,
			expected: "",
		},
	}

	for _, v := range cases {
		actual := NormalizeNilable(v.input)
		if v.expected != actual {
			t.Fatalf("Expected %q but got %q", v.expected, actual)
		}
	}
}
