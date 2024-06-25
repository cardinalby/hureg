package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestSetMaxBodyBytes(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		src      int64
		dst      int64
		override bool
		expected int64
	}{
		{
			name:     "has src, dst, dont override",
			src:      1,
			dst:      2,
			override: false,
			expected: 1,
		},
		{
			name:     "has src, dst, override",
			src:      1,
			dst:      2,
			override: true,
			expected: 2,
		},
		{
			name:     "no src, dont override",
			dst:      2,
			override: false,
			expected: 2,
		},
		{
			name:     "no src, override",
			dst:      2,
			override: true,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				MaxBodyBytes: tc.src,
			}
			SetMaxBodyBytes(tc.dst, tc.override)(&op)
			require.Equal(t, tc.expected, op.MaxBodyBytes)
		})
	}
}
