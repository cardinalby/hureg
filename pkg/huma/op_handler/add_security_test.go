package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestAddSecurity(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		src             []map[string][]string
		securityEntries []map[string][]string
		expected        []map[string][]string
	}{
		{
			name:            "has src, securityEntries",
			src:             []map[string][]string{{"key": {"value"}}},
			securityEntries: []map[string][]string{{"key2": {"value2"}}},
			expected:        []map[string][]string{{"key": {"value"}}, {"key2": {"value2"}}},
		},
		{
			name:            "no src, securityEntries",
			securityEntries: []map[string][]string{{"key": {"value"}}},
			expected:        []map[string][]string{{"key": {"value"}}},
		},
		{
			name:     "has src, no securityEntries",
			src:      []map[string][]string{{"key": {"value"}}},
			expected: []map[string][]string{{"key": {"value"}}},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				Security: tc.src,
			}
			AddSecurity(tc.securityEntries...)(&op)
			require.Equal(t, tc.expected, op.Security)
		})
	}
}
