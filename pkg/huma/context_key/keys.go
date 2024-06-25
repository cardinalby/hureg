package context_key

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/internal"
)

// KeyTransformers is a key for huma.Context that stores transformers that should be applied
// after the response is generated. It's filled with a slice of transformers stored in APIGen
// instance that was used for operation registration.
// It's for internal library usage, and you are not supposed to modify it.
const KeyTransformers = internal.PkgID + "transformers"

// GetTransformers returns transformers stored in the context.
func GetTransformers(ctx huma.Context) []huma.Transformer {
	transformers, _ := ctx.Context().Value(KeyTransformers).([]huma.Transformer)
	return transformers
}
