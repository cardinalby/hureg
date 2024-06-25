package op_handler

import (
	"github.com/danielgtaylor/huma/v2"
)

// SetExtensionsKey creates an OperationHandler that adds a key to the operation`s Extensions field.
// If override is `false`, it will not override existing key.
func SetExtensionsKey(key string, value any, override bool) OperationHandler {
	return func(op *huma.Operation) {
		if op.Extensions == nil {
			op.Extensions = make(map[string]any)
		} else if _, hasKey := op.Extensions[key]; hasKey && !override {
			return
		}
		op.Extensions[key] = value
	}
}
