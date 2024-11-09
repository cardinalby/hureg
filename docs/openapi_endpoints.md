# OpenAPI endpoints

## OpenAPI endpoints in Huma

Huma generates OpenAPI endpoints by default:
- **OpenAPI spec** endpoints (if `Config.OpenAPIPath` is set):
  - `/openapi.yaml`
  - `/openapi.json`
  - `/openapi-3.0.json`
  - `/openapi-3.0.yaml`
- **Docs** endpoint returning HTML page that uses data from `/openapi.yaml` to render docs (if `Config.DocsPath` is set):
  - `/docs`
- Schema endpoint (if `Config.SchemaPath` is set):
  - `/schemas`

### The limitations

- These endpoints are registered directly in `Adapter` and you can't add middlewares or transformers
  to them.
- For example, if you want to add any authentication to protect spec endpoints the only way is to use
  the underlying router middlewares
- You can't control which spec versions to expose and their paths (only the path prefix)

### Hureg solution

Hureg provides [oapi_handlers](./../pkg/huma/oapi_handlers) package with functions that return 
handlers for OpenAPI endpoints. They repeat the behavior of the default Huma OpenAPI endpoints.

You can use these handlers to register OpenAPI endpoints manually having full control over them.

#### Example

Let's add basic auth to OpenAPI endpoints.

**Step 1**: define a group for convenience that
- makes operation in it hidden
- adds basic auth middleware

```go
import (
	"github.com/cardinalby/hureg/pkg/huma/middlewares"
    "github.com/cardinalby/hureg/pkg/huma/oapi_handlers"
)

basicAuthMiddleware := middlewares.BasicAuth(
	func(ctx huma.Context, username, password string) (huma.Context, bool) {
		return ctx, username == "test" && password == "test"
	},
	"Hello!",
)

oapiGroup := api.  // api is APIGen instance
	AddOpHandler(op_handler.SetHidden(true, true)).
	AddMiddlewares(basicAuthMiddleware)
```

**Step 2**: register OpenAPI endpoints manually using library-provided handlers

```go
openAPI := api.GetHumaAPI().GetOpenAPI() // OpenAPI object

yaml31Handler, err := oapi_handlers.GetOpenAPISpecHandler(
    openAPI, oapi_handlers.OpenAPIVersion3dot1, oapi_handlers.OpenAPIFormatYAML,
)

docsHandler := oapi_handlers.GetDocsHandler(
	api.GetHumaAPI(), 
	"/openapi.yaml",     // path to the OpenAPI spec HTML page will request
)

schemaHandler := oapi_handlers.GetSchemaHandler(openApi, "")

Get(oapiGroup, "/openapi.yaml", yaml31Handler)
Get(oapiGroup, "/docs", docsHandler)
Get(oapiGroup, "/schemas/{schemaPath}", schemaHandler)
```

### Advanced: separate spec for some endpoints

Sometimes we need to have a separate spec that contains only a subset of the operations. 

Add scoped `OpenAPI` object to your group:

```go
api := hureg.NewAPIGen(humaApi)
experimentalApi, experimentalOpenAPI := api.AddOwnOpenAPI(huma.DefaultConfig("experimental", "1.0.0"))
stableApi, stableOpenAPI := api.AddOwnOpenAPI(huma.DefaultConfig("stable", "1.0.0"))
```

`experimentalOpenAPI` and `stableOpenAPI` are `*huma.OpenAPI` objects bind to the corresponding `APIGen` instances.

All registered operations will be added to the main `api` instance and its `OpenAPI` object. 

But `experimentalOpenAPI` and `stableOpenAPI` will contain only operations registered with 
`experimentalApi` and `stableApi` correspondingly.

You can use the handlers from the previous example to expose them separately.


