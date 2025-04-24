// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package convert_test

import (
	"github.com/hashicorp/go-azure-helpers/framework/typehelpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type TestFrameworkModelComplex struct {
	BoolProperty      types.Bool                                                    `tfsdk:"bool_property"`
	FloatProperty     types.Float64                                                 `tfsdk:"float_property"`
	IntProperty       types.Int64                                                   `tfsdk:"int_property"`
	StringProperty    types.String                                                  `tfsdk:"string_property"`
	ListProperty      typehelpers.ListNestedObjectValueOf[TestFrameworkNestedModel] `tfsdk:"list_property"`
	ListOfPrimitives  typehelpers.ListValueOf[types.String]                         `tfsdk:"list_of_primitives"`
	MapBoolProperty   typehelpers.MapValueOf[basetypes.BoolValue]                   `tfsdk:"map_bool_property"`
	MapFloatProperty  typehelpers.MapValueOf[basetypes.Float64Value]                `tfsdk:"map_float_property"`
	MapInt64Property  typehelpers.MapValueOf[basetypes.Int64Value]                  `tfsdk:"map_int64_property"`
	MapStringProperty typehelpers.MapValueOf[basetypes.StringValue]                 `tfsdk:"map_string_property"`
	SetProperty       typehelpers.SetNestedObjectValueOf[TestFrameworkNestedModel]  `tfsdk:"set_property"`
}

type TestAPIModelComplex struct {
	BoolProperty      bool
	FloatProperty     float64
	IntProperty       int64
	StringProperty    string
	ListOfPrimitives  []string
	ListProperty      []TestAPINestedModel
	MapBoolProperty   map[string]bool
	MapFloatProperty  map[string]float64
	MapInt64Property  map[string]int64
	MapStringProperty map[string]string
	SetProperty       []TestAPINestedModel
	// MapObjectProperty map[string]TestAPINestedModel // v6 protocol only - Future requirement?
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

type TestBoolOnlyFWModel struct {
	BoolProperty types.Bool `tfsdk:"bool_property"`
}

type TestBoolRequiredModel struct {
	BoolProperty bool
}

type TestBoolOptionalModel struct {
	BoolProperty *bool
}

type TestStringOnlyFWModel struct {
	StringProperty types.String `tfsdk:"string_property"`
}

type TestStringRequiredModel struct {
	StringProperty string
}
type TestStringOptionalModel struct {
	StringProperty *string
}
type TestInt64OnlyFWModel struct {
	Int64Property types.Int64 `tfsdk:"int64_property"`
}

type TestInt64RequiredModel struct {
	Int64Property int64
}
type TestInt64OptionalModel struct {
	Int64Property *int64
}

type TestFloat64OnlyFWModel struct {
	Float64Property types.Float64 `tfsdk:"float64_property"`
}

type TestFloat64RequiredModel struct {
	Float64Property float64
}
type TestFloat64OptionalModel struct {
	Float64Property *float64
}

type TestMapOfStringFWModel struct {
	MapStringProperty typehelpers.MapValueOf[basetypes.StringValue] `tfsdk:"map_float_property"`
}
type TestMapOfStringModel struct {
	MapStringProperty map[string]string
}

type TestMapOfStringPtrModel struct {
	MapStringProperty map[string]*string
}

type TestMapOfFloatFWModel struct {
	MapFloatProperty typehelpers.MapValueOf[basetypes.Float64Value] `tfsdk:"map_float_property"`
}
type TestMapOfFloatModel struct {
	MapFloatProperty map[string]float64
}

type TestMapOfFloatPtrModel struct {
	MapFloatProperty map[string]*float64
}

type TestMapOfInt64FWModel struct {
	MapInt64Property typehelpers.MapValueOf[basetypes.Int64Value] `tfsdk:"map_float_property"`
}
type TestMapOfInt64Model struct {
	MapInt64Property map[string]int64 `tfsdk:"map_float_property"`
}

type TestMapOfInt64PtrModel struct {
	MapInt64Property map[string]*int64 `tfsdk:"map_float_property"`
}

type TestMapOfBoolFWModel struct {
	MapBoolProperty typehelpers.MapValueOf[basetypes.BoolValue] `tfsdk:"map_float_property"`
}
type TestMapOfBoolModel struct {
	MapBoolProperty map[string]bool `tfsdk:"map_float_property"`
}

type TestMapOfBoolPtrModel struct {
	MapBoolProperty map[string]*bool `tfsdk:"map_float_property"`
}

// Lists and Sets
type TestListOfStringFWModel struct {
	ListStringProperty typehelpers.ListValueOf[types.String] `tfsdk:"list_string_property"`
}

type TestSetOfStringFWModel struct {
	SetStringProperty typehelpers.SetValueOf[types.String] `tfsdk:"set_string_property"`
}

type TestListOfStringModel struct {
	ListStringProperty []string
}

type TestSetOfStringModel struct {
	SetStringProperty []string
}

type TestListOfBoolFWModel struct {
	ListBoolProperty typehelpers.ListValueOf[types.Bool] `tfsdk:"list_bool_property"`
}

type TestListOfBoolModel struct {
	ListBoolProperty []bool
}

type TestSetOfBoolFWModel struct {
	SetBoolProperty typehelpers.SetValueOf[types.Bool] `tfsdk:"set_bool_property"`
}

type TestSetOfBoolModel struct {
	SetBoolProperty []bool
}
type TestListOfFloatFWModel struct {
	ListFloatProperty typehelpers.ListValueOf[types.Float64] `tfsdk:"list_float_property"`
}

type TestListOfFloatModel struct {
	ListFloatProperty []float64
}

type TestSetOfFloatFWModel struct {
	SetFloatProperty typehelpers.SetValueOf[types.Float64] `tfsdk:"set_float_property"`
}

type TestSetOfFloatModel struct {
	SetFloatProperty []float64
}

type TestListOfIntFWModel struct {
	ListIntProperty typehelpers.ListValueOf[types.Int64] `tfsdk:"list_int_property"`
}

type TestListOfIntModel struct {
	ListIntProperty []int64
}

type TestSetOfIntFWModel struct {
	SetIntProperty typehelpers.SetValueOf[types.Int64] `tfsdk:"set_int_property"`
}

type TestSetOfIntModel struct {
	SetIntProperty []int64
}

// Complex models
type TestNestedOneLevelFWModel struct {
	TopLevelString types.String                                               `tfsdk:"top_level_string"`
	NestedModel    typehelpers.ListNestedObjectValueOf[TestStringOnlyFWModel] `tfsdk:"nested_model"`
}

type TestNestedOneLevelModel struct {
	TopLevelString string
	NestedModel    []TestStringOptionalModel
}

type TestNestedTwoLevelFWModel struct {
	TopLevelString types.String
	NestedModel    typehelpers.ListNestedObjectValueOf[TestNestedOneLevelFWModel]
}

type TestNestedTwoLevelModel struct {
	TopLevelString string
	NestedModel    []TestNestedOneLevelModel
}

type TestNestedComplexFWModel struct {
	TopLevelString      types.String                                                   `tfsdk:"top_level_string"`
	TopLevelOptionalInt types.Int64                                                    `tfsdk:"top_level_optional_int"`
	NestedSimpleModel   typehelpers.ListNestedObjectValueOf[TestStringOnlyFWModel]     `tfsdk:"nested_simple_model"`
	NestedComplexModel  typehelpers.ListNestedObjectValueOf[TestFrameworkModelComplex] `tfsdk:"nested_complex_model"`
}

type TestNestedComplexModel struct {
	TopLevelString      string
	TopLevelOptionalInt *int64
	NestedSimpleModel   []TestStringOptionalModel
	NestedComplexModel  []TestAPIModelComplex
}
