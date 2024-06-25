package metadata

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestSetKey(t *testing.T) {
	t.Parallel()
	op := &huma.Operation{}
	SetKey(op, "key", "value")
	require.Equal(t, "value", op.Metadata["key"])
	SetKey(op, "key", "value2")
	require.Equal(t, "value2", op.Metadata["key"])
	SetKey(op, "key2", "value3")
	require.Equal(t, "value3", op.Metadata["key2"])
}
