package oapi_handlers

import (
	"encoding/json"
	"path"
	"regexp"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

var rxSchema = regexp.MustCompile(`#/components/schemas/([^"]+)`)

// GetSchemaAdapterHandler returns a handler that will return a schema from the OpenAPI spec.
// The handler format is suitable for passing it directly
// into Adapter.Handle() method or using with huma.StreamResponse
func GetSchemaAdapterHandler(humaApi huma.API, schemasPath string) func(ctx huma.Context) {
	return func(ctx huma.Context) {
		// Some routers dislike a path param+suffix, so we strip it here instead.
		schema := strings.TrimSuffix(ctx.Param("schema"), ".json")
		ctx.SetHeader("Content-Type", "application/json")
		openApi := humaApi.OpenAPI()
		b, _ := json.Marshal(openApi.Components.Schemas.Map()[schema])
		b = rxSchema.ReplaceAll(b, []byte(path.Join(schemasPath, `/$1.json`)))
		_, _ = ctx.BodyWriter().Write(b)
	}
}

// GetSchemaTypedHandler returns a handler that will return a schema from the OpenAPI spec.
// The handler format is suitable for passing it to huma or hureg registration functions.
func GetSchemaTypedHandler(humaApi huma.API, schemasPath string) TypedStreamHandler {
	adapterHandler := GetSchemaAdapterHandler(humaApi, schemasPath)
	return getTypedStreamHandler(adapterHandler)
}
