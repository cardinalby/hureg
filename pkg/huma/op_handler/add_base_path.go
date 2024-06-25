package op_handler

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
	"github.com/cardinalby/hureg/pkg/routepath"
)

// AddBasePath creates an OperationHandler that sets the base path (path prefix) for the operation.
// It updates operation`s Path and metadata.KeyBasePath.
func AddBasePath(basePath string) OperationHandler {
	return func(op *huma.Operation) {
		if basePath == "" || basePath == "/" {
			return
		}
		existingBasePath, hasInitPath := metadata.GetBasePath(op)
		if hasInitPath && existingBasePath != "" {
			existingBasePath = routepath.Join(existingBasePath, basePath)
		} else {
			existingBasePath = basePath
		}
		metadata.SetKey(op, metadata.KeyBasePath, existingBasePath)

		if initPath, hasInitPath := metadata.GetInitPath(op); hasInitPath {
			op.Path = routepath.Join(existingBasePath, initPath)
		} else {
			// Fallback, should not happen
			op.Path = existingBasePath
		}
	}
}
