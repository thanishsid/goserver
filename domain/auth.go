package domain

import (
	"context"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type AuthService interface {
	// Verify user with provided email and password credentials if valid will create and return a new session.
	PasswordLogin(ctx context.Context, input PasswordLoginInput) (*Session, error)

	// Generate and return the google login url.
	GenerateGoogleUrl(ctx context.Context, input RedirectLinkInput) (string, error)

	// Generate google access token and fetch user info, if user does not exists a new account
	// will be created and then a new session will be created and returned along with the redirect link
	// provided to the url generate function.
	GoogleLogin(ctx context.Context, input OauthLoginInput) (*Session, string, error)
}

//--------- Inputs ----------

type PasswordLoginInput struct {
	Email     string `json:"string"`
	Password  string `json:"password"`
	UserAgent string `json:"-"`
}

func (i PasswordLoginInput) Validate() error {
	return vd.ValidateStruct(&i,
		vd.Field(&i.Email, vd.Required, IsEmail),
		vd.Field(&i.Password, vd.Required),
		vd.Field(&i.UserAgent, vd.Required),
	)
}

type RedirectLinkInput struct {
	SuccessRedirect string `schema:"success-redirect"`
	FailRedirect    string `schema:"fail-redirect"`
}

func (i RedirectLinkInput) Validate() error {
	return vd.ValidateStruct(&i,
		vd.Field(&i.SuccessRedirect, vd.Required, is.Base64),
		vd.Field(&i.FailRedirect, vd.Required, is.Base64),
	)
}

type OauthLoginInput struct {
	State     string `schema:"state"`
	AuthCode  string `schema:"code"`
	UserAgent string `schema:"-"`
}

func (i OauthLoginInput) Validate() error {
	return vd.ValidateStruct(&i,
		vd.Field(&i.State, vd.Required),
		vd.Field(&i.AuthCode, vd.Required),
		vd.Field(&i.UserAgent, vd.Required),
	)
}
