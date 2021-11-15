package identity

import (
	"encoding/json"
	"testing"
)

func TestSystemAssignedMarshal(t *testing.T) {
	testData := []struct{
		input         *SystemAssigned
		expectedValue string
	}{
		{
			input: nil,
			expectedValue: "None",
		},
		{
			input: &SystemAssigned{},
			expectedValue: "None",
		},
		{
			input: &SystemAssigned{
				Type: TypeNone,
			},
			expectedValue: "None",
		},
		{
			input: &SystemAssigned{
				Type: TypeSystemAssignedUserAssigned,
			},
			expectedValue: "None",
		},
		{
			input: &SystemAssigned{
				Type: TypeUserAssigned,
			},
			expectedValue: "None",
		},
		{
			input: &SystemAssigned{
				Type: TypeSystemAssigned,
			},
			expectedValue: "SystemAssigned",
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

		actualValue := out["type"].(string)
		if v.expectedValue != actualValue {
			t.Fatalf("expected %q but got %q", v.expectedValue, actualValue)
		}
	}
}
