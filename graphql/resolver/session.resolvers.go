package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/cookiemanager"
	"github.com/thanishsid/goserver/graphql/generated"
	"github.com/thanishsid/goserver/graphql/model"
)

// Login is the resolver for the Login field.
func (r *mutationsResolver) Login(ctx context.Context, email string, password string) (*model.Message, error) {
	useragent, err := UseragentFor(ctx)
	if err != nil {
		return nil, err
	}

	session, err := r.AuthService.PasswordLogin(ctx, domain.PasswordLoginInput{
		Email:     email,
		Password:  password,
		UserAgent: useragent,
	})
	if err != nil {
		return nil, err
	}

	cookiemanager.For(ctx).SetCookie(config.SESSION_COOKIE_NAME, session.ID.String(), config.SESSION_TTL)

	return &model.Message{
		Message: "Login Successful",
	}, nil
}

// Logout is the resolver for the Logout field.
func (r *mutationsResolver) Logout(ctx context.Context) (*model.Message, error) {
	session, err := domain.SessionFor(ctx)
	if err != nil {
		return nil, err
	}

	if err := r.SessionService.Delete(ctx, session.ID); err != nil {
		return nil, err
	}

	cookiemanager.For(ctx).RemoveCookie(session.ID.String())

	return &model.Message{
		Message: "Logout Successful",
	}, nil
}

// LogoutFromAllDevices is the resolver for the LogoutFromAllDevices field.
func (r *mutationsResolver) LogoutFromAllDevices(ctx context.Context) (*model.Message, error) {
	session, err := domain.SessionFor(ctx)
	if err != nil {
		return nil, err
	}

	if err := r.SessionService.DeleteAllByUserID(ctx, session.UserID); err != nil {
		return nil, err
	}

	return &model.Message{
		Message: "logged out from all devices successfully",
	}, nil
}

// ID is the resolver for the id field.
func (r *sessionResolver) ID(ctx context.Context, obj *domain.Session) (string, error) {
	return obj.ID.String(), nil
}

// Session returns generated.SessionResolver implementation.
func (r *Resolver) Session() generated.SessionResolver { return &sessionResolver{r} }

type sessionResolver struct{ *Resolver }
