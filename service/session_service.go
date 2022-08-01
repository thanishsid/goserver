package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/google/uuid"
	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/rediscache"
)

func NewSessionService(cacheStore *rediscache.CacheStore) domain.SessionService {
	return &sessionService{cacheStore}
}

type sessionService struct {
	*rediscache.CacheStore
}

var _ domain.SessionService = (*sessionService)(nil)

// Create a new seesion and returns a session id.
func (s *sessionService) Create(ctx context.Context, userID uuid.UUID, userAgent string, data map[string]any) (domain.SID, error) {

	sid := domain.NewSID(userID)

	now := time.Now()

	session := &domain.Session{
		ID:         sid,
		UserID:     userID,
		UserAgent:  userAgent,
		CreatedAt:  now,
		AccessedAt: now,
		Data:       data,
	}

	cacheItem := &cache.Item{
		Ctx:   ctx,
		Key:   sid.String(),
		Value: session,
		TTL:   config.SESSION_TTL,
	}

	if err := s.Cache.Set(cacheItem); err != nil {
		return "", nil
	}

	return sid, nil
}

// Delete a session by id.
func (s *sessionService) Delete(ctx context.Context, id domain.SID) error {
	return s.Cache.Delete(ctx, id.String())
}

// Delete all sessions by UserID.
func (s *sessionService) DeleteAllByUserID(ctx context.Context, userID uuid.UUID) error {

	keys, err := s.Client.Keys(ctx, userID.String()+"*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		if err := s.Cache.Delete(ctx, key); err != nil {
			return err
		}
	}

	return nil
}

// Get a session, takes in a session id and a fresh useragent string from client
// if the useragent has changed from existing value it will be updated.
// Everytime the get function is called on a valid session the AcessedAt property on
// a session will be updated.
func (s *sessionService) Get(ctx context.Context, id domain.SID, userAgent string) (*domain.Session, error) {

	session := new(domain.Session)

	if err := s.Cache.Get(ctx, id.String(), session); err != nil {
		if err == cache.ErrCacheMiss {
			return nil, ErrNotFound
		}

		return nil, err
	}

	if session.UserAgent != userAgent {
		session.UserAgent = userAgent

		if err := s.Cache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   id.String(),
			Value: session,
			TTL:   session.CreatedAt.Add(config.SESSION_TTL).Sub(time.Now()),
		}); err != nil {
			return nil, fmt.Errorf("SessionService.Get: failed to set updated useragent: %w", err)
		}
	}

	return session, nil
}

// Get all sessions by userID sorted by accessedAt time.
func (s *sessionService) GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Session, error) {

	keys, err := s.Client.Keys(ctx, userID.String()+"*").Result()
	if err != nil {
		return nil, err
	}

	sessions := make([]*domain.Session, len(keys))

	for _, key := range keys {
		session := new(domain.Session)
		if err := s.Cache.Get(ctx, key, session); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

// Add Data to a session by key and value.
func (s *sessionService) AddData(ctx context.Context, id domain.SID, key string, value any) error {
	session := new(domain.Session)

	if err := s.Cache.Get(ctx, id.String(), session); err != nil {
		if err == cache.ErrCacheMiss {
			return ErrNotFound
		}

		return err
	}

	if session.Data == nil {
		session.Data = make(map[string]any)
	}

	session.Data[key] = value

	if err := s.Cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   id.String(),
		Value: session,
		TTL:   session.CreatedAt.Add(config.SESSION_TTL).Sub(time.Now()),
	}); err != nil {
		return fmt.Errorf("SessionService.AddData: failed to set updated session: %w", err)
	}

	return nil
}

// Remove data from a session by key.
func (s *sessionService) RemoveData(ctx context.Context, id domain.SID, key string) error {
	session := new(domain.Session)

	if err := s.Cache.Get(ctx, id.String(), session); err != nil {
		if err == cache.ErrCacheMiss {
			return ErrNotFound
		}

		return err
	}

	delete(session.Data, key)

	if err := s.Cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   id.String(),
		Value: session,
		TTL:   session.CreatedAt.Add(config.SESSION_TTL).Sub(time.Now()),
	}); err != nil {
		return fmt.Errorf("SessionService.RemoveData: failed to set updated session: %w", err)
	}

	return nil
}

// Clear all data in a session.
func (s *sessionService) ClearData(ctx context.Context, id domain.SID) error {
	session := new(domain.Session)

	if err := s.Cache.Get(ctx, id.String(), session); err != nil {
		if err == cache.ErrCacheMiss {
			return ErrNotFound
		}

		return err
	}

	session.Data = nil

	if err := s.Cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   id.String(),
		Value: session,
		TTL:   session.CreatedAt.Add(config.SESSION_TTL).Sub(time.Now()),
	}); err != nil {
		return fmt.Errorf("SessionService.ClearData: failed to set updated session: %w", err)
	}

	return nil
}
