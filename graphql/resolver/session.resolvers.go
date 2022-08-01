package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/cookiemanager"
	"github.com/thanishsid/goserver/graphql/generated"
	"github.com/thanishsid/goserver/graphql/model"
)

// Login is the resolver for the Login field.
func (r *mutationsResolver) Login(ctx context.Context, email string, password string) (*model.Message, error) {
	user, err := r.UserService.VerifyCredentials(ctx, domain.VerifyCredentialsInput{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	sessionID, err := r.SessionService.Create(ctx, user.ID, "abcdefg", nil)
	if err != nil {
		return nil, err
	}

	cookieManager, err := cookiemanager.For(ctx)
	if err != nil {
		return nil, err
	}

	cookieManager.SetCookie(config.SESSION_COOKIE_NAME, sessionID.String(), config.SESSION_TTL)

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

	return &model.Message{
		Message: "Logout Successful",
	}, nil
}

// LogoutFromAllDevices is the resolver for the LogoutFromAllDevices field.
func (r *mutationsResolver) LogoutFromAllDevices(ctx context.Context) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetCurrentSession is the resolver for the GetCurrentSession field.
func (r *queriesResolver) GetCurrentSession(ctx context.Context) (*domain.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetMySessions is the resolver for the GetMySessions field.
func (r *queriesResolver) GetMySessions(ctx context.Context) ([]*domain.Session, error) {

	session, err := domain.SessionFor(ctx)
	if err != nil {
		return nil, err
	}

	sessions, err := r.SessionService.GetAllByUserID(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

// GetSessionsByUser is the resolver for the GetSessionsByUser field.
func (r *queriesResolver) GetSessionsByUser(ctx context.Context, id string) ([]*domain.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

// ID is the resolver for the id field.
func (r *sessionResolver) ID(ctx context.Context, obj *domain.Session) (string, error) {
	return obj.ID.String(), nil
}

// Session returns generated.SessionResolver implementation.
func (r *Resolver) Session() generated.SessionResolver { return &sessionResolver{r} }

type sessionResolver struct{ *Resolver }
