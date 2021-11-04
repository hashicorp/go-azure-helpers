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
			input:    pointer.ToStringPointer("West US"),
			expected: "westus",
		},
		{
			input:    pointer.ToStringPointer("South East Asia"),
			expected: "southeastasia",
		},
		{
			input:    pointer.ToStringPointer("southeastasia"),
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
