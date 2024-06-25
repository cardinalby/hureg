package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

type testMiddlewareFactory struct {
	received []int
}

func (f *testMiddlewareFactory) create(vals ...int) (res []func(_ huma.Context, next func(huma.Context))) {
	for _, val := range vals {
		val := val
		res = append(res, func(_ huma.Context, _ func(huma.Context)) {
			f.received = append(f.received, val)
		})
	}
	return res
}

func (f *testMiddlewareFactory) requireValues(
	t *testing.T,
	expectedValues []int,
	middlewares ...func(_ huma.Context, next func(huma.Context)),
) {
	f.received = nil
	for _, m := range middlewares {
		m(nil, nil)
	}
	require.Equal(t, expectedValues, f.received)
}

func TestAddMiddlewares(t *testing.T) {
	t.Parallel()

	f := &testMiddlewareFactory{}

	testCases := []struct {
		name         string
		src          []func(_ huma.Context, next func(huma.Context))
		add          []func(_ huma.Context, next func(huma.Context))
		expectedVals []int
	}{
		{
			name:         "no src, no add",
			src:          nil,
			add:          nil,
			expectedVals: nil,
		},
		{
			name:         "no src, has add",
			src:          nil,
			add:          f.create(1, 2),
			expectedVals: []int{1, 2},
		},
		{
			name:         "has src, no add",
			src:          f.create(1, 2),
			add:          nil,
			expectedVals: []int{1, 2},
		},
		{
			name:         "has src, has add",
			src:          f.create(1, 2),
			add:          f.create(3, 4),
			expectedVals: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			op := huma.Operation{
				Middlewares: tc.src,
			}
			AddMiddlewares(tc.add...)(&op)
			f.requireValues(t, tc.expectedVals, op.Middlewares...)
		})
	}
}
