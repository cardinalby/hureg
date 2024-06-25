## Operation Handlers package

### Introduction

This package provides a set of _Operation Handlers_ that plays well with `hureg` library.

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

hureg.Get(derivedApi, "/cat", catHandler)  // oh1 and oh2 will be applied to the operation
```

## Package contents

Package provides a set of _Operation Handlers_ that can be used both in registration function (`hureg.Get()`, 
`hureg.Post()`, ..., `hureg.Register()`) and in `APIGen.AddOpHandler()` method to apply them to all operations
registered with the `APIGen` instance.

Even though the handlers follow the standard Huma `func(op *huma.Operation)` signature, some of them require
operation [metadata keys](./../../../docs/metadata.md) specific to `hureg` library.

### Index

What can you do with provided _Operation Handlers_ is:

- [Add Security entries to the operation](./add_security.go)
- [Add Tags to the operation](./add_tags.go)
- [Generate operationID](./generate_operation_id.go) - can be useful for explicitly defined operations
- [Generate operation Summary](./generate_summary.go) - can be useful for explicitly defined operations
- [Conditionally apply other handlers](./if.go)
- [Add middlewares to the operation](./middlewares.go) - there is a shortcut in `ApiGEN` for this
- [Set BodyReadTimeout field](./set_body_read_timeout.go)
- [Set Deprecated field](./set_deprecated.go)
- [Set Description field](./set_description.go)
- [Set Extension key](./set_extensions_key.go)
- [Set ExternalDocs field](./set_external_docs.go)
- [Set Hidden field](./set_hidden.go)
- [Set MaxBodyBytes field](./set_max_body_bytes.go)
- [Set Metadata keys](./set_metadata_key.go)
- [Set Responses key](./set_response.go)
- [Set SkipValidateBody field](./set_skip_validate_body.go)
- [Set Summary field](./set_summary.go)
- [Update generated Summary field](./update_generated_summary.go) after operation modification