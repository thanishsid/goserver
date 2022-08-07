package config

type ContextKey string

const (
	SESSION_KEY        ContextKey = "session"
	COOKIE_MANAGER_KEY ContextKey = "cookie-manager"
	USERAGENT_KEY      ContextKey = "user-agent"
)
