package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestSetDeprecated(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		src      bool
		dst      bool
		override bool
		expected bool
	}{
		{
			name:     "has src, dst, dont override",
			src:      true,
			dst:      false,
			override: false,
			expected: true,
		},
		{
			name:     "has src, dst, override",
			src:      true,
			dst:      false,
			override: true,
			expected: false,
		},
		{
			name:     "no src, dont override",
			src:      false,
			dst:      false,
			override: false,
			expected: false,
		},
		{
			name:     "no src, override",
			src:      false,
			dst:      true,
			override: true,
			expected: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				Deprecated: tc.src,
			}
			SetDeprecated(tc.dst, tc.override)(&op)
			require.Equal(t, tc.expected, op.Deprecated)
		})
	}
}
