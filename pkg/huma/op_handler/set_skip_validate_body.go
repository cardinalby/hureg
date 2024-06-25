package op_handler

import (
	"github.com/danielgtaylor/huma/v2"
)

// SetSkipValidateBody creates an OperationHandler that sets the SkipValidateBody field on the operation.
// If override is `false`, it will not override the existing `true` value.
func SetSkipValidateBody(skipValidateBody bool, override bool) OperationHandler {
	return func(op *huma.Operation) {
		if override || !op.SkipValidateBody {
			op.SkipValidateBody = skipValidateBody
		}
	}
}
