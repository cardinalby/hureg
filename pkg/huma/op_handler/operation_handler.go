package op_handler

import "github.com/danielgtaylor/huma/v2"

type OperationHandler = func(op *huma.Operation)

// ApplyToOperation is a helper method to apply multiple OperationHandlers to an Operation,
// returning the modified Operation.
func ApplyToOperation(op huma.Operation, handlers ...OperationHandler) huma.Operation {
	for _, h := range handlers {
		h(&op)
	}
	return op
}

// ApplyToOperationPtr is a helper method to apply multiple OperationHandlers to an Operation pointer.
func ApplyToOperationPtr(op *huma.Operation, handlers ...OperationHandler) {
	for _, h := range handlers {
		h(op)
	}
}
