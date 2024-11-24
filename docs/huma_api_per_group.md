# Use a dedicated `huma.API` for a group

## Recap

HUMA doesn't allow you to use [router-specific middlewares](https://huma.rocks/features/middleware/) for 
individual operations, you can only assign them to your router before creating `huma.API` instance. 
Therefore, they will be applied to all operations registered with this `huma.API` instance.

This way, HUMA pushes you towards using "router-agnostic" own middleware format.

## The problem
However, sometimes you already have some middlewares compatible with your router and want to re-use them.

## The solution

Hureg allows you to define a dedicated `huma.API` instance (based on router's group) for a group of operations
using `APIGen.ReplaceHumaAPI()` method.

### Example

Let's say we have a simple setup with **chi router** and "v1" group defined using Hureg:
```go
mux := chi.NewRouter()

cfg := huma.DefaultConfig("My API", "1.0.0")
humaApi := humachi.New(mux, cfg)
api := hureg.NewAPIGen(humaApi)

v1api := api.AddBasePath("/v1")
hureg.Get(v1api, "/cat", catHandler)    // GET /v1/cat
```

Now we want to add **"dog"** endpoint that will use chi's `middleware.Logger`:
```go
loggedMux := mux.With(middleware.Logger)       // chi group with logger middleware

loggedCfg := cfg            // The key point here is using the same OpenAPI pointer from the original cfg
loggedCfg.OpenAPIPath = ""  // Don't overwrite OpenAPI endpoints defined in the original huma.API
loggedCfg.DocsPath = ""     
loggedCfg.SchemasPath = ""

// Create a separate huma.API instance based on chi group.
loggedHumaApi := humachi.New(loggedMux, loggedCfg)

// Create a new inherited APIGen instance (that already has '/v1' prefix from `api`) with 
// the dedicated huma.API (that uses chi's logger middleware) 
loggedApi := api.ReplaceHumaAPI(loggedHumaApi)  

hureg.Get(loggedApi, "/dog", dogHandler)  // GET /v1/dog   (with chi's logger middleware)
```

The trick works because both `huma.API` instances share the same OpenAPI pointer, so both `/v1/cat` and `/v1/dog`
will appear in the same OpenAPI spec that is served by the original `humaApi`.

## See also

- [OpenAPI endpoints](./openapi_endpoints.md) and multiple scoped specs