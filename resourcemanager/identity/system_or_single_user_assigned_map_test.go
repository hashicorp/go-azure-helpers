// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identity

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestSystemOrSingleUserAssignedMapMarshal(t *testing.T) {
	testData := []struct {
		input                           *SystemOrSingleUserAssignedMap
		expectedIdentityType            string
		expectedUserAssignedIdentityIds []string
	}{
		{
			input:                           nil,
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input:                           &SystemOrSingleUserAssignedMap{},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &SystemOrSingleUserAssignedMap{
				Type: TypeNone,
			},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &SystemOrSingleUserAssignedMap{
				Type: TypeNone,
				IdentityIds: map[string]UserAssignedIdentityDetails{
					"first": {},
				},
			},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{
				// intentionally empty since this is bad data
			},
		},
		{
			input: &SystemOrSingleUserAssignedMap{
				Type:        TypeSystemAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{},
			},
			expectedIdentityType:            "SystemAssigned",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &SystemOrSingleUserAssignedMap{
				Type:        TypeSystemAssignedUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{},
			},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &SystemOrSingleUserAssignedMap{
				Type:        TypeUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{},
			},
			expectedIdentityType:            "UserAssigned",
			expectedUserAssignedIdentityIds: []string{},
		},

		{
			input: &SystemOrSingleUserAssignedMap{
				Type: TypeSystemAssignedUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{
					"first":  {},
					"second": {},
				},
			},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{
				// bad data
			},
		},
		{
			input: &SystemOrSingleUserAssignedMap{
				Type: TypeUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{
					"first":  {},
					"second": {},
				},
			},
			expectedIdentityType: "UserAssigned",
			expectedUserAssignedIdentityIds: []string{
				"first",
				"second",
			},
		},
	}
	for i, v := range testData {
		t.Logf("step %d..", i)

		encoded, err := v.input.MarshalJSON()
		if err != nil {
			t.Fatalf("marshaling: %+v", err)
		}

		var out map[string]interface{}
		if err := json.Unmarshal(encoded, &out); err != nil {
			t.Fatalf("decoding: %+v", err)
		}

		actualIdentityValue := out["type"].(string)
		if v.expectedIdentityType != actualIdentityValue {
			t.Fatalf("expected %q but got %q", v.expectedIdentityType, actualIdentityValue)
		}

		actualUserAssignedIdentityIdsRaw, ok := out["userAssignedIdentities"].(map[string]interface{})
		if !ok {
			if len(v.expectedUserAssignedIdentityIds) == 0 {
				continue
			}

			t.Fatalf("`userAssignedIdentities` was nil")
		}
		actualUserAssignedIdentityIds := make([]string, 0)
		for k := range actualUserAssignedIdentityIdsRaw {
			actualUserAssignedIdentityIds = append(actualUserAssignedIdentityIds, k)
		}
		sort.Strings(v.expectedUserAssignedIdentityIds)
		sort.Strings(actualUserAssignedIdentityIds)

		if !reflect.DeepEqual(v.expectedUserAssignedIdentityIds, actualUserAssignedIdentityIds) {
			t.Fatalf("expected %q but got %q", strings.Join(v.expectedUserAssignedIdentityIds, ", "), strings.Join(actualUserAssignedIdentityIds, ", "))
		}
	}
}
