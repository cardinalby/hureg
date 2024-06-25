# Basic usage

## APIGen

The core of the library is [`APIGen`](./../api_gen.go) type.

It wraps HUMA's `API` instance together with custom _Registration middlewares_
(including _Operation Handlers_) and _Transformers_ that will be applied to all operations
registered with the particular instance of `APIGen`.

`APIGen` is immutable and all methods return a new instance of `APIGen` with the applied changes.

```go
import "github.com/cardinalby/hureg"

httpServeMux := http.NewServeMux()            // with Go 1.22 
cfg := huma.DefaultConfig("My API", "1.0.0")  // 
humaApi := humago.New(httpServeMux, cfg)      // It's default Huma initialization

api := hureg.NewAPIGen(humaApi)               // That's how APIGen instance is created

derived = api.AddBasePath("/v1")              // All operations registered with `derived` will have 
                                              // "/v1" base path
												
derived2 = derived.AddBasePath("/abc")        // All operations registered with `derived2` will have 
                                              // "/v1/abc" base path
```

- The main difference from Huma is that you use `APIGen` instance instead of `huma.API` in the registration functions.
- You can always get wrapped `huma.API` instance with `apiGen.GetHumaAPI()` method.

## Registration of operations

Similar to Huma, the library provides a set of [registration methods](../register.go):

- `Get()`,
- `Post()`
- ...
- `Register()`

The methods:
- Accept `APIGen` instance instead of `huma.API`
- Apply all the changes that are configured in the `APIGen` instance
- Finally, call the original Huma `Register` method with the modified operation.

Registration happens as usual, but you use `APIGen` instance and `hureg` package registration functions:

```go
hureg.Get(api, "/cat", catHandler)

hureg.Register(api, huma.Operation{...}, otherHandler)
```

---
[Registration Middlewares →](./reg_middlewares.md)