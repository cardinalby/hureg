package hureg

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"testing"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/cardinalby/hureg/pkg/huma/middlewares"
	"github.com/cardinalby/hureg/pkg/huma/oapi_handlers"
	"github.com/cardinalby/hureg/pkg/huma/op_handler"
)

func TestHttpServer(t *testing.T) {
	handler := createTestServer(t)
	addr, stop := listenAndServe(t, handler)
	defer stop()

	_ = addr
	testServerEndpoints(t, addr)
	testOpenApiSpec(t, addr)

	// uncomment to play with the server
	//waitSigInt(stop)
}

func createTestServer(t *testing.T) http.Handler {
	httpServeMux := http.NewServeMux()

	cfg := huma.DefaultConfig("My API", "1.0.0")
	cfg.OpenAPIPath = ""
	cfg.DocsPath = ""
	cfg.SchemasPath = ""

	humaApi := humago.New(httpServeMux, cfg)

	api := NewAPIGen(humaApi)

	defineAnimalEndpoints(api)

	apiWithBasicAuth := api.AddMiddlewares(newTestBasicAuthMiddleware())
	defineManualOpenApiEndpoints(t, apiWithBasicAuth)

	return httpServeMux
}

func defineAnimalEndpoints(api APIGen) {
	type testResponseDto struct {
		Body string
	}

	beasts := api.
		AddOpHandler(op_handler.AddTags("beasts")).
		AddTransformers(duplicateResponseStringTransformer)

	v1gr := beasts.AddBasePath("/v1")

	Get(v1gr, "/cat", func(ctx context.Context, _ *struct{}) (*testResponseDto, error) {
		return &testResponseDto{Body: "Meow"}, nil
	})

	v2gr := beasts.AddBasePath("/v2")

	Get(v2gr, "/dog", func(ctx context.Context, _ *struct{}) (*testResponseDto, error) {
		return &testResponseDto{Body: "Woof"}, nil
	})

	multiGr := api.
		AddOpHandler(op_handler.AddTags("birds")).
		AddMultiBasePaths(nil, "/v3", "/v4")

	Get(multiGr, "/sparrow", func(ctx context.Context, _ *struct{}) (*testResponseDto, error) {
		return &testResponseDto{Body: "Tweet"}, nil
	})
}

func duplicateResponseStringTransformer(_ huma.Context, _ string, v any) (any, error) {
	if str, ok := v.(string); ok {
		return str + str, nil
	}
	return v, nil
}

func newTestBasicAuthMiddleware() func(ctx huma.Context, next func(huma.Context)) {
	return middlewares.BasicAuth(
		func(ctx huma.Context, username, password string) (huma.Context, bool) {
			return ctx, username == "test" && password == "test"
		},
		"enter test:test",
	)
}

func defineManualOpenApiEndpoints(t *testing.T, api APIGen) {
	api = api.AddOpHandler(op_handler.SetHidden(true, true))
	humaApi := api.GetHumaAPI()

	yaml31Handler, err := oapi_handlers.GetOpenAPITypedHandler(
		humaApi, oapi_handlers.OpenAPIVersion3dot1, oapi_handlers.OpenAPIFormatYAML,
	)
	require.NoError(t, err)

	Get(api, "/openapi.yaml", yaml31Handler)
	Get(api, "/docs", oapi_handlers.GetDocsTypedHandler(humaApi, "/openapi.yaml"))
}

func testServerEndpoints(t *testing.T, addr string) {
	require.Equal(t, "MeowMeow", getStrResponse(t, addr, "/v1/cat"))
	require.Equal(t, "WoofWoof", getStrResponse(t, addr, "/v2/dog"))
	require.Equal(t, "Tweet", getStrResponse(t, addr, "/v3/sparrow"))
	require.Equal(t, "Tweet", getStrResponse(t, addr, "/v4/sparrow"))
}

func testOpenApiSpec(t *testing.T, addr string) {
	oa := getYamlResponse(t, addr, "/openapi.yaml", "test:test")
	require.Len(t, getMapsKey(oa, "paths"), 4)

	catOp := getMapsKey(oa, "paths", "/v1/cat", "get")
	require.Equal(t, "get-v1-cat", getMapsKey(catOp, "operationId"))
	require.Equal(t, "Get v1 cat", getMapsKey(catOp, "summary"))
	require.Equal(t, []any{"beasts"}, getMapsKey(catOp, "tags"))

	dogOp := getMapsKey(oa, "paths", "/v2/dog", "get")
	require.Equal(t, "get-v2-dog", getMapsKey(dogOp, "operationId"))
	require.Equal(t, "Get v2 dog", getMapsKey(dogOp, "summary"))
	require.Equal(t, []any{"beasts"}, getMapsKey(dogOp, "tags"))

	v3sparrowOp := getMapsKey(oa, "paths", "/v3/sparrow", "get")
	require.Equal(t, "get-v3-sparrow", getMapsKey(v3sparrowOp, "operationId"))
	require.Equal(t, "Get v3 sparrow", getMapsKey(v3sparrowOp, "summary"))
	require.Equal(t, []any{"birds"}, getMapsKey(v3sparrowOp, "tags"))

	v4sparrowOp := getMapsKey(oa, "paths", "/v4/sparrow", "get")
	require.Equal(t, "get-v4-sparrow", getMapsKey(v4sparrowOp, "operationId"))
	require.Equal(t, "Get v4 sparrow", getMapsKey(v4sparrowOp, "summary"))
	require.Equal(t, []any{"birds"}, getMapsKey(v4sparrowOp, "tags"))
}

func getMapsKey(data any, paths ...string) any {
	var ok bool
	for _, path := range paths {
		switch v := data.(type) {
		case map[string]any:
			data, ok = v[path]
			if !ok {
				return nil
			}
		default:
			return nil
		}
	}
	return data
}

func listenAndServe(t *testing.T, handler http.Handler) (addr string, stop func()) {
	addr, err := getFreePort(8080)
	require.NoError(t, err)

	server := &http.Server{Addr: addr, Handler: handler}
	go func() {
		t.Log("Starting server at", addr)
		if err := server.ListenAndServe(); err != nil {
			require.ErrorIs(t, err, http.ErrServerClosed)
		}
	}()

	stopCh := make(chan struct{})
	go func() {
		<-stopCh
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		require.NoError(t, err)
	}()

	var stopped atomic.Bool
	return addr, func() {
		if !stopped.Swap(true) {
			close(stopCh)
		}
	}
}

func getFreePort(desiredPort int) (addr string, err error) {
	localHost := "127.0.0.1"
	addr = fmt.Sprintf("%s:%d", localHost, desiredPort)
	ln, err := net.Listen("tcp", addr)
	if err == nil {
		return addr, ln.Close()
	}

	var a *net.TCPAddr
	if a, err = net.ResolveTCPAddr("tcp", localHost+":0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			port := l.Addr().(*net.TCPAddr).Port
			err = l.Close()
			return fmt.Sprintf("%s:%d", localHost, port), err
		}
	}
	return
}

func getStrResponse(t *testing.T, addr, path string) string {
	data := getBytesResponse(t, addr, path, nil)
	var strResp string
	require.NoError(t, json.Unmarshal(data, &strResp))
	return strResp
}

func getYamlResponse(t *testing.T, addr, path string, basicAuth string) any {
	data := getBytesResponse(t, addr, path, http.Header{
		"Authorization": {
			"Basic " + base64.StdEncoding.EncodeToString([]byte(basicAuth)),
		},
	})
	var resp any
	require.NoError(t, yaml.Unmarshal(data, &resp))
	return resp
}

func getBytesResponse(t *testing.T, addr, path string, headers http.Header) []byte {
	//goland:noinspection HttpUrlsUsage
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s%s", addr, path), nil)
	require.NoError(t, err)
	req.Header = headers
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer func() {
		require.NoError(t, resp.Body.Close())
	}()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return data
}

//goland:noinspection GoUnusedFunction
func waitSigInt(stop func()) {
	onInterrupt := make(chan os.Signal, 1)
	signal.Notify(onInterrupt, os.Interrupt)
	<-onInterrupt
	stop()
}
