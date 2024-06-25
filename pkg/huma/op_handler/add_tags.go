package op_handler

import (
	"github.com/danielgtaylor/huma/v2"
	"golang.org/x/exp/slices"
)

// AddTags creates an OperationHandler that appends the given tags to the operation's Tags field.
// If a tag is already present, it will not be duplicated.
func AddTags(tags ...string) OperationHandler {
	return func(op *huma.Operation) {
		for _, tag := range tags {
			if !slices.Contains(op.Tags, tag) {
				op.Tags = append(op.Tags, tag)
			}
		}
	}
}
