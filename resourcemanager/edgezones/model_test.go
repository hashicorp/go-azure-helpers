// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package edgezones

import (
	"encoding/json"
	"testing"
)

func TestMarshalModel(t *testing.T) {
	testData := []struct {
		Input        *Model
		ExpectedName *string
	}{
		{
			// empty name would be bad data, so this can be cleared?
			Input: &Model{
				Name: "",
			},
			ExpectedName: nil,
		},
		{
			Input: &Model{
				Name: "Locutus",
			},
			ExpectedName: ptr("Locutus"),
		},
	}
	for i, v := range testData {
		t.Logf("item %d", i)

		out, err := json.Marshal(v.Input)
		if err != nil {
			t.Fatalf("marshaling: %+v", err)
		}

		var decoded map[string]interface{}
		if err := json.Unmarshal(out, &decoded); err != nil {
			t.Fatalf("unmarshaling test output: %+v", err)
		}

		name, ok := decoded["name"].(string)
		if !ok {
			if v.ExpectedName == nil {
				continue
			}

			t.Fatalf("expected the encoded name to be %q but it was nil", *v.ExpectedName)
		}

		if v.ExpectedName == nil {
			t.Fatalf("expected there to be no encoded name but got %q", name)
		}

		if *v.ExpectedName != name {
			t.Fatalf("expected the encoded name to be %q but got %q", *v.ExpectedName, name)
		}
	}
}

func TestUnmarshalModel(t *testing.T) {
	testData := []struct {
		ExpectedName *string
		Payload      string
	}{
		{
			// Invalid
			ExpectedName: nil,
			Payload:      `{}`,
		},
		{
			// Invalid
			ExpectedName: nil,
			Payload:      `{"name": "Locutus"}`,
		},
		{
			// Invalid
			ExpectedName: nil,
			Payload:      `{"name": "Locutus", "type": "Borg"}`,
		},
		{
			// Valid
			ExpectedName: ptr("Bob"),
			Payload:      `{"name": "Bob", "type": "EdgeZone"}`,
		},
	}
	for i, v := range testData {
		t.Logf("item %d", i)

		var model Model
		if err := json.Unmarshal([]byte(v.Payload), &model); err != nil {
			t.Fatalf("unmarshaling: %+v", err)
		}

		if model.Name == "" && v.ExpectedName == nil {
			// expected failure
			continue
		}

		if model.Name != "" && v.ExpectedName == nil {
			t.Fatalf("expected model to be nil but got %q", model.Name)
		}
		if model.Name == "" && v.ExpectedName != nil {
			t.Fatalf("expected to get a model named %q but didn't", *v.ExpectedName)
		}

		if model.Name != *v.ExpectedName {
			t.Fatalf("expected the deserialized name to be %q but got %q", *v.ExpectedName, model.Name)
		}
	}
}

func ptr(in string) *string {
	return &in
}
