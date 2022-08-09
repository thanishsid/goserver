package middleware

import (
	"context"
	"net/http"

	"github.com/thanishsid/goserver/config"
)

func LoadRequestMetadataMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctxWithUseragent := context.WithValue(r.Context(), config.CTX_USERAGENT_KEY, r.UserAgent())
			next.ServeHTTP(w, r.WithContext(ctxWithUseragent))
		})
	}
}
