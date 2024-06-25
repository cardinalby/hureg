package op_handler

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"
)

func TestApplyToOperation(t *testing.T) {
	t.Parallel()
	op := ApplyToOperation(huma.Operation{},
		func(op *huma.Operation) {
			op.Summary += "a"
		},
		func(op *huma.Operation) {
			op.Summary += "b"
		},
	)
	require.Equal(t, "ab", op.Summary)
}

func TestApplyToOperationPtr(t *testing.T) {
	t.Parallel()
	op := huma.Operation{}
	ApplyToOperationPtr(&op,
		func(op *huma.Operation) {
			op.Summary += "a"
		},
		func(op *huma.Operation) {
			op.Summary += "b"
		},
	)
	require.Equal(t, "ab", op.Summary)
}
