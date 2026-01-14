package validators

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestResourceId_ValidateString(t *testing.T) {
	v := AzureResourceManagerId(&commonids.AppServiceId{})

	req := validator.StringRequest{ // TODO - more ID cases?
		ConfigValue: types.StringValue("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Web/sites/site1"),
	}
	var resp validator.StringResponse

	v.ValidateString(context.Background(), req, &resp)

	if len(resp.Diagnostics) != 0 {
		for _, d := range resp.Diagnostics {
			t.Errorf("unexpected diagnostic: %s: %s", d.Summary(), d.Detail())
		}
	}
}

func TestResourceId_ValidateString_IncorrectProvider(t *testing.T) {
	v := AzureResourceManagerId(&commonids.AppServiceId{})

	req := validator.StringRequest{
		ConfigValue: types.StringValue("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/sites/site1"),
	}
	var resp validator.StringResponse

	v.ValidateString(context.Background(), req, &resp)

	if len(resp.Diagnostics) == 0 {
		t.Fatalf("expected diagnostics for invalid provider, got none")
	}

	found := false
	for _, d := range resp.Diagnostics {
		if d.Severity().String() == "Error" && d.Summary() == "ID validation error" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected an error diagnostic with summary 'ID validation error' but got: %#v", resp.Diagnostics)
	}
}

func TestResourceId_ValidateString_Unknown(t *testing.T) {
	v := AzureResourceManagerId(&commonids.AppServiceId{})
	req := validator.StringRequest{ConfigValue: types.StringUnknown()}
	var resp validator.StringResponse
	v.ValidateString(context.Background(), req, &resp)
	if len(resp.Diagnostics) != 0 {
		t.Fatalf("expected no diagnostics for unknown, got %d", resp.Diagnostics.ErrorsCount())
	}
}

func TestResourceId_ValidateString_Null(t *testing.T) {
	v := AzureResourceManagerId(&commonids.AppServiceId{})
	req := validator.StringRequest{ConfigValue: types.StringNull()}
	var resp validator.StringResponse
	v.ValidateString(context.Background(), req, &resp)
	if len(resp.Diagnostics) != 0 {
		t.Fatalf("expected no diagnostics for null, got: %d", resp.Diagnostics.ErrorsCount())
	}
}

func TestResourceId_ValidateParameterString_Valid(t *testing.T) {
	v := AzureResourceManagerId(&commonids.AppServiceId{})

	req := function.StringParameterValidatorRequest{
		ArgumentPosition: 0,
		Value:            types.StringValue("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Web/sites/site1"),
	}
	var resp function.StringParameterValidatorResponse

	v.ValidateParameterString(context.Background(), req, &resp)

	if resp.Error != nil {
		t.Fatalf("unexpected function error: %s", resp.Error.Error())
	}
}

func TestResourceId_ValidateParameterString_IncorrectProvider(t *testing.T) {
	v := AzureResourceManagerId(&commonids.AppServiceId{})

	req := function.StringParameterValidatorRequest{
		ArgumentPosition: 2,
		Value:            types.StringValue("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/sites/site1"),
	}
	var resp function.StringParameterValidatorResponse

	v.ValidateParameterString(context.Background(), req, &resp)

	if resp.Error == nil {
		t.Fatalf("expected function error for invalid provider, got none")
	}
	if resp.Error.FunctionArgument == nil || *resp.Error.FunctionArgument != req.ArgumentPosition {
		t.Fatalf("expected function argument position %d in error, got %#v", req.ArgumentPosition, resp.Error.FunctionArgument)
	}
	if !strings.Contains(resp.Error.Error(), "Invalid Parameter Value: ID validation error") {
		t.Fatalf("unexpected error text: %q", resp.Error.Error())
	}
}

func TestResourceId_ValidateParameterString_Unknown(t *testing.T) {
	v := AzureResourceManagerId(&commonids.AppServiceId{})

	req := function.StringParameterValidatorRequest{ArgumentPosition: 1, Value: types.StringUnknown()}
	var resp function.StringParameterValidatorResponse

	v.ValidateParameterString(context.Background(), req, &resp)

	if resp.Error != nil {
		t.Fatalf("expected no error for unknown, got: %s", resp.Error.Error())
	}
}
func TestResourceId_ValidateParameterString_Null(t *testing.T) {
	v := AzureResourceManagerId(&commonids.AppServiceId{})

	req := function.StringParameterValidatorRequest{ArgumentPosition: 1, Value: types.StringNull()}
	var resp function.StringParameterValidatorResponse
	v.ValidateParameterString(context.Background(), req, &resp)
	if resp.Error != nil {
		t.Fatalf("expected no error for null, got: %s", resp.Error.Error())
	}
}
