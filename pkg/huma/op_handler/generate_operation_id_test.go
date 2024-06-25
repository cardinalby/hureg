package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

func TestSetDefaultOperationID(t *testing.T) {
	t.Parallel()

	type testOutput struct {
		Body string
	}
	var testOutputPtr *testOutput

	testCases := []struct {
		name         string
		method       string
		path         string
		operationID  string
		outputPtr    any
		override     bool
		expectedOpId string
	}{
		{
			name:         "GET /",
			method:       "GET",
			path:         "/",
			operationID:  "",
			outputPtr:    testOutputPtr,
			override:     false,
			expectedOpId: "get",
		},
		{
			name:         "GET /test",
			method:       "GET",
			path:         "/test",
			operationID:  "",
			outputPtr:    testOutputPtr,
			override:     false,
			expectedOpId: "get-test",
		},
		{
			name:         "override GET /test",
			method:       "GET",
			path:         "/test",
			operationID:  "some",
			outputPtr:    testOutputPtr,
			override:     true,
			expectedOpId: "get-test",
		},
		{
			name:         "non-override GET /test",
			method:       "GET",
			path:         "/test",
			operationID:  "some",
			outputPtr:    testOutputPtr,
			override:     false,
			expectedOpId: "some",
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
				OperationID: tc.operationID,
			}
			GenerateOperationID(tc.override)(&op)
			require.Equal(t, tc.expectedOpId, op.OperationID)
		})
	}
}
