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

func TestLegacySystemUserAssignedMapMarshal(t *testing.T) {
	testData := []struct {
		input                           *LegacySystemAndUserAssignedMap
		expectedIdentityType            string
		expectedUserAssignedIdentityIds []string
		expectError                     bool
	}{
		{
			input:                           nil,
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input:                           &LegacySystemAndUserAssignedMap{},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &LegacySystemAndUserAssignedMap{
				Type: TypeNone,
			},
			expectedIdentityType:            "None",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &LegacySystemAndUserAssignedMap{
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
			input: &LegacySystemAndUserAssignedMap{
				Type:        TypeSystemAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{},
			},
			expectedIdentityType:            "SystemAssigned",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &LegacySystemAndUserAssignedMap{
				Type:        TypeSystemAssignedUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{},
			},
			expectedIdentityType:            "SystemAssigned,UserAssigned",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &LegacySystemAndUserAssignedMap{
				Type:        typeLegacySystemAssignedUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{},
			},
			expectError: true,
		},
		{
			input: &LegacySystemAndUserAssignedMap{
				Type:        TypeUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{},
			},
			expectedIdentityType:            "UserAssigned",
			expectedUserAssignedIdentityIds: []string{},
		},
		{
			input: &LegacySystemAndUserAssignedMap{
				Type: TypeSystemAssignedUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{
					"first":  {},
					"second": {},
				},
			},
			expectedIdentityType: "SystemAssigned,UserAssigned",
			expectedUserAssignedIdentityIds: []string{
				"first",
				"second",
			},
		},
		{
			input: &LegacySystemAndUserAssignedMap{
				Type: typeLegacySystemAssignedUserAssigned,
				IdentityIds: map[string]UserAssignedIdentityDetails{
					"first":  {},
					"second": {},
				},
			},
			expectError: true,
		},
		{
			input: &LegacySystemAndUserAssignedMap{
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
			if v.expectError {
				continue
			}

			t.Fatalf("encoding: %+v", err)
		}
		if v.expectError {
			t.Fatalf("expected an error but didn't get one")
		}

		var out map[string]interface{}
		err = json.Unmarshal(encoded, &out)
		if err != nil {
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

func TestLegacySystemAndUserAssignedMapUnmarshal(t *testing.T) {
	testData := []struct {
		input                string
		expectedIdentityType string
		expectError          bool
	}{
		{
			input: `{
			  "userAssignedIdentities":{
				 "/subscriptions/xxx/resourcegroups/testRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUAI1":{
					"principalId":"00000000-00000-0000-0000-000000000000",
					"clientId":"00000000-00000-0000-0000-000000000000"
				 }
			  },
			  "principalId":"00000000-00000-0000-0000-000000000000",
			  "type":"SystemAssigned, UserAssigned",
			  "tenantId":"00000000-00000-0000-0000-000000000000"
			}`,
			expectedIdentityType: "SystemAssigned, UserAssigned",
		},
		{
			input: `{
			  "userAssignedIdentities":{
				 "/subscriptions/xxx/resourcegroups/testRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUAI1":{
					"principalId":"00000000-00000-0000-0000-000000000000",
					"clientId":"00000000-00000-0000-0000-000000000000"
				 }
			  },
			  "principalId":"00000000-00000-0000-0000-000000000000",
			  "type":"SystemAssigned,UserAssigned",
			  "tenantId":"00000000-00000-0000-0000-000000000000"
			}`,
			expectedIdentityType: "SystemAssigned, UserAssigned",
		},
		{
			input: `{
			  "principalId":"00000000-00000-0000-0000-000000000000",
			  "type":"SystemAssigned",
			  "tenantId":"00000000-00000-0000-0000-000000000000"
			}`,
			expectedIdentityType: "SystemAssigned",
		},
		{
			input: `{
			  "userAssignedIdentities":{
				 "/subscriptions/xxx/resourcegroups/testRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUAI1":{
					"principalId":"00000000-00000-0000-0000-000000000000",
					"clientId":"00000000-00000-0000-0000-000000000000"
				 }
			  },
			  "type":"UserAssigned"
			}`,
			expectedIdentityType: "UserAssigned",
		},
		{
			input:                `{"type":"None"}`,
			expectedIdentityType: "None",
		},
		{
			input:                `{"type":"unknown"}`,
			expectedIdentityType: "None",
		},
		{
			// input is invalid JSON
			input:       `{"type": UserAssigned}`,
			expectError: true,
		},
	}

	var s LegacySystemAndUserAssignedMap
	err := s.UnmarshalJSON(nil)
	if err != nil {
		t.Errorf("expected nil error, got: %v", err)
	}

	for i, v := range testData {
		t.Logf("step %d..", i)

		err = s.UnmarshalJSON([]byte(v.input))
		if err != nil {
			if v.expectError {
				continue
			}

			t.Errorf("expected nil error, got: %v", err)
		}

		if string(s.Type) != v.expectedIdentityType {
			t.Errorf("expected type to be %s, got: %s", v.expectedIdentityType, string(s.Type))
		}
	}
}
