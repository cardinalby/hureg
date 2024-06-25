## Metadata keys

`Metadata` is a field of `huma.Operation` (`map[string]any`) 
that is useful set arbitrary data assigned with the operation. 

`hureg` utilises it to make built-in functions and _Operation Handlers_ work, 
but you can find it useful for your custom 
[_Operation Handlers_](./op_handlers.md) and 
[_Registration Middlewares_](./reg_middlewares.md).

## Pre-defined keys

If you use `hureg` registration functions, every operation will have `Metadata` map filled with the 
library-specific keys. 

You can use this data in own _Operation Handlers_ to:
- Get input/output types of a registered handler
- Access OpenAPI object
- Find out if operation was defined explicitly (via `Register` function) or implicitly via convenience methods.

[Look at](./../pkg/huma/metadata/keys.go) `metadata` package for details.

You are not supposed to modify these keys directly, it can break the library functionality.

---

[Per-group Transformers →](./transformers.md)