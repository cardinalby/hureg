package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestSetResponse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		statusCode        int
		response          *huma.Response
		override          bool
		existingResponses map[string]*huma.Response
		expectedResponses map[string]*huma.Response
	}{
		{
			name:       "no existing responses",
			statusCode: 200,
			response: &huma.Response{
				Description: "OK",
			},
			override:          false,
			existingResponses: nil,
			expectedResponses: map[string]*huma.Response{
				"200": {
					Description: "OK",
				},
			},
		},
		{
			name:       "existing response, dont override",
			statusCode: 200,
			response: &huma.Response{
				Description: "OK",
			},
			override: false,
			existingResponses: map[string]*huma.Response{
				"200": {
					Description: "Existing",
				},
			},
			expectedResponses: map[string]*huma.Response{
				"200": {
					Description: "Existing",
				},
			},
		},
		{
			name:       "existing response, override",
			statusCode: 200,
			response: &huma.Response{
				Description: "OK",
			},
			override: true,
			existingResponses: map[string]*huma.Response{
				"200": {
					Description: "Existing",
				},
			},
			expectedResponses: map[string]*huma.Response{
				"200": {
					Description: "OK",
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				Responses: tc.existingResponses,
			}
			SetResponse(tc.statusCode, tc.response, tc.override)(&op)
			require.Equal(t, tc.expectedResponses, op.Responses)
		})
	}
}
