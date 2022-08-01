package security

import (
	"context"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/google/uuid"

	"github.com/thanishsid/goserver/domain"
)

const SESSION_COOKIE_USER_ID_KEY = "user-id"
const CONTEXT_USER_KEY = "user"

// Get user from the database using the id found in a session cookie and load it into the context.
func LoadCtxUserMiddleware(us domain.UserService, sm *scs.SessionManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userIDStr := sm.GetString(r.Context(), SESSION_COOKIE_USER_ID_KEY)
			if userIDStr == "" {
				next.ServeHTTP(w, r)
				return
			}

			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				http.Error(w, "invalid cookie", http.StatusForbidden)
				return
			}

			user, err := us.One(r.Context(), userID)
			if err != nil {
				http.Error(w, "invalid cookie", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), CONTEXT_USER_KEY, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Gets and returns the user object from context and an error if user does not exists in the context.
func GetUserFromCtx(ctx context.Context) (*domain.User, error) {
	user, ok := ctx.Value(CONTEXT_USER_KEY).(*domain.User)
	if !ok {
		return nil, ErrNotLoggedIn
	}

	return user, nil
}
