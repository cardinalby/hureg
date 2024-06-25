package op_handler

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

// GenerateSummary creates an OperationHandler that generates a summary for the operation in a way
// Huma does it by default.
// If override is `false`, it will not override existing Summary.
// It's supposed to be used internally by the library, but you can use it in your registration pipeline
// with override == false to initialize explicitly provided operations with not specified Summary.
func GenerateSummary(override bool) OperationHandler {
	return func(op *huma.Operation) {
		if override || op.Summary == "" {
			op.Summary = huma.GenerateSummary(
				op.Method,
				op.Path,
				op.Metadata[metadata.KeyOutputObjPtr],
			)
		}
	}
}
