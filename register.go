package hureg

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
	"github.com/cardinalby/hureg/pkg/huma/middlewares"
	"github.com/cardinalby/hureg/pkg/huma/op_handler"
)

var testHumaRegisterer func(api huma.API, operation huma.Operation, handler any)

// API is wrapper of huma.API that stores RegMiddlewares that should be applied to operations
// before registration in Huma.
type API interface {
	GetHumaAPI() huma.API
	GetRegMiddlewares() RegMiddlewares
	GetTransformers() []huma.Transformer
}

// Register registers a handler for an explicitly provided operation.
// It will apply all RegMiddlewares stored in API to the operation and passed operationHandlers
// after that. Depending on RegMiddlewares, an operation can be modified, registered multiple times with
// different base paths or not registered at all.
// Before passing an operation to the RegMiddlewares, it will be initialized with the metadata keys required
// for the library functionality.
// metadata.KeyIsExplicitOperationID, metadata.KeyIsExplicitSummary keys will be set to `true`
// See other keys in metadata package for details.
func Register[I, O any](
	api API,
	operation huma.Operation,
	handler func(context.Context, *I) (*O, error),
	operationHandlers ...func(o *huma.Operation),
) {
	registerImpl(api, &operation, true, handler, operationHandlers)
}

// Get is a shortcut for Register method that implicitly generates Operation object
// metadata.KeyIsExplicitOperationID, metadata.KeyIsExplicitSummary keys will not be set.
func Get[I, O any](
	api API,
	path string,
	handler func(context.Context, *I) (*O, error),
	operationHandlers ...func(o *huma.Operation),
) {
	convenience(api, http.MethodGet, path, handler, operationHandlers...)
}

// Post is a shortcut for Register method that implicitly generates Operation object
// metadata.KeyIsExplicitOperationID, metadata.KeyIsExplicitSummary keys will not be set.
//
//goland:noinspection GoUnusedExportedFunction
func Post[I, O any](
	api API,
	path string,
	handler func(context.Context, *I) (*O, error),
	operationHandlers ...func(o *huma.Operation),
) {
	convenience(api, http.MethodPost, path, handler, operationHandlers...)
}

// Put is a shortcut for Register method that implicitly generates Operation object
// metadata.KeyIsExplicitOperationID, metadata.KeyIsExplicitSummary keys will not be set.
//
//goland:noinspection GoUnusedExportedFunction
func Put[I, O any](
	api API,
	path string,
	handler func(context.Context, *I) (*O, error),
	operationHandlers ...func(o *huma.Operation),
) {
	convenience(api, http.MethodPut, path, handler, operationHandlers...)
}

// Patch is a shortcut for Register method that implicitly generates Operation object
// metadata.KeyIsExplicitOperationID, metadata.KeyIsExplicitSummary keys will not be set.
//
//goland:noinspection GoUnusedExportedFunction
func Patch[I, O any](
	api API,
	path string,
	handler func(context.Context, *I) (*O, error),
	operationHandlers ...func(o *huma.Operation),
) {
	convenience(api, http.MethodPatch, path, handler, operationHandlers...)
}

// Delete is a shortcut for Register method that implicitly generates Operation object
// metadata.KeyIsExplicitOperationID, metadata.KeyIsExplicitSummary keys will not be set.
//
//goland:noinspection GoUnusedExportedFunction
func Delete[I, O any](
	api API,
	path string,
	handler func(
		context.Context,
		*I,
	) (*O, error), operationHandlers ...func(o *huma.Operation)) {
	convenience(api, http.MethodDelete, path, handler, operationHandlers...)
}

// Options is a shortcut for Register method that implicitly generates Operation object
// metadata.KeyIsExplicitOperationID, metadata.KeyIsExplicitSummary keys will not be set.
//
//goland:noinspection GoUnusedExportedFunction
func Options[I, O any](
	api API,
	path string,
	handler func(
		context.Context,
		*I,
	) (*O, error), operationHandlers ...func(o *huma.Operation)) {
	convenience(api, http.MethodOptions, path, handler, operationHandlers...)
}

// Head is a shortcut for Register method that implicitly generates Operation object
// metadata.KeyIsExplicitOperationID, metadata.KeyIsExplicitSummary keys will not be set.
//
//goland:noinspection GoUnusedExportedFunction
func Head[I, O any](
	api API,
	path string,
	handler func(
		context.Context,
		*I,
	) (*O, error), operationHandlers ...func(o *huma.Operation)) {
	convenience(api, http.MethodHead, path, handler, operationHandlers...)
}

func convenience[I, O any](
	api API,
	method,
	path string,
	handler func(context.Context, *I) (*O, error),
	operationHandlers ...func(o *huma.Operation),
) {
	operation := huma.Operation{
		Method: method,
		Path:   path,
	}
	registerImpl(api, &operation, false, handler, operationHandlers)
}

func registerImpl[I, O any](
	api API,
	operation *huma.Operation,
	isExplicit bool,
	handler func(context.Context, *I) (*O, error),
	operationHandlers []func(o *huma.Operation),
) {
	humaAPI := api.GetHumaAPI()
	initOpMetadata[I, O](humaAPI, operation, isExplicit)

	humaRegister := func(op huma.Operation) {
		// Pass transformers saved in api to the ctx key, that will be later picked up by
		// humaApiWrapper.Transform method
		if apiTransformers := api.GetTransformers(); len(apiTransformers) > 0 {
			op.Middlewares = append(op.Middlewares, middlewares.SetCtxTransformers(apiTransformers))
		}

		op_handler.ApplyToOperationPtr(&op, operationHandlers...)
		if testHumaRegisterer == nil {
			huma.Register(humaAPI, op, handler)
		} else {
			testHumaRegisterer(humaAPI, op, handler)
		}
	}
	api.GetRegMiddlewares().Handler(humaRegister)(*operation)
}

func initOpMetadata[I, O any](humaApi huma.API, op *huma.Operation, isExplicit bool) {
	if op.Metadata == nil {
		op.Metadata = make(map[string]any)
	}
	op.Metadata[metadata.KeyOpenApiObj] = humaApi.OpenAPI()
	op.Metadata[metadata.KeyInitPath] = op.Path
	var input *I
	var output *O
	op.Metadata[metadata.KeyInputObjPtr] = input
	op.Metadata[metadata.KeyOutputObjPtr] = output
	if isExplicit {
		op.Metadata[metadata.KeyIsExplicitOperationID] = true
		op.Metadata[metadata.KeyIsExplicitSummary] = true
	} else {
		op_handler.GenerateOperationID(true)(op)
		op_handler.GenerateSummary(true)(op)
	}
	op.Metadata[metadata.KeyInitOperationID] = op.OperationID
}
