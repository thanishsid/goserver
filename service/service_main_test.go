package service

import (
	"math/rand"
	"os"
	"testing"
	"time"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
)

type TestUser struct {
	User     *domain.User
	Password string
}

type TestUsersCollection struct {
	TestUsers []*TestUser
	Users     []*domain.User
	UserIDS   []uuid.UUID
}

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	config.ReadConfig("../.env")
	code := m.Run()
	os.Exit(code)
}

func getRandRole() domain.Role {
	randIndex := rand.Intn(len(domain.Roles))
	return domain.Roles[randIndex].ID
}

func getFakeUser(t *testing.T) *TestUser {
	password := fake.Password(true, true, true, true, false, 16)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)

	user := &domain.User{
		ID:           uuid.New(),
		Email:        fake.Email(),
		Username:     fake.Username(),
		FullName:     fake.Name(),
		PasswordHash: null.StringFrom(string(passwordHash)),
		Role:         getRandRole(),
		PictureID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		CreatedAt: time.Now().Add(-time.Hour * 24 * 200),
		UpdatedAt: time.Now().Add(-time.Hour * 36),
	}

	return &TestUser{
		User:     user,
		Password: password,
	}
}

func getManyFakeUsers(t *testing.T, count int) *TestUsersCollection {
	testUsers := make([]*TestUser, count)
	users := make([]*domain.User, count)
	userIds := make([]uuid.UUID, count)

	for i := 0; i < count; i++ {
		testUser := getFakeUser(t)
		testUsers[i] = testUser
		users[i] = testUser.User
		userIds[i] = testUser.User.ID
	}

	return &TestUsersCollection{
		TestUsers: testUsers,
		Users:     users,
	}
}

func getFakeSession(t *testing.T, userID uuid.UUID, userRole domain.Role, useragent string, extPic *string, data map[string]any) *domain.Session {
	now := time.Now()

	return &domain.Session{
		ID:              domain.NewSID(userID),
		UserID:          userID,
		UserRole:        userRole,
		UserAgent:       useragent,
		ExternalPicture: null.StringFromPtr(extPic).Ptr(),
		CreatedAt:       now,
		AccessedAt:      now,
		Data:            data,
	}
}
