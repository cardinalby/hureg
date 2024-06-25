package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestAppendDescription(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		separator    string
		parts        []string
		existingDesc string
		expectedDesc string
	}{
		{
			name:         "no existing description",
			separator:    ", ",
			parts:        []string{"New", "description"},
			existingDesc: "",
			expectedDesc: "New, description",
		},
		{
			name:         "existing description",
			separator:    ", ",
			parts:        []string{"New", "description"},
			existingDesc: "Existing",
			expectedDesc: "Existing, New, description",
		},
		{
			name:         "existing description, empty parts",
			separator:    ", ",
			parts:        []string{},
			existingDesc: "Existing",
			expectedDesc: "Existing",
		},
		{
			name:         "existing description, empty separator",
			separator:    "",
			parts:        []string{"New", "description"},
			existingDesc: "Existing",
			expectedDesc: "ExistingNewdescription",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			op := &huma.Operation{
				Description: tc.existingDesc,
			}
			AppendDescription(tc.separator, tc.parts...)(op)
			require.Equal(t, tc.expectedDesc, op.Description)
		})
	}
}
