package operation

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestClone(t *testing.T) {
	t.Parallel()
	op := huma.Operation{
		Method: "GET",
		Path:   "/test",
		Middlewares: []func(huma.Context, func(huma.Context)){
			func(ctx huma.Context, next func(huma.Context)) { next(ctx) },
		},
		Metadata: map[string]any{
			"key": "value",
		},
		Tags: []string{"tag"},
	}
	cloned := Clone(op)
	require.Equal(t, op.Method, cloned.Method)
	require.Equal(t, op.Path, cloned.Path)
	require.Equal(t, len(op.Middlewares), len(cloned.Middlewares))
	require.Equal(t, op.Metadata, cloned.Metadata)
	require.Equal(t, op.Tags, cloned.Tags)

	cloned.Middlewares = append(cloned.Middlewares, func(ctx huma.Context, next func(huma.Context)) { next(ctx) })
	cloned.Metadata["key"] = "new value"
	cloned.Tags = append(cloned.Tags, "new tag")

	require.Equal(t, 1, len(op.Middlewares))
	require.Equal(t, "value", op.Metadata["key"])
	require.Equal(t, []string{"tag"}, op.Tags)
}
