package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestIf(t *testing.T) {
	t.Parallel()

	oh := If(
		func(op *huma.Operation) bool {
			return op.OperationID == "test"
		},
		func(op *huma.Operation) {
			op.Summary = "effect"
		},
	)

	op := &huma.Operation{
		OperationID: "test",
	}
	oh(op)
	require.Equal(t, "effect", op.Summary)

	op.OperationID = "not-test"
	op.Summary = ""
	oh(op)
	require.Equal(t, "", op.Summary)

}
