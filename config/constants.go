package config

import "time"

const (
	SESSION_COOKIE_NAME = "session-token"

	REGISTRATION_TOKEN_TTL = time.Hour * 24 * 30
	SESSION_TTL            = time.Hour * 24 * 30

	DEFAULT_USERS_LIST_LIMIT = 40
)
