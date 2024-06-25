package op_handler

import "github.com/danielgtaylor/huma/v2"

// SetExternalDocs returns an OperationHandler that sets ExternalDocs field of an operation.
// If override is `false` and the operation already has ExternalDocs, it will not be overridden.
func SetExternalDocs(externalDocs *huma.ExternalDocs, override bool) OperationHandler {
	return func(op *huma.Operation) {
		if !override && op.ExternalDocs != nil {
			return
		}
		op.ExternalDocs = externalDocs
	}
}
