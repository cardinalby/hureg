package op_handler

import "github.com/danielgtaylor/huma/v2"

// SetDescription creates an OperationHandler that sets the Description field of an operation
// If override is `false`, it will not override the existing description.
func SetDescription(description string, override bool) OperationHandler {
	return func(op *huma.Operation) {
		if op.Description != "" && !override {
			return
		}
		op.Description = description
	}
}
