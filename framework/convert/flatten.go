package convert

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Expand converts a Terraform / Framework object into a go-azure-sdk (native Go) object
// it will write any diagnostics back to the supplied diag.Diagnostics pointer
func Flatten(ctx context.Context, fwObject any, apiObject any, diags *diag.Diagnostics) {

}

// expand does the heavy lifting via reflection to convert the tfObject into
func flatten() {

}
