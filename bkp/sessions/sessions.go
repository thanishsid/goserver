package sessions

import (
	"github.com/alexedwards/scs/v2"
	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/infrastructure/rediscache"
)

func NewSessionManager(cachStore *rediscache.CacheStore, prefix string) *scs.SessionManager {
	sm := scs.New()

	sm.Store = &redisStore{rc: cachStore, prefix: prefix}

	sm.Lifetime = config.SESSION_TTL
	sm.IdleTimeout = 0
	sm.Cookie.Name = config.SESSION_COOKIE_NAME
	sm.Cookie.HttpOnly = true
	sm.Cookie.Secure = config.C.Environment == "production"

	return sm
}
