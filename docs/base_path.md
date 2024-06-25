# Base path

One of the most common use cases for routing libraries is to define a base path for a group of endpoints.

Huma doesn't have this feature out of the box, requiring you to path the full operation path to 
the registration functions.

Hureg implement this feature with the help of [_Registration Middlewares_](./reg_middlewares.md).

Since the implementation involves some tricky parts, Hureg exposes dedicated methods of `APIGen` instance
to make it easier to use.

Some of the operation modifications require complex logic and exposed as dedicated `APIGen` methods:

### 🔻 `AddBasePath()`

Adds base path segment to the operations and takes care of regenerating operation's `OperationID` and `Summary` fields
for operations registered with convenience methods like `Get()`, `Post()`, etc.

> The method returns a new `APIGen` instance with the base path applied, the original instance is not modified.

OperationIDs of the operations registered with `Register()` method are not regenerated.

#### Example

```go
// all operations will be registered with "/v1" base path
v1gr := api.AddBasePath("/v1")

hureg.Get(v1gr, "/cat", ...) // "/v1/cat" 
hureg.Get(v1gr, "/dog", ...) // "/v1/dog"

// all operations will be registered with "/v1/rodents" base path
rodentsGr := api.AddBasePath("/rodents")

hureg.Get(rodentsGr, "/mouse", ...) // "/v1/rodents/mouse" 
hureg.Get(rodentsGr, "/rat", ...)   // "/v1/rodents/rat"
```

### 🔻 `AddMultiBasePaths()`

Allows you to register the same operation with multiple alternative base paths.

> The method returns a new `APIGen` instance with the specified multiple base paths applied, 
the original instance is not modified.

#### OperationID regeneration

Since it leads to multiple registration calls to Huma for the same operation, the library takes care of
regenerating unique `OperationID` and `Summary` fields for derived operations to avoid conflicts.

For implicit operations (registered via convenience functions like `Get()`, `Post()`, etc.) the library
just re-generates `OperationID` and `Summary` for each instance (with own base path) of the operation.

If the operation is set explicitly (via `Register()` method), the library uses a provided `explicitOpIDBuilder` 
from the first argument of `AddMultiBasePaths()` method. 

If `explicitOpIDBuilder` is `nil`, the default implementation is used - 
it adds base path prefix to the original OperationID.

#### Example

```go
multiGr := api.AddMultiBasePaths(nil, "/v1", "/v2")

hureg.Get(multiGr, "/sparrow", ...) // "/v1/sparrow"
                                    // "/v2/sparrow"
```

---

[Operation Handlers →](./op_handlers.md) 