package middleware

import (
	"context"
	"net/http"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/cookiemanager"
	"github.com/thanishsid/goserver/service"
)

func LoadSessionMiddleware(sessionSvc domain.SessionService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the session cookie.
			sessionCookie, err := r.Cookie(config.SESSION_COOKIE_NAME)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			sessionID := domain.SID(sessionCookie.Value)

			// Get the session from redis cache.
			session, err := sessionSvc.Get(r.Context(), sessionID, r.UserAgent())
			if err != nil {

				// If session is not found in the cache then the session cookie is removed and next handler is invoked.
				if err == service.ErrNotFound {
					cookiemanager.For(r.Context()).RemoveCookie(config.SESSION_COOKIE_NAME)
				}

				next.ServeHTTP(w, r)
				return
			}

			// Create new context with existing request context and the user session.
			ctxWithSession := context.WithValue(r.Context(), domain.CTX_SESSION_KEY, session)

			// Invoke next handlers with the user session.
			next.ServeHTTP(w, r.WithContext(ctxWithSession))
		})
	}
}
