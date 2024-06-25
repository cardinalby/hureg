package op_handler

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

// GenerateOperationID creates an OperationHandler that sets the OperationID for the operation in a way
// Huma does it by default.
// If override is `false`, it will not override existing OperationID.
// It's supposed to be used internally by the library, but you can use it in your registration pipeline
// with override == false to initialize explicitly provided operations with not specified OperationID.
func GenerateOperationID(override bool) OperationHandler {
	return func(op *huma.Operation) {
		if override || op.OperationID == "" {
			op.OperationID = huma.GenerateOperationID(
				op.Method,
				op.Path,
				op.Metadata[metadata.KeyOutputObjPtr],
			)
		}
	}
}
