package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/generated"
	"github.com/thanishsid/goserver/graphql/model"
)

// ID is the resolver for the id field.
func (r *imageResolver) ID(ctx context.Context, obj *domain.Image) (string, error) {
	return obj.ID.String(), nil
}

// Title is the resolver for the title field.
func (r *imageResolver) Title(ctx context.Context, obj *domain.Image) (*string, error) {
	return obj.Title.Ptr(), nil
}

// Account is the resolver for the account field.
func (r *myInfoResolver) Account(ctx context.Context, obj *model.MyInfo) (*domain.User, error) {
	session, err := domain.SessionFor(ctx)
	if err != nil {
		return nil, err
	}

	user, err := r.UserService.One(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Image returns generated.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

// MyInfo returns generated.MyInfoResolver implementation.
func (r *Resolver) MyInfo() generated.MyInfoResolver { return &myInfoResolver{r} }

type imageResolver struct{ *Resolver }
type myInfoResolver struct{ *Resolver }
