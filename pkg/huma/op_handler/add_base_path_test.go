package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

func TestAddBasePath(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		opInitPath       string
		opBasePath       string
		opPath           string
		addBasePath      string
		expectedPath     string
		expectedBasePath any
	}{
		{
			name:             "[]/ to [abc]/",
			opInitPath:       "/",
			opBasePath:       "",
			opPath:           "/",
			addBasePath:      "/abc",
			expectedPath:     "/abc/",
			expectedBasePath: "/abc",
		},
		{
			name:             "[abc]/ to [abc/def]/",
			opInitPath:       "/",
			opBasePath:       "/abc",
			opPath:           "/abc/",
			addBasePath:      "/def",
			expectedPath:     "/abc/def/",
			expectedBasePath: "/abc/def",
		},
		{
			name:             "[a]/c to [a/b]/c",
			opInitPath:       "/c",
			opBasePath:       "/a",
			opPath:           "/a/c",
			addBasePath:      "/b",
			expectedPath:     "/a/b/c",
			expectedBasePath: "/a/b",
		},
		{
			name:             "add empty base path",
			opInitPath:       "/b",
			opBasePath:       "/a",
			opPath:           "/a/b",
			addBasePath:      "",
			expectedPath:     "/a/b",
			expectedBasePath: "/a",
		},
		{
			name:             "add slash base path",
			opInitPath:       "/b",
			opBasePath:       "/a",
			opPath:           "/a/b",
			addBasePath:      "/",
			expectedPath:     "/a/b",
			expectedBasePath: "/a",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			op := huma.Operation{
				Path: tc.opPath,
				Metadata: map[string]interface{}{
					metadata.KeyInitPath: tc.opInitPath,
				},
			}
			if tc.opBasePath != "" {
				op.Metadata[metadata.KeyBasePath] = tc.opBasePath
			}
			AddBasePath(tc.addBasePath)(&op)
			require.Equal(t, tc.expectedPath, op.Path)
			require.Equal(t, tc.expectedBasePath, op.Metadata[metadata.KeyBasePath])
		})
	}
}
