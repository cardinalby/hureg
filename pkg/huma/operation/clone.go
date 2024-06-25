package operation

import (
	"maps"
	"slices"

	"github.com/danielgtaylor/huma/v2"
)

// Clone creates a 1-level deep copy of an operation. It clones slices and maps in the operation,
// but doesn't go deeper. It's used for fan-out registration middlewares to pass a copy of the
// operation to branch where it can be modified without affecting the siblings.
// The depth of the cloning is enough for the modifications made by library-provided operation handlers.
func Clone(op huma.Operation) huma.Operation {
	op.Errors = slices.Clone(op.Errors)
	op.Metadata = maps.Clone(op.Metadata)
	op.Middlewares = slices.Clone(op.Middlewares)
	op.Tags = slices.Clone(op.Tags)
	if op.ExternalDocs != nil {
		externalDocsCopy := *op.ExternalDocs
		op.ExternalDocs = &externalDocsCopy
	}
	op.Parameters = slices.Clone(op.Parameters)
	if op.RequestBody != nil {
		requestBodyCopy := *op.RequestBody
		op.RequestBody = &requestBodyCopy
	}
	op.Responses = maps.Clone(op.Responses)
	op.Callbacks = maps.Clone(op.Callbacks)
	op.Security = slices.Clone(op.Security)
	op.Servers = slices.Clone(op.Servers)
	op.Extensions = maps.Clone(op.Extensions)
	return op
}
