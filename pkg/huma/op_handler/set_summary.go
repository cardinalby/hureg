package op_handler

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

// SetSummary creates an OperationHandler that sets the Summary field on the operation.
// If override is `false`, it will not override the existing summary.
// If summary is set, it also sets the metadata key metadata.KeyIsExplicitSummary to true
// to inform next handlers that summary was set explicitly and should not be re-generated
func SetSummary(summary string, override bool) func(o *huma.Operation) {
	return func(o *huma.Operation) {
		if override || o.Summary == "" {
			o.Summary = summary
			metadata.SetKey(o, metadata.KeyIsExplicitSummary, true)
		}
	}
}
