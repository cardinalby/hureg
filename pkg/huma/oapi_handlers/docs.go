package oapi_handlers

import "github.com/danielgtaylor/huma/v2"

func getOpenApiTitle(humaApi huma.API) string {
	openApi := humaApi.OpenAPI()
	if openApi == nil {
		return ""
	}
	if openApi.Info == nil {
		return ""
	}
	return openApi.Info.Title
}

// GetDocsAdapterHandler returns a handler that will return HTML page that renders OpenAPI spec from the specified
// `openAPIYamlPath` URL.
// The handler format is suitable for passing it directly
// into Adapter.Handle() method or using with huma.StreamResponse
func GetDocsAdapterHandler(humaApi huma.API, openAPIYamlPath string) func(ctx huma.Context) {
	return func(ctx huma.Context) {
		ctx.SetHeader("Content-Type", "text/html")
		title := "Elements in HTML"

		if oaTitle := getOpenApiTitle(humaApi); oaTitle != "" {
			title = oaTitle + " Reference"
		}
		_, _ = ctx.BodyWriter().Write([]byte(`<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="referrer" content="same-origin" />
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
    <title>` + title + `</title>
    <!-- Embed elements Elements via Web Component -->
    <link href="https://unpkg.com/@stoplight/elements@8.1.0/styles.min.css" rel="stylesheet" />
    <script src="https://unpkg.com/@stoplight/elements@8.1.0/web-components.min.js"
            integrity="sha256-985sDMZYbGa0LDS8jYmC4VbkVlh7DZ0TWejFv+raZII="
            crossorigin="anonymous"></script>
  </head>
  <body style="height: 100vh;">

    <elements-api
      apiDescriptionUrl="` + openAPIYamlPath + `"
      router="hash"
      layout="sidebar"
      tryItCredentialsPolicy="same-origin"
    />

  </body>
</html>`))
	}
}

// GetDocsTypedHandler returns a handler that will return HTML page that renders OpenAPI spec from the specified
// `openAPIYamlPath` URL.
// The handler format is suitable for passing it to huma or hureg registration functions.
func GetDocsTypedHandler(humaApi huma.API, openAPIYamlPath string) TypedStreamHandler {
	adapterHandler := GetDocsAdapterHandler(humaApi, openAPIYamlPath)
	return getTypedStreamHandler(adapterHandler)
}
