package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestSetMetadataKey(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		src      map[string]any
		key      string
		value    any
		override bool
		expected map[string]any
	}{
		{
			name:     "has src, key, value, dont override",
			src:      map[string]any{"key": "value"},
			key:      "key",
			value:    "value2",
			override: false,
			expected: map[string]any{"key": "value"},
		},
		{
			name:     "has src, key, value, override",
			src:      map[string]any{"key": "value"},
			key:      "key",
			value:    "value2",
			override: true,
			expected: map[string]any{"key": "value2"},
		},
		{
			name:     "no src, dont override",
			key:      "key",
			value:    "value",
			override: false,
			expected: map[string]any{"key": "value"},
		},
		{
			name:     "no src, override",
			key:      "key",
			value:    "value",
			override: true,
			expected: map[string]any{"key": "value"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				Metadata: tc.src,
			}
			SetMetadataKey(tc.key, tc.value, tc.override)(&op)
			require.Equal(t, tc.expected, op.Metadata)
		})
	}
}
