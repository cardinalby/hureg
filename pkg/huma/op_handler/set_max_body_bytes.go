package op_handler

import (
	"github.com/danielgtaylor/huma/v2"
)

// SetMaxBodyBytes creates an OperationHandler that sets the operation's MaxBodyBytes field to the given value.
// If override is `false`, it will not override existing non-zero MaxBodyBytes.
func SetMaxBodyBytes(n int64, override bool) OperationHandler {
	return func(op *huma.Operation) {
		if override || op.MaxBodyBytes == 0 {
			op.MaxBodyBytes = n
		}
	}
}
