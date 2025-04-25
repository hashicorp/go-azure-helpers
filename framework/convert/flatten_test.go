// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert_test

import (
	"context"
	"encoding/base64"
	"math/rand"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/go-azure-helpers/framework/convert"
	"github.com/hashicorp/go-azure-helpers/framework/typehelpers"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	intVal   = rand.Int63()
	floatVal = rand.Float64()
	strVal   = randomString(rand.Intn(60))
)

func randomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func TestFlatten_complex(t *testing.T) {
	ctx := context.Background()
	input := &TestAPIModelComplex{
		BoolProperty:   true,
		StringProperty: "foo",
		IntProperty:    365,
		FloatProperty:  3.14,
		ListProperty: []TestAPINestedModel{
			{
				SubPropertyBool: true,
			},
		},
	}

	expected := &TestFrameworkModelComplex{
		BoolProperty:     types.BoolValue(true),
		StringProperty:   types.StringValue("foo"),
		IntProperty:      types.Int64Value(365),
		FloatProperty:    types.Float64Value(3.14),
		ListOfPrimitives: typehelpers.NewListValueOfNull[types.String](ctx),
		ListProperty: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []TestFrameworkNestedModel{
			{
				SubPropertyBool:   types.BoolValue(true),
				SubPropertyFloat:  types.Float64Value(0),
				SubPropertyInt:    types.Int64Value(0),
				SubPropertyString: types.StringValue(""),
			},
		}),
		SetProperty:       typehelpers.NewSetNestedObjectValueOfNull[TestFrameworkNestedModel](ctx),
		MapStringProperty: typehelpers.NewMapValueOfNull[types.String](ctx),
		MapFloatProperty:  typehelpers.NewMapValueOfNull[types.Float64](ctx),
		MapBoolProperty:   typehelpers.NewMapValueOfNull[types.Bool](ctx),
		MapInt64Property:  typehelpers.NewMapValueOfNull[types.Int64](ctx),
	}

	result := &TestFrameworkModelComplex{}
	diags := &diag.Diagnostics{}

	convert.Flatten(context.Background(), input, result, diags)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nexpected:\n %+v\ngot:\n %+v", expected, result)
	}

	if diags.HasError() {
		t.Errorf("diags: %+v", diags)
	}
}

func TestFlatten_boolRequiredProperty(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestBoolRequiredModel
		expected    *TestBoolOnlyFWModel
		expectError bool
	}{
		{
			name:        "empty input",
			input:       &TestBoolRequiredModel{},
			expectError: true,
		},
		{
			name: "true input",
			input: &TestBoolRequiredModel{
				BoolProperty: true,
			},
			expected: &TestBoolOnlyFWModel{
				BoolProperty: types.BoolValue(true),
			},
		},
		{
			name: "false input",
			input: &TestBoolRequiredModel{
				BoolProperty: false,
			},
			expected: &TestBoolOnlyFWModel{
				BoolProperty: types.BoolValue(false),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestBoolOnlyFWModel{}

		t.Logf("testing flatten Bool Required: %s", c.name)
		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_boolOptionalProperty(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestBoolOptionalModel
		expected    *TestBoolOnlyFWModel
		expectError bool
	}{
		{
			name:  "empty input",
			input: &TestBoolOptionalModel{},
			expected: &TestBoolOnlyFWModel{
				BoolProperty: types.BoolNull(),
			},
		},
		{
			name: "true input",
			input: &TestBoolOptionalModel{
				BoolProperty: pointer.To(true),
			},
			expected: &TestBoolOnlyFWModel{
				BoolProperty: types.BoolValue(true),
			},
		},
		{
			name: "false input",
			input: &TestBoolOptionalModel{
				BoolProperty: pointer.To(false),
			},
			expected: &TestBoolOnlyFWModel{
				BoolProperty: types.BoolValue(false),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestBoolOnlyFWModel{}

		t.Logf("testing flatten Bool Optional: %s", c.name)
		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_stringRequiredProperty(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestStringRequiredModel
		expected    *TestStringOnlyFWModel
		expectError bool
	}{
		{
			name:        "omitted property",
			input:       &TestStringRequiredModel{},
			expectError: true,
		},
		{
			name: "zero length input",
			input: &TestStringRequiredModel{
				StringProperty: "",
			},
			expected: &TestStringOnlyFWModel{
				StringProperty: types.StringValue(""),
			},
		},
		{
			name: "valid input",
			input: &TestStringRequiredModel{
				StringProperty: strVal,
			},
			expected: &TestStringOnlyFWModel{
				StringProperty: types.StringValue(strVal),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestStringOnlyFWModel{}

		t.Logf("testing flatten String Required: %s", c.name)
		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_stringOptionalProperty(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestStringOptionalModel
		expected    *TestStringOnlyFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestStringOptionalModel{},
			expected: &TestStringOnlyFWModel{
				StringProperty: types.StringNull(),
			},
		},
		{
			name: "zero length input",
			input: &TestStringOptionalModel{
				StringProperty: pointer.To(""),
			},
			expected: &TestStringOnlyFWModel{
				StringProperty: types.StringValue(""),
			},
		},
		{
			name: "valid input",
			input: &TestStringOptionalModel{
				StringProperty: pointer.To(strVal),
			},
			expected: &TestStringOnlyFWModel{
				StringProperty: types.StringValue(strVal),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestStringOnlyFWModel{}

		t.Logf("testing flatten String Optional: %s", c.name)
		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_int64RequiredProperty(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestInt64RequiredModel
		expected    *TestInt64OnlyFWModel
		expectError bool
	}{
		{
			name:        "omitted property",
			input:       &TestInt64RequiredModel{},
			expectError: true,
		},
		{
			name: "valid input",
			input: &TestInt64RequiredModel{
				Int64Property: intVal,
			},
			expected: &TestInt64OnlyFWModel{
				Int64Property: types.Int64Value(intVal),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestInt64OnlyFWModel{}

		t.Logf("testing flatten Int64 Required: %s", c.name)
		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_int64OptionalProperty(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestInt64OptionalModel
		expected    *TestInt64OnlyFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestInt64OptionalModel{},
			expected: &TestInt64OnlyFWModel{
				Int64Property: types.Int64Null(),
			},
		},
		{
			name: "valid input",
			input: &TestInt64OptionalModel{
				Int64Property: pointer.To(intVal),
			},
			expected: &TestInt64OnlyFWModel{
				Int64Property: types.Int64Value(intVal),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestInt64OnlyFWModel{}

		t.Logf("testing flatten Int64 Optional: %s", c.name)
		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_float64RequiredProperty(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestFloat64RequiredModel
		expected    *TestFloat64OnlyFWModel
		expectError bool
	}{
		{
			name:        "omitted property",
			input:       &TestFloat64RequiredModel{},
			expectError: true,
		},
		{
			name: "valid input",
			input: &TestFloat64RequiredModel{
				Float64Property: floatVal,
			},
			expected: &TestFloat64OnlyFWModel{
				Float64Property: types.Float64Value(floatVal),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestFloat64OnlyFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_float64OptionalProperty(t *testing.T) {
	cases := []struct {
		name        string
		input       *TestFloat64OptionalModel
		expected    *TestFloat64OnlyFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestFloat64OptionalModel{},
			expected: &TestFloat64OnlyFWModel{
				Float64Property: types.Float64Null(),
			},
		},
		{
			name: "valid input",
			input: &TestFloat64OptionalModel{
				Float64Property: pointer.To(floatVal),
			},
			expected: &TestFloat64OnlyFWModel{
				Float64Property: types.Float64Value(floatVal),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestFloat64OnlyFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_mapOfFloat(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfFloatModel
		expected    *TestMapOfFloatFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestMapOfFloatModel{},
			expected: &TestMapOfFloatFWModel{
				MapFloatProperty: typehelpers.NewMapValueOfNull[types.Float64](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestMapOfFloatModel{
				MapFloatProperty: map[string]float64{},
			},
			expected: &TestMapOfFloatFWModel{
				MapFloatProperty: typehelpers.NewMapValueOfMust[types.Float64](ctx, map[string]attr.Value{}),
			},
		},
		{
			name: "length 1 input",
			input: &TestMapOfFloatModel{
				MapFloatProperty: map[string]float64{
					"foo": floatVal,
				},
			},
			expected: &TestMapOfFloatFWModel{
				MapFloatProperty: typehelpers.NewMapValueOfMust[types.Float64](ctx, map[string]attr.Value{
					"foo": types.Float64Value(floatVal),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestMapOfFloatFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_mapOfFloatPtr(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfFloatPtrModel
		expected    *TestMapOfFloatFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestMapOfFloatPtrModel{},
			expected: &TestMapOfFloatFWModel{
				MapFloatProperty: typehelpers.NewMapValueOfNull[types.Float64](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestMapOfFloatPtrModel{
				MapFloatProperty: map[string]*float64{},
			},
			expected: &TestMapOfFloatFWModel{
				MapFloatProperty: typehelpers.NewMapValueOfMust[types.Float64](ctx, map[string]attr.Value{}),
			},
		},
		{
			name: "length 1 input",
			input: &TestMapOfFloatPtrModel{
				MapFloatProperty: map[string]*float64{
					"foo": pointer.To(floatVal),
				},
			},
			expected: &TestMapOfFloatFWModel{
				MapFloatProperty: typehelpers.NewMapValueOfMust[types.Float64](ctx, map[string]attr.Value{
					"foo": types.Float64Value(floatVal),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestMapOfFloatFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_mapOfString(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfStringModel
		expected    *TestMapOfStringFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestMapOfStringModel{},
			expected: &TestMapOfStringFWModel{
				MapStringProperty: typehelpers.NewMapValueOfNull[types.String](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestMapOfStringModel{
				MapStringProperty: map[string]string{},
			},
			expected: &TestMapOfStringFWModel{
				MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{}),
			},
		},
		{
			name: "length 1 input",
			input: &TestMapOfStringModel{
				MapStringProperty: map[string]string{
					"foo": strVal,
				},
			},
			expected: &TestMapOfStringFWModel{
				MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{
					"foo": types.StringValue(strVal),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestMapOfStringFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_mapOfStringPtr(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfStringPtrModel
		expected    *TestMapOfStringFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestMapOfStringPtrModel{},
			expected: &TestMapOfStringFWModel{
				MapStringProperty: typehelpers.NewMapValueOfNull[types.String](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestMapOfStringPtrModel{
				MapStringProperty: map[string]*string{},
			},
			expected: &TestMapOfStringFWModel{
				MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{}),
			},
		},
		{
			name: "length 1 input",
			input: &TestMapOfStringPtrModel{
				MapStringProperty: map[string]*string{
					"foo": pointer.To(strVal),
				},
			},
			expected: &TestMapOfStringFWModel{
				MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{
					"foo": types.StringValue(strVal),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestMapOfStringFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_mapOfInt64(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfInt64Model
		expected    *TestMapOfInt64FWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestMapOfInt64Model{},
			expected: &TestMapOfInt64FWModel{
				MapInt64Property: typehelpers.NewMapValueOfNull[types.Int64](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestMapOfInt64Model{
				MapInt64Property: map[string]int64{},
			},
			expected: &TestMapOfInt64FWModel{
				MapInt64Property: typehelpers.NewMapValueOfMust[types.Int64](ctx, map[string]attr.Value{}),
			},
		},
		{
			name: "length 1 input",
			input: &TestMapOfInt64Model{
				MapInt64Property: map[string]int64{
					"days": intVal,
				},
			},
			expected: &TestMapOfInt64FWModel{
				MapInt64Property: typehelpers.NewMapValueOfMust[types.Int64](ctx, map[string]attr.Value{
					"days": types.Int64Value(intVal),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestMapOfInt64FWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_mapOfInt64Ptr(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfInt64PtrModel
		expected    *TestMapOfInt64FWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestMapOfInt64PtrModel{},
			expected: &TestMapOfInt64FWModel{
				MapInt64Property: typehelpers.NewMapValueOfNull[types.Int64](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestMapOfInt64PtrModel{
				MapInt64Property: map[string]*int64{},
			},
			expected: &TestMapOfInt64FWModel{
				MapInt64Property: typehelpers.NewMapValueOfMust[types.Int64](ctx, map[string]attr.Value{}),
			},
		},
		{
			name: "length 1 input",
			input: &TestMapOfInt64PtrModel{
				MapInt64Property: map[string]*int64{
					"foo": pointer.To(intVal),
				},
			},
			expected: &TestMapOfInt64FWModel{
				MapInt64Property: typehelpers.NewMapValueOfMust[types.Int64](ctx, map[string]attr.Value{
					"foo": types.Int64Value(intVal),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestMapOfInt64FWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_mapOfBool(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfBoolModel
		expected    *TestMapOfBoolFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestMapOfBoolModel{},
			expected: &TestMapOfBoolFWModel{
				MapBoolProperty: typehelpers.NewMapValueOfNull[types.Bool](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestMapOfBoolModel{
				MapBoolProperty: map[string]bool{},
			},
			expected: &TestMapOfBoolFWModel{
				MapBoolProperty: typehelpers.NewMapValueOfMust[types.Bool](ctx, map[string]attr.Value{}),
			},
		},
		{
			name: "length 1 input",
			input: &TestMapOfBoolModel{
				MapBoolProperty: map[string]bool{
					"true": true,
				},
			},
			expected: &TestMapOfBoolFWModel{
				MapBoolProperty: typehelpers.NewMapValueOfMust[types.Bool](ctx, map[string]attr.Value{
					"true": types.BoolValue(true),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestMapOfBoolFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_listOfBool(t *testing.T) {
	ctx := context.Background()
	emptyList := make([]bool, 0)
	cases := []struct {
		name        string
		input       *TestListOfBoolModel
		expected    *TestListOfBoolFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestListOfBoolModel{},
			expected: &TestListOfBoolFWModel{
				ListBoolProperty: typehelpers.NewListValueOfNull[types.Bool](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestListOfBoolModel{
				ListBoolProperty: emptyList,
			},
			expected: &TestListOfBoolFWModel{
				ListBoolProperty: typehelpers.NewListValueOfMust[types.Bool](ctx, []attr.Value{}),
			},
		},
		{
			name: "single Element",
			input: &TestListOfBoolModel{
				ListBoolProperty: []bool{
					true,
				},
			},
			expected: &TestListOfBoolFWModel{
				ListBoolProperty: typehelpers.NewListValueOfMust[types.Bool](ctx, []attr.Value{
					types.BoolValue(true),
				}),
			},
		},
		{
			name: "multiple elements",
			input: &TestListOfBoolModel{
				ListBoolProperty: []bool{
					true,
					false,
					true,
				},
			},
			expected: &TestListOfBoolFWModel{
				ListBoolProperty: typehelpers.NewListValueOfMust[types.Bool](ctx, []attr.Value{
					types.BoolValue(true),
					types.BoolValue(false),
					types.BoolValue(true),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestListOfBoolFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_listOfFloat(t *testing.T) {
	ctx := context.Background()
	emptyList := make([]float64, 0)
	cases := []struct {
		name        string
		input       *TestListOfFloatModel
		expected    *TestListOfFloatFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestListOfFloatModel{},
			expected: &TestListOfFloatFWModel{
				ListFloatProperty: typehelpers.NewListValueOfNull[types.Float64](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestListOfFloatModel{
				ListFloatProperty: emptyList,
			},
			expected: &TestListOfFloatFWModel{
				ListFloatProperty: typehelpers.NewListValueOfMust[types.Float64](ctx, []attr.Value{}),
			},
		},
		{
			name: "single Element",
			input: &TestListOfFloatModel{
				ListFloatProperty: []float64{
					floatVal,
				},
			},
			expected: &TestListOfFloatFWModel{
				ListFloatProperty: typehelpers.NewListValueOfMust[types.Float64](ctx, []attr.Value{
					types.Float64Value(floatVal),
				}),
			},
		},
		{
			name: "multiple elements",
			input: &TestListOfFloatModel{
				ListFloatProperty: []float64{
					floatVal + 1,
					floatVal + 2,
					floatVal + 3,
				},
			},
			expected: &TestListOfFloatFWModel{
				ListFloatProperty: typehelpers.NewListValueOfMust[types.Float64](ctx, []attr.Value{
					types.Float64Value(floatVal + 1),
					types.Float64Value(floatVal + 2),
					types.Float64Value(floatVal + 3),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestListOfFloatFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_listOfInt(t *testing.T) {
	ctx := context.Background()
	emptyList := make([]int64, 0)
	cases := []struct {
		name        string
		input       *TestListOfIntModel
		expected    *TestListOfIntFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestListOfIntModel{},
			expected: &TestListOfIntFWModel{
				ListIntProperty: typehelpers.NewListValueOfNull[types.Int64](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestListOfIntModel{
				ListIntProperty: emptyList,
			},
			expected: &TestListOfIntFWModel{
				ListIntProperty: typehelpers.NewListValueOfMust[types.Int64](ctx, []attr.Value{}),
			},
		},
		{
			name: "single Element",
			input: &TestListOfIntModel{
				ListIntProperty: []int64{
					intVal,
				},
			},
			expected: &TestListOfIntFWModel{
				ListIntProperty: typehelpers.NewListValueOfMust[types.Int64](ctx, []attr.Value{
					types.Int64Value(intVal),
				}),
			},
		},
		{
			name: "multiple elements",
			input: &TestListOfIntModel{
				ListIntProperty: []int64{
					intVal + 1,
					intVal + 2,
					intVal + 3,
				},
			},
			expected: &TestListOfIntFWModel{
				ListIntProperty: typehelpers.NewListValueOfMust[types.Int64](ctx, []attr.Value{
					types.Int64Value(intVal + 1),
					types.Int64Value(intVal + 2),
					types.Int64Value(intVal + 3),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestListOfIntFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_listOfString(t *testing.T) {
	ctx := context.Background()
	emptyList := make([]string, 0)
	cases := []struct {
		name        string
		input       *TestListOfStringModel
		expected    *TestListOfStringFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestListOfStringModel{},
			expected: &TestListOfStringFWModel{
				ListStringProperty: typehelpers.NewListValueOfNull[types.String](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestListOfStringModel{
				ListStringProperty: emptyList,
			},
			expected: &TestListOfStringFWModel{
				ListStringProperty: typehelpers.NewListValueOfMust[types.String](ctx, []attr.Value{}),
			},
		},
		{
			name: "single Element",
			input: &TestListOfStringModel{
				ListStringProperty: []string{
					strVal,
				},
			},
			expected: &TestListOfStringFWModel{
				ListStringProperty: typehelpers.NewListValueOfMust[types.String](ctx, []attr.Value{
					types.StringValue(strVal),
				}),
			},
		},
		{
			name: "multiple elements",
			input: &TestListOfStringModel{
				ListStringProperty: []string{
					strVal + "1",
					strVal + "2",
					strVal + "3",
				},
			},
			expected: &TestListOfStringFWModel{
				ListStringProperty: typehelpers.NewListValueOfMust[types.String](ctx, []attr.Value{
					types.StringValue(strVal + "1"),
					types.StringValue(strVal + "2"),
					types.StringValue(strVal + "3"),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestListOfStringFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_setOfBool(t *testing.T) {
	ctx := context.Background()
	emptySet := make([]bool, 0)
	cases := []struct {
		name        string
		input       *TestSetOfBoolModel
		expected    *TestSetOfBoolFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestSetOfBoolModel{},
			expected: &TestSetOfBoolFWModel{
				SetBoolProperty: typehelpers.NewSetValueOfNull[types.Bool](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestSetOfBoolModel{
				SetBoolProperty: emptySet,
			},
			expected: &TestSetOfBoolFWModel{
				SetBoolProperty: typehelpers.NewSetValueOfMust[types.Bool](ctx, []attr.Value{}),
			},
		},
		{
			name: "single Element",
			input: &TestSetOfBoolModel{
				SetBoolProperty: []bool{
					true,
				},
			},
			expected: &TestSetOfBoolFWModel{
				SetBoolProperty: typehelpers.NewSetValueOfMust[types.Bool](ctx, []attr.Value{
					types.BoolValue(true),
				}),
			},
		},
		{
			name: "multiple elements",
			input: &TestSetOfBoolModel{
				SetBoolProperty: []bool{
					true,
					false,
					true,
				},
			},
			expected: &TestSetOfBoolFWModel{
				SetBoolProperty: typehelpers.NewSetValueOfMust[types.Bool](ctx, []attr.Value{
					types.BoolValue(true),
					types.BoolValue(false),
					types.BoolValue(true),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestSetOfBoolFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_setOfFloat(t *testing.T) {
	ctx := context.Background()
	emptySet := make([]float64, 0)
	cases := []struct {
		name        string
		input       *TestSetOfFloatModel
		expected    *TestSetOfFloatFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestSetOfFloatModel{},
			expected: &TestSetOfFloatFWModel{
				SetFloatProperty: typehelpers.NewSetValueOfNull[types.Float64](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestSetOfFloatModel{
				SetFloatProperty: emptySet,
			},
			expected: &TestSetOfFloatFWModel{
				SetFloatProperty: typehelpers.NewSetValueOfMust[types.Float64](ctx, []attr.Value{}),
			},
		},
		{
			name: "single Element",
			input: &TestSetOfFloatModel{
				SetFloatProperty: []float64{
					floatVal,
				},
			},
			expected: &TestSetOfFloatFWModel{
				SetFloatProperty: typehelpers.NewSetValueOfMust[types.Float64](ctx, []attr.Value{
					types.Float64Value(floatVal),
				}),
			},
		},
		{
			name: "multiple elements",
			input: &TestSetOfFloatModel{
				SetFloatProperty: []float64{
					floatVal + 1,
					floatVal + 2,
					floatVal + 3,
				},
			},
			expected: &TestSetOfFloatFWModel{
				SetFloatProperty: typehelpers.NewSetValueOfMust[types.Float64](ctx, []attr.Value{
					types.Float64Value(floatVal + 1),
					types.Float64Value(floatVal + 2),
					types.Float64Value(floatVal + 3),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestSetOfFloatFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_setOfInt(t *testing.T) {
	ctx := context.Background()
	emptySet := make([]int64, 0)
	cases := []struct {
		name        string
		input       *TestSetOfIntModel
		expected    *TestSetOfIntFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestSetOfIntModel{},
			expected: &TestSetOfIntFWModel{
				SetIntProperty: typehelpers.NewSetValueOfNull[types.Int64](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestSetOfIntModel{
				SetIntProperty: emptySet,
			},
			expected: &TestSetOfIntFWModel{
				SetIntProperty: typehelpers.NewSetValueOfMust[types.Int64](ctx, []attr.Value{}),
			},
		},
		{
			name: "single Element",
			input: &TestSetOfIntModel{
				SetIntProperty: []int64{
					intVal,
				},
			},
			expected: &TestSetOfIntFWModel{
				SetIntProperty: typehelpers.NewSetValueOfMust[types.Int64](ctx, []attr.Value{
					types.Int64Value(intVal),
				}),
			},
		},
		{
			name: "multiple elements",
			input: &TestSetOfIntModel{
				SetIntProperty: []int64{
					intVal + 1,
					intVal + 2,
					intVal + 3,
				},
			},
			expected: &TestSetOfIntFWModel{
				SetIntProperty: typehelpers.NewSetValueOfMust[types.Int64](ctx, []attr.Value{
					types.Int64Value(intVal + 1),
					types.Int64Value(intVal + 2),
					types.Int64Value(intVal + 3),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestSetOfIntFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_setOfString(t *testing.T) {
	ctx := context.Background()
	emptySet := make([]string, 0)
	cases := []struct {
		name        string
		input       *TestSetOfStringModel
		expected    *TestSetOfStringFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestSetOfStringModel{},
			expected: &TestSetOfStringFWModel{
				SetStringProperty: typehelpers.NewSetValueOfNull[types.String](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestSetOfStringModel{
				SetStringProperty: emptySet,
			},
			expected: &TestSetOfStringFWModel{
				SetStringProperty: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{}),
			},
		},
		{
			name: "single Element",
			input: &TestSetOfStringModel{
				SetStringProperty: []string{
					strVal,
				},
			},
			expected: &TestSetOfStringFWModel{
				SetStringProperty: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{
					types.StringValue(strVal),
				}),
			},
		},
		{
			name: "multiple elements",
			input: &TestSetOfStringModel{
				SetStringProperty: []string{
					strVal + "1",
					strVal + "2",
					strVal + "3",
				},
			},
			expected: &TestSetOfStringFWModel{
				SetStringProperty: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{
					types.StringValue(strVal + "1"),
					types.StringValue(strVal + "2"),
					types.StringValue(strVal + "3"),
				}),
			},
		},
		{
			name: "multiple elements",
			input: &TestSetOfStringModel{
				SetStringProperty: []string{
					strVal + "1",
					strVal + "2",
					strVal + "3",
				},
			},
			expected: &TestSetOfStringFWModel{
				SetStringProperty: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{
					types.StringValue(strVal + "1"),
					types.StringValue(strVal + "2"),
					types.StringValue(strVal + "3"),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestSetOfStringFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_mapOfBoolPtr(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		input       *TestMapOfBoolPtrModel
		expected    *TestMapOfBoolFWModel
		expectError bool
	}{
		{
			name:  "omitted property",
			input: &TestMapOfBoolPtrModel{},
			expected: &TestMapOfBoolFWModel{
				MapBoolProperty: typehelpers.NewMapValueOfNull[types.Bool](ctx),
			},
		},
		{
			name: "zero length input",
			input: &TestMapOfBoolPtrModel{
				MapBoolProperty: map[string]*bool{},
			},
			expected: &TestMapOfBoolFWModel{
				MapBoolProperty: typehelpers.NewMapValueOfMust[types.Bool](ctx, map[string]attr.Value{}),
			},
		},
		{
			name: "length 1 input",
			input: &TestMapOfBoolPtrModel{
				MapBoolProperty: map[string]*bool{
					"true": pointer.To(true),
				},
			},
			expected: &TestMapOfBoolFWModel{
				MapBoolProperty: typehelpers.NewMapValueOfMust[types.Bool](ctx, map[string]attr.Value{
					"true": types.BoolValue(true),
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestMapOfBoolFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() && !c.expectError {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) && !c.expectError {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_nestedOneLevel(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name     string
		input    *TestNestedOneLevelModel
		expected *TestNestedOneLevelFWModel
	}{
		{
			name: "empty",
			input: &TestNestedOneLevelModel{
				TopLevelString: strVal,
			},
			expected: &TestNestedOneLevelFWModel{
				TopLevelString: types.StringValue(strVal),
				NestedModel:    typehelpers.NewListNestedObjectValueOfNull[TestStringOnlyFWModel](ctx),
			},
		},
		{
			name: "complete",
			input: &TestNestedOneLevelModel{
				TopLevelString: strVal,
				NestedModel: []TestStringOptionalModel{
					{
						StringProperty: &strVal,
					},
					{
						StringProperty: pointer.To(strVal + strVal),
					},
				},
			},
			expected: &TestNestedOneLevelFWModel{
				TopLevelString: types.StringValue(strVal),
				NestedModel: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []TestStringOnlyFWModel{
					{
						StringProperty: types.StringValue(strVal),
					},
					{
						StringProperty: types.StringValue(strVal + strVal),
					},
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestNestedOneLevelFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_nestedTwoLevel(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name     string
		input    *TestNestedTwoLevelModel
		expected *TestNestedTwoLevelFWModel
	}{
		{
			name:  "empty",
			input: &TestNestedTwoLevelModel{},
			expected: &TestNestedTwoLevelFWModel{
				TopLevelString: types.StringValue(""),
				NestedModel:    typehelpers.NewListNestedObjectValueOfNull[TestNestedOneLevelFWModel](ctx),
			},
		},
		{
			name: "complete",
			input: &TestNestedTwoLevelModel{
				TopLevelString: strVal,
				NestedModel: []TestNestedOneLevelModel{
					{
						TopLevelString: strVal,
					},
					{
						TopLevelString: strVal + strVal,
					},
				},
			},
			expected: &TestNestedTwoLevelFWModel{
				TopLevelString: types.StringValue(strVal),
				NestedModel: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []TestNestedOneLevelFWModel{
					{
						TopLevelString: types.StringValue(strVal),
						NestedModel:    typehelpers.NewListNestedObjectValueOfNull[TestStringOnlyFWModel](ctx),
					},
					{
						TopLevelString: types.StringValue(strVal + strVal),
						NestedModel:    typehelpers.NewListNestedObjectValueOfNull[TestStringOnlyFWModel](ctx),
					},
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestNestedTwoLevelFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}

func TestFlatten_complexNested(t *testing.T) {
	// t.Skip() // TODO - panics
	ctx := context.Background()
	cases := []struct {
		name     string
		input    *TestNestedComplexModel
		expected *TestNestedComplexFWModel
	}{
		// {
		// 	name: "empty",
		// 	input: &TestNestedComplexModel{
		// 		TopLevelString: strVal,
		// 	},
		// 	expected: &TestNestedComplexFWModel{
		// 		TopLevelString:      types.StringValue(strVal),
		// 		TopLevelOptionalInt: types.Int64Null(),
		// 		NestedSimpleModel:   typehelpers.NewListNestedObjectValueOfNull[TestStringOnlyFWModel](ctx),
		// 		NestedComplexModel:  typehelpers.NewListNestedObjectValueOfNull[TestFrameworkModelComplex](ctx),
		// 	},
		// },
		// {
		// 	name: "With Simple Model, no Complex Model",
		// 	input: &TestNestedComplexModel{
		// 		TopLevelString:      strVal,
		// 		TopLevelOptionalInt: &intVal,
		// 		NestedSimpleModel: []TestStringOptionalModel{
		// 			{
		// 				StringProperty: &strVal,
		// 			},
		// 		},
		// 	},
		// 	expected: &TestNestedComplexFWModel{
		// 		TopLevelString:      types.StringValue(strVal),
		// 		TopLevelOptionalInt: types.Int64Value(intVal),
		// 		NestedSimpleModel: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []TestStringOnlyFWModel{
		// 			{
		// 				StringProperty: types.StringValue(strVal),
		// 			},
		// 		}),
		// 		NestedComplexModel: typehelpers.NewListNestedObjectValueOfNull[TestFrameworkModelComplex](ctx),
		// 	},
		// },
		// {
		// 	name: "With Multiple Simple Model, no Complex Model",
		// 	input: &TestNestedComplexModel{
		// 		TopLevelString:      strVal,
		// 		TopLevelOptionalInt: &intVal,
		// 		NestedSimpleModel: []TestStringOptionalModel{
		// 			{
		// 				StringProperty: &strVal,
		// 			},
		// 			{
		// 				StringProperty: &strVal,
		// 			},
		// 		},
		// 	},
		// 	expected: &TestNestedComplexFWModel{
		// 		TopLevelString:      types.StringValue(strVal),
		// 		TopLevelOptionalInt: types.Int64Value(intVal),
		// 		NestedSimpleModel: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []TestStringOnlyFWModel{
		// 			{
		// 				StringProperty: types.StringValue(strVal),
		// 			},
		// 			{
		// 				StringProperty: types.StringValue(strVal),
		// 			},
		// 		}),
		// 		NestedComplexModel: typehelpers.NewListNestedObjectValueOfNull[TestFrameworkModelComplex](ctx),
		// 	},
		// },
		{
			name: "Complete",
			input: &TestNestedComplexModel{
				TopLevelString:      strVal,
				TopLevelOptionalInt: &intVal,
				NestedSimpleModel: []TestStringOptionalModel{
					{
						StringProperty: &strVal,
					},
				},
				NestedComplexModel: []TestAPIModelComplex{
					{
						BoolProperty:   true,
						StringProperty: strVal,
						IntProperty:    intVal,
						FloatProperty:  floatVal,
						ListOfPrimitives: []string{
							strVal,
						},
						ListProperty: []TestAPINestedModel{
							{
								SubPropertyBool:   true,
								SubPropertyFloat:  floatVal + 1,
								SubPropertyInt:    intVal + 1,
								SubPropertyString: strings.ToUpper(strVal),
							},
							{
								SubPropertyBool:   false,
								SubPropertyFloat:  floatVal + 2,
								SubPropertyInt:    intVal + 2,
								SubPropertyString: strings.ToLower(strVal),
							},
						},
						SetProperty: []TestAPINestedModel{
							{
								SubPropertyBool:   false,
								SubPropertyFloat:  floatVal + 3,
								SubPropertyInt:    intVal + 3,
								SubPropertyString: strings.ToUpper(strVal),
							},
							{
								SubPropertyBool:   true,
								SubPropertyFloat:  floatVal + 4,
								SubPropertyInt:    intVal + 4,
								SubPropertyString: strings.ToLower(strVal),
							},
							{
								SubPropertyBool:   true,
								SubPropertyFloat:  floatVal + 5,
								SubPropertyInt:    intVal + 5,
								SubPropertyString: strVal,
							},
						},
						MapBoolProperty: map[string]bool{
							"true":  true,
							"false": false,
						},
						MapFloatProperty: map[string]float64{
							"floatVal": floatVal,
							"plusOne":  floatVal + 1,
						},
						MapInt64Property: map[string]int64{
							"intVal":    intVal,
							"plusTwo":   intVal + 2,
							"plusThree": intVal + 3,
						},
						MapStringProperty: map[string]string{
							"strVal": strVal,
							"concat": strVal + strVal,
						},
					},
				},
			},
			expected: &TestNestedComplexFWModel{
				TopLevelString:      types.StringValue(strVal),
				TopLevelOptionalInt: types.Int64Value(intVal),
				NestedSimpleModel: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []TestStringOnlyFWModel{
					{
						StringProperty: types.StringValue(strVal),
					},
				}),
				NestedComplexModel: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []TestFrameworkModelComplex{
					{
						BoolProperty:   types.BoolValue(true),
						StringProperty: types.StringValue(strVal),
						IntProperty:    types.Int64Value(intVal),
						FloatProperty:  types.Float64Value(floatVal),
						ListOfPrimitives: typehelpers.NewListValueOfMust[types.String](ctx, []attr.Value{
							types.StringValue(strVal),
						}),
						ListProperty: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []TestFrameworkNestedModel{
							{
								SubPropertyBool:   types.BoolValue(true),
								SubPropertyFloat:  types.Float64Value(floatVal + 1),
								SubPropertyInt:    types.Int64Value(intVal + 1),
								SubPropertyString: types.StringValue(strings.ToUpper(strVal)),
							},
							{
								SubPropertyBool:   types.BoolValue(false),
								SubPropertyFloat:  types.Float64Value(floatVal + 2),
								SubPropertyInt:    types.Int64Value(intVal + 2),
								SubPropertyString: types.StringValue(strings.ToLower(strVal)),
							},
						}),
						SetProperty: typehelpers.NewSetNestedObjectValueOfValueSliceMust(ctx, []TestFrameworkNestedModel{
							{
								SubPropertyBool:   types.BoolValue(false),
								SubPropertyFloat:  types.Float64Value(floatVal + 3),
								SubPropertyInt:    types.Int64Value(intVal + 3),
								SubPropertyString: types.StringValue(strings.ToUpper(strVal)),
							},
							{
								SubPropertyBool:   types.BoolValue(true),
								SubPropertyFloat:  types.Float64Value(floatVal + 4),
								SubPropertyInt:    types.Int64Value(intVal + 4),
								SubPropertyString: types.StringValue(strings.ToLower(strVal)),
							},
							{
								SubPropertyBool:   types.BoolValue(true),
								SubPropertyFloat:  types.Float64Value(floatVal + 5),
								SubPropertyInt:    types.Int64Value(intVal + 5),
								SubPropertyString: types.StringValue(strVal),
							},
						}),
						MapBoolProperty: typehelpers.NewMapValueOfMust[types.Bool](ctx, map[string]attr.Value{
							"true":  types.BoolValue(true),
							"false": types.BoolValue(false),
						}),
						MapFloatProperty: typehelpers.NewMapValueOfMust[types.Float64](ctx, map[string]attr.Value{
							"floatVal": types.Float64Value(floatVal),
							"plusOne":  types.Float64Value(floatVal + 1),
						}),
						MapInt64Property: typehelpers.NewMapValueOfMust[types.Int64](ctx, map[string]attr.Value{
							"intVal":    types.Int64Value(intVal),
							"plusTwo":   types.Int64Value(intVal + 2),
							"plusThree": types.Int64Value(intVal + 3),
						}),
						MapStringProperty: typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{
							"strVal": types.StringValue(strVal),
							"concat": types.StringValue(strVal + strVal),
						}),
					},
				}),
			},
		},
	}

	for _, c := range cases {
		diags := &diag.Diagnostics{}
		result := &TestNestedComplexFWModel{}

		convert.Flatten(context.Background(), c.input, result, diags)

		if diags.HasError() {
			t.Errorf("Test: %s \ndiags: %+v", c.name, diags)
		}

		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("\nTest: %s\nexpected:\n %+v\ngot:\n %+v", c.name, c.expected, result)
		}

	}
}
