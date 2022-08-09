package middleware

import (
	"context"
	"net/http"

	"github.com/thanishsid/goserver/graphql/cookiemanager"
)

// Middleware to load cookie manager into the context and set any cookies in the cookie stack
// at the end of the handler chain.
func LoadCookieManager(cfg cookiemanager.CookieConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mgr := cookiemanager.NewCookieManager(cfg, w)
			ctxWithMgr := context.WithValue(r.Context(), cookiemanager.CTX_KEY, mgr)
			next.ServeHTTP(w, r.WithContext(ctxWithMgr))
		})
	}
}
