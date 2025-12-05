// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package location

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
			input:    "West US",
			expected: "westus",
		},
		{
			input:    "South East Asia",
			expected: "southeastasia",
		},
		{
			input:    "southeastasia",
			expected: "southeastasia",
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
			input:    pointer.To("West US"),
			expected: "westus",
		},
		{
			input:    pointer.To("South East Asia"),
			expected: "southeastasia",
		},
		{
			input:    pointer.To("southeastasia"),
			expected: "southeastasia",
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
