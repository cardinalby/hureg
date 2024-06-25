package oapi_handlers

import (
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

// GetOpenAPIAdapterHandler returns a handler that will return OpenAPI spec in the requested format and version.
// The handler format is suitable for passing it directly
// into Adapter.Handle() method or using with huma.StreamResponse
func GetOpenAPIAdapterHandler(
	humaApi huma.API,
	version OpenAPIVersion,
	format OpenAPIFormat,
) (func(ctx huma.Context), error) {
	var contentType string
	switch format {
	case OpenAPIFormatJSON:
		contentType = "application/vnd.oai.openapi+json"
	case OpenAPIFormatYAML:
		contentType = "application/vnd.oai.openapi+yaml"
	default:
		return nil, fmt.Errorf("unsupported OpenAPI format: %s", format)
	}

	if version != OpenAPIVersion3dot0 && version != OpenAPIVersion3dot1 {
		return nil, fmt.Errorf("unsupported OpenAPI version: %s", version)
	}

	return func(ctx huma.Context) {
		ctx.SetHeader("Content-Type", contentType)

		var specBytes []byte
		var err error
		switch {
		case version == OpenAPIVersion3dot0 && format == OpenAPIFormatJSON:
			specBytes, err = humaApi.OpenAPI().Downgrade()
		case version == OpenAPIVersion3dot0 && format == OpenAPIFormatYAML:
			specBytes, err = humaApi.OpenAPI().DowngradeYAML()
		case version == OpenAPIVersion3dot1 && format == OpenAPIFormatJSON:
			specBytes, err = json.Marshal(humaApi.OpenAPI())
		case version == OpenAPIVersion3dot1 && format == OpenAPIFormatYAML:
			specBytes, err = humaApi.OpenAPI().YAML()
		}
		if err != nil {
			_ = huma.WriteErr(
				humaApi, ctx, 500, "Internal Server Error", huma.Error500InternalServerError(err.Error()),
			)
			return
		}
		_, _ = ctx.BodyWriter().Write(specBytes)
	}, nil
}

// GetOpenAPITypedHandler returns a handler that will return OpenAPI spec in the requested format and version.
// The handler format is suitable for passing it to huma or hureg registration functions.
func GetOpenAPITypedHandler(
	humaApi huma.API,
	version OpenAPIVersion,
	format OpenAPIFormat,
) (TypedStreamHandler, error) {
	adapterHandler, err := GetOpenAPIAdapterHandler(humaApi, version, format)
	if err != nil {
		return nil, err
	}
	return getTypedStreamHandler(adapterHandler), nil
}
