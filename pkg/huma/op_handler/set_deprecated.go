package op_handler

import (
	"github.com/danielgtaylor/huma/v2"
)

// SetDeprecated creates an OperationHandler that sets operation`s Deprecated field.
// If override is `false`, it will not override existing `true` value.
func SetDeprecated(deprecated bool, override bool) OperationHandler {
	return func(op *huma.Operation) {
		if override || !op.Deprecated {
			op.Deprecated = deprecated
		}
	}
}
