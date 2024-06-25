package op_handler

import (
	"testing"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestSetBodyReadTimeout(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		src      time.Duration
		dst      time.Duration
		override bool
		expected time.Duration
	}{
		{
			name:     "has src, dst, dont override",
			src:      time.Second,
			dst:      time.Hour,
			override: false,
			expected: time.Second,
		},
		{
			name:     "has src, dst, override",
			src:      time.Second,
			dst:      time.Hour,
			override: true,
			expected: time.Hour,
		},
		{
			name:     "no src, dont override",
			dst:      time.Hour,
			override: false,
			expected: time.Hour,
		},
		{
			name:     "no src, override",
			dst:      time.Hour,
			override: true,
			expected: time.Hour,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				BodyReadTimeout: tc.src,
			}
			SetBodyReadTimeout(tc.dst, tc.override)(&op)
			require.Equal(t, tc.expected, op.BodyReadTimeout)
		})
	}

}
