package input

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/internal/security"
)

// Registration initialization input.
type InitRegistration struct {
	Username               string        `json:"username"`
	Email                  string        `json:"email"`
	FullName               string        `json:"fullName"`
	Role                   security.Role `json:"role"`
	PictureID              uuid.NullUUID `json:"pictureId"`
	ClientRegistrationLink string        `json:"clientRegistrationLink"`
}

func (f InitRegistration) Validate() error {
	return vd.ValidateStruct(&f,
		vd.Field(&f.Username, vd.Required),
		vd.Field(&f.Email, vd.Required, is.Email),
		vd.Field(&f.FullName, vd.Required),
		vd.Field(&f.Role),
		vd.Field(&f.ClientRegistrationLink, vd.Required, is.URL),
	)
}

// Registration Completion input.
type CompleteRegistration struct {
	RegistrationToken string `json:"registrationToken"`
	Password          string `json:"password"`
}

func (f CompleteRegistration) Validate() error {
	return vd.ValidateStruct(&f,
		vd.Field(&f.RegistrationToken, vd.Required),
		vd.Field(&f.Password, vd.Required),
	)
}

// Update user input.
type UserUpdate struct {
	Username  string        `json:"username"`
	FullName  string        `json:"fullName"`
	PictureID uuid.NullUUID `json:"pictureId"`
}

func (f UserUpdate) Validate() error {
	return vd.ValidateStruct(&f,
		vd.Field(&f.Username, vd.Required),
		vd.Field(&f.FullName, vd.Required),
	)
}

// User role change input.
type RoleChange struct {
	UserID uuid.UUID     `json:"userId"`
	Role   security.Role `json:"role"`
}

func (f RoleChange) Validate() error {
	return vd.ValidateStruct(&f,
		vd.Field(&f.UserID, is.UUIDv4),
		vd.Field(&f.Role, vd.By(func(_ interface{}) error {
			return f.Role.ValidateRole()
		})),
	)
}

// User Filteration input.
type UserFilter struct {
	UserFilterBase
	Cursor null.String `schema:"cursor"`
}

func (f UserFilter) GetFilterBaseFromCursor() (UserFilterBase, error) {
	var userFilter UserFilterBase

	if !f.Cursor.Valid {
		return userFilter, errors.New("unable to get filter from null cursor")
	}

	jsn := make([]byte, len(f.Cursor.String))

	_, err := base64.URLEncoding.Decode(jsn, []byte(f.Cursor.String))
	if err != nil {
		return userFilter, err
	}

	if err := json.Unmarshal(jsn, &userFilter); err != nil {
		return userFilter, err
	}

	return userFilter, nil
}

// User Filteration input base.
type UserFilterBase struct {
	Query       null.String `schema:"search" json:"query"`
	Role        null.String `schema:"role" json:"role"`
	ShowDeleted null.Bool   `schema:"showDeleted" json:"showDeleted"`
	Limit       null.Int    `schema:"limit" json:"limit"`

	// Non client params
	UpdatedAfter null.Time `schema:"-" json:"updatedAfter"`
}

func (f UserFilterBase) Validate() error {
	return vd.ValidateStruct(&f,
		vd.Field(&f.Role, vd.When(f.Role.Valid, vd.By(func(_ interface{}) error {
			role := security.Role(f.Role.ValueOrZero())
			return role.ValidateRole()
		}))),
	)
}

func (f UserFilterBase) CreateCursor() (string, error) {
	jsn, err := json.Marshal(f)
	if err != nil {
		return "", err
	}

	encString := base64.URLEncoding.EncodeToString(jsn)

	return encString, nil
}
