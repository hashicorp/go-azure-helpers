package convert

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// convert handles the conversion between Framework types and Go native types used by go-azure-sdk
// returns reflect.Values for the order from -> to and a diag.Diagnostics for warnings/errors
func convert(fwType, sdkType any) (reflect.Value, reflect.Value, diag.Diagnostics) {
	diags := diag.Diagnostics{}
	fwVal := reflect.ValueOf(fwType)
	sdkVal := reflect.ValueOf(sdkType)

	if fwVal.Kind() == reflect.Ptr {
		fwVal = fwVal.Elem()
	}

	if kind := sdkVal.Kind(); kind != reflect.Ptr {
		diags.AddError("convert", fmt.Sprintf("target (%T): %s is not a pointer", sdkType, kind))
		return reflect.Value{}, reflect.Value{}, diags
	}

	sdkVal = sdkVal.Elem()

	return fwVal, sdkVal, diags
}
