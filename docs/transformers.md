# Transformers

Huma has a concept of transformers that allow you to modify the response before it's sent to the client.

In Huma, _Transformers_ are defined in `Config` upon `huma.API` creation and get applied to all responses.

The library allows you to define transformers **per `APIGen` instance**:

```go
hureg.Get(api, "/mouse", ...)      // no transformers applied

trGr := api.AddTransformers(tr1)   // transformers will be applied only to the operations 
                                   // registered in this group

hureg.Get(trGr, "/crocodile", ...) // tr1 will be applied 
```

---

[OpenAPI endpoints →](./openapi_endpoints.md)