package op_handler

import (
	"time"

	"github.com/danielgtaylor/huma/v2"
)

// SetBodyReadTimeout creates an OperationHandler that sets operation`s BodyReadTimeout field.
// If override is `false`, it will not override existing non-zero value.
func SetBodyReadTimeout(timeout time.Duration, override bool) OperationHandler {
	return func(op *huma.Operation) {
		if override || op.BodyReadTimeout == 0 {
			op.BodyReadTimeout = timeout
		}
	}
}
