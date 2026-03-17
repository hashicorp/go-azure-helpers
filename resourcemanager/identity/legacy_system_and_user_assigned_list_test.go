// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package identity

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestLegacySystemUserAssignedListMarshal(t *testing.T) {
	testData := []struct {
		input                           any
		expect                          map[string]any
		expectedIdentityType            string
		expectedUserAssignedIdentityIds []string
		expectError                     bool
	}{
		{
			input: &LegacySystemAndUserAssignedList{},
			expect: map[string]any{
				"type":                   "None",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &LegacySystemAndUserAssignedList{
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
			input: &LegacySystemAndUserAssignedList{
				Type: TypeNone,
				IdentityIds: []string{
					"first",
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
			input: &LegacySystemAndUserAssignedList{
				Type:        TypeSystemAssigned,
				IdentityIds: []string{},
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
			input: LegacySystemAndUserAssignedList{
				Type:        TypeSystemAssigned,
				IdentityIds: []string{},
			},
			expect: map[string]any{
				"type":                   "SystemAssigned",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "SystemAssigned",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &LegacySystemAndUserAssignedList{
				Type:        TypeSystemAssignedUserAssigned,
				IdentityIds: []string{},
			},
			expect: map[string]any{
				"type":                   "SystemAssigned,UserAssigned",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "SystemAssigned,UserAssigned",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &LegacySystemAndUserAssignedList{
				Type:        typeLegacySystemAssignedUserAssigned,
				IdentityIds: []string{},
			},
			expectError: true,
		},
		{
			input: &LegacySystemAndUserAssignedList{
				Type:        TypeUserAssigned,
				IdentityIds: []string{},
			},
			expect: map[string]any{
				"type":                   "UserAssigned",
				"userAssignedIdentities": nil,
			},
			expectedIdentityType:            "UserAssigned",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &LegacySystemAndUserAssignedList{
				Type: TypeSystemAssignedUserAssigned,
				IdentityIds: []string{
					"first",
					"second",
				},
			},
			expect: map[string]any{
				"type": "SystemAssigned,UserAssigned",
				"userAssignedIdentities": []any{
					"first",
					"second",
				},
			},
			expectedIdentityType: "SystemAssigned,UserAssigned",
			expectedUserAssignedIdentityIds: []string{
				"first",
				"second",
			},
		},
		{
			input: &LegacySystemAndUserAssignedList{
				Type: typeLegacySystemAssignedUserAssigned,
				IdentityIds: []string{
					"first",
					"second",
				},
			},
			expectError: true,
		},
		{
			input: &LegacySystemAndUserAssignedList{
				Type: TypeUserAssigned,
				IdentityIds: []string{
					"first",
					"second",
				},
			},
			expect: map[string]any{
				"type": "UserAssigned",
				"userAssignedIdentities": []any{
					"first",
					"second",
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
			if v.expectError {
				continue
			}

			t.Fatalf("encoding: %+v", err)
		}
		if v.expectError {
			t.Fatalf("expected an error but didn't get one")
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

		actualUserAssignedIdentityIdsRaw, ok := out["userAssignedIdentities"].([]interface{})
		if !ok {
			if len(v.expectedUserAssignedIdentityIds) == 0 {
				continue
			}

			t.Fatalf("`userAssignedIdentities` was nil")
		}
		actualUserAssignedIdentityIds := make([]string, 0)
		for _, v := range actualUserAssignedIdentityIdsRaw {
			actualUserAssignedIdentityIds = append(actualUserAssignedIdentityIds, v.(string))
		}
		if !reflect.DeepEqual(v.expectedUserAssignedIdentityIds, actualUserAssignedIdentityIds) {
			t.Fatalf("expected %q but got %q", strings.Join(v.expectedUserAssignedIdentityIds, ", "), strings.Join(actualUserAssignedIdentityIds, ", "))
		}
	}
}
