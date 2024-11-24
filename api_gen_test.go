package hureg

import (
	"net/http"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/stretchr/testify/require"

	"github.com/cardinalby/hureg/pkg/huma/op_handler"
)

func newTestApiGen() APIGen {
	humaAPI := humago.New(http.NewServeMux(), huma.DefaultConfig("test_api", "1.0.1"))
	return NewAPIGen(humaAPI)
}

func testRegMiddleware(
	t *testing.T,
	rm RegMiddleware,
	op huma.Operation,
	testFn func(huma.Operation)) {
	wasCalled := false
	rm(op, func(op huma.Operation) {
		wasCalled = true
		testFn(op)
	})
	require.True(t, wasCalled)
}

func TestAPIGen_GetHumaAPI(t *testing.T) {
	t.Parallel()
	cfg := huma.DefaultConfig("test_api", "1.0.1")
	humaAPI := humago.New(http.NewServeMux(), cfg)
	api := NewAPIGen(humaAPI).AddRegMiddleware(func(op huma.Operation, next func(huma.Operation)) {
		next(op)
	})
	require.Same(t, humaAPI, api.GetHumaAPI().(humaApiWrapper).API)
	require.Len(t, api.GetRegMiddlewares(), 1)

	cfg.OpenAPIPath = ""
	cfg.DocsPath = ""
	cfg.SchemasPath = ""
	humaAPI2 := humago.New(http.NewServeMux(), cfg)
	api2 := api.ReplaceHumaAPI(humaAPI2)
	require.NotSame(t, api.GetHumaAPI().(humaApiWrapper).API, api2.GetHumaAPI().(humaApiWrapper).API)
	require.Len(t, api.GetRegMiddlewares(), 1)
}

func TestAPIGen_GetRegMiddlewares(t *testing.T) {
	t.Parallel()
	api := newTestApiGen()
	rm1 := NewRegMiddleware(op_handler.SetSummary("a", true))
	rm2 := NewRegMiddleware(op_handler.SetSummary("b", true))
	require.Empty(t, api.GetRegMiddlewares())
	derived := api.AddRegMiddleware(rm1, rm2)
	require.Len(t, api.GetRegMiddlewares(), 0)

	resRegMiddlewares := derived.GetRegMiddlewares()
	require.Len(t, resRegMiddlewares, 2)
	testRegMiddleware(t, resRegMiddlewares[0], huma.Operation{}, func(op huma.Operation) {
		require.Equal(t, "a", op.Summary)
	})
	testRegMiddleware(t, resRegMiddlewares[1], huma.Operation{}, func(op huma.Operation) {
		require.Equal(t, "b", op.Summary)
	})
}
