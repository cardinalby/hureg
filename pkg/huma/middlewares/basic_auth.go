package middlewares

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"

	"github.com/cardinalby/hureg/pkg/http_util"
)

// BasicAuth is a middleware that checks for basic auth credentials in the Authorization header.
// If the credentials are not present or are incorrect, it will return a 401 Unauthorized response.
// If the credentials are correct, it will call the next middleware in the chain.
// The authHandler function should return the next context and a boolean indicating if the request should be allowed.
// The basicRealm is the realm to use in the WWW-Authenticate header.
func BasicAuth(
	authHandler func(ctx huma.Context, username, password string) (nextCtx huma.Context, allow bool),
	basicRealm string,
) func(ctx huma.Context, next func(huma.Context)) {
	answerUnauthorized := func(ctx huma.Context) {
		ctx.SetHeader(
			"WWW-Authenticate",
			"Basic realm="+http_util.CleanQuotedString(basicRealm),
		)
		ctx.SetStatus(http.StatusUnauthorized)
	}

	return func(ctx huma.Context, next func(huma.Context)) {
		var username, password string

		headerValue := ctx.Header("Authorization")
		basicAuthPrefix := "Basic "

		if strings.HasPrefix(headerValue, basicAuthPrefix) {
			encodedCreds := headerValue[len(basicAuthPrefix):]
			creds, err := base64.StdEncoding.DecodeString(encodedCreds)
			if err != nil {
				answerUnauthorized(ctx)
				return
			}
			credsParts := strings.SplitN(string(creds), ":", 2)
			if len(credsParts) < 2 {
				answerUnauthorized(ctx)
				return
			}
			username, password = credsParts[0], credsParts[1]
		}
		authCtx, allow := authHandler(ctx, username, password)
		if authCtx == nil {
			ctx.SetStatus(http.StatusInternalServerError)
		}
		if !allow {
			answerUnauthorized(authCtx)
			return
		}

		next(authCtx)
	}
}
