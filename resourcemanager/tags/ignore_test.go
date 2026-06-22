// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package tags

import (
	"testing"
)

func strPtr(s string) *string {
	return &s
}

func TestIgnoreConfigApplyPtrMap(t *testing.T) {
	testCases := []struct {
		name     string
		config   *IgnoreConfig
		input    map[string]string
		expected map[string]string
	}{
		{
			name:     "nil config is a no-op",
			config:   nil,
			input:    map[string]string{"environment": "prod", "createdBy": "policy"},
			expected: map[string]string{"environment": "prod", "createdBy": "policy"},
		},
		{
			name:     "empty config is a no-op",
			config:   &IgnoreConfig{},
			input:    map[string]string{"environment": "prod", "createdBy": "policy"},
			expected: map[string]string{"environment": "prod", "createdBy": "policy"},
		},
		{
			name:     "exact key is removed",
			config:   &IgnoreConfig{Keys: []string{"createdBy"}},
			input:    map[string]string{"environment": "prod", "createdBy": "policy"},
			expected: map[string]string{"environment": "prod"},
		},
		{
			name:     "exact key match is case-sensitive",
			config:   &IgnoreConfig{Keys: []string{"createdby"}},
			input:    map[string]string{"environment": "prod", "createdBy": "policy"},
			expected: map[string]string{"environment": "prod", "createdBy": "policy"},
		},
		{
			name:     "key prefix is removed",
			config:   &IgnoreConfig{KeyPrefixes: []string{"azure-policy-"}},
			input:    map[string]string{"environment": "prod", "azure-policy-id": "abc", "azure-policy-name": "def"},
			expected: map[string]string{"environment": "prod"},
		},
		{
			name:     "key prefix match is case-sensitive",
			config:   &IgnoreConfig{KeyPrefixes: []string{"Azure-Policy-"}},
			input:    map[string]string{"azure-policy-id": "abc"},
			expected: map[string]string{"azure-policy-id": "abc"},
		},
		{
			name:     "union of keys and key_prefixes",
			config:   &IgnoreConfig{Keys: []string{"createdBy"}, KeyPrefixes: []string{"internal:"}},
			input:    map[string]string{"environment": "prod", "createdBy": "policy", "internal:owner": "x", "internal:cost": "y"},
			expected: map[string]string{"environment": "prod"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.config.ApplyPtrMap(&tc.input)
			if result == nil {
				t.Fatalf("expected non-nil result")
			}
			if len(*result) != len(tc.expected) {
				t.Fatalf("expected %d tags, got %d: %v", len(tc.expected), len(*result), *result)
			}
			for k, v := range tc.expected {
				if (*result)[k] != v {
					t.Fatalf("expected %q=%q, got %q", k, v, (*result)[k])
				}
			}
		})
	}
}

func TestIgnoreConfigApplyPtrMapNilInput(t *testing.T) {
	config := &IgnoreConfig{Keys: []string{"createdBy"}}
	if result := config.ApplyPtrMap(nil); result != nil {
		t.Fatalf("expected nil result for nil input, got %v", result)
	}
}

func TestIgnoreConfigApplyMap(t *testing.T) {
	config := &IgnoreConfig{Keys: []string{"createdBy"}, KeyPrefixes: []string{"internal:"}}
	input := map[string]*string{
		"environment":    strPtr("prod"),
		"createdBy":      strPtr("policy"),
		"internal:owner": strPtr("x"),
	}

	result := config.ApplyMap(input)
	if len(result) != 1 {
		t.Fatalf("expected 1 tag, got %d: %v", len(result), result)
	}
	if result["environment"] == nil || *result["environment"] != "prod" {
		t.Fatalf("expected environment=prod to be retained, got %v", result["environment"])
	}
}

func TestIgnoreConfigApplyMapNilReceiver(t *testing.T) {
	var config *IgnoreConfig
	input := map[string]*string{"createdBy": strPtr("policy")}
	result := config.ApplyMap(input)
	if len(result) != 1 {
		t.Fatalf("expected nil receiver to be a no-op, got %v", result)
	}
}

func TestSetIgnoreAndIgnoreRoundTrip(t *testing.T) {
	t.Cleanup(func() { SetIgnore(nil) })

	if Ignore() != nil {
		t.Fatalf("expected nil ignore config by default")
	}

	cfg := &IgnoreConfig{Keys: []string{"createdBy"}}
	SetIgnore(cfg)
	if Ignore() != cfg {
		t.Fatalf("expected Ignore() to return the configured value")
	}
}

func TestExpandHonoursActiveIgnoreConfig(t *testing.T) {
	t.Cleanup(func() { SetIgnore(nil) })
	SetIgnore(&IgnoreConfig{Keys: []string{"createdBy"}})

	result := Expand(map[string]interface{}{
		"environment": "prod",
		"createdBy":   "policy",
	})

	if len(*result) != 1 {
		t.Fatalf("expected ignored key to be excluded on expand, got %v", *result)
	}
	if (*result)["environment"] != "prod" {
		t.Fatalf("expected environment=prod, got %v", *result)
	}
}

func TestFlattenHonoursActiveIgnoreConfig(t *testing.T) {
	t.Cleanup(func() { SetIgnore(nil) })
	SetIgnore(&IgnoreConfig{KeyPrefixes: []string{"azure-policy-"}})

	input := map[string]string{
		"environment":     "prod",
		"azure-policy-id": "abc",
	}
	result := Flatten(&input)

	if len(result) != 1 {
		t.Fatalf("expected ignored prefix key to be scrubbed on flatten, got %v", result)
	}
	if result["environment"] != "prod" {
		t.Fatalf("expected environment=prod, got %v", result)
	}
}
