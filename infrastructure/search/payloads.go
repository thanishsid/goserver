package search

import (
	"time"

	"github.com/google/uuid"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/security"
	"gopkg.in/guregu/null.v4"
)

type UserPayload struct {
	ID        uuid.UUID     `json:"id"`
	Email     string        `json:"email"`
	Username  string        `json:"username"`
	FullName  string        `json:"fullName"`
	Role      security.Role `json:"role"`
	PictureID uuid.NullUUID `json:"-"`
	CreatedAt int64         `json:"createdAt"`
	UpdatedAt int64         `json:"updatedAt"`
	DeletedAt int64         `json:"deletedAt"`
}

func (u UserPayload) ToUser() domain.User {
	user := domain.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		FullName:  u.FullName,
		Role:      u.Role,
		PictureID: u.PictureID,
		CreatedAt: time.Unix(u.CreatedAt, 0),
		UpdatedAt: time.Unix(u.UpdatedAt, 0),
	}

	if u.DeletedAt > 0 {
		user.DeletedAt = null.TimeFrom(time.Unix(u.DeletedAt, 0))
	}

	return user
}

func (u *UserPayload) LoadFromUser(user *domain.User) {
	u.ID = user.ID
	u.Email = user.Email
	u.Username = user.Username
	u.FullName = user.FullName
	u.Role = user.Role
	u.PictureID = user.PictureID
	u.CreatedAt = user.CreatedAt.Unix()
	u.UpdatedAt = user.CreatedAt.Unix()

	if user.DeletedAt.Valid {
		u.DeletedAt = user.DeletedAt.ValueOrZero().Unix()
	}
}
