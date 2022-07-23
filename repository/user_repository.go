package repository

import (
	"context"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/postgres"
	"github.com/thanishsid/goserver/infrastructure/search"
	"github.com/thanishsid/goserver/infrastructure/security"
	"github.com/thanishsid/goserver/input"
)

type userRepository struct {
	db          postgres.Querier
	searchIndex search.UserSearcher
}

var _ domain.UserRepository = (*userRepository)(nil)

func (u *userRepository) SaveOrUpdate(ctx context.Context, user *domain.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := u.db.InsertOrUpdateUser(ctx, postgres.InsertOrUpdateUserParams{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		FullName:     user.FullName,
		UserRole:     string(user.Role),
		PasswordHash: user.PasswordHash,
		PictureID:    user.PictureID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}); err != nil {
		return err
	}

	return u.searchIndex.AddOrUpdateUser(ctx, user)
}

func (u *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := u.db.SoftDeleteUser(ctx, id); err != nil {
		return err
	}

	return u.searchIndex.RemoveUser(ctx, id)
}

func (u *userRepository) OneByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	userRow, err := u.db.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           userRow.ID,
		Email:        userRow.Email,
		Username:     userRow.Username,
		FullName:     userRow.FullName,
		Role:         security.Role(userRow.UserRole),
		PasswordHash: userRow.PasswordHash,
		PictureID:    userRow.PictureID,
		CreatedAt:    userRow.CreatedAt,
		UpdatedAt:    userRow.UpdatedAt,
		DeletedAt:    null.NewTime(userRow.DeletedAt.Time, userRow.DeletedAt.Valid),
	}, nil
}

func (u *userRepository) OneByEmail(ctx context.Context, email string) (*domain.User, error) {
	userRow, err := u.db.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           userRow.ID,
		Email:        userRow.Email,
		Username:     userRow.Username,
		FullName:     userRow.FullName,
		Role:         security.Role(userRow.UserRole),
		PasswordHash: userRow.PasswordHash,
		PictureID:    userRow.PictureID,
		CreatedAt:    userRow.CreatedAt,
		UpdatedAt:    userRow.UpdatedAt,
		DeletedAt:    null.NewTime(userRow.DeletedAt.Time, userRow.DeletedAt.Valid),
	}, nil
}

func (u *userRepository) Many(ctx context.Context, params input.UserFilterBase) ([]domain.User, error) {
	users, err := u.searchIndex.SearchUsers(ctx, params)
	if err != nil {
		return nil, err
	}

	return users, nil
}
