package hureg

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/huma/op_handler"
	"github.com/cardinalby/hureg/pkg/huma/operation"
)

// RegMiddleware is core concept of the library that provides flexibility in setting up the registration pipeline:
// using route groups, pre-specified middlewares and other properties of operations for an instance of API
// RegMiddlewares are functions that are called in chain (similar to normal HTTP middlewares) during the
// operation registration (before the operation will be registered in Huma).
// RegMiddleware can produce side effects or modify the operation and pass it to the next RegMiddleware in the chain.
// - Calling `next` more than one time will lead to multiple registrations of the same operation,
// see RegMiddlewares.FanOut method
// - By not calling `next` at all you can prevent the operation from being registered
type RegMiddleware func(op huma.Operation, next func(huma.Operation))

// NewRegMiddleware creates a new RegMiddleware that applies the given OperationHandlers to the operation.
func NewRegMiddleware(opHandlers ...op_handler.OperationHandler) RegMiddleware {
	return func(op huma.Operation, next func(huma.Operation)) {
		for _, oh := range opHandlers {
			oh(&op)
		}
		next(op)
	}
}

func (h RegMiddleware) wrap(next func(huma.Operation)) func(huma.Operation) {
	return func(op huma.Operation) {
		h(op, next)
	}
}

type RegMiddlewares []RegMiddleware

// Chain creates a new RegMiddleware that chains RegMiddlewares together returning a single RegMiddleware.
func (hs RegMiddlewares) Chain() RegMiddleware {
	return func(op huma.Operation, next func(huma.Operation)) {
		hs.chainWithRegisterer(next)(op)
	}
}

// FanOut creates a new RegMiddleware that calls all RegMiddlewares one by one passing the same `next` function
// to them and a copy of the received operation.
// It leads to "multiplication" of the registration pipeline, creating multiple branches of registration
// pipeline after this RegMiddleware.
// It can be used to register multiple operations in Huma out of one registration call.
// Normally you are not supposed to use it in your own registration pipelines since:
//   - it requires a special handling of OperationID in order not to get multiple operations with the same OperationID
//   - it uses internal operation cloning that is sufficient for the library-provided operation handlers
//     but may not be sufficient for your custom operation handlers that modify the nested structures of the operation.
//
// It's used to create multiple alternative base paths for the same operation.
func (hs RegMiddlewares) FanOut() RegMiddleware {
	return func(op huma.Operation, next func(huma.Operation)) {
		for _, regMiddleware := range hs {
			regMiddleware(operation.Clone(op), next)
		}
	}
}

// Handler creates a final function that accepts Operation, passes it to the chain
// of RegMiddlewares and then to the provided registerer function.
func (hs RegMiddlewares) Handler(registerer func(huma.Operation)) func(huma.Operation) {
	return hs.chainWithRegisterer(registerer)
}

func (hs RegMiddlewares) chainWithRegisterer(registerer func(huma.Operation)) func(huma.Operation) {
	if len(hs) == 0 {
		return registerer
	}

	w := hs[len(hs)-1].wrap(registerer)
	for i := len(hs) - 2; i >= 0; i-- {
		w = hs[i].wrap(w)
	}
	return w
}
