# Registration Middlewares

Similar to familiar middlewares that are used in request handling pipeline, 
the library introduces the ["**Registration Middleware**"](./../reg_middleware.go) concept.

```go
type RegMiddleware func(op huma.Operation, next func(huma.Operation))
```

_Registration Middlewares_ are stored in each `APIGen` instance and are chained and called upon registration of each
operation with the `APIGen` instance.

They are used **only during the registration process** and don't slow down request handling.

This powerful concept allows you to control the registration process flexibly:

- **Modify** operation properties before registration
- **Prevent** registration of some operations
- Register operation **multiple** times with different properties
- Produce **side effects** during registration (e.g. maintain own permission registry)

## Add a Registration Middleware

To create a new `APIGen` instance with given _Registration Middlewares_ attached:

```go
myRegMiddleware := func(op huma.Operation, next func(huma.Operation)) {
    // do something with the operation
    next(op)
}

myApi := api.AddRegMiddleware(myRegMiddleware)
```

## Simple operation modifications

For most of the cases when you need to modify an operation and produce some side effects, 
you can use _Operation Handlers_ (see [library provided ones](./../pkg/huma/op_handler)) 
and add them as _Registration Middlewares_.

Actually, when you do the following:
```go
myApi := api.AddOpHandler(op_handlers.AddTags("tag1", "tag2"))
```

It's the same as:
```go   
myRegMiddleware := func(op huma.Operation, next func(huma.Operation)) {
    op_handlers.AddTags("tag1", "tag2")(&op)
    next(op)
}
```

## Advanced: bubbling registration middlewares

If you want to see the changes made by all registration middlewares
in the chain (in derived `APIGen` instances) before an operation will be actually registered in Huma, here is the way:

```go
bubblingRegMiddleware := func(op huma.Operation, next func(huma.Operation)) {
    // do something with the operation
    next(op)
}
api = api.AddBubblingRegMiddleware(bubblingRegMiddleware)
derived = api.AddOpHandler(op_handlers.AddTags("tag1"))
```

In this example `bubblingRegMiddleware` will receive an operation registered with `derived` back being able
to observe added tags.

---
[Create a group with base path →](./base_path.md)
