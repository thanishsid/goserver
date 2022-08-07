package cookiemanager

import (
	"context"
	"net/http"
	"time"

	"github.com/thanishsid/goserver/config"
)

func newCookieManager(config CookieConfig, rw http.ResponseWriter) *CookieManager {
	return &CookieManager{
		cfg: config,
		rw:  rw,
	}
}

type CookieConfig struct {
	HttpOnly bool
	Secure   bool
	SameSite http.SameSite
	Path     string
	Domain   string
}

type CookieManager struct {
	rw  http.ResponseWriter
	cfg CookieConfig
}

// Add a cookie into the cookie stack.
func (c *CookieManager) SetCookie(name string, value string, maxAge time.Duration) {

	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(maxAge.Seconds()),
		Expires:  time.Now().Add(maxAge),
		HttpOnly: c.cfg.HttpOnly,
		Secure:   c.cfg.Secure,
		SameSite: c.cfg.SameSite,
		Path:     c.cfg.Path,
		Domain:   c.cfg.Domain,
	}

	if err := cookie.Valid(); err != nil {
		panic("invalid cookie " + err.Error())
	}

	http.SetCookie(c.rw, cookie)
}

// Remove a cookie by name (will expire the cookie).
func (c *CookieManager) RemoveCookie(name string) {
	http.SetCookie(c.rw, &http.Cookie{
		Name:     name,
		Value:    "",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: c.cfg.HttpOnly,
		Secure:   c.cfg.Secure,
		SameSite: c.cfg.SameSite,
		Path:     c.cfg.Path,
		Domain:   c.cfg.Domain,
	})
}

// Middleware to load cookie manager into the context and set any cookies in the cookie stack
// at the end of the handler chain.
func LoadManager(cfg CookieConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mgr := newCookieManager(cfg, w)
			ctxWithMgr := context.WithValue(r.Context(), config.COOKIE_MANAGER_KEY, mgr)

			next.ServeHTTP(w, r.WithContext(ctxWithMgr))
		})
	}
}

// Get the cookie manager from context.
func For(ctx context.Context) *CookieManager {
	cookieMgr, ok := ctx.Value(config.COOKIE_MANAGER_KEY).(*CookieManager)
	if !ok {
		panic("cookie manager not found in context")
	}

	return cookieMgr
}
