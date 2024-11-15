package convert_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/go-azure-helpers/framework/convert"
	"github.com/hashicorp/go-azure-helpers/framework/typehelpers"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type TestFrameworkModel struct {
	BoolProperty      types.Bool                                                    `tfsdk:"bool_property"`
	StringProperty    types.String                                                  `tfsdk:"string_property"`
	IntProperty       types.Int64                                                   `tfsdk:"int_property"`
	ListProperty      typehelpers.ListNestedObjectValueOf[TestFrameworkNestedModel] `tfsdk:"list_property"`
	FloatProperty     types.Float64                                                 `tfsdk:"float_property"`
	SetProperty       typehelpers.SetNestedObjectValueOf[TestFrameworkNestedModel]  `tfsdk:"set_property"` // TODO Sets
	MapStringProperty typehelpers.MapValueOf[basetypes.StringValue]                 `tfsdk:"map_string_property"`
	// MapObjectProperty   types.Map     `tfsdk:"map_property"` // TODO
}

type TestAPIModel struct {
	BoolProperty      bool
	StringProperty    string
	IntProperty       int64
	FloatProperty     float64
	ListProperty      []TestAPINestedModel
	SetProperty       []TestAPINestedModel
	MapStringProperty map[string]string
	//MapObjectProperty map[string]TestAPINestedModel
}

type TestFrameworkNestedModel struct {
	SubPropertyBool   types.Bool    `tfsdk:"sub_property_bool"`
	SubPropertyFloat  types.Float64 `tfsdk:"sub_property_float"`
	SubPropertyInt    types.Int64   `tfsdk:"sub_property_int"`
	SubPropertyString types.String  `tfsdk:"sub_property_string"`
}

type TestAPINestedModel struct {
	SubPropertyBool   bool
	SubPropertyFloat  float64
	SubPropertyInt    int64
	SubPropertyString string
}

func TestExpand(t *testing.T) {
	input := &TestFrameworkModel{
		BoolProperty:   types.BoolValue(true),
		StringProperty: types.StringValue("foo"),
	}

	expectedOutput := &TestAPIModel{
		BoolProperty:   true,
		StringProperty: "foo",
	}

	result := &TestAPIModel{}

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
		input       *TestFrameworkModel
		expected    *TestAPIModel
		expectError bool
	}{
		{
			name:        "empty",
			input:       &TestFrameworkModel{},
			expected:    &TestAPIModel{},
			expectError: false,
		},
		{
			name: "single bool",
			input: &TestFrameworkModel{
				BoolProperty: types.BoolValue(true),
			},
			expected: &TestAPIModel{
				BoolProperty: true,
			},
			expectError: false,
		},
		{
			name: "single string",
			input: &TestFrameworkModel{
				StringProperty: types.StringValue("foo"),
			},
			expected: &TestAPIModel{
				StringProperty: "foo",
			},
			expectError: false,
		},
		{
			name: "multiple list entries",
			input: &TestFrameworkModel{
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
			expected: &TestAPIModel{
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
			input: &TestFrameworkModel{
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
			expected: &TestAPIModel{
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

		result := &TestAPIModel{}

		diags := &diag.Diagnostics{}

		convert.Expand(context.Background(), testCase.input, result, diags)
		if diags.HasError() != testCase.expectError {
			t.Fatalf("Expand failed for %s: %+v", testCase.name, diags.Errors())
		}

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Fatalf("Expand failed for case %s, expected: %+v \n Got: %+v", testCase.name, testCase.expected, result)
		}
	}
}
