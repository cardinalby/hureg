package op_handler

import (
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

// AppendDescription creates an OperationHandler that appends provided `parts` to
// the Description field of an operation separating them with `separator`.
func AppendDescription(separator string, parts ...string) OperationHandler {
	return func(op *huma.Operation) {
		if len(parts) == 0 {
			return
		}
		if op.Description != "" {
			op.Description += separator
		}
		op.Description += strings.Join(parts, separator)
	}
}
