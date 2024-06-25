![hureg logo](./docs/hureg.png)

[HUMA](https://github.com/danielgtaylor/huma) is a great Go framework that enables you to 
expose generated OpenAPI spec in the best way possible. Unfortunately, it lacks some features from other routers.

This library wraps [HUMA framework](https://github.com/danielgtaylor/huma) endpoints 
registration pipeline to provide the missing features:

### ❤️ Create registration **groups**

Similar to other routers you can create a derived `api` (i.e. group) that has pre-defined:
- [**Base path**](./docs/base_path.md) (same as `Group`, `Route` methods in other routers) 
- [**Multiple**](./docs/base_path.md) alternative **base paths**
- [**Middlewares**](./pkg/huma/op_handler/middlewares.go)
- [**Transformers**](./docs/transformers.md)
- [**Tags**](./pkg/huma/op_handler/add_tags.go) and [other](./pkg/huma/op_handler) Huma Operation properties 
  that will be applied to all endpoints in a group.
- [**Control**](./docs/reg_middlewares.md) the registration pipeline preventing operation from 
  registration or registering it multiple times with different properties

### ❤️ Control over OpenAPI endpoints

Now you have [manual control](./docs/openapi_endpoints.md) over exposing the spec, docs and schemas:
- Expose only needed spec versions
- Add own middlewares (e.g. authentication to protect the spec on public APIs)

### ❤️ Access more metadata in Operation Handlers

The library [provides](./docs/metadata.md) additional information via `Metadata` field to your 
own _Operation Handlers_:
- Input/Output types of a handler
- OpenAPI object from `huma.API` instance
- Whether an operation was defined explicitly or implicitly via convenience methods
- etc.

## Installation

```bash
go get github.com/cardinalby/hureg
```

## Documentation

### Key concepts

- [Basic usage](./docs/basic_usage.md)
- [Registration Middlewares](./docs/reg_middlewares.md)

### Common use-cases

- [Create a group with base path](./docs/base_path.md)
- [Operation Handlers](./docs/op_handlers.md)

### Additional features

- [Operation metadata](./docs/metadata.md)
- [Per-group Transformers](./docs/transformers.md)
- [OpenAPI endpoints](./docs/openapi_endpoints.md)

## Examples

### 🔻 Initialization

```go
import "github.com/cardinalby/hureg"

chiRouter := chi.NewRouter()                    // --
cfg := huma.DefaultConfig("My API", "1.0.0")    // default HUMA initialization
humaApi := humachi.New(chiRouter, cfg)          // --

api := hureg.NewAPIGen(humaApi)    // The new line
```

### 🔻 "Base path + tags + middlewares" group

```go
v1gr := api.            // all operations registered with v1gr will have:
	AddBasePath("/v1").                            // - "/v1" base path
	AddOpHandler(op_handler.AddTags("some_tag")).  // - "some_tag" tag
	AddMiddlewares(m1, m2)                         // - m1, m2 middlewares
	
hureg.Get(v1gr, "/cat", ...) // "/v1/cat" with "some_tag" tag and m1, m2 middlewares
hureg.Get(v1gr, "/dog", ...) // "/v1/dog" with "some_tag" tag and m1, m2 middlewares
```

### 🔻 Multiple base paths

Sometimes we need to register the same endpoint with multiple base paths (e.g. `/v1` and `/v2`).

```go
multiPathGr := api.AddMultiBasePaths(nil, "/v1", "/v2")

hureg.Get(multiPathGr, "/sparrow", ...) // "/v1/sparrow"
                                        // "/v2/sparrow"
```

### 🔻 Transformers per group

```go
trGr := api.AddTransformers(...) // transformers will be applied only to the operations 
                                 // registered in this group

hureg.Get(trGr, "/crocodile", ...)
```

### 🔻 Complete server setup

Check out [integration_test.go](./integration_test.go) for a complete example of how to use the library:
- create `huma.API` from `chi` router
- create `APIGen` instance on top of `huma.API`
- register operations with `APIGen` instance
  - use base paths, tags and _Transformers_ to the groups
  - register OpenAPI endpoints manually with Basic Auth middleware

Uncommenting one line you can run the server and play with it in live mode.