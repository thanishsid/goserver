package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
)

type Auth struct {
	UserService    domain.UserService
	SessionService domain.SessionService
	GoogleConfig   *oauth2.Config
}

var _ domain.AuthService = (*Auth)(nil)

// Verify user with provided email and password credentials if valid will create and return a new session.
func (a *Auth) PasswordLogin(ctx context.Context, input domain.PasswordLoginInput) (*domain.Session, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	user, err := a.UserService.OneByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	if !user.PasswordHash.Valid {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.ValueOrZero()), []byte(input.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	session, err := a.SessionService.Create(ctx, domain.CreateSessionInput{
		UserID:    user.ID,
		UserRole:  domain.Role(user.Role),
		UserAgent: input.UserAgent,
	})
	if err != nil {
		return nil, err
	}

	return session, nil
}

// Generate and return the google login url.
func (a *Auth) GenerateGoogleUrl(ctx context.Context, input domain.RedirectLinkInput) (string, error) {
	if err := input.Validate(); err != nil {
		return "", err
	}

	state := config.C.OauthState + input.SuccessRedirect + "#" + input.FailRedirect

	return a.GoogleConfig.AuthCodeURL(state), nil
}

// Generate google access token and fetch user info, if user does not exists a new account
// will be created and then a new session will be created and returned along with the redirect link
// provided to the url generate function.
func (a *Auth) GoogleLogin(ctx context.Context, input domain.OauthLoginInput) (*domain.Session, string, error) {

	if err := input.Validate(); err != nil {
		return nil, "", err
	}

	redirects := strings.Split(strings.TrimPrefix(input.State, config.C.OauthState), "#")

	if len(redirects) < 2 {
		return nil, "", errors.New("redirect links are invalid or unavailable")
	}

	successRedirBytes, err := base64.URLEncoding.DecodeString(redirects[0])
	if err != nil {
		return nil, "", fmt.Errorf("failed to parse success redirect url")
	}

	failRedirBytes, err := base64.URLEncoding.DecodeString(redirects[1])
	if err != nil {
		return nil, "", fmt.Errorf("failed to parse fail redirect url")
	}

	successRedirect := string(successRedirBytes)
	failRedirect := string(failRedirBytes)

	type GoogleUserInfo struct {
		Sub           string      `json:"sub"`
		Name          string      `json:"name"`
		GivenName     string      `json:"given_name"`
		FamilyName    string      `json:"family_name"`
		Picture       null.String `json:"picture"`
		Email         string      `json:"email"`
		EmailVerified bool        `json:"email_verified"`
		Locale        string      `json:"locale"`
	}

	token, err := a.GoogleConfig.Exchange(ctx, input.AuthCode)
	if err != nil {
		return nil, failRedirect, err
	}

	client := a.GoogleConfig.Client(ctx, token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token" + token.AccessToken)
	if err != nil {
		return nil, failRedirect, err
	}

	userInfo := new(GoogleUserInfo)

	if err := json.NewDecoder(resp.Body).Decode(userInfo); err != nil {
		return nil, failRedirect, err
	}

	user, err := a.UserService.OneByEmail(ctx, userInfo.Email)
	if err != nil && err != ErrNotFound {
		return nil, failRedirect, err
	}

	if err == ErrNotFound {
		user, err = a.UserService.Create(ctx, domain.CreateUserInput{
			Username: userInfo.GivenName,
			Email:    userInfo.Email,
			FullName: userInfo.Name,
			Role:     domain.RoleClient,
		})
		if err != nil {
			return nil, failRedirect, err
		}
	}

	session, err := a.SessionService.Create(ctx, domain.CreateSessionInput{
		UserID:          user.ID,
		UserRole:        user.Role,
		UserAgent:       input.UserAgent,
		ExternalPicture: userInfo.Picture,
	})
	if err != nil {
		return nil, failRedirect, err
	}

	return session, successRedirect, nil
}
