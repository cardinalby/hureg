package routepath

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJoin(t *testing.T) {
	t.Parallel()
	require.Equal(t, "/a/b", Join("/a", "b"))
	require.Equal(t, "/a/b", Join("/a", "/b"))
	require.Equal(t, "/a/b/", Join("/a", "b/"))
	require.Equal(t, "/a", Join("/a", ""))
}
