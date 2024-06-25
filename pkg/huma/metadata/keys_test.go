package metadata

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestIsExplicitOperationID(t *testing.T) {
	t.Parallel()
	op := huma.Operation{}
	require.False(t, IsExplicitOperationID(&op))
	SetKey(&op, KeyIsExplicitOperationID, true)
	require.True(t, IsExplicitOperationID(&op))
	SetKey(&op, KeyIsExplicitOperationID, false)
	require.False(t, IsExplicitOperationID(&op))
	SetKey(&op, KeyIsExplicitOperationID, "unknown")
	require.False(t, IsExplicitOperationID(&op))
}

func TestIsExplicitSummary(t *testing.T) {
	t.Parallel()
	op := huma.Operation{}
	require.False(t, IsExplicitSummary(&op))
	SetKey(&op, KeyIsExplicitSummary, true)
	require.True(t, IsExplicitSummary(&op))
	SetKey(&op, KeyIsExplicitSummary, false)
	require.False(t, IsExplicitSummary(&op))
	SetKey(&op, KeyIsExplicitSummary, "unknown")
	require.False(t, IsExplicitSummary(&op))
}

func TestGetInitOperationID(t *testing.T) {
	t.Parallel()
	op := huma.Operation{}
	id, ok := GetInitOperationID(&op)
	require.False(t, ok)
	require.Empty(t, id)
	SetKey(&op, KeyInitOperationID, "test")
	id, ok = GetInitOperationID(&op)
	require.True(t, ok)
	require.Equal(t, "test", id)
}

func TestGetInitPath(t *testing.T) {
	t.Parallel()
	op := huma.Operation{}
	path, ok := GetInitPath(&op)
	require.False(t, ok)
	require.Empty(t, path)
	SetKey(&op, KeyInitPath, "test")
	path, ok = GetInitPath(&op)
	require.True(t, ok)
	require.Equal(t, "test", path)
}

func TestGetBasePath(t *testing.T) {
	t.Parallel()
	op := huma.Operation{}
	path, ok := GetBasePath(&op)
	require.False(t, ok)
	require.Empty(t, path)
	SetKey(&op, KeyBasePath, "test")
	path, ok = GetBasePath(&op)
	require.True(t, ok)
	require.Equal(t, "test", path)
}
