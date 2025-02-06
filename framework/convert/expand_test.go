// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/go-azure-helpers/framework/convert"
	"github.com/hashicorp/go-azure-helpers/framework/typehelpers"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpand(t *testing.T) {
	input := &TestFrameworkModelComplex{
		BoolProperty:   types.BoolValue(true),
		StringProperty: types.StringValue("foo"),
	}

	expectedOutput := &TestAPIModelComplex{
		BoolProperty:   true,
		StringProperty: "foo",
	}

	result := &TestAPIModelComplex{}

	diags := &diag.Diagnostics{}

	convert.Expand(context.Background(), input, result, diags)
	if diags.HasError() {
		t.Fatalf("Expand failed: %+v", diags.Errors())
	}
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Fatalf("Expand failed, expected: %+v \n Got: %+v", expectedOutput, result)
	}
}

func TestExpands(t *testing.T) {
	ctx := context.Background()
	testCases := []struct {
		name        string
		input       *TestFrameworkModelComplex
		expected    *TestAPIModelComplex
		expectError bool
	}{
		{
			name:        "empty",
			input:       &TestFrameworkModelComplex{},
			expected:    &TestAPIModelComplex{},
			expectError: false,
		},
		{
			name: "single bool",
			input: &TestFrameworkModelComplex{
				BoolProperty: types.BoolValue(true),
			},
			expected: &TestAPIModelComplex{
				BoolProperty: true,
			},
			expectError: false,
		},
		{
			name: "single string",
			input: &TestFrameworkModelComplex{
				StringProperty: types.StringValue("foo"),
			},
			expected: &TestAPIModelComplex{
				StringProperty: "foo",
			},
			expectError: false,
		},
		{
			name: "multiple list entries",
			input: &TestFrameworkModelComplex{
				ListProperty: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []TestFrameworkNestedModel{
					{
						SubPropertyBool:   types.BoolValue(true),
						SubPropertyFloat:  types.Float64Value(3.142),
						SubPropertyInt:    types.Int64Value(1984),
						SubPropertyString: types.StringValue("foo"),
					},
					{
						SubPropertyBool:   types.BoolValue(false),
						SubPropertyFloat:  types.Float64Value(2.54),
						SubPropertyInt:    types.Int64Value(9000),
						SubPropertyString: types.StringValue("bar"),
					},
				}),
			},
			expected: &TestAPIModelComplex{
				ListProperty: []TestAPINestedModel{
					{
						SubPropertyBool:   true,
						SubPropertyFloat:  3.142,
						SubPropertyInt:    1984,
						SubPropertyString: "foo",
					},
					{
						SubPropertyBool:   false,
						SubPropertyFloat:  2.54,
						SubPropertyInt:    9000,
						SubPropertyString: "bar",
					},
				},
			},
		},
		{
			name: "complete",
			input: &TestFrameworkModelComplex{
				BoolProperty:   types.BoolValue(true),
				StringProperty: types.StringValue("foo"),
				FloatProperty:  types.Float64Value(3.142),
				IntProperty:    types.Int64Value(1984),
				ListProperty: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []TestFrameworkNestedModel{
					{
						SubPropertyBool:   types.BoolValue(true),
						SubPropertyFloat:  types.Float64Value(3.142),
						SubPropertyInt:    types.Int64Value(1984),
						SubPropertyString: types.StringValue("foo"),
					},
				}),
				SetProperty: typehelpers.NewSetNestedObjectValueOfValueSliceMust(ctx, []TestFrameworkNestedModel{
					{
						SubPropertyBool:   types.BoolValue(true),
						SubPropertyFloat:  types.Float64Value(3.1425),
						SubPropertyInt:    types.Int64Value(1985),
						SubPropertyString: types.StringValue("bar"),
					},
				}),
				MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{
					"foo": types.StringValue("bar"),
				}),
			},
			expected: &TestAPIModelComplex{
				BoolProperty:   true,
				StringProperty: "foo",
				FloatProperty:  3.142,
				IntProperty:    1984,
				ListProperty: []TestAPINestedModel{
					{
						SubPropertyBool:   true,
						SubPropertyFloat:  3.142,
						SubPropertyInt:    1984,
						SubPropertyString: "foo",
					},
				},
				SetProperty: []TestAPINestedModel{
					{
						SubPropertyBool:   true,
						SubPropertyFloat:  3.1425,
						SubPropertyInt:    1985,
						SubPropertyString: "bar",
					},
				},
				MapStringProperty: map[string]string{
					"foo": "bar",
				},
			},
			expectError: false,
		},
	}

	for _, testCase := range testCases {

		result := &TestAPIModelComplex{}

		diags := &diag.Diagnostics{}

		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() != testCase.expectError {
			t.Fatalf("Expand failed for %s: %+v", testCase.name, diags.Errors())
		}

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Fatalf("Expand failed for case %s, \nExpected: %+v \nGot: %+v", testCase.name, testCase.expected, result)
		}
	}
}

func TestExpand_boolRequired(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestBoolOnlyFWModel
		expect      *TestBoolRequiredModel
		expectError bool
	}{
		{
			name:  "empty",
			input: &TestBoolOnlyFWModel{},
			expect: &TestBoolRequiredModel{
				BoolProperty: false,
			},
			expectError: false,
		},
		{
			name: "true",
			input: &TestBoolOnlyFWModel{
				BoolProperty: types.BoolValue(true),
			},
			expect: &TestBoolRequiredModel{
				BoolProperty: true,
			},
		},
		{
			name: "false",
			input: &TestBoolOnlyFWModel{
				BoolProperty: types.BoolValue(false),
			},
			expect: &TestBoolRequiredModel{
				BoolProperty: false,
			},
		},
	}

	for _, testCase := range cases {
		result := &TestBoolRequiredModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand required bool: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_boolOptional(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestBoolOnlyFWModel
		expect      *TestBoolOptionalModel
		expectError bool
	}{
		{
			name:   "empty",
			input:  &TestBoolOnlyFWModel{},
			expect: &TestBoolOptionalModel{},
		},
		{
			name: "true",
			input: &TestBoolOnlyFWModel{
				BoolProperty: types.BoolValue(true),
			},
			expect: &TestBoolOptionalModel{
				BoolProperty: pointer.To(true),
			},
		},
		{
			name: "false",
			input: &TestBoolOnlyFWModel{
				BoolProperty: types.BoolValue(false),
			},
			expect: &TestBoolOptionalModel{
				BoolProperty: pointer.To(false),
			},
		},
	}

	for _, testCase := range cases {
		result := &TestBoolOptionalModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand optional bool: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_stringRequired(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestStringOnlyFWModel
		expect      *TestStringRequiredModel
		expectError bool
	}{
		{
			name:  "null",
			input: &TestStringOnlyFWModel{},
			expect: &TestStringRequiredModel{
				StringProperty: "",
			},
			expectError: false,
		},
		{
			name: "zero",
			input: &TestStringOnlyFWModel{
				StringProperty: types.StringValue(""),
			},
			expect: &TestStringRequiredModel{
				StringProperty: "",
			},
		},
		{
			name: "value",
			input: &TestStringOnlyFWModel{
				StringProperty: types.StringValue("foo"),
			},
			expect: &TestStringRequiredModel{
				StringProperty: "foo",
			},
		},
	}

	for _, testCase := range cases {
		result := &TestStringRequiredModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand required string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_stringOptional(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestStringOnlyFWModel
		expect      *TestStringOptionalModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestStringOnlyFWModel{},
			expect: &TestStringOptionalModel{},
		},
		{
			name: "zero",
			input: &TestStringOnlyFWModel{
				StringProperty: types.StringValue(""),
			},
			expect: &TestStringOptionalModel{
				StringProperty: pointer.To(""),
			},
		},
		{
			name: "value",
			input: &TestStringOnlyFWModel{
				StringProperty: types.StringValue("foo"),
			},
			expect: &TestStringOptionalModel{
				StringProperty: pointer.To("foo"),
			},
		},
	}

	for _, testCase := range cases {
		result := &TestStringOptionalModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand optional string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_int64Required(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestInt64OnlyFWModel
		expect      *TestInt64RequiredModel
		expectError bool
	}{
		{
			name:  "null",
			input: &TestInt64OnlyFWModel{},
			expect: &TestInt64RequiredModel{
				Int64Property: 0,
			},
			expectError: false,
		},
		{
			name: "zero",
			input: &TestInt64OnlyFWModel{
				Int64Property: types.Int64Value(0),
			},
			expect: &TestInt64RequiredModel{
				Int64Property: 0,
			},
		},
		{
			name: "value",
			input: &TestInt64OnlyFWModel{
				Int64Property: types.Int64Value(365),
			},
			expect: &TestInt64RequiredModel{
				Int64Property: 365,
			},
		},
	}

	for _, testCase := range cases {
		result := &TestInt64RequiredModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand required int64: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_int64Optional(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestInt64OnlyFWModel
		expect      *TestInt64OptionalModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestInt64OnlyFWModel{},
			expect: &TestInt64OptionalModel{},
		},
		{
			name: "zero",
			input: &TestInt64OnlyFWModel{
				Int64Property: types.Int64Value(0),
			},
			expect: &TestInt64OptionalModel{
				Int64Property: pointer.To(int64(0)),
			},
		},
		{
			name: "value",
			input: &TestInt64OnlyFWModel{
				Int64Property: types.Int64Value(365),
			},
			expect: &TestInt64OptionalModel{
				Int64Property: pointer.To(int64(365)),
			},
		},
	}

	for _, testCase := range cases {
		result := &TestInt64OptionalModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand optional int64: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_float64Required(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestFloat64OnlyFWModel
		expect      *TestFloat64RequiredModel
		expectError bool
	}{
		{
			name:  "null",
			input: &TestFloat64OnlyFWModel{},
			expect: &TestFloat64RequiredModel{
				Float64Property: 0,
			},
			expectError: false,
		},
		{
			name: "zero",
			input: &TestFloat64OnlyFWModel{
				Float64Property: types.Float64Value(0),
			},
			expect: &TestFloat64RequiredModel{
				Float64Property: 0,
			},
		},
		{
			name: "value",
			input: &TestFloat64OnlyFWModel{
				Float64Property: types.Float64Value(3.142),
			},
			expect: &TestFloat64RequiredModel{
				Float64Property: 3.142,
			},
		},
	}

	for _, testCase := range cases {
		result := &TestFloat64RequiredModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand required float64: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_float64Optional(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestFloat64OnlyFWModel
		expect      *TestFloat64OptionalModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestFloat64OnlyFWModel{},
			expect: &TestFloat64OptionalModel{},
		},
		{
			name: "zero",
			input: &TestFloat64OnlyFWModel{
				Float64Property: types.Float64Value(0),
			},
			expect: &TestFloat64OptionalModel{
				Float64Property: pointer.To(float64(0)),
			},
		},
		{
			name: "value",
			input: &TestFloat64OnlyFWModel{
				Float64Property: types.Float64Value(3.142),
			},
			expect: &TestFloat64OptionalModel{
				Float64Property: pointer.To(3.142),
			},
		},
	}

	for _, testCase := range cases {
		result := &TestFloat64OptionalModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand optional float64: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_mapOfBoolRequired(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfBoolFWModel
		expect      *TestMapOfBoolModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestMapOfBoolFWModel{},
			expect: &TestMapOfBoolModel{},
		},
		{
			name: "zero",
			input: &TestMapOfBoolFWModel{
				MapBoolProperty: typehelpers.NewMapValueOfMust[types.Bool](ctx, map[string]attr.Value{}),
			},
			expect: &TestMapOfBoolModel{
				MapBoolProperty: map[string]bool{},
			},
		},
		{
			name: "value",
			input: &TestMapOfBoolFWModel{
				MapBoolProperty: typehelpers.NewMapValueOfMust[types.Bool](ctx, map[string]attr.Value{
					"true": types.BoolValue(true),
				}),
			},
			expect: &TestMapOfBoolModel{
				MapBoolProperty: map[string]bool{
					"true": true,
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestMapOfBoolModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand map of bool: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_mapOfBoolOptional(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfBoolFWModel
		expect      *TestMapOfBoolPtrModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestMapOfBoolFWModel{},
			expect: &TestMapOfBoolPtrModel{},
		},
		{
			name: "zero",
			input: &TestMapOfBoolFWModel{
				MapBoolProperty: typehelpers.NewMapValueOfMust[types.Bool](ctx, map[string]attr.Value{}),
				// MapBoolProperty: typehelpers.NewMapValueOfNull[types.Bool](ctx),
			},
			expect: &TestMapOfBoolPtrModel{
				MapBoolProperty: map[string]*bool{},
			},
		},
		{
			name: "value",
			input: &TestMapOfBoolFWModel{
				MapBoolProperty: typehelpers.NewMapValueOfMust[types.Bool](ctx, map[string]attr.Value{
					"true": types.BoolValue(true),
				}),
			},
			expect: &TestMapOfBoolPtrModel{
				MapBoolProperty: map[string]*bool{
					"true": pointer.To(true),
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestMapOfBoolPtrModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand map of bool: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_mapOfStringRequired(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfStringFWModel
		expect      *TestMapOfStringModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestMapOfStringFWModel{},
			expect: &TestMapOfStringModel{},
		},
		{
			name: "zero",
			input: &TestMapOfStringFWModel{
				MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{}),
			},
			expect: &TestMapOfStringModel{
				MapStringProperty: map[string]string{},
			},
		},
		{
			name: "value",
			input: &TestMapOfStringFWModel{
				MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{
					"foo": types.StringValue("bar"),
				}),
			},
			expect: &TestMapOfStringModel{
				MapStringProperty: map[string]string{
					"foo": "bar",
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestMapOfStringModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand map of string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_mapOfStringOptional(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfStringFWModel
		expect      *TestMapOfStringPtrModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestMapOfStringFWModel{},
			expect: &TestMapOfStringPtrModel{},
		},
		{
			name: "zero",
			input: &TestMapOfStringFWModel{
				MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{}),
			},
			expect: &TestMapOfStringPtrModel{
				MapStringProperty: map[string]*string{},
			},
		},
		{
			name: "value",
			input: &TestMapOfStringFWModel{
				MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{
					"foo": types.StringValue("bar"),
				}),
			},
			expect: &TestMapOfStringPtrModel{
				MapStringProperty: map[string]*string{
					"foo": pointer.To("bar"),
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestMapOfStringPtrModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand map of string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_mapOfInt64Required(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfInt64FWModel
		expect      *TestMapOfInt64Model
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestMapOfInt64FWModel{},
			expect: &TestMapOfInt64Model{},
		},
		{
			name: "zero",
			input: &TestMapOfInt64FWModel{
				MapInt64Property: typehelpers.NewMapValueOfMust[types.Int64](ctx, map[string]attr.Value{}),
			},
			expect: &TestMapOfInt64Model{
				MapInt64Property: map[string]int64{},
			},
		},
		{
			name: "value",
			input: &TestMapOfInt64FWModel{
				MapInt64Property: typehelpers.NewMapValueOfMust[types.Int64](ctx, map[string]attr.Value{
					"days_later": types.Int64Value(28),
				}),
			},
			expect: &TestMapOfInt64Model{
				MapInt64Property: map[string]int64{
					"days_later": 28,
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestMapOfInt64Model{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand map of Int64: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_mapOfInt64Optional(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfInt64FWModel
		expect      *TestMapOfInt64PtrModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestMapOfInt64FWModel{},
			expect: &TestMapOfInt64PtrModel{},
		},
		{
			name: "zero",
			input: &TestMapOfInt64FWModel{
				MapInt64Property: typehelpers.NewMapValueOfMust[types.Int64](ctx, map[string]attr.Value{}),
			},
			expect: &TestMapOfInt64PtrModel{
				MapInt64Property: map[string]*int64{},
			},
		},
		{
			name: "value",
			input: &TestMapOfInt64FWModel{
				MapInt64Property: typehelpers.NewMapValueOfMust[types.Int64](ctx, map[string]attr.Value{
					"days_later": types.Int64Value(28),
				}),
			},
			expect: &TestMapOfInt64PtrModel{
				MapInt64Property: map[string]*int64{
					"days_later": pointer.To(int64(28)),
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestMapOfInt64PtrModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand map of Int64: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_mapOfFloat64Required(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfFloatFWModel
		expect      *TestMapOfFloatModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestMapOfFloatFWModel{},
			expect: &TestMapOfFloatModel{},
		},
		{
			name: "zero",
			input: &TestMapOfFloatFWModel{
				MapFloatProperty: typehelpers.NewMapValueOfMust[types.Float64](ctx, map[string]attr.Value{}),
			},
			expect: &TestMapOfFloatModel{
				MapFloatProperty: map[string]float64{},
			},
		},
		{
			name: "value",
			input: &TestMapOfFloatFWModel{
				MapFloatProperty: typehelpers.NewMapValueOfMust[types.Float64](ctx, map[string]attr.Value{
					"avogadro": types.Float64Value(6.023),
				}),
			},
			expect: &TestMapOfFloatModel{
				MapFloatProperty: map[string]float64{
					"avogadro": 6.023,
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestMapOfFloatModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand map of Float64: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_mapOfFloatOptional(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfFloatFWModel
		expect      *TestMapOfFloatPtrModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestMapOfFloatFWModel{},
			expect: &TestMapOfFloatPtrModel{},
		},
		{
			name: "zero",
			input: &TestMapOfFloatFWModel{
				MapFloatProperty: typehelpers.NewMapValueOfMust[types.Float64](ctx, map[string]attr.Value{}),
			},
			expect: &TestMapOfFloatPtrModel{
				MapFloatProperty: map[string]*float64{},
			},
		},
		{
			name: "value",
			input: &TestMapOfFloatFWModel{
				MapFloatProperty: typehelpers.NewMapValueOfMust[types.Float64](ctx, map[string]attr.Value{
					"avogadro": types.Float64Value(6.023),
				}),
			},
			expect: &TestMapOfFloatPtrModel{
				MapFloatProperty: map[string]*float64{
					"avogadro": pointer.To(6.023),
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestMapOfFloatPtrModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand map of Float64: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_listOfBool(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestListOfBoolFWModel
		expect      *TestListOfBoolModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestListOfBoolFWModel{},
			expect: &TestListOfBoolModel{},
		},
		{
			name: "explicit null list",
			input: &TestListOfBoolFWModel{
				ListBoolProperty: typehelpers.NewListValueOfNull[types.Bool](ctx),
			},
			expect: &TestListOfBoolModel{},
		},
		{
			name: "empty list",
			input: &TestListOfBoolFWModel{
				ListBoolProperty: typehelpers.NewListValueOfMust[types.Bool](ctx, []attr.Value{}),
			},
			expect: &TestListOfBoolModel{
				ListBoolProperty: []bool{},
			},
		},
		{
			name: "single element",
			input: &TestListOfBoolFWModel{
				ListBoolProperty: typehelpers.NewListValueOfMust[types.Bool](ctx, []attr.Value{
					types.BoolValue(true),
				}),
			},
			expect: &TestListOfBoolModel{
				ListBoolProperty: []bool{
					true,
				},
			},
		},
		{
			name: "multiple elements",
			input: &TestListOfBoolFWModel{
				ListBoolProperty: typehelpers.NewListValueOfMust[types.Bool](ctx, []attr.Value{
					types.BoolValue(true),
					types.BoolValue(true),
					types.BoolValue(false),
				}),
			},
			expect: &TestListOfBoolModel{
				ListBoolProperty: []bool{
					true,
					true,
					false,
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestListOfBoolModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand list of string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_listOfString(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestListOfStringFWModel
		expect      *TestListOfStringModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestListOfStringFWModel{},
			expect: &TestListOfStringModel{},
		},
		{
			name: "explicit null list",
			input: &TestListOfStringFWModel{
				ListStringProperty: typehelpers.NewListValueOfNull[types.String](ctx),
			},
			expect: &TestListOfStringModel{},
		},
		{
			name: "empty list",
			input: &TestListOfStringFWModel{
				ListStringProperty: typehelpers.NewListValueOfMust[types.String](ctx, []attr.Value{}),
			},
			expect: &TestListOfStringModel{
				ListStringProperty: []string{},
			},
		},
		{
			name: "single element",
			input: &TestListOfStringFWModel{
				ListStringProperty: typehelpers.NewListValueOfMust[types.String](ctx, []attr.Value{
					types.StringValue("foo"),
				}),
			},
			expect: &TestListOfStringModel{
				ListStringProperty: []string{
					"foo",
				},
			},
		},
		{
			name: "multiple elements",
			input: &TestListOfStringFWModel{
				ListStringProperty: typehelpers.NewListValueOfMust[types.String](ctx, []attr.Value{
					types.StringValue("foo"),
					types.StringValue("bar"),
					types.StringValue("peb"),
				}),
			},
			expect: &TestListOfStringModel{
				ListStringProperty: []string{
					"foo",
					"bar",
					"peb",
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestListOfStringModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand list of string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_listOfFloat(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestListOfFloatFWModel
		expect      *TestListOfFloatModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestListOfFloatFWModel{},
			expect: &TestListOfFloatModel{},
		},
		{
			name: "explicit null list",
			input: &TestListOfFloatFWModel{
				ListFloatProperty: typehelpers.NewListValueOfNull[types.Float64](ctx),
			},
			expect: &TestListOfFloatModel{},
		},
		{
			name: "empty list",
			input: &TestListOfFloatFWModel{
				ListFloatProperty: typehelpers.NewListValueOfMust[types.Float64](ctx, []attr.Value{}),
			},
			expect: &TestListOfFloatModel{
				ListFloatProperty: []float64{},
			},
		},
		{
			name: "single element",
			input: &TestListOfFloatFWModel{
				ListFloatProperty: typehelpers.NewListValueOfMust[types.Float64](ctx, []attr.Value{
					types.Float64Value(3.142),
				}),
			},
			expect: &TestListOfFloatModel{
				ListFloatProperty: []float64{
					3.142,
				},
			},
		},
		{
			name: "multiple elements",
			input: &TestListOfFloatFWModel{
				ListFloatProperty: typehelpers.NewListValueOfMust[types.Float64](ctx, []attr.Value{
					types.Float64Value(3.142),
					types.Float64Value(1.1),
					types.Float64Value(9.009),
				}),
			},
			expect: &TestListOfFloatModel{
				ListFloatProperty: []float64{
					3.142,
					1.1,
					9.009,
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestListOfFloatModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand list of string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_listOfInt(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestListOfIntFWModel
		expect      *TestListOfIntModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestListOfIntFWModel{},
			expect: &TestListOfIntModel{},
		},
		{
			name: "explicit null list",
			input: &TestListOfIntFWModel{
				ListIntProperty: typehelpers.NewListValueOfNull[types.Int64](ctx),
			},
			expect: &TestListOfIntModel{},
		},
		{
			name: "empty list",
			input: &TestListOfIntFWModel{
				ListIntProperty: typehelpers.NewListValueOfMust[types.Int64](ctx, []attr.Value{}),
			},
			expect: &TestListOfIntModel{
				ListIntProperty: []int64{},
			},
		},
		{
			name: "single element",
			input: &TestListOfIntFWModel{
				ListIntProperty: typehelpers.NewListValueOfMust[types.Int64](ctx, []attr.Value{
					types.Int64Value(101),
				}),
			},
			expect: &TestListOfIntModel{
				ListIntProperty: []int64{
					101,
				},
			},
		},
		{
			name: "multiple elements",
			input: &TestListOfIntFWModel{
				ListIntProperty: typehelpers.NewListValueOfMust[types.Int64](ctx, []attr.Value{
					types.Int64Value(101),
					types.Int64Value(202),
					types.Int64Value(303),
				}),
			},
			expect: &TestListOfIntModel{
				ListIntProperty: []int64{
					101,
					202,
					303,
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestListOfIntModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand list of string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

// TODO - Sets
func TestExpand_setOfBool(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestSetOfBoolFWModel
		expect      *TestSetOfBoolModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestSetOfBoolFWModel{},
			expect: &TestSetOfBoolModel{},
		},
		{
			name: "explicit null set",
			input: &TestSetOfBoolFWModel{
				SetBoolProperty: typehelpers.NewSetValueOfNull[types.Bool](ctx),
			},
			expect: &TestSetOfBoolModel{},
		},
		{
			name: "empty set",
			input: &TestSetOfBoolFWModel{
				SetBoolProperty: typehelpers.NewSetValueOfMust[types.Bool](ctx, []attr.Value{}),
			},
			expect: &TestSetOfBoolModel{
				SetBoolProperty: []bool{},
			},
		},
		{
			name: "single element",
			input: &TestSetOfBoolFWModel{
				SetBoolProperty: typehelpers.NewSetValueOfMust[types.Bool](ctx, []attr.Value{
					types.BoolValue(true),
				}),
			},
			expect: &TestSetOfBoolModel{
				SetBoolProperty: []bool{
					true,
				},
			},
		},
		{
			name: "multiple elements",
			input: &TestSetOfBoolFWModel{
				SetBoolProperty: typehelpers.NewSetValueOfMust[types.Bool](ctx, []attr.Value{
					types.BoolValue(true),
					types.BoolValue(true),
					types.BoolValue(false),
				}),
			},
			expect: &TestSetOfBoolModel{
				SetBoolProperty: []bool{
					true,
					true,
					false,
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestSetOfBoolModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand set of string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_setOfString(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestSetOfStringFWModel
		expect      *TestSetOfStringModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestSetOfStringFWModel{},
			expect: &TestSetOfStringModel{},
		},
		{
			name: "explicit null set",
			input: &TestSetOfStringFWModel{
				SetStringProperty: typehelpers.NewSetValueOfNull[types.String](ctx),
			},
			expect: &TestSetOfStringModel{},
		},
		{
			name: "empty set",
			input: &TestSetOfStringFWModel{
				SetStringProperty: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{}),
			},
			expect: &TestSetOfStringModel{
				SetStringProperty: []string{},
			},
		},
		{
			name: "single element",
			input: &TestSetOfStringFWModel{
				SetStringProperty: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{
					types.StringValue("foo"),
				}),
			},
			expect: &TestSetOfStringModel{
				SetStringProperty: []string{
					"foo",
				},
			},
		},
		{
			name: "multiple elements",
			input: &TestSetOfStringFWModel{
				SetStringProperty: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{
					types.StringValue("foo"),
					types.StringValue("bar"),
					types.StringValue("peb"),
				}),
			},
			expect: &TestSetOfStringModel{
				SetStringProperty: []string{
					"foo",
					"bar",
					"peb",
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestSetOfStringModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand set of string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_setOfFloat(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestSetOfFloatFWModel
		expect      *TestSetOfFloatModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestSetOfFloatFWModel{},
			expect: &TestSetOfFloatModel{},
		},
		{
			name: "explicit null set",
			input: &TestSetOfFloatFWModel{
				SetFloatProperty: typehelpers.NewSetValueOfNull[types.Float64](ctx),
			},
			expect: &TestSetOfFloatModel{},
		},
		{
			name: "empty set",
			input: &TestSetOfFloatFWModel{
				SetFloatProperty: typehelpers.NewSetValueOfMust[types.Float64](ctx, []attr.Value{}),
			},
			expect: &TestSetOfFloatModel{
				SetFloatProperty: []float64{},
			},
		},
		{
			name: "single element",
			input: &TestSetOfFloatFWModel{
				SetFloatProperty: typehelpers.NewSetValueOfMust[types.Float64](ctx, []attr.Value{
					types.Float64Value(3.142),
				}),
			},
			expect: &TestSetOfFloatModel{
				SetFloatProperty: []float64{
					3.142,
				},
			},
		},
		{
			name: "multiple elements",
			input: &TestSetOfFloatFWModel{
				SetFloatProperty: typehelpers.NewSetValueOfMust[types.Float64](ctx, []attr.Value{
					types.Float64Value(3.142),
					types.Float64Value(1.1),
					types.Float64Value(9.009),
				}),
			},
			expect: &TestSetOfFloatModel{
				SetFloatProperty: []float64{
					3.142,
					1.1,
					9.009,
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestSetOfFloatModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand set of string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_setOfInt(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestSetOfIntFWModel
		expect      *TestSetOfIntModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestSetOfIntFWModel{},
			expect: &TestSetOfIntModel{},
		},
		{
			name: "explicit null set",
			input: &TestSetOfIntFWModel{
				SetIntProperty: typehelpers.NewSetValueOfNull[types.Int64](ctx),
			},
			expect: &TestSetOfIntModel{},
		},
		{
			name: "empty set",
			input: &TestSetOfIntFWModel{
				SetIntProperty: typehelpers.NewSetValueOfMust[types.Int64](ctx, []attr.Value{}),
			},
			expect: &TestSetOfIntModel{
				SetIntProperty: []int64{},
			},
		},
		{
			name: "single element",
			input: &TestSetOfIntFWModel{
				SetIntProperty: typehelpers.NewSetValueOfMust[types.Int64](ctx, []attr.Value{
					types.Int64Value(101),
				}),
			},
			expect: &TestSetOfIntModel{
				SetIntProperty: []int64{
					101,
				},
			},
		},
		{
			name: "multiple elements",
			input: &TestSetOfIntFWModel{
				SetIntProperty: typehelpers.NewSetValueOfMust[types.Int64](ctx, []attr.Value{
					types.Int64Value(101),
					types.Int64Value(202),
					types.Int64Value(303),
				}),
			},
			expect: &TestSetOfIntModel{
				SetIntProperty: []int64{
					101,
					202,
					303,
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestSetOfIntModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand set of string: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_nestedOneLevel(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestNestedOneLevelFWModel
		expect      *TestNestedOneLevelModel
		expectError bool
	}{
		{
			name:  "null",
			input: &TestNestedOneLevelFWModel{},
			expect: &TestNestedOneLevelModel{
				TopLevelString: "",
			},
		},
		{
			name: "zero",
			input: &TestNestedOneLevelFWModel{
				NestedModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestStringOnlyFWModel](ctx, []TestStringOnlyFWModel{}),
			},
			expect: &TestNestedOneLevelModel{
				TopLevelString: "",
				NestedModel:    []TestStringOptionalModel{},
			},
		},
		{
			name: "value",
			input: &TestNestedOneLevelFWModel{
				NestedModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestStringOnlyFWModel](ctx, []TestStringOnlyFWModel{
					{
						StringProperty: types.StringValue("foo"),
					},
				}),
			},
			expect: &TestNestedOneLevelModel{
				TopLevelString: "",
				NestedModel: []TestStringOptionalModel{
					{
						StringProperty: pointer.To("foo"),
					},
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestNestedOneLevelModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand nestedOneLevel: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_nestedTwoLevel(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestNestedTwoLevelFWModel
		expect      *TestNestedTwoLevelModel
		expectError bool
	}{
		{
			name:  "null",
			input: &TestNestedTwoLevelFWModel{},
			expect: &TestNestedTwoLevelModel{
				TopLevelString: "",
			},
		},
		{
			name: "zero",
			input: &TestNestedTwoLevelFWModel{
				NestedModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestNestedOneLevelFWModel](ctx, []TestNestedOneLevelFWModel{}),
			},
			expect: &TestNestedTwoLevelModel{
				TopLevelString: "",
				NestedModel:    []TestNestedOneLevelModel{},
			},
		},
		{
			name: "value",
			input: &TestNestedTwoLevelFWModel{
				NestedModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestNestedOneLevelFWModel](ctx, []TestNestedOneLevelFWModel{
					{
						NestedModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestStringOnlyFWModel](ctx, []TestStringOnlyFWModel{
							{
								StringProperty: types.StringValue("foo"),
							},
						}),
					},
				}),
			},
			expect: &TestNestedTwoLevelModel{
				TopLevelString: "",
				NestedModel: []TestNestedOneLevelModel{
					{
						NestedModel: []TestStringOptionalModel{
							{
								StringProperty: pointer.To("foo"),
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestNestedTwoLevelModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand nestedOneLevel: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: %+v \n Got: %+v", testCase.expect, result)
		}
	}
}

func TestExpand_complexModel(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestNestedComplexFWModel
		expect      *TestNestedComplexModel
		expectError bool
	}{
		{
			name:   "null",
			input:  &TestNestedComplexFWModel{},
			expect: &TestNestedComplexModel{},
		},
		{
			name: "zero",
			input: &TestNestedComplexFWModel{
				TopLevelString:      types.StringValue(""),
				TopLevelOptionalInt: types.Int64Value(0),
				NestedSimpleModel:   typehelpers.NewListNestedObjectValueOfValueSliceMust[TestStringOnlyFWModel](ctx, []TestStringOnlyFWModel{}),
				NestedComplexModel:  typehelpers.NewListNestedObjectValueOfValueSliceMust[TestFrameworkModelComplex](ctx, []TestFrameworkModelComplex{}),
			},
			expect: &TestNestedComplexModel{
				TopLevelString:      "",
				TopLevelOptionalInt: pointer.To(int64(0)),
				NestedSimpleModel:   []TestStringOptionalModel{},
				NestedComplexModel:  []TestAPIModelComplex{},
			},
		},
		{
			name: "partial simple",
			input: &TestNestedComplexFWModel{
				TopLevelString:      types.StringValue(""),
				TopLevelOptionalInt: types.Int64Value(0),
				NestedSimpleModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestStringOnlyFWModel](ctx, []TestStringOnlyFWModel{
					{
						StringProperty: types.StringValue("foo"),
					},
				}),
				NestedComplexModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestFrameworkModelComplex](ctx, []TestFrameworkModelComplex{}),
			},
			expect: &TestNestedComplexModel{
				TopLevelString:      "",
				TopLevelOptionalInt: pointer.To(int64(0)),
				NestedSimpleModel: []TestStringOptionalModel{
					{
						StringProperty: pointer.To("foo"),
					},
				},
				NestedComplexModel: []TestAPIModelComplex{},
			},
		},
		{
			name: "partial simple multiple",
			input: &TestNestedComplexFWModel{
				TopLevelString:      types.StringValue(""),
				TopLevelOptionalInt: types.Int64Value(0),
				NestedSimpleModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestStringOnlyFWModel](ctx, []TestStringOnlyFWModel{
					{
						StringProperty: types.StringValue("foo"),
					},
					{
						StringProperty: types.StringValue("bar"),
					},
					{
						StringProperty: types.StringValue("blah"),
					},
				}),
				NestedComplexModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestFrameworkModelComplex](ctx, []TestFrameworkModelComplex{}),
			},
			expect: &TestNestedComplexModel{
				TopLevelString:      "",
				TopLevelOptionalInt: pointer.To(int64(0)),
				NestedSimpleModel: []TestStringOptionalModel{
					{
						StringProperty: pointer.To("foo"),
					},
					{
						StringProperty: pointer.To("bar"),
					},
					{
						StringProperty: pointer.To("blah"),
					},
				},
				NestedComplexModel: []TestAPIModelComplex{},
			},
		},
		{
			name: "partial Complex",
			input: &TestNestedComplexFWModel{
				TopLevelString:      types.StringValue(""),
				TopLevelOptionalInt: types.Int64Value(0),
				NestedSimpleModel:   typehelpers.NewListNestedObjectValueOfValueSliceMust[TestStringOnlyFWModel](ctx, []TestStringOnlyFWModel{}),
				NestedComplexModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestFrameworkModelComplex](ctx, []TestFrameworkModelComplex{
					{
						BoolProperty:   types.BoolValue(true),
						IntProperty:    types.Int64Value(1),
						FloatProperty:  types.Float64Value(3.142),
						StringProperty: types.StringValue("foo"),
						ListProperty: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestFrameworkNestedModel](ctx, []TestFrameworkNestedModel{
							{
								SubPropertyBool:   types.BoolValue(true),
								SubPropertyInt:    types.Int64Value(2),
								SubPropertyFloat:  types.Float64Value(3.1415),
								SubPropertyString: types.StringValue("bar"),
							},
						}),
						MapBoolProperty: typehelpers.NewMapValueOfMust[types.Bool](ctx, map[string]attr.Value{
							"is_true":  types.BoolValue(true),
							"not_true": types.BoolValue(false),
						}),
						MapInt64Property: typehelpers.NewMapValueOfMust[types.Int64](ctx, map[string]attr.Value{
							"two":                  types.Int64Value(2),
							"one_hundred_and_four": types.Int64Value(104),
						}),
						MapFloatProperty: typehelpers.NewMapValueOfMust[types.Float64](ctx, map[string]attr.Value{
							"pi":       types.Float64Value(3.14159),
							"avogadro": types.Float64Value(6.0221407),
						}),
						MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{
							"foo": types.StringValue("bar"),
							"peb": types.StringValue("cak"),
						}),
						SetProperty: typehelpers.NewSetNestedObjectValueOfValueSliceMust[TestFrameworkNestedModel](ctx, []TestFrameworkNestedModel{
							{
								SubPropertyBool:   types.BoolValue(true),
								SubPropertyInt:    types.Int64Value(2),
								SubPropertyFloat:  types.Float64Value(9.1093837139),
								SubPropertyString: types.StringValue("bar"),
							},
						}),
					},
				}),
			},
			expect: &TestNestedComplexModel{
				TopLevelString:      "",
				TopLevelOptionalInt: pointer.To(int64(0)),
				NestedSimpleModel:   []TestStringOptionalModel{},
				NestedComplexModel: []TestAPIModelComplex{
					{
						BoolProperty:   true,
						IntProperty:    int64(1),
						FloatProperty:  3.142,
						StringProperty: "foo",
						ListProperty: []TestAPINestedModel{
							{
								SubPropertyBool:   true,
								SubPropertyInt:    int64(2),
								SubPropertyFloat:  3.1415,
								SubPropertyString: "bar",
							},
						},
						MapBoolProperty: map[string]bool{
							"is_true":  true,
							"not_true": false,
						},
						MapInt64Property: map[string]int64{
							"two":                  int64(2),
							"one_hundred_and_four": int64(104),
						},
						MapFloatProperty: map[string]float64{
							"pi":       3.14159,
							"avogadro": 6.0221407,
						},
						MapStringProperty: map[string]string{
							"foo": "bar",
							"peb": "cak",
						},
						SetProperty: []TestAPINestedModel{
							{
								SubPropertyBool:   true,
								SubPropertyInt:    int64(2),
								SubPropertyFloat:  9.1093837139,
								SubPropertyString: "bar",
							},
						},
					},
				},
			},
		},
		{
			name: "Full Complex",
			input: &TestNestedComplexFWModel{
				TopLevelString:      types.StringValue(""),
				TopLevelOptionalInt: types.Int64Value(0),
				NestedSimpleModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestStringOnlyFWModel](ctx, []TestStringOnlyFWModel{
					{
						StringProperty: types.StringValue("foo"),
					},
					{
						StringProperty: types.StringValue("bar"),
					},
					{
						StringProperty: types.StringValue("blah"),
					},
				}),
				NestedComplexModel: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestFrameworkModelComplex](ctx, []TestFrameworkModelComplex{
					{
						BoolProperty:   types.BoolValue(true),
						IntProperty:    types.Int64Value(1),
						FloatProperty:  types.Float64Value(3.142),
						StringProperty: types.StringValue("foo"),
						// ListOfPrimitives: typehelpers.NewListValueOfNull[types.String](ctx),
						ListProperty: typehelpers.NewListNestedObjectValueOfValueSliceMust[TestFrameworkNestedModel](ctx, []TestFrameworkNestedModel{
							{
								SubPropertyBool:   types.BoolValue(true),
								SubPropertyInt:    types.Int64Value(2),
								SubPropertyFloat:  types.Float64Value(3.1415),
								SubPropertyString: types.StringValue("bar"),
							},
							{
								SubPropertyBool:   types.BoolValue(false),
								SubPropertyInt:    types.Int64Value(3),
								SubPropertyFloat:  types.Float64Value(1.1),
								SubPropertyString: types.StringValue("peb"),
							},
						}),
						MapBoolProperty: typehelpers.NewMapValueOfMust[types.Bool](ctx, map[string]attr.Value{
							"is_true":  types.BoolValue(true),
							"not_true": types.BoolValue(false),
						}),
						MapInt64Property: typehelpers.NewMapValueOfMust[types.Int64](ctx, map[string]attr.Value{
							"two":                  types.Int64Value(2),
							"one_hundred_and_four": types.Int64Value(104),
						}),
						MapFloatProperty: typehelpers.NewMapValueOfMust[types.Float64](ctx, map[string]attr.Value{
							"pi":       types.Float64Value(3.14159),
							"avogadro": types.Float64Value(6.0221407),
						}),
						MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{
							"foo": types.StringValue("bar"),
							"peb": types.StringValue("cak"),
						}),
						SetProperty: typehelpers.NewSetNestedObjectValueOfValueSliceMust[TestFrameworkNestedModel](ctx, []TestFrameworkNestedModel{
							{
								SubPropertyBool:   types.BoolValue(true),
								SubPropertyInt:    types.Int64Value(2),
								SubPropertyFloat:  types.Float64Value(9.1093837139),
								SubPropertyString: types.StringValue("bar"),
							},
							{
								SubPropertyBool:   types.BoolValue(false),
								SubPropertyInt:    types.Int64Value(4),
								SubPropertyFloat:  types.Float64Value(9.1),
								SubPropertyString: types.StringValue("the cake is a lie"),
							},
						}),
					},
				}),
			},
			expect: &TestNestedComplexModel{
				TopLevelString:      "",
				TopLevelOptionalInt: pointer.To(int64(0)),
				NestedSimpleModel: []TestStringOptionalModel{
					{
						StringProperty: pointer.To("foo"),
					},
					{
						StringProperty: pointer.To("bar"),
					},
					{
						StringProperty: pointer.To("blah"),
					},
				},
				NestedComplexModel: []TestAPIModelComplex{
					{
						BoolProperty:   true,
						IntProperty:    int64(1),
						FloatProperty:  3.142,
						StringProperty: "foo",
						ListProperty: []TestAPINestedModel{
							{
								SubPropertyBool:   true,
								SubPropertyInt:    int64(2),
								SubPropertyFloat:  3.1415,
								SubPropertyString: "bar",
							},
							{
								SubPropertyBool:   false,
								SubPropertyInt:    int64(3),
								SubPropertyFloat:  1.1,
								SubPropertyString: "peb",
							},
						},
						MapBoolProperty: map[string]bool{
							"is_true":  true,
							"not_true": false,
						},
						MapInt64Property: map[string]int64{
							"two":                  int64(2),
							"one_hundred_and_four": int64(104),
						},
						MapFloatProperty: map[string]float64{
							"pi":       3.14159,
							"avogadro": 6.0221407,
						},
						MapStringProperty: map[string]string{
							"foo": "bar",
							"peb": "cak",
						},
						SetProperty: []TestAPINestedModel{
							{
								SubPropertyBool:   true,
								SubPropertyInt:    int64(2),
								SubPropertyFloat:  9.1093837139,
								SubPropertyString: "bar",
							},
							{
								SubPropertyBool:   false,
								SubPropertyInt:    int64(4),
								SubPropertyFloat:  9.1,
								SubPropertyString: "the cake is a lie",
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range cases {
		result := &TestNestedComplexModel{}

		diags := &diag.Diagnostics{}

		t.Logf("testing expand complexModel: %s", testCase.name)
		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() && !testCase.expectError {
			t.Errorf("Expand failed: %+v", diags.Errors())
		}
		if !diags.HasError() && testCase.expectError {
			t.Errorf("Expand failed, expected error but didn't get one")
		}
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Errorf("Expand failed, expected: \n%+v \n Got: \n%+v", testCase.expect, result)
		}
	}
}
