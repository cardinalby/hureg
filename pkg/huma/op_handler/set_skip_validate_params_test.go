package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestSetSkipValidateParams(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		src      bool
		new      bool
		override bool
		expected bool
	}{
		{
			name:     "has src, new, dont override",
			src:      true,
			new:      false,
			override: false,
			expected: true,
		},
		{
			name:     "has src, new, override",
			src:      true,
			new:      false,
			override: true,
			expected: false,
		},
		{
			name:     "no src, dont override",
			new:      true,
			override: false,
			expected: true,
		},
		{
			name:     "no src, override",
			new:      true,
			override: true,
			expected: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				SkipValidateParams: tc.src,
			}
			SetSkipValidateParams(tc.new, tc.override)(&op)
			require.Equal(t, tc.expected, op.SkipValidateParams)
		})
	}
}
