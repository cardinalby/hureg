package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

func TestUpdateOperationID(t *testing.T) {
	t.Parallel()

	type testOutput struct {
		Body string
	}
	var testOutputPtr *testOutput

	testCases := []struct {
		name                string
		method              string
		path                string
		metadata            map[string]any
		operationID         string
		explicitOpIDBuilder func(*huma.Operation) string
		expectedOpId        string
	}{
		{
			name:   "implicit GET /",
			method: "GET",
			path:   "/",
			metadata: map[string]any{
				metadata.KeyIsExplicitOperationID: false,
				metadata.KeyOutputObjPtr:          testOutputPtr,
				metadata.KeyInitOperationID:       "get",
			},
			operationID:         "get",
			explicitOpIDBuilder: nil,
			expectedOpId:        "get",
		},
		{
			name:   "implicit GET /test to /test/abc",
			method: "GET",
			path:   "/test/abc",
			metadata: map[string]any{
				metadata.KeyIsExplicitOperationID: false,
				metadata.KeyOutputObjPtr:          testOutputPtr,
				metadata.KeyInitOperationID:       "custom",
				metadata.KeyBasePath:              "/test/abc",
			},
			operationID:         "get-test",
			explicitOpIDBuilder: nil,
			expectedOpId:        "get-test-abc",
		},
		{
			name:   "explicit GET /test to /test/abc with /test basePath with nil builder",
			method: "GET",
			path:   "/test/abc",
			metadata: map[string]any{
				metadata.KeyIsExplicitOperationID: true,
				metadata.KeyOutputObjPtr:          testOutputPtr,
				metadata.KeyInitOperationID:       "custom",
				metadata.KeyBasePath:              "/test",
			},
			operationID:         "abc_custom",
			explicitOpIDBuilder: nil,
			expectedOpId:        "test-custom",
		},
		{
			name:   "explicit GET /test to /test/abc with /test basePath with builder",
			method: "GET",
			path:   "/test/abc",
			metadata: map[string]any{
				metadata.KeyIsExplicitOperationID: true,
				metadata.KeyOutputObjPtr:          testOutputPtr,
				metadata.KeyInitOperationID:       "custom",
				metadata.KeyBasePath:              "/test",
			},
			operationID: "abc-custom",
			explicitOpIDBuilder: func(o *huma.Operation) string {
				return o.Metadata[metadata.KeyBasePath].(string) +
					"_" +
					o.Metadata[metadata.KeyInitOperationID].(string)
			},
			expectedOpId: "/test_custom",
		},
		{
			name:   "explicit GET /test to /test/abc with no basePath with nil builder",
			method: "GET",
			path:   "/test/abc",
			metadata: map[string]any{
				metadata.KeyIsExplicitOperationID: true,
				metadata.KeyOutputObjPtr:          testOutputPtr,
				metadata.KeyInitOperationID:       "custom",
			},
			operationID:         "custom",
			explicitOpIDBuilder: nil,
			expectedOpId:        "test-abc-custom",
		},
		{
			name:   "explicit with empty InitOperationID with nil builder",
			method: "GET",
			path:   "/test/abc",
			metadata: map[string]any{
				metadata.KeyIsExplicitOperationID: true,
				metadata.KeyOutputObjPtr:          testOutputPtr,
				metadata.KeyInitOperationID:       "",
				metadata.KeyBasePath:              "/test",
			},
			operationID:         "",
			explicitOpIDBuilder: nil,
			expectedOpId:        "",
		},
		{
			name:   "explicit with empty InitOperationID with set OperationID with nil builder",
			method: "GET",
			path:   "/test/abc",
			metadata: map[string]any{
				metadata.KeyIsExplicitOperationID: true,
				metadata.KeyOutputObjPtr:          testOutputPtr,
				metadata.KeyInitOperationID:       "",
				metadata.KeyBasePath:              "/test",
			},
			operationID:         "some",
			explicitOpIDBuilder: nil,
			expectedOpId:        "test-some",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				OperationID: tc.operationID,
				Method:      tc.method,
				Path:        tc.path,
				Metadata:    tc.metadata,
			}
			UpdateOperationID(tc.explicitOpIDBuilder)(&op)
			require.Equal(t, tc.expectedOpId, op.OperationID)
		})
	}
}
