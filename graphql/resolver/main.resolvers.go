package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/generated"
	"github.com/thanishsid/goserver/graphql/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

// MyInfo is the resolver for the myInfo field.
func (r *queriesResolver) MyInfo(ctx context.Context) (*model.MyInfo, error) {
	session, err := domain.SessionFor(ctx)
	if err != nil {
		return nil, err
	}

	return &model.MyInfo{
		Session: session,
	}, nil
}

// Roles is the resolver for the roles field.
func (r *queriesResolver) Roles(ctx context.Context) ([]*model.Role, error) {
	roles := make([]*model.Role, len(domain.AllRoles))

	caser := cases.Title(language.English)

	for i, role := range domain.AllRoles {
		roles[i] = &model.Role{
			ID:   string(role),
			Name: caser.String(role.String()),
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
