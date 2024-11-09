package oapi_handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

type OpenAPIVersion string

const (
	OpenAPIVersion3dot0 OpenAPIVersion = "3.0"
	OpenAPIVersion3dot1 OpenAPIVersion = "3.1"
)

type OpenAPIFormat string

const (
	OpenAPIFormatJSON OpenAPIFormat = "json"
	OpenAPIFormatYAML OpenAPIFormat = "yaml"
)

const openApiJsonContentType = "application/vnd.oai.openapi+json"
const openApiYamlContentType = "application/vnd.oai.openapi+yaml"

// GetOpenAPISpecHandler returns a handler that will return OpenAPI spec in the requested format and version.
// The handler format is suitable for passing it to huma or hureg registration functions
func GetOpenAPISpecHandler(
	openApi *huma.OpenAPI,
	version OpenAPIVersion,
	format OpenAPIFormat,
) (StreamResponseHandler[*struct{}], error) {
	marshaller, contentType, err := GetOpenAPISpecMarshaller(openApi, version, format)
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context, _ *struct{}) (*huma.StreamResponse, error) {
		specBytes, err := marshaller()
		if err != nil {
			return nil, huma.Error500InternalServerError("failed to generate spec", err)
		}

		return &huma.StreamResponse{
			Body: func(ctx huma.Context) {
				ctx.SetHeader("Content-Type", contentType)
				_, _ = ctx.BodyWriter().Write(specBytes)
			},
		}, nil
	}, nil
}

// GetOpenAPISpecMarshaller returns a marshaller function that will return OpenAPI spec
// in the requested format and version.
func GetOpenAPISpecMarshaller(
	openApi *huma.OpenAPI,
	version OpenAPIVersion,
	format OpenAPIFormat,
) (marshaller func() ([]byte, error), contentType string, err error) {
	switch format {
	case OpenAPIFormatJSON:
		contentType = openApiJsonContentType
	case OpenAPIFormatYAML:
		contentType = openApiYamlContentType
	default:
		return nil, "", fmt.Errorf("unsupported OpenAPI format: %s", format)
	}

	switch version {
	case OpenAPIVersion3dot0:
		if format == OpenAPIFormatJSON {
			marshaller = openApi.Downgrade
		} else {
			marshaller = openApi.DowngradeYAML
		}

	case OpenAPIVersion3dot1:
		if format == OpenAPIFormatJSON {
			marshaller = func() ([]byte, error) {
				return json.Marshal(openApi)
			}
		} else {
			marshaller = openApi.YAML
		}
	}

	return marshaller, contentType, err
}
