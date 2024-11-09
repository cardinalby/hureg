package oapi_handlers

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

type StreamResponseHandler[I any] func(ctx context.Context, req I) (*huma.StreamResponse, error)
