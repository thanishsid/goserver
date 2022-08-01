package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/generated"
)

// ID is the resolver for the id field.
func (r *imageResolver) ID(ctx context.Context, obj *domain.Image) (string, error) {
	return obj.ID.String(), nil
}

// Title is the resolver for the title field.
func (r *imageResolver) Title(ctx context.Context, obj *domain.Image) (*string, error) {
	return obj.Title.Ptr(), nil
}

// Image returns generated.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }
