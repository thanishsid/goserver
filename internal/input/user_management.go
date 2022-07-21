package input

import (
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
	return nil
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
	Search      null.String   `schema:"search"`
	Role        security.Role `schema:"role"`
	ShowDeleted null.Bool     `schema:"showDeleted"`
	Limit       null.Int      `schema:"limit"`
	Page        null.Int      `schema:"page"`
}

func (f UserFilter) Validate() error {
	return vd.ValidateStruct(&f,
		vd.Field(&f.Role, vd.When(!vd.IsEmpty(f.Role), vd.By(func(_ interface{}) error {
			return f.Role.ValidateRole()
		}))),
	)
}
