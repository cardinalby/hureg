package hureg

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"

	"github.com/cardinalby/hureg/pkg/huma/op_handler"
)

func TestNewRegMiddleware(t *testing.T) {
	t.Parallel()
	op := huma.Operation{
		Tags: []string{"tag1"},
	}
	rm := NewRegMiddleware(
		op_handler.AddTags("tag2"),
		op_handler.AddTags("tag3"),
	)
	receiverWasCalled := false
	rm(op, func(op huma.Operation) {
		receiverWasCalled = true
		require.Equal(t, []string{"tag1", "tag2", "tag3"}, op.Tags)
	})
	require.True(t, receiverWasCalled)
}

func TestRegMiddlewares_Chain(t *testing.T) {
	t.Parallel()
	op := huma.Operation{
		Tags: []string{"tag1"},
	}
	rms := RegMiddlewares{
		NewRegMiddleware(op_handler.AddTags("tag2")),
		NewRegMiddleware(op_handler.AddTags("tag3")),
	}
	chained := rms.Chain()
	receiverWasCalled := false
	chained(op, func(op huma.Operation) {
		receiverWasCalled = true
		require.Equal(t, []string{"tag1", "tag2", "tag3"}, op.Tags)
	})
	require.True(t, receiverWasCalled)
}

func TestRegMiddlewares_FanOut(t *testing.T) {
	t.Parallel()
	op := huma.Operation{
		Method: "GET",
		Path:   "/",
		Tags:   []string{"tag1"},
	}
	rms := RegMiddlewares{
		NewRegMiddleware(op_handler.AddTags("tag2")),
		NewRegMiddleware(op_handler.AddTags("tag3")),
	}
	fanOut := rms.FanOut()
	receiverCallsNumber := 0
	fanOut(op, func(op huma.Operation) {
		receiverCallsNumber++
		switch receiverCallsNumber {
		case 1:
			require.Equal(t, []string{"tag1", "tag2"}, op.Tags)
		case 2:
			require.Equal(t, []string{"tag1", "tag3"}, op.Tags)
		default:
			require.Fail(t, "unexpected receiver call: %d", receiverCallsNumber)
		}
	})
	require.Equal(t, 2, receiverCallsNumber)
}
