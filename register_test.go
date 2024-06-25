package hureg

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/stretchr/testify/require"

	"github.com/cardinalby/hureg/pkg/huma/metadata"
)

type testInputObj struct {
	Body string
}

type testOutputObj struct {
	Body string
}

type testHandler func(context.Context, *testInputObj) (*testOutputObj, error)

func testHumaRegistration(
	t *testing.T,
	apiGen APIGen,
	registrationFn func(handler testHandler),
	operationCheckFns ...func(op huma.Operation),
) {
	testHumaRegistererCallsNumber := 0
	defer func() {
		testHumaRegisterer = nil
	}()

	testHumaRegisterer = func(api huma.API, operation huma.Operation, handler any) {
		require.Equal(t, apiGen.GetHumaAPI(), api)
		expHandler, ok := handler.(func(context.Context, *testInputObj) (*testOutputObj, error))
		require.True(t, ok)
		testOut, err := expHandler(context.Background(), &testInputObj{})
		require.Nil(t, err)
		require.Equal(t, "test_out", testOut.Body)
		require.Equal(t, reflect.TypeOf(new(testInputObj)), reflect.TypeOf(operation.Metadata[metadata.KeyInputObjPtr]))
		require.Equal(t, reflect.TypeOf(new(testOutputObj)), reflect.TypeOf(operation.Metadata[metadata.KeyOutputObjPtr]))

		require.Less(t, testHumaRegistererCallsNumber, len(operationCheckFns))
		operationCheckFns[testHumaRegistererCallsNumber](operation)
		testHumaRegistererCallsNumber++
	}

	handler := func(ctx context.Context, _ *testInputObj) (*testOutputObj, error) {
		return &testOutputObj{Body: "test_out"}, nil
	}

	registrationFn(handler)
	require.Equal(t, len(operationCheckFns), testHumaRegistererCallsNumber)
}

func TestAPIGen_AddRegMiddleware(t *testing.T) {
	t.Run("set summary", func(t *testing.T) {
		api := newTestApiGen().
			AddRegMiddleware(func(op huma.Operation, next func(huma.Operation)) {
				op.Summary = "test_summary"
				next(op)
			})
		testHumaRegistration(
			t,
			api,
			func(handler testHandler) {
				Get(api, "/a", handler)
			},
			func(op huma.Operation) {
				require.Equal(t, http.MethodGet, op.Method)
				require.Equal(t, "test_summary", op.Summary)
			},
		)
	})

	t.Run("stop registration", func(t *testing.T) {
		api := newTestApiGen().
			AddRegMiddleware(func(_ huma.Operation, _ func(huma.Operation)) {
				// do nothing
			})
		testHumaRegistration(
			t,
			api,
			func(handler testHandler) {
				Get(api, "/a", handler)
			},
			// no calls expected
		)
	})
}

func TestAPIGen_AddOpHandler(t *testing.T) {
	api := newTestApiGen().
		AddOpHandler(
			func(op *huma.Operation) {
				op.Summary = "test_summary"
			},
			func(op *huma.Operation) {
				op.Summary += "2"
			},
		)

	testHumaRegistration(
		t,
		api,
		func(handler testHandler) {
			Delete(api, "/a", handler)
		},
		func(op huma.Operation) {
			require.Equal(t, http.MethodDelete, op.Method)
			require.Equal(t, "test_summary2", op.Summary)
		},
	)
}

func TestAPIGen_AddBasePath(t *testing.T) {
	api := newTestApiGen().
		AddBasePath("/a").
		AddBasePath("/b")

	testHumaRegistration(
		t,
		api,
		func(handler testHandler) {
			Get(api, "/c", handler)
		},
		func(op huma.Operation) {
			require.Equal(t, "get-a-b-c", op.OperationID)
			require.Equal(t, "Get a b c", op.Summary)
			require.Equal(t, "/a/b/c", op.Path)
			require.Equal(t, "/a/b", op.Metadata[metadata.KeyBasePath])
			require.Equal(t, "/c", op.Metadata[metadata.KeyInitPath])
			require.Equal(t, false, metadata.IsExplicitOperationID(&op))
			require.Equal(t, false, metadata.IsExplicitSummary(&op))
		},
	)
}

func TestAPIGen_AddMultiBasePaths(t *testing.T) {
	t.Run("default id builder, implicit op", func(t *testing.T) {
		api := newTestApiGen().
			AddBasePath("/a").
			AddMultiBasePaths(nil, "/b1", "/b2").
			AddBasePath("/c")

		testHumaRegistration(
			t,
			api,
			func(handler testHandler) {
				Get(api, "/d", handler)
			},
			func(op huma.Operation) {
				require.Equal(t, "get-a-b1-c-d", op.OperationID)
				require.Equal(t, "Get a b1 c d", op.Summary)
				require.Equal(t, "/a/b1/c/d", op.Path)
				require.Equal(t, "/a/b1/c", op.Metadata[metadata.KeyBasePath])
				require.Equal(t, "/d", op.Metadata[metadata.KeyInitPath])
				require.Equal(t, false, metadata.IsExplicitOperationID(&op))
				require.Equal(t, false, metadata.IsExplicitSummary(&op))
			},
			func(op huma.Operation) {
				require.Equal(t, "get-a-b2-c-d", op.OperationID)
				require.Equal(t, "Get a b2 c d", op.Summary)
				require.Equal(t, "/a/b2/c/d", op.Path)
				require.Equal(t, "/a/b2/c", op.Metadata[metadata.KeyBasePath])
				require.Equal(t, "/d", op.Metadata[metadata.KeyInitPath])
				require.Equal(t, false, metadata.IsExplicitOperationID(&op))
				require.Equal(t, false, metadata.IsExplicitSummary(&op))
			},
		)
	})

	t.Run("default id builder, explicit op", func(t *testing.T) {
		api := newTestApiGen().
			AddBasePath("/a").
			AddMultiBasePaths(nil, "/b1", "/b2").
			AddBasePath("/c")

		testHumaRegistration(
			t,
			api,
			func(handler testHandler) {
				Register(api, huma.Operation{
					Method:      http.MethodPost,
					Path:        "/d",
					OperationID: "explicit-op",
					Summary:     "Explicit op",
				}, handler)
			},
			func(op huma.Operation) {
				require.Equal(t, "a-b1-explicit-op", op.OperationID)
				require.Equal(t, "Explicit op", op.Summary)
				require.Equal(t, "/a/b1/c/d", op.Path)
				require.Equal(t, "/a/b1/c", op.Metadata[metadata.KeyBasePath])
				require.Equal(t, "/d", op.Metadata[metadata.KeyInitPath])
				require.True(t, metadata.IsExplicitOperationID(&op))
				require.True(t, metadata.IsExplicitSummary(&op))
			},
			func(op huma.Operation) {
				require.Equal(t, "a-b2-explicit-op", op.OperationID)
				require.Equal(t, "Explicit op", op.Summary)
				require.Equal(t, "/a/b2/c/d", op.Path)
				require.Equal(t, "/a/b2/c", op.Metadata[metadata.KeyBasePath])
				require.Equal(t, "/d", op.Metadata[metadata.KeyInitPath])
				require.True(t, metadata.IsExplicitOperationID(&op))
				require.True(t, metadata.IsExplicitSummary(&op))
			},
		)
	})

	t.Run("custom id builder, explicit op", func(t *testing.T) {
		opIdBuilder := func(op *huma.Operation) string {
			return op.Metadata[metadata.KeyBasePath].(string) + "__" + op.Metadata[metadata.KeyInitOperationID].(string)
		}

		api := newTestApiGen().
			AddBasePath("/a").
			AddMultiBasePaths(opIdBuilder, "/b1", "/b2").
			AddBasePath("/c")

		testHumaRegistration(
			t,
			api,
			func(handler testHandler) {
				Register(api, huma.Operation{
					Method:      http.MethodPost,
					Path:        "/d",
					OperationID: "explicit-op",
					Summary:     "Explicit op",
				}, handler)
			},
			func(op huma.Operation) {
				require.Equal(t, "/a/b1__explicit-op", op.OperationID)
				require.Equal(t, "Explicit op", op.Summary)
				require.Equal(t, "/a/b1/c/d", op.Path)
				require.Equal(t, "/a/b1/c", op.Metadata[metadata.KeyBasePath])
				require.Equal(t, "/d", op.Metadata[metadata.KeyInitPath])
				require.True(t, metadata.IsExplicitOperationID(&op))
				require.True(t, metadata.IsExplicitSummary(&op))
			},
			func(op huma.Operation) {
				require.Equal(t, "/a/b2__explicit-op", op.OperationID)
				require.Equal(t, "Explicit op", op.Summary)
				require.Equal(t, "/a/b2/c/d", op.Path)
				require.Equal(t, "/a/b2/c", op.Metadata[metadata.KeyBasePath])
				require.Equal(t, "/d", op.Metadata[metadata.KeyInitPath])
				require.True(t, metadata.IsExplicitOperationID(&op))
				require.True(t, metadata.IsExplicitSummary(&op))
			},
		)
	})
}

func TestAPIGen_AddMiddlewares(t *testing.T) {
	var receivers []int
	api := newTestApiGen().
		AddMiddlewares(
			func(ctx huma.Context, next func(huma.Context)) {
				receivers = append(receivers, 2)
				next(ctx)
			},
			func(ctx huma.Context, next func(huma.Context)) {
				receivers = append(receivers, 3)
				next(ctx)
			},
		)

	testHumaRegistration(
		t,
		api,
		func(handler testHandler) {
			Register(
				api,
				huma.Operation{
					Method: http.MethodGet,
					Path:   "/a",
					Middlewares: huma.Middlewares{
						func(ctx huma.Context, next func(huma.Context)) {
							receivers = append(receivers, 1)
							next(ctx)
						},
					},
				},
				handler,
			)
		},
		func(op huma.Operation) {
			require.Equal(t, http.MethodGet, op.Method)
			require.Len(t, op.Middlewares, 3)
			for i, m := range op.Middlewares {
				m(nil, func(huma.Context) {})
				require.Equal(t, i+1, receivers[i])
			}
		},
	)
}

func TestInitOpMetadata(t *testing.T) {
	t.Parallel()
	type testInput struct {
	}
	type testOutput struct {
		Body string
	}
	var testInputPtr *testInput
	var testOutputPtr *testOutput

	testCommonMetadataKeys := func(t *testing.T, op huma.Operation) {
		require.Equal(t, "/test", op.Metadata[metadata.KeyInitPath])
		require.Equal(t, testInputPtr, op.Metadata[metadata.KeyInputObjPtr])
		require.Equal(t, testOutputPtr, op.Metadata[metadata.KeyOutputObjPtr])
		require.Equal(t, testInputPtr, op.Metadata[metadata.KeyInputObjPtr])
		require.Equal(t, testOutputPtr, op.Metadata[metadata.KeyOutputObjPtr])

		require.Equal(t, "test_api", metadata.GetOpenApiObj(&op).Info.Title)
		require.Equal(t, "1.0.1", metadata.GetOpenApiObj(&op).Info.Version)
	}

	t.Run("explicit op", func(t *testing.T) {
		t.Parallel()
		op := huma.Operation{
			Method: "GET",
			Path:   "/test",
		}
		initOpMetadata[testInput, testOutput](newTestApiGen().GetHumaAPI(), &op, true)
		require.Equal(t, "", op.OperationID)
		require.Equal(t, "", op.Summary)
		require.True(t, metadata.IsExplicitOperationID(&op))
		require.True(t, metadata.IsExplicitSummary(&op))
		require.Equal(t, "", op.Metadata[metadata.KeyInitOperationID])
		testCommonMetadataKeys(t, op)
	})

	t.Run("implicit op", func(t *testing.T) {
		t.Parallel()
		op := huma.Operation{
			Method: "GET",
			Path:   "/test",
		}
		initOpMetadata[testInput, testOutput](newTestApiGen().GetHumaAPI(), &op, false)
		require.Equal(t, "get-test", op.OperationID)
		require.Equal(t, "Get test", op.Summary)
		require.False(t, metadata.IsExplicitOperationID(&op))
		require.False(t, metadata.IsExplicitSummary(&op))
		require.Equal(t, "get-test", op.Metadata[metadata.KeyInitOperationID])
		testCommonMetadataKeys(t, op)
	})
}
