package op_handler

import (
	"github.com/danielgtaylor/huma/v2"
)

// AddMiddlewares creates an OperationHandler that appends the given middlewares to the operation's Middlewares field.
func AddMiddlewares(middlewares ...func(huma.Context, func(huma.Context))) OperationHandler {
	return func(op *huma.Operation) {
		op.Middlewares = append(op.Middlewares, middlewares...)
	}
}
