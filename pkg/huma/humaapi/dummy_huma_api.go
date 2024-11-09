package humaapi

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

// NewDummyHumaAPI creates a new Huma API with dummy adapter to be used only for OpenAPI object generation
// combined with manually added endpoints for docs, schemas and spec.
func NewDummyHumaAPI(config huma.Config) huma.API {
	config.OpenAPIPath = ""
	config.DocsPath = ""
	config.SchemasPath = ""
	return huma.NewAPI(config, dummyAdapter{})
}

type dummyAdapter struct{}

func (d dummyAdapter) Handle(*huma.Operation, func(ctx huma.Context)) {
}

func (d dummyAdapter) ServeHTTP(http.ResponseWriter, *http.Request) {
}
