package op_handler

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

// SetMetadataKey creates an OperationHandler that sets the metadata key for the operation
// If override is `false`, it will not override the existing key.
func SetMetadataKey(key string, value any, override bool) OperationHandler {
	return func(op *huma.Operation) {
		if _, hasKey := op.Metadata[key]; hasKey && !override {
			return
		}
		metadata.SetKey(op, key, value)
	}
}
