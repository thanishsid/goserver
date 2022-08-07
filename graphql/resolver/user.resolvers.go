package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/dataloader"
	"github.com/thanishsid/goserver/graphql/generated"
	"github.com/thanishsid/goserver/graphql/model"
	null "gopkg.in/guregu/null.v4"
)

// StartAccountRegistration is the resolver for the StartAccountRegistration field.
func (r *mutationsResolver) StartAccountRegistration(ctx context.Context, input model.StartRegistrationInput) (*model.Message, error) {
	if err := r.UserService.InitRegistration(ctx, domain.InitRegistrationInput{
		Email:                  input.Email,
		FullName:               input.FullName,
		Role:                   domain.RoleClient,
		ClientRegistrationLink: input.CallbackURL,
	}); err != nil {
		return nil, err
	}

	return &model.Message{
		Message: "Please check your email to complete the account registration.",
	}, nil
}

// StartAccountCreation is the resolver for the StartAccountCreation field.
func (r *mutationsResolver) StartAccountCreation(ctx context.Context, input model.StartUserCreationInput) (*model.Message, error) {
	if err := r.UserService.InitRegistration(ctx, domain.InitRegistrationInput{
		Email:                  input.Email,
		FullName:               input.FullName,
		Role:                   domain.Role(input.Role),
		ClientRegistrationLink: input.CallbackURL,
	}); err != nil {
		return nil, err
	}

	return &model.Message{
		Message: "Please check your email to complete the account registration.",
	}, nil
}

// CompleteRegistration is the resolver for the CompleteRegistration field.
func (r *mutationsResolver) CompleteRegistration(ctx context.Context, input model.CompleteRegistrationInput) (*model.Message, error) {
	_, err := r.UserService.CompleteRegistration(ctx, domain.CompleteRegistrationInput{
		Token:     input.Token,
		Username:  input.Username,
		PictureID: UUIDFromPtr(input.PictureID),
		Password:  input.Password,
	})
	if err != nil {
		return nil, err
	}

	return &model.Message{
		Message: "Registration successful",
	}, nil
}

// UpdateProfile is the resolver for the UpdateProfile field.
func (r *mutationsResolver) UpdateProfile(ctx context.Context, input model.UpdateProfileInput) (*model.Message, error) {
	session, err := domain.SessionFor(ctx)
	if err != nil {
		return nil, err
	}

	if err := r.UserService.Update(ctx, session.UserID, domain.UserUpdateInput{
		Username:  input.Username,
		FullName:  input.FullName,
		PictureID: UUIDFromPtr(input.PictureID),
	}); err != nil {
		return nil, err
	}

	return &model.Message{
		Message: "Account updated successfully",
	}, nil
}

// DeleteOwnAccount is the resolver for the DeleteOwnAccount field.
func (r *mutationsResolver) DeleteOwnAccount(ctx context.Context) (*model.Message, error) {
	session, err := domain.SessionFor(ctx)
	if err != nil {
		return nil, err
	}

	if err := r.UserService.Delete(ctx, session.UserID); err != nil {
		return nil, err
	}

	return &model.Message{
		Message: "Account deleted successfully",
	}, nil
}

// DeleteAnotherAccount is the resolver for the DeleteAnotherAccount field.
func (r *mutationsResolver) DeleteAnotherAccount(ctx context.Context, id string) (*model.Message, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	if err := r.UserService.Delete(ctx, userID); err != nil {
		return nil, err
	}

	return &model.Message{
		Message: "Account deleted successfully",
	}, nil
}

// RecoverAccount is the resolver for the RecoverAccount field.
func (r *mutationsResolver) RecoverAccount(ctx context.Context, id string) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// User is the resolver for the user field.
func (r *queriesResolver) User(ctx context.Context, id string) (*domain.User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	user, err := r.UserService.One(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Users is the resolver for the users field.
func (r *queriesResolver) Users(ctx context.Context, query *string, role *string, showDeleted *bool, limit *int, cursor *string) (*model.UserCollection, error) {
	usersData, err := r.UserService.Many(ctx, domain.UserFilterInput{
		Query:       null.StringFromPtr(query),
		Role:        null.StringFromPtr(role),
		ShowDeleted: null.BoolFromPtr(showDeleted),
		Limit:       NullIntFromPtr(limit),
		Cursor:      null.StringFromPtr(cursor),
	})
	if err != nil {
		return nil, err
	}

	userCount := len(usersData.Nodes)

	userCollection := new(model.UserCollection)
	userCollection.Nodes = make([]*domain.User, userCount)
	userCollection.Edges = make([]*model.UserEdge, userCount)

	for i := 0; i < userCount; i++ {
		userCollection.Nodes[i] = usersData.Nodes[i]
		userCollection.Edges[i] = &model.UserEdge{
			Cursor: usersData.Cursors[i],
			Node:   usersData.Nodes[i],
		}
	}

	userCollection.PageInfo = &model.PageInfo{
		StartCursor: usersData.StartCursor.ValueOrZero(),
		EndCursor:   usersData.EndCursor.ValueOrZero(),
		HasNextPage: usersData.HasMore,
	}

	return userCollection, nil
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *domain.User) (string, error) {
	return obj.ID.String(), nil
}

// Role is the resolver for the role field.
func (r *userResolver) Role(ctx context.Context, obj *domain.User) (string, error) {
	return string(obj.Role), nil
}

// Picture is the resolver for the picture field.
func (r *userResolver) Picture(ctx context.Context, obj *domain.User) (*domain.Image, error) {
	if obj.PictureID.Valid {
		return dataloader.For(ctx).GetImage(ctx, obj.PictureID.UUID)
	}

	return nil, nil
}

// Sessions is the resolver for the sessions field.
func (r *userResolver) Sessions(ctx context.Context, obj *domain.User) ([]*domain.Session, error) {
	sessions, err := r.SessionService.GetAllByUserID(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *userResolver) DeletedAt(ctx context.Context, obj *domain.User) (*time.Time, error) {
	return obj.DeletedAt.Ptr(), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
