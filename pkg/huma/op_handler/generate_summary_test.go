package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

func TestSetDefaultSummary(t *testing.T) {
	t.Parallel()

	type testOutput struct {
		Body string
	}
	var testOutputPtr *testOutput

	testCases := []struct {
		name            string
		method          string
		path            string
		summary         string
		outputPtr       any
		override        bool
		expectedSummary string
	}{
		{
			name:            "GET /",
			method:          "GET",
			path:            "/",
			summary:         "",
			outputPtr:       testOutputPtr,
			override:        false,
			expectedSummary: "Get",
		},
		{
			name:            "GET /test",
			method:          "GET",
			path:            "/test",
			summary:         "",
			outputPtr:       testOutputPtr,
			override:        false,
			expectedSummary: "Get test",
		},
		{
			name:            "override GET /test",
			method:          "GET",
			path:            "/test",
			summary:         "some",
			outputPtr:       testOutputPtr,
			override:        true,
			expectedSummary: "Get test",
		},
		{
			name:            "non-override GET /test",
			method:          "GET",
			path:            "/test",
			summary:         "some",
			outputPtr:       testOutputPtr,
			override:        false,
			expectedSummary: "some",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				Method: tc.method,
				Path:   tc.path,
				Metadata: map[string]any{
					metadata.KeyOutputObjPtr: tc.outputPtr,
				},
				Summary: tc.summary,
			}
			GenerateSummary(tc.override)(&op)
			require.Equal(t, tc.expectedSummary, op.Summary)
		})
	}
}
