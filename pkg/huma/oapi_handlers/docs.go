package oapi_handlers

import (
	"context"
	"embed"
	"html/template"

	"github.com/danielgtaylor/huma/v2"
)

//go:embed docs_page.gohtml
var tmplFS embed.FS

type pageData struct {
	Title           string
	OpenAPIYamlPath string
}

// GetDocsHandler returns a handler that will return HTML page that renders OpenAPI spec from the specified
// `openAPIYamlPath` URL.
// The handler format is suitable for passing it to huma or hureg registration functions.
func GetDocsHandler(openAPI *huma.OpenAPI, openAPIYamlPath string) StreamResponseHandler[*struct{}] {
	tmpl, err := template.ParseFS(tmplFS, "docs_page.gohtml")
	if err != nil {
		panic(err)
	}

	return func(ctx context.Context, _ *struct{}) (*huma.StreamResponse, error) {
		pageData := getPageData(openAPI, openAPIYamlPath)

		return &huma.StreamResponse{
			Body: func(ctx huma.Context) {
				ctx.SetHeader("Content-Type", "text/html")
				_ = tmpl.Execute(ctx.BodyWriter(), pageData)
			},
		}, nil
	}
}

func getPageData(openAPI *huma.OpenAPI, openAPIYamlPath string) (res pageData) {
	res.Title = "Elements in HTML"
	if openAPI.Info != nil && openAPI.Info.Title != "" {
		res.Title = openAPI.Info.Title + " Reference"
	}
	res.OpenAPIYamlPath = openAPIYamlPath
	return res
}
