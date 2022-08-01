package domain

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/thanishsid/goserver/config"
)

type SessionService interface {
	// Create a new seesion and returns a session id.
	Create(ctx context.Context, userID uuid.UUID, userAgent string, data map[string]any) (SID, error)

	// Delete a session by id.
	Delete(ctx context.Context, id SID) error

	// Delete all sessions by UserID.
	DeleteAllByUserID(ctx context.Context, userID uuid.UUID) error

	// Get a session, takes in a session id and a fresh useragent string from client
	// if the useragent has changed from existing value it will be updated.
	// Everytime the get function is called on a valid session the AcessedAt property on
	// a session will be updated.
	Get(ctx context.Context, id SID, userAgent string) (*Session, error)

	// Get all sessions by userID sorted by accessedAt time.
	GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]*Session, error)

	// Add Data to a session by key and value.
	AddData(ctx context.Context, id SID, key string, value any) error

	// Remove data from a session by key.
	RemoveData(ctx context.Context, id SID, key string) error

	// Clear all data in a session.
	ClearData(ctx context.Context, id SID) error
}

type SID string

func NewSID(userID uuid.UUID) SID {
	return SID(fmt.Sprintf("%s#%s", userID.String(), uuid.New().String()))
}

func (s SID) GetUserID() (uuid.UUID, error) {
	idSplit := strings.Split(string(s), "#")
	userIDStr := idSplit[0]
	return uuid.Parse(userIDStr)
}

func (s SID) String() string {
	return string(s)
}

type Session struct {
	ID         SID            `json:"id"`
	UserID     uuid.UUID      `json:"userId"`
	UserAgent  string         `json:"userAgent"`
	CreatedAt  time.Time      `json:"createdAt"`
	AccessedAt time.Time      `json:"accessedAt"`
	Data       map[string]any `json:"data"`
}

// Get session from context.
func SessionFor(ctx context.Context) (*Session, error) {
	session, ok := ctx.Value(config.SESSION_KEY).(*Session)

	if !ok {
		return nil, errors.New("session not found in context")
	}

	return session, nil
}
