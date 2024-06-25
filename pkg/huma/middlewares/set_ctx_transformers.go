package middlewares

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/context_key"
)

// SetCtxTransformers sets transformers slice to the context key. These transformers will be applied
// to the response in the humaApiWrapper.Transform method.
// It's for internal library usage, and it will work only if you use `hureg` lib registration functions.
func SetCtxTransformers(transformers []huma.Transformer) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		next(huma.WithValue(ctx, context_key.KeyTransformers, transformers))
	}
}
