package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/dataloader"
	"github.com/thanishsid/goserver/graphql/generated"
	"github.com/thanishsid/goserver/graphql/model"
)

// ID is the resolver for the id field.
func (r *imageResolver) ID(ctx context.Context, obj *domain.Image) (string, error) {
	return obj.ID.String(), nil
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

// ID is the resolver for the id field.
func (r *videoResolver) ID(ctx context.Context, obj *domain.Video) (string, error) {
	return obj.ID.String(), nil
}

// Thumbnail is the resolver for the thumbnail field.
func (r *videoResolver) Thumbnail(ctx context.Context, obj *domain.Video) (*domain.Image, error) {
	return dataloader.For(ctx).GetImage(ctx, obj.ThumbnailID)
}

// Image returns generated.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

// MyInfo returns generated.MyInfoResolver implementation.
func (r *Resolver) MyInfo() generated.MyInfoResolver { return &myInfoResolver{r} }

// Video returns generated.VideoResolver implementation.
func (r *Resolver) Video() generated.VideoResolver { return &videoResolver{r} }

type imageResolver struct{ *Resolver }
type myInfoResolver struct{ *Resolver }
type videoResolver struct{ *Resolver }
