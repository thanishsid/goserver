package sessions

import (
	"context"
	"fmt"
	"time"

	scs "github.com/alexedwards/scs/v2"
	"github.com/go-redis/cache/v8"

	"github.com/thanishsid/goserver/infrastructure/rediscache"
)

type redisStore struct {
	prefix string
	rc     *rediscache.CacheStore
}

var _ scs.Store = (*redisStore)(nil)
var _ scs.IterableStore = (*redisStore)(nil)

// Find should return the data for a session token from the store. If the
// session token is not found or is expired, the found return value should
// be false (and the err return value should be nil). Similarly, tampered
// or malformed tokens should result in a found return value of false and a
// nil err value. The err return value should be used for system errors only.
func (s *redisStore) Find(token string) (b []byte, found bool, err error) {
	if err := s.rc.Cache.Get(context.Background(), s.prefix+token, &b); err != nil {
		if err == cache.ErrCacheMiss {
			return nil, false, nil
		}

		return nil, false, err
	}

	return b, true, nil
}

// Commit should add the session token and data to the store, with the given
// expiry time. If the session token already exists, then the data and
// expiry time should be overwritten.
func (s *redisStore) Commit(token string, b []byte, expiry time.Time) (err error) {
	if time.Now().After(expiry) {
		return fmt.Errorf("invalid expiry time, expiry time is before current time")
	}

	return s.rc.Cache.Set(&cache.Item{
		Ctx:   context.Background(),
		Key:   s.prefix + token,
		Value: b,
		TTL:   expiry.Sub(time.Now()),
	})
}

// Delete should remove the session token and corresponding data from the
// session store. If the token does not exist then Delete should be a no-op
// and return nil (not an error).
func (s *redisStore) Delete(token string) (err error) {
	return s.rc.Cache.Delete(context.Background(), s.prefix+token)

}

// All should return a map containing data for all active sessions (i.e.
// sessions which have not expired). The map key should be the session
// token and the map value should be the session data. If no active
// sessions exist this should return an empty (not nil) map.
func (s *redisStore) All() (map[string][]byte, error) {

	keys := s.rc.Client.Keys(context.Background(), s.prefix+"*")
	if keys.Err() != nil {
		return nil, keys.Err()
	}

	sessions := make(map[string][]byte)

	for _, key := range keys.Val() {
		token := key[len(s.prefix):]

		data, exists, err := s.Find(token)
		if err != nil {
			if err == cache.ErrCacheMiss {
				return nil, nil
			}

			return nil, err
		}

		if exists {
			sessions[token] = data
		}
	}

	return sessions, nil
}
