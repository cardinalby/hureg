package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

func TestUpdateGeneratedSummary(t *testing.T) {
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
		metadata        map[string]any
		expectedSummary string
	}{
		{
			name:    "implicit GET /",
			method:  "GET",
			path:    "/",
			summary: "Get",
			metadata: map[string]any{
				metadata.KeyOutputObjPtr: testOutputPtr,
			},
			expectedSummary: "Get",
		},
		{
			name:    "explicit GET /test",
			method:  "GET",
			path:    "/test",
			summary: "Custom",
			metadata: map[string]any{
				metadata.KeyOutputObjPtr:      testOutputPtr,
				metadata.KeyIsExplicitSummary: true,
				metadata.KeyBasePath:          "/test",
			},
			expectedSummary: "Custom",
		},
		{
			name:    "implicit GET /test/abc",
			method:  "GET",
			path:    "/test/abc",
			summary: "Get abc",
			metadata: map[string]any{
				metadata.KeyOutputObjPtr: testOutputPtr,
				metadata.KeyBasePath:     "/test",
			},
			expectedSummary: "Get test abc",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				Method:   tc.method,
				Path:     tc.path,
				Summary:  tc.summary,
				Metadata: tc.metadata,
			}
			UpdateGeneratedSummary(&op)
			require.Equal(t, tc.expectedSummary, op.Summary)
		})
	}
}
