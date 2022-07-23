package repository

import (
	"context"
	"math/rand"
	"testing"
	"time"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/postgres"
	"github.com/thanishsid/goserver/internal/input"
	"github.com/thanishsid/goserver/internal/security"
	"github.com/thanishsid/goserver/mock/mockpostgres"
	"github.com/thanishsid/goserver/mock/mocksearch"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandRole() security.Role {
	randIndex := rand.Intn(len(security.AllRoles))

	return security.AllRoles[randIndex]
}

func createRandomUser() *domain.User {
	now := time.Now()

	return &domain.User{
		ID:           uuid.New(),
		Email:        fake.Email(),
		Username:     fake.Username(),
		FullName:     fake.Name(),
		PasswordHash: fake.Password(true, true, true, true, false, 16),
		Role:         getRandRole(),
		PictureID:    uuid.NullUUID{UUID: uuid.New(), Valid: true},
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func TestSaveOrUpdateUser(t *testing.T) {

	newUser := createRandomUser()

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
		UserRole:     string(newUser.Role),
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

	user := createRandomUser()

	mockQuerier := mockpostgres.NewMockQuerier(ctrl)

	mockQuerier.EXPECT().
		GetUserById(gomock.Any(), gomock.Eq(user.ID)).
		Return(postgres.User{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			FullName:     user.FullName,
			UserRole:     string(user.Role),
			PasswordHash: user.PasswordHash,
			PictureID:    user.PictureID,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
			DeletedAt:    user.DeletedAt,
		}, nil)

	userRepo := &userRepository{
		db: mockQuerier,
	}

	gotUser, err := userRepo.OneByID(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, user, gotUser)
}

func TestGetOneUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := createRandomUser()

	mockQuerier := mockpostgres.NewMockQuerier(ctrl)

	mockQuerier.EXPECT().
		GetUserByEmail(gomock.Any(), gomock.Eq(user.Email)).
		Return(postgres.User{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			FullName:     user.FullName,
			UserRole:     string(user.Role),
			PasswordHash: user.PasswordHash,
			PictureID:    user.PictureID,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
			DeletedAt:    user.DeletedAt,
		}, nil)

	userRepo := &userRepository{
		db: mockQuerier,
	}

	gotUser, err := userRepo.OneByEmail(context.Background(), user.Email)
	require.NoError(t, err)
	require.Equal(t, user, gotUser)
}

func TestGetManyUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	users := make([]domain.User, 20)

	for i := range users {
		users[i] = *createRandomUser()
		users[i].PasswordHash = ""
	}

	manyUsersParams := input.UserFilterBase{
		Query:       null.StringFrom(fake.Name()),
		Role:        null.StringFrom(string(getRandRole())),
		ShowDeleted: null.BoolFrom(true),
		Limit:       null.IntFrom(40),
	}

	mockUserSearcher := mocksearch.NewMockUserSearcher(ctrl)

	mockUserSearcher.EXPECT().SearchUsers(gomock.Any(), gomock.Eq(manyUsersParams)).Return(users, nil)

	userRepo := &userRepository{
		searchIndex: mockUserSearcher,
	}

	gotUsers, err := userRepo.Many(context.Background(), manyUsersParams)
	require.NoError(t, err)

	require.Equal(t, gotUsers, users)
}
