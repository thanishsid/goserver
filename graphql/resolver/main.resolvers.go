package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/generated"
	"github.com/thanishsid/goserver/graphql/model"
)

// UploadImage is the resolver for the UploadImage field.
func (r *mutationsResolver) UploadImage(ctx context.Context, file graphql.Upload) (*domain.Image, error) {
	image, err := r.ImageService.Save(ctx, domain.ImageUploadInput{
		FileUploadData: domain.FileUploadData{
			File:        file.File,
			FileName:    file.Filename,
			Size:        file.Size,
			ContentType: file.ContentType,
		},
	})
	if err != nil {
		return nil, err
	}

	return image, nil
}

// UploadVideo is the resolver for the UploadVideo field.
func (r *mutationsResolver) UploadVideo(ctx context.Context, file graphql.Upload) (*domain.Video, error) {
	video, err := r.VideoService.Save(ctx, domain.VideoUploadInput{
		FileUploadData: domain.FileUploadData{
			File:        file.File,
			FileName:    file.Filename,
			Size:        file.Size,
			ContentType: file.ContentType,
		},
	})
	if err != nil {
		return nil, err
	}

	return video, nil
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
	roles := make([]*model.Role, len(domain.Roles))

	for i, role := range domain.Roles {
		roles[i] = &model.Role{
			ID:   role.ID.String(),
			Name: role.Name,
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
