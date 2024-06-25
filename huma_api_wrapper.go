package hureg

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/context_key"
)

// humaApiWrapper is a wrapper around huma.API that allows to apply transformers from the context key
// to the response
type humaApiWrapper struct {
	huma.API
}

func newHumaApiWrapper(api huma.API) humaApiWrapper {
	return humaApiWrapper{api}
}

// Transform applies transformers from the context key to the response and then calls the wrapped API.Transform
func (w humaApiWrapper) Transform(ctx huma.Context, status string, v any) (res any, err error) {
	// Apply transformers from the context that was set to a particular ApiGen instance
	for _, t := range context_key.GetTransformers(ctx) {
		v, err = t(ctx, status, v)
		if err != nil {
			return nil, err
		}
	}
	return w.API.Transform(ctx, status, v)
}
