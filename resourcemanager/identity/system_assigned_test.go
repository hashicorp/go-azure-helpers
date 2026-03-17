// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package identity

import (
	"encoding/json"
	"testing"
)

func TestSystemAssignedMarshal(t *testing.T) {
	testData := []struct {
		input         any
		expect        map[string]any
		expectedValue string
	}{
		{
			input: &SystemAssigned{},
			expect: map[string]any{
				"type": "None",
			},
			expectedValue: "None",
		},
		{
			input: &SystemAssigned{
				Type: TypeNone,
			},
			expect: map[string]any{
				"type": "None",
			},
			expectedValue: "None",
		},
		{
			input: &SystemAssigned{
				Type: TypeSystemAssignedUserAssigned,
			},
			expect: map[string]any{
				"type": "None",
			},
			expectedValue: "None",
		},
		{
			input: &SystemAssigned{
				Type: TypeUserAssigned,
			},
			expect: map[string]any{
				"type": "None",
			},
			expectedValue: "None",
		},
		{
			input: &SystemAssigned{
				Type: TypeSystemAssigned,
			},
			expect: map[string]any{
				"type": "SystemAssigned",
			},
			expectedValue: "SystemAssigned",
		},
		{
			// Value type (instead of pointer type)
			input: SystemAssigned{
				Type: TypeSystemAssigned,
			},
			expect: map[string]any{
				"type": "SystemAssigned",
			},
			expectedValue: "SystemAssigned",
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

		actualValue := out["type"].(string)
		if v.expectedValue != actualValue {
			t.Fatalf("expected %q but got %q", v.expectedValue, actualValue)
		}
	}
}
