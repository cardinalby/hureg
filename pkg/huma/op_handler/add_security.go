package op_handler

import (
	"github.com/danielgtaylor/huma/v2"
)

// AddSecurity creates an OperationHandler that appends the given security entries to the operation's AddSecurity field.
func AddSecurity(securityEntries ...map[string][]string) OperationHandler {
	return func(op *huma.Operation) {
		if op.Security == nil {
			op.Security = make([]map[string][]string, 0)
		}
		op.Security = append(op.Security, securityEntries...)
	}
}
