package op_handler

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

// UpdateGeneratedSummary creates an OperationHandler that re-generates the Summary field on the operation
// if it is not explicitly set.
func UpdateGeneratedSummary(op *huma.Operation) {
	if !metadata.IsExplicitSummary(op) {
		GenerateSummary(true)(op)
	}
}
