package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestSetDescription(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		description  string
		override     bool
		existingDesc string
		expectedDesc string
	}{
		{
			name:         "no existing description",
			description:  "New description",
			override:     false,
			existingDesc: "",
			expectedDesc: "New description",
		},
		{
			name:         "existing description, dont override",
			description:  "New description",
			override:     false,
			existingDesc: "Existing",
			expectedDesc: "Existing",
		},
		{
			name:         "existing description, override",
			description:  "New description",
			override:     true,
			existingDesc: "Existing",
			expectedDesc: "New description",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			op := &huma.Operation{}
			op.Description = tc.existingDesc

			SetDescription(tc.description, tc.override)(op)
			require.Equal(t, tc.expectedDesc, op.Description)
		})
	}
}
