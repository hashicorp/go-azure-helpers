// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pointer_test

import (
	"reflect"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
)

func TestFrom_ListOfStrings(t *testing.T) {
	testCases := []struct {
		Input    *[]string
		Expected []string
	}{
		{
			Input:    &[]string{},
			Expected: []string{},
		},
		{
			Input:    &[]string{"1", "2"},
			Expected: []string{"1", "2"},
		},
	}

	for _, v := range testCases {
		actual := pointer.From(v.Input)
		if !reflect.DeepEqual(actual, v.Expected) {
			t.Fatalf("expectd %#v, got %#v", v.Expected, actual)
		}
	}
}

func TestFrom_NilTypes(t *testing.T) {
	var stringInput *string
	var intInput *int64
	var floatInput *float64
	type customType struct{}
	if actual := pointer.From(stringInput); actual != "" {
		t.Fatal("stringInput")
	}
	if actual := pointer.From(intInput); actual != 0 {
		t.Fatal("intInput")
	}
	if actual := pointer.From(floatInput); actual != 0 {
		t.Fatal("floatInput")
	}
	var custom *customType
	customExpected := customType{}
	if actual := pointer.From(custom); actual != customExpected {
		t.Fatal("customType")
	}
	var complex *map[string]customType
	complexExpected := map[string]customType{}
	if actual := pointer.From(complex); reflect.DeepEqual(actual, complexExpected) {
		t.Fatal("complexType")
	}
}

func TestFrom_MapOfInterface(t *testing.T) {
	testCases := []struct {
		Input    *map[string]interface{}
		Expected map[string]interface{}
	}{
		{
			Input:    &map[string]interface{}{},
			Expected: map[string]interface{}{},
		},
		{
			Input: &map[string]interface{}{
				"foo": "bar",
			},
			Expected: map[string]interface{}{
				"foo": "bar",
			},
		},
	}

	for _, v := range testCases {
		actual := pointer.From(v.Input)
		if !reflect.DeepEqual(actual, v.Expected) {
			t.Fatalf("expectd %#v, got %#v", v.Expected, actual)
		}
	}
}

// TODO - More comprehensive type coverage?
