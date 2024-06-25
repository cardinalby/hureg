package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

func TestSetSummary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                    string
		summary                 string
		isExplicitSummary       bool
		newSummary              string
		override                bool
		expectedSummary         string
		expectedExplicitSummary bool
	}{
		{
			name:                    "implicit set no override",
			summary:                 "Abc",
			isExplicitSummary:       false,
			newSummary:              "Def",
			override:                false,
			expectedSummary:         "Abc",
			expectedExplicitSummary: false,
		},
		{
			name:                    "implicit set override",
			summary:                 "Abc",
			isExplicitSummary:       false,
			newSummary:              "Def",
			override:                true,
			expectedSummary:         "Def",
			expectedExplicitSummary: true,
		},
		{
			name:                    "explicit set no override",
			summary:                 "Abc",
			isExplicitSummary:       true,
			newSummary:              "Def",
			override:                false,
			expectedSummary:         "Abc",
			expectedExplicitSummary: true,
		},
		{
			name:                    "explicit set override",
			summary:                 "Abc",
			isExplicitSummary:       true,
			newSummary:              "Def",
			override:                true,
			expectedSummary:         "Def",
			expectedExplicitSummary: true,
		},
		{
			name:                    "empty set no override",
			summary:                 "",
			isExplicitSummary:       false,
			newSummary:              "Def",
			override:                false,
			expectedSummary:         "Def",
			expectedExplicitSummary: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				Summary: tc.summary,
			}
			if tc.isExplicitSummary {
				metadata.SetKey(&op, metadata.KeyIsExplicitSummary, true)
			}
			SetSummary(tc.newSummary, tc.override)(&op)
			require.Equal(t, tc.expectedSummary, op.Summary)
			require.Equal(t, tc.expectedExplicitSummary, metadata.IsExplicitSummary(&op))
		})
	}
}
