package domain

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/thanishsid/goserver/config"
	"gopkg.in/guregu/null.v4"
)

type SessionService interface {
	// Create a new seesion and returns a session id.
	Create(ctx context.Context, input CreateSessionInput) (*Session, error)

	// Delete a session by id.
	Delete(ctx context.Context, id SID) error

	// Delete all sessions by UserID.
	DeleteAllByUserID(ctx context.Context, userID uuid.UUID) error

	// Update the role of in all sessions with the given userID.
	UpdateRoleByUserID(ctx context.Context, userID uuid.UUID, role Role) error

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

// SessionID type
type SID string

// Create a new sessionID by joining userID and a new UUID sessionID with a '#' between.
// this allows easy pattern search through redis keys for all sessions for a userID.
func NewSID(userID uuid.UUID) SID {
	return SID(fmt.Sprintf("%s#%s", userID.String(), uuid.New().String()))
}

// Get the userID from a sessionID.
func (s SID) GetUserID() (uuid.UUID, error) {
	idSplit := strings.Split(string(s), "#")
	userIDStr := idSplit[0]
	return uuid.Parse(userIDStr)
}

// Connvert SessionID to string.
func (s SID) String() string {
	return string(s)
}

type Session struct {
	ID              SID            `json:"id"`
	UserID          uuid.UUID      `json:"userId"`
	UserRole        Role           `json:"userRole"`
	UserAgent       string         `json:"userAgent"`
	ExternalPicture *string        `json:"externalPicture"`
	CreatedAt       time.Time      `json:"createdAt"`
	AccessedAt      time.Time      `json:"accessedAt"`
	Data            map[string]any `json:"data"`
}

// Check is userID is the owner of the session
func (s *Session) IsOwner(userID uuid.UUID) bool {
	return s.UserID == userID
}

// Get session from context.
func SessionFor(ctx context.Context) (*Session, error) {
	session, ok := ctx.Value(config.SESSION_KEY).(*Session)

	if !ok {
		return nil, errors.New("session not found in context")
	}

	return session, nil
}

//----- INPUTS --------

type CreateSessionInput struct {
	UserID          uuid.UUID
	UserRole        Role
	UserAgent       string
	ExternalPicture null.String
	Data            map[string]any
}

func (i CreateSessionInput) Validate() error {
	return vd.ValidateStruct(&i,
		vd.Field(&i.UserID, is.UUIDv4),
		vd.Field(&i.UserRole, IsRole),
		vd.Field(&i.UserAgent, vd.Required),
	)
}
