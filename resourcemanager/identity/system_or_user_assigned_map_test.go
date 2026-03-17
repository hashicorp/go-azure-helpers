// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package identity

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestSystemOrUserAssignedMapMarshal(t *testing.T) {
	testData := []struct {
		input                           any
		expect                          map[string]any
		expectedIdentityType            string
		expectedUserAssignedIdentityIds []string
	}{
		{
			input: &SystemOrUserAssignedMap{},
			expect: map[string]any{
				"type":                   "None",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &SystemOrUserAssignedMap{
				Type: TypeNone,
			},
			expect: map[string]any{
				"type":                   "None",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &SystemOrUserAssignedMap{
				Type: TypeNone,
				IdentityIds: map[string]UserAssignedIdentityDetails{
					"first": {},
				},
			},
			expect: map[string]any{
				"type":                   "None",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{
				// intentionally empty since this is bad data
			},
		},
		{
			input: &SystemOrUserAssignedMap{
				Type:        TypeSystemAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{},
			},
			expect: map[string]any{
				"type":                   "SystemAssigned",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "SystemAssigned",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			// Value type (instead of pointer type)
			input: SystemOrUserAssignedMap{
				Type:        TypeSystemAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{},
			},
			expect: map[string]any{
				"type":                   "SystemAssigned",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "SystemAssigned",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &SystemOrUserAssignedMap{
				Type:        TypeSystemAssignedUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{},
			},
			expect: map[string]any{
				"type":                   "None",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &SystemOrUserAssignedMap{
				Type:        TypeUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{},
			},
			expect: map[string]any{
				"type":                   "UserAssigned",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "UserAssigned",
			expectedUserAssignedIdentityIds: []string{},
		},

		{
			input: &SystemOrUserAssignedMap{
				Type: TypeSystemAssignedUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{
					"first":  {},
					"second": {},
				},
			},
			expect: map[string]any{
				"type":                   "None",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{
				// bad data
			},
		},
		{
			input: &SystemOrUserAssignedMap{
				Type: TypeUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{
					"first":  {},
					"second": {},
				},
			},
			expect: map[string]any{
				"type": "UserAssigned",
				"userAssignedIdentities": map[string]any{
					"first":  map[string]any{},
					"second": map[string]any{},
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

		encoded, err := json.Marshal(v.input)
		if err != nil {
			t.Fatalf("marshaling: %+v", err)
		}

		expectEncoded, _ := json.Marshal(v.expect)
		if string(encoded) != string(expectEncoded) {
			t.Fatalf("marshaled JSON is not as expected. got=%v, expect=%v", string(encoded), string(expectEncoded))
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
