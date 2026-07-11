// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package commonschema

import (
	"context"
	"testing"

	"github.com/hashicorp/go-azure-helpers/framework/typehelpers"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/tags"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExpandTagsHonoursIgnoreConfig(t *testing.T) {
	t.Cleanup(func() { tags.SetIgnore(nil) })
	tags.SetIgnore(&tags.IgnoreConfig{Keys: []string{"createdBy"}})

	ctx := context.Background()
	input := typehelpers.NewMapValueOfMust[types.String](ctx, map[string]attr.Value{
		"environment": types.StringValue("prod"),
		"createdBy":   types.StringValue("policy"),
	})

	var diags diag.Diagnostics
	result := ExpandTags(ctx, input, &diags)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics from ExpandTags: %v", diags)
	}
	if result == nil {
		t.Fatalf("expected non-nil result")
	}
	if len(*result) != 1 {
		t.Fatalf("expected ignored key to be excluded, got %v", *result)
	}
	if (*result)["environment"] != "prod" {
		t.Fatalf("expected environment=prod, got %v", *result)
	}
}

func TestFlattenTagsHonoursIgnoreConfig(t *testing.T) {
	t.Cleanup(func() { tags.SetIgnore(nil) })
	tags.SetIgnore(&tags.IgnoreConfig{KeyPrefixes: []string{"azure-policy-"}})

	ctx := context.Background()
	input := map[string]string{
		"environment":     "prod",
		"azure-policy-id": "abc",
	}

	var diags diag.Diagnostics
	result := FlattenTags(ctx, &input, &diags)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics from FlattenTags: %v", diags)
	}

	elems := result.Elements()
	if len(elems) != 1 {
		t.Fatalf("expected ignored prefix key to be scrubbed, got %v", elems)
	}
	if _, ok := elems["environment"]; !ok {
		t.Fatalf("expected environment to be retained, got %v", elems)
	}
}
