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
