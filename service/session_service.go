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

type Session struct {
	*rediscache.CacheStore
}

var _ domain.SessionService = (*Session)(nil)

// Create a new seesion and returns a session id.
func (s *Session) Create(ctx context.Context, input domain.CreateSessionInput) (*domain.Session, error) {
	if err := input.Validate(); err != nil {
		return nil, nil
	}

	sid := domain.NewSID(input.UserID)

	now := time.Now()

	session := &domain.Session{
		ID:         sid,
		UserID:     input.UserID,
		UserRole:   input.UserRole,
		UserAgent:  input.UserAgent,
		CreatedAt:  now,
		AccessedAt: now,
		Data:       input.Data,
	}

	cacheItem := &cache.Item{
		Ctx:   ctx,
		Key:   sid.String(),
		Value: session,
		TTL:   config.SESSION_TTL,
	}

	if err := s.Cache.Set(cacheItem); err != nil {
		return nil, err
	}

	return session, nil
}

// Delete a session by id.
func (s *Session) Delete(ctx context.Context, id domain.SID) error {
	return s.Cache.Delete(ctx, id.String())
}

// Delete all sessions by UserID.
func (s *Session) DeleteAllByUserID(ctx context.Context, userID uuid.UUID) error {
	var keys []string

	iter := s.Client.Scan(ctx, 0, userID.String()+"*", 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	for _, key := range keys {
		if err := s.Cache.Delete(ctx, key); err != nil {
			return err
		}
	}

	return nil
}

// Update the role of in all sessions with the given userID.
func (s *Session) UpdateRoleByUserID(ctx context.Context, userID uuid.UUID, role domain.Role) error {
	sessions, err := s.GetAllByUserID(ctx, userID)
	if err != nil {
		return err
	}

	for _, session := range sessions {

		session.UserRole = role

		if err := s.Cache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   session.ID.String(),
			Value: session,
			TTL:   time.Since(session.CreatedAt.Add(config.SESSION_TTL)),
		}); err != nil {
			return err
		}
	}

	return nil
}

// Get a session, takes in a session id and a fresh useragent string from client
// if the useragent has changed from existing value it will be updated.
// Everytime the get function is called on a valid session the AcessedAt property on
// a session will be updated.
func (s *Session) Get(ctx context.Context, id domain.SID, userAgent string) (*domain.Session, error) {

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
			TTL:   time.Until(session.CreatedAt.Add(config.SESSION_TTL)),
		}); err != nil {
			return nil, fmt.Errorf("SessionService.Get: failed to set updated useragent: %w", err)
		}
	}

	return session, nil
}

// Get all sessions by userID sorted by accessedAt time.
func (s *Session) GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Session, error) {
	var keys []string

	iter := s.Client.Scan(ctx, 0, userID.String()+"*", 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	sessions := make([]*domain.Session, len(keys))

	for i, key := range keys {
		session := new(domain.Session)
		if err := s.Cache.Get(ctx, key, session); err != nil {
			return nil, err
		}
		sessions[i] = session
	}

	return sessions, nil
}

// Add Data to a session by key and value.
func (s *Session) AddData(ctx context.Context, id domain.SID, key string, value any) error {
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
		TTL:   time.Until(session.CreatedAt.Add(config.SESSION_TTL)),
	}); err != nil {
		return fmt.Errorf("SessionService.AddData: failed to set updated session: %w", err)
	}

	return nil
}

// Remove data from a session by key.
func (s *Session) RemoveData(ctx context.Context, id domain.SID, key string) error {
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
		TTL:   time.Until(session.CreatedAt.Add(config.SESSION_TTL)),
	}); err != nil {
		return fmt.Errorf("SessionService.RemoveData: failed to set updated session: %w", err)
	}

	return nil
}

// Clear all data in a session.
func (s *Session) ClearData(ctx context.Context, id domain.SID) error {
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
		TTL:   time.Until(session.CreatedAt.Add(config.SESSION_TTL)),
	}); err != nil {
		return fmt.Errorf("SessionService.ClearData: failed to set updated session: %w", err)
	}

	return nil
}
