package op_handler

import (
	"github.com/danielgtaylor/huma/v2"
)

// SetSkipValidateParams creates an OperationHandler that sets the SkipValidateParams field on the operation.
// If override is `false`, it will not override the existing `true` value.
func SetSkipValidateParams(skipValidateParams bool, override bool) OperationHandler {
	return func(op *huma.Operation) {
		if override || !op.SkipValidateParams {
			op.SkipValidateParams = skipValidateParams
		}
	}
}
