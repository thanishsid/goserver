package tokenizer

import (
	"errors"
	"time"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/thanishsid/goserver/internal/security"
)

var ErrInvalidExpiry = errors.New("token must have a valid expiry time")
var ErrTokenExpired = errors.New("this token has expired")

type RegistrationClaims struct {
	Username  string        `json:"uname"`
	Email     string        `json:"email"`
	FullName  string        `json:"fname"`
	Role      security.Role `json:"role"`
	PictureID uuid.NullUUID `json:"pictureId"`
	Expiry    time.Time     `json:"exp"`
}

func (c RegistrationClaims) Validate() error {
	return vd.ValidateStruct(&c,
		vd.Field(&c.Username, vd.Required),
		vd.Field(&c.Email, vd.Required, is.Email),
		vd.Field(&c.FullName, vd.Required),
		vd.Field(&c.Role, vd.By(func(value interface{}) error {
			return c.Role.ValidateRole()
		})),
		vd.Field(&c.Expiry, vd.By(func(value interface{}) error {
			if c.Expiry.IsZero() {
				return ErrInvalidExpiry
			}

			if time.Now().After(c.Expiry) {
				return ErrTokenExpired
			}

			return nil
		})),
	)
}
