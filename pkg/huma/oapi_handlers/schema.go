package oapi_handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

var rxSchema = regexp.MustCompile(`#/components/schemas/([^"]+)`)

type SchemaRequest struct {
	SchemaPath string `path:"schemaPath"`
}

// GetSchemaHandler returns a handler that will return a schema from the OpenAPI spec.
// The handler format is suitable for passing it to huma or hureg registration functions.
// `schemasPath` is a prefix to that will be added to all schema refs in the returned schema.
func GetSchemaHandler(openAPI *huma.OpenAPI, schemasPath string) StreamResponseHandler[*SchemaRequest] {
	return func(ctx context.Context, req *SchemaRequest) (*huma.StreamResponse, error) {
		// Some routers dislike a path param+suffix, so we strip it here instead.
		schemaName := strings.TrimSuffix(req.SchemaPath, ".json")
		schema, ok := openAPI.Components.Schemas.Map()[schemaName]
		if !ok {
			return nil, huma.Error404NotFound(fmt.Sprintf("schema '%q' not found", schemaName))
		}
		schemaBytes, err := json.Marshal(schema)
		if err != nil {
			return nil, huma.Error500InternalServerError("failed to marshal schema", err)
		}
		schemaBytes = rxSchema.ReplaceAll(schemaBytes, []byte(path.Join(schemasPath, `/$1.json`)))

		return &huma.StreamResponse{
			Body: func(ctx huma.Context) {
				ctx.SetHeader("Content-Type", "application/json")
				_, _ = ctx.BodyWriter().Write(schemaBytes)
			},
		}, nil
	}
}
