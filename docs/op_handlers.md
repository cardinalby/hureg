# Operation Handlers

## Recall

_Operation Handlers_ are used in Huma to modify an operation before registration.

```go
type OperationHandler func(op *huma.Operation)
```

Normally, they are passed to `huma.Get()`, `huma.Post()`, etc. functions as the last arguments.

## Hureg usage

With `hureg` you still can use them as usual, but you can also add them to the `APIGen` instance to apply
them to all operations registered with this instance.

```go
derivedApi := api.AddOpHandler(oh1, oh2)
```

When you do so, `oh1` and `oh2` will be converted to [_Registration Middlewares_](./reg_middlewares.md) 
and added to the new `APIGen` instance.

Normally, you don't need the full power of _Registration Middlewares_ and can use _Operation Handlers_ for
most of the cases where you need to modify an operation before registration but don't need to perform
multiple registration or prevent registration at all.

## Built-in Operation Handlers

The library provides the most commonly needed _Operation Handlers_ in the 
[`op_handlers`](./../pkg/huma/op_handler) package.

---

[Extended operation metadata →](./metadata.md)
