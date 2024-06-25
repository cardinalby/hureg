package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestAddTags(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		src      []string
		add      []string
		expected []string
	}{
		{
			name:     "no src, no add",
			src:      nil,
			add:      nil,
			expected: nil,
		},
		{
			name:     "no src, has add",
			src:      nil,
			add:      []string{"tag1", "tag2"},
			expected: []string{"tag1", "tag2"},
		},
		{
			name:     "has src, no add",
			src:      []string{"tag1", "tag2"},
			add:      nil,
			expected: []string{"tag1", "tag2"},
		},
		{
			name:     "has src, has add",
			src:      []string{"tag1", "tag2"},
			add:      []string{"tag3", "tag4"},
			expected: []string{"tag1", "tag2", "tag3", "tag4"},
		},
		{
			name:     "duplicated tags",
			src:      []string{"tag1", "tag2"},
			add:      []string{"tag2", "tag3"},
			expected: []string{"tag1", "tag2", "tag3"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				Tags: tc.src,
			}
			AddTags(tc.add...)(&op)
			require.Equal(t, tc.expected, op.Tags)
		})
	}
}
