package op_handler

import (
	"regexp"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/casing"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

var reRemoveIDs = regexp.MustCompile(`\{([^}]+)}`)

// UpdateOperationID creates an OperationHandler that re-generates the OperationID field on the operation.
// It does not re-generates OperationID if it's empty
// If OperationID is not explicitly set, it generates it from the path (in the same way as huma does it by default).
// If OperationID is explicitly set, it will be updated by the provided `explicitOpIDBuilder`.
// If `explicitOpIDBuilder` is nil, the following approach will be used for operations with explicitly set
// OperationID:
// It will take metadata.KeyBasePath as a prefix for the OperationID, turning it into kebab-case and
// will append it to the metadata.KeyInitOperationID
// It's used internally by the library after registration of multiple base paths to
// make OperationID unique for different base paths and normally should not be used directly.
func UpdateOperationID(
	explicitOpIDBuilder func(*huma.Operation) string,
) OperationHandler {
	return func(op *huma.Operation) {
		if op.OperationID == "" {
			// Should not update operation ID that is empty
			return
		}

		if !metadata.IsExplicitOperationID(op) {
			GenerateOperationID(true)(op)
			return
		}

		if explicitOpIDBuilder != nil {
			op.OperationID = explicitOpIDBuilder(op)
			return
		}

		prefix, hasBasePath := metadata.GetBasePath(op)
		if !hasBasePath || prefix == "" {
			// Fallback, should not normally happen
			prefix = op.Path
		}
		prefix = casing.Kebab(reRemoveIDs.ReplaceAllString(prefix, "by-$1"))

		initOperationID, hasInitOperationID := metadata.GetInitOperationID(op)
		// Normally hasInitOperationID should be true if an operation passed out registration pipeline.
		// initOperationID == "" is more valid situation, it means that explicit operation had an empty OperationID,
		// than it was assigned in one of registration middlewares. Use it to append to the prefix
		// in this case (we already checked that op.OperationID is not empty).
		if !hasInitOperationID || initOperationID == "" {
			initOperationID = op.OperationID
		}

		if prefix == "" {
			op.OperationID = initOperationID
		} else {
			op.OperationID = prefix + "-" + initOperationID
		}
	}
}
