package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestSetExtensionsKey(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		src      map[string]any
		addKey   string
		addValue any
		override bool
		expected map[string]any
	}{
		{
			name:     "has src, key, value, dont override",
			src:      map[string]any{"key": "value"},
			addKey:   "key",
			addValue: "value2",
			override: false,
			expected: map[string]any{"key": "value"},
		},
		{
			name:     "has src, key, value, override",
			src:      map[string]any{"key": "value"},
			addKey:   "key",
			addValue: "value2",
			override: true,
			expected: map[string]any{"key": "value2"},
		},
		{
			name:     "no src, dont override",
			addKey:   "key",
			addValue: "value",
			override: false,
			expected: map[string]any{"key": "value"},
		},
		{
			name:     "no src, override",
			addKey:   "key",
			addValue: "value",
			override: true,
			expected: map[string]any{"key": "value"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				Extensions: tc.src,
			}
			SetExtensionsKey(tc.addKey, tc.addValue, tc.override)(&op)
			require.Equal(t, tc.expected, op.Extensions)
		})
	}
}
