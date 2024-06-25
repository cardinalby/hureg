package oapi_handlers

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

type TypedStreamHandler = func(ctx context.Context, _ *struct{}) (*huma.StreamResponse, error)

func getTypedStreamHandler(
	adapterHandler func(ctx huma.Context),
) TypedStreamHandler {
	return func(ctx context.Context, _ *struct{}) (*huma.StreamResponse, error) {
		return &huma.StreamResponse{
			Body: adapterHandler,
		}, nil
	}
}
