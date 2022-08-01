package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/generated"
	"github.com/thanishsid/goserver/graphql/model"
	"github.com/thanishsid/goserver/infrastructure/security"
	null "gopkg.in/guregu/null.v4"
)

// UploadImage is the resolver for the UploadImage field.
func (r *mutationsResolver) UploadImage(ctx context.Context, file graphql.Upload, title *string) (*domain.Image, error) {
	image, err := r.ImageService.Save(ctx, domain.ImageUploadInput{
		Title:       null.StringFromPtr(title),
		File:        file.File,
		ContentType: file.ContentType,
	})
	if err != nil {
		return nil, err
	}

	return image, nil
}

// Me is the resolver for the me field.
func (r *queriesResolver) Me(ctx context.Context) (*domain.User, error) {
	user, err := security.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Roles is the resolver for the roles field.
func (r *queriesResolver) Roles(ctx context.Context) ([]*model.Role, error) {
	roles := make([]*model.Role, len(domain.AllRoles))

	for i, role := range domain.AllRoles {
		roles[i] = &model.Role{
			ID:   string(role),
			Name: strings.Title(string(role)),
		}
	}

	return roles, nil
}

// Mutations returns generated.MutationsResolver implementation.
func (r *Resolver) Mutations() generated.MutationsResolver { return &mutationsResolver{r} }

// Queries returns generated.QueriesResolver implementation.
func (r *Resolver) Queries() generated.QueriesResolver { return &queriesResolver{r} }

type mutationsResolver struct{ *Resolver }
type queriesResolver struct{ *Resolver }
