package metadata

// This file contains keys for metadata in huma.Operation.Metadata map that are used by the `hureg` library

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/internal"
)

// KeyInputObjPtr stores handler's input object pointer (typed nil)
const KeyInputObjPtr = internal.PkgID + "input_obj_ptr"

// KeyOutputObjPtr stores handler's output object pointer (typed nil)
const KeyOutputObjPtr = internal.PkgID + "output_obj_ptr"

// KeyIsExplicitOperationID stores `true` if the operation ID is explicitly set by the user:
// - operation was passed to Register method, not to a shortcut method
const KeyIsExplicitOperationID = internal.PkgID + "is_explicit_operation_id"

// KeyIsExplicitSummary stores `true` if the summary is explicitly set by the user:
// - operation was passed to Register method, not to a shortcut method OR
// - summary was set in the registration pipeline via ophanlder.Summary
// If it's `true`, the summary will not be implicitly modified by other library functions
const KeyIsExplicitSummary = internal.PkgID + "is_explicit_summary"

// KeyInitOperationID stores the initial operation ID (before any modifications in the registration pipeline)
// - for operations passed to Register method it's OperationID provided by the user (even if empty)
// - for operations passed to a shortcut method it's generated OperationID
const KeyInitOperationID = internal.PkgID + "init_operation_id"

// KeyInitPath stores the initial path of the operation specified by the user
// (before any modifications in the registration pipeline)
const KeyInitPath = internal.PkgID + "init_path"

// KeyBasePath stores the base path of the operation made of joined base paths added by BasePath operation handler
// calls. Normally, operation Path should be equal to joined KeyBasePath + KeyInitPath
const KeyBasePath = internal.PkgID + "base_path"

const KeyOpenApiObj = internal.PkgID + "openapi_obj"

// IsExplicitOperationID is a shortcut function to check if the operation ID is explicitly set by the user
func IsExplicitOperationID(op *huma.Operation) bool {
	isExplicit, hasBool := op.Metadata[KeyIsExplicitOperationID].(bool)
	return hasBool && isExplicit
}

// IsExplicitSummary is a shortcut function to check if the summary is explicitly set by the user
func IsExplicitSummary(op *huma.Operation) bool {
	isExplicit, hasBool := op.Metadata[KeyIsExplicitSummary].(bool)
	return hasBool && isExplicit
}

// GetInitOperationID reads KeyInitOperationID from the operation metadata
func GetInitOperationID(op *huma.Operation) (string, bool) {
	initOperationID, hasID := op.Metadata[KeyInitOperationID].(string)
	return initOperationID, hasID
}

// GetInitPath reads KeyInitPath from the operation metadata
func GetInitPath(op *huma.Operation) (string, bool) {
	initPath, hasPath := op.Metadata[KeyInitPath].(string)
	return initPath, hasPath
}

// GetBasePath reads KeyBasePath from the operation metadata
func GetBasePath(op *huma.Operation) (string, bool) {
	basePath, hasPath := op.Metadata[KeyBasePath].(string)
	return basePath, hasPath
}

// GetOpenApiObj reads KeyOpenApiObj from the operation metadata
func GetOpenApiObj(op *huma.Operation) *huma.OpenAPI {
	if openApiObj, hasObj := op.Metadata[KeyOpenApiObj].(*huma.OpenAPI); hasObj {
		return openApiObj
	}
	return nil
}
