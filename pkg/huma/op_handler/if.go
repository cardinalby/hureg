package op_handler

import (
	"github.com/danielgtaylor/huma/v2"
)

// If creates an OperationHandler that applies the given handlers if the condition is true.
// It's a helper method to avoid creating a separate OperationHandlers with condition checks.
func If(
	condition func(op *huma.Operation) bool,
	handlers ...OperationHandler,
) OperationHandler {
	return func(op *huma.Operation) {
		if condition(op) {
			for _, handler := range handlers {
				handler(op)
			}
		}
	}
}
