package graphql

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/cookiemanager"
)

// var decoder = schema.NewDecoder()

func LoadOauthRoutes(authService domain.AuthService) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/google-login", GoogleLoginHandler(authService))
	r.Get("/google-callback", GoogleCallbackHandler(authService))

	return r
}

func GoogleLoginHandler(authService domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "failed to parse form", http.StatusInternalServerError)
			return
		}

		input := domain.RedirectLinkInput{
			SuccessRedirect: r.FormValue("success-redirect"),
			FailRedirect:    r.FormValue("fail-redirect"),
		}

		redirectUrl, err := authService.GenerateGoogleUrl(r.Context(), input)
		if err != nil {
			http.Error(w, "login failed, "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
	}
}

func GoogleCallbackHandler(authService domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "failed to parse form", http.StatusInternalServerError)
			return
		}

		input := domain.OauthLoginInput{
			State:     r.FormValue("state"),
			AuthCode:  r.FormValue("code"),
			UserAgent: r.UserAgent(),
		}

		session, redirectUrl, err := authService.GoogleLogin(r.Context(), input)
		if err != nil {
			if redirectUrl != "" {
				http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
				return
			}

			http.Error(w, "login failed", http.StatusInternalServerError)
			return
		}

		cookieManager := cookiemanager.For(r.Context())

		cookieManager.SetCookie(config.SESSION_COOKIE_NAME, session.ID.String(), config.SESSION_TTL)

		fmt.Println(redirectUrl)

		http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
	}
}
