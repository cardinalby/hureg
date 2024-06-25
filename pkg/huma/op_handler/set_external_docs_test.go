package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestSetExternalDocs(t *testing.T) {
	t.Parallel()

	//goland:noinspection HttpUrlsUsage
	testCases := []struct {
		name                 string
		externalDocs         *huma.ExternalDocs
		override             bool
		existingExternalDocs *huma.ExternalDocs
		expectedExternalDocs *huma.ExternalDocs
	}{
		{
			name: "override existing",
			externalDocs: &huma.ExternalDocs{
				Description: "external docs",
				URL:         "http://example.com",
			},
			override: true,
			existingExternalDocs: &huma.ExternalDocs{
				Description: "existing external docs",
				URL:         "http://example.com/existing",
			},
			expectedExternalDocs: &huma.ExternalDocs{
				Description: "external docs",
				URL:         "http://example.com",
			},
		},
		{
			name: "do not override existing",
			externalDocs: &huma.ExternalDocs{
				Description: "external docs",
				URL:         "http://example.com",
			},
			override: false,
			existingExternalDocs: &huma.ExternalDocs{
				Description: "existing external docs",
				URL:         "http://example.com/existing",
			},
			expectedExternalDocs: &huma.ExternalDocs{
				Description: "existing external docs",
				URL:         "http://example.com/existing",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := &huma.Operation{
				ExternalDocs: tc.existingExternalDocs,
			}
			SetExternalDocs(tc.externalDocs, tc.override)(op)
			require.Equal(t, tc.expectedExternalDocs, op.ExternalDocs)
		})
	}
}
