package op_handler

import (
	"strconv"

	"github.com/danielgtaylor/huma/v2"
)

func SetResponse(statusCode int, response *huma.Response, override bool) OperationHandler {
	return func(op *huma.Operation) {
		strStatusCode := strconv.Itoa(statusCode)

		if op.Responses == nil {
			op.Responses = map[string]*huma.Response{}
		} else if _, hasKey := op.Responses[strStatusCode]; hasKey && !override {
			return
		}
		op.Responses[strStatusCode] = response
	}
}
