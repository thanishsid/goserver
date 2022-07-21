package repository

import (
	"context"
	"testing"
	"time"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/postgres"
	"github.com/thanishsid/goserver/internal/security"
	"github.com/thanishsid/goserver/mock/mockpostgres"
	"github.com/thanishsid/goserver/mock/mocksearch"
)

func TestSaveOrUpdateUser(t *testing.T) {
	now := time.Now()

	newUser := &domain.User{
		ID:           uuid.New(),
		Email:        fake.Email(),
		Username:     fake.Username(),
		FullName:     fake.Name(),
		PasswordHash: fake.Password(true, true, true, true, false, 16),
		RoleID:       security.Administrator,
		PictureID:    uuid.NullUUID{UUID: uuid.New(), Valid: true},
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	userWithoutEmail := *newUser
	userWithoutEmail.Email = ""

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuerier := mockpostgres.NewMockQuerier(ctrl)

	// passing valid InsertOrUpdate params.
	mockQuerier.EXPECT().InsertOrUpdateUser(gomock.Any(), gomock.Eq(postgres.InsertOrUpdateUserParams{
		ID:           newUser.ID,
		Email:        newUser.Email,
		Username:     newUser.Username,
		FullName:     newUser.FullName,
		RoleID:       int32(newUser.RoleID),
		PasswordHash: newUser.PasswordHash,
		PictureID:    newUser.PictureID,
		CreatedAt:    newUser.CreatedAt,
		UpdatedAt:    newUser.UpdatedAt,
	})).Return(nil)

	mockUserSearcher := mocksearch.NewMockUserSearcher(ctrl)

	mockUserSearcher.EXPECT().AddOrUpdateUser(gomock.Any(), gomock.Eq(newUser)).Return(nil)

	userRepo := &userRepository{
		db:          mockQuerier,
		searchIndex: mockUserSearcher,
	}

	err := userRepo.SaveOrUpdate(context.Background(), newUser)
	require.NoError(t, err)

	err = userRepo.SaveOrUpdate(context.Background(), &userWithoutEmail)
	require.Error(t, err)
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()

	mockQuerier := mockpostgres.NewMockQuerier(ctrl)

	mockQuerier.EXPECT().SoftDeleteUser(gomock.Any(), gomock.Eq(id)).Return(nil)

	mockUserSearcher := mocksearch.NewMockUserSearcher(ctrl)

	mockUserSearcher.EXPECT().RemoveUser(gomock.Any(), gomock.Eq(id)).Return(nil)

	userRepo := &userRepository{
		db:          mockQuerier,
		searchIndex: mockUserSearcher,
	}

	err := userRepo.Delete(context.Background(), id)
	require.NoError(t, err)
}

func TestGetOneUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userID := uuid.New()

	now := time.Now()

	user := &domain.User{
		ID:           userID,
		Email:        fake.Email(),
		Username:     fake.Username(),
		FullName:     fake.Name(),
		PasswordHash: fake.Password(true, true, true, true, false, 16),
		RoleID:       security.Administrator,
		PictureID:    uuid.NullUUID{UUID: uuid.New(), Valid: true},
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	mockQuerier := mockpostgres.NewMockQuerier(ctrl)

	mockQuerier.EXPECT().
		GetUserById(gomock.Any(), gomock.Eq(userID)).
		Return(postgres.User{
			ID:           userID,
			Username:     user.Username,
			Email:        user.Email,
			FullName:     user.FullName,
			RoleID:       int32(user.RoleID),
			PasswordHash: user.PasswordHash,
			PictureID:    user.PictureID,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
			DeletedAt:    user.DeletedAt,
		}, nil)

	userRepo := &userRepository{
		db: mockQuerier,
	}

	gotUser, err := userRepo.OneByID(context.Background(), userID)
	require.NoError(t, err)
	require.Equal(t, user, gotUser)
}
