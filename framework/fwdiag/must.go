package fwdiag

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func Must[T any](x T, diags diag.Diagnostics) T {
	return ErrMust(x, DiagAsError(diags))
}

func ErrMust[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func DiagAsError(diags diag.Diagnostics) error {
	errs := make([]error, 0)

	for _, err := range diags.Errors() {
		errStr := err.Summary()
		if err.Detail() != "" {
			errStr += ": " + err.Detail()
		}
		errs = append(errs, errors.New(errStr))
	}

	return errors.Join(errs...)
}
