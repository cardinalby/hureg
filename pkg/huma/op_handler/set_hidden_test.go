package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestSetHidden(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		isHidden         bool
		override         bool
		existingIsHidden bool
		expected         bool
	}{
		{
			name:             "non-override true, existing isHidden false",
			isHidden:         true,
			override:         true,
			existingIsHidden: false,
			expected:         true,
		},
		{
			name:             "non-override false, existing isHidden true",
			isHidden:         true,
			override:         false,
			existingIsHidden: true,
			expected:         true,
		},
		{
			name:             "override false, existing isHidden true",
			isHidden:         false,
			override:         true,
			existingIsHidden: true,
			expected:         false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := &huma.Operation{Hidden: tc.existingIsHidden}
			SetHidden(tc.isHidden, tc.override)(op)
			require.Equal(t, tc.expected, op.Hidden)
		})
	}
}
