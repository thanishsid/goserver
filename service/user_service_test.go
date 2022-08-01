package service

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/db"
	"github.com/thanishsid/goserver/infrastructure/mailer"
	"github.com/thanishsid/goserver/infrastructure/tokenizer"
	"github.com/thanishsid/goserver/mock/mockdb"
	"github.com/thanishsid/goserver/mock/mockmailer"
	"github.com/thanishsid/goserver/mock/mocktoken"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type UserServiceTestSuite struct {
	Token         *mocktoken.MockTokenizer
	Mail          *mockmailer.MockMailer
	DB            *mockdb.MockDB
	Querier       *mockdb.MockQuerier
	Transactioner *mockdb.MockTransactioner
}

func (ust *UserServiceTestSuite) GetService() domain.UserService {
	return &userService{
		Tokens: ust.Token,
		Mail:   ust.Mail,
		DB:     ust.DB,
	}
}

func getRandRole() domain.Role {
	randIndex := rand.Intn(len(domain.AllRoles))
	return domain.AllRoles[randIndex]
}

func CreateUserServiceTestSuite(t *testing.T) *UserServiceTestSuite {
	t.Parallel()
	ctrl := gomock.NewController(t)
	tokens := mocktoken.NewMockTokenizer(ctrl)
	mail := mockmailer.NewMockMailer(ctrl)
	dbs := mockdb.NewMockDB(ctrl)
	qry := mockdb.NewMockQuerier(ctrl)
	trx := mockdb.NewMockTransactioner(ctrl)

	return &UserServiceTestSuite{
		Token:         tokens,
		Mail:          mail,
		DB:            dbs,
		Querier:       qry,
		Transactioner: trx,
	}
}

func Test_InitRegistration(t *testing.T) {
	suite := CreateUserServiceTestSuite(t)

	regForm := domain.InitRegistrationInput{
		Email:                  "john@gmail.com",
		FullName:               fake.Name(),
		Role:                   getRandRole(),
		ClientRegistrationLink: fake.URL(),
	}

	fakeToken := "abcd"

	linkMailData := mailer.LinkMailData{
		To:               regForm.Email,
		Subject:          "Account Registration",
		Title:            "New Account Registration",
		PrimaryMessage:   fmt.Sprintf("Hi %s thank you for registering with us.", regForm.FullName),
		SecondaryMessage: "Please click the link below to complete your registration.",
		Link:             regForm.ClientRegistrationLink + "/" + fakeToken,
	}

	suite.Token.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Return(fakeToken, nil)
	suite.Mail.EXPECT().SendLinkMail(gomock.Any(), linkMailData).Return(nil)

	userSvc := NewUserService(suite.Token, suite.Mail, suite.DB)

	err := userSvc.InitRegistration(context.Background(), regForm)
	require.NoError(t, err)
}

func Test_CompleteRegistration(t *testing.T) {
	suite := CreateUserServiceTestSuite(t)

	fakeToken := "abcd"

	completeRegForm := domain.CompleteRegistrationInput{
		Token:    fakeToken,
		Username: fake.Name(),
		PictureID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		Password: fake.Password(true, true, true, true, false, 16),
	}

	claims := tokenizer.RegistrationClaims{
		Email:    "john@gmail.com",
		FullName: fake.Name(),
		Role:     getRandRole(),
		Expiry:   time.Now().Add(config.REGISTRATION_TOKEN_TTL),
	}

	suite.Token.EXPECT().GetClaims(gomock.Any(), fakeToken, gomock.Any()).SetArg(2, claims).Return(nil)
	suite.DB.EXPECT().InsertOrUpdateUser(gomock.Any(), gomock.Any()).Return(nil)

	userSvc := NewUserService(suite.Token, suite.Mail, suite.DB)
	user, err := userSvc.CompleteRegistration(context.Background(), completeRegForm)
	require.NoError(t, err)
	require.Equal(t, user.Email, claims.Email)
	require.Equal(t, user.Username, completeRegForm.Username)
	require.Equal(t, user.FullName, claims.FullName)
	require.Equal(t, user.Role, claims.Role)
	require.Equal(t, user.PictureID, completeRegForm.PictureID)
}

type UserUpdateParamWithMatcher struct {
	db.InsertOrUpdateUserParams
}

// Matches returns whether x is a match.
func (u UserUpdateParamWithMatcher) Matches(x interface{}) bool {
	m, ok := x.(db.InsertOrUpdateUserParams)
	if !ok {
		return false
	}

	return u.ID == m.ID && u.Username == m.Username &&
		u.Email == m.Email && u.FullName == m.FullName &&
		u.PasswordHash == m.PasswordHash && u.PictureID == m.PictureID &&
		u.Role == m.Role && u.CreatedAt == m.CreatedAt &&
		!u.UpdatedAt.Equal(m.UpdatedAt)
}

// String describes what the matcher matches.
func (u UserUpdateParamWithMatcher) String() string {
	return "user update params match"
}

func Test_UpdateUser(t *testing.T) {
	suite := CreateUserServiceTestSuite(t)

	gotUserRow := db.GetUserByIdRow{
		ID:           uuid.New(),
		Email:        fake.Email(),
		Username:     fake.Username(),
		FullName:     fake.Name(),
		Role:         string(domain.RoleAdministrator),
		PasswordHash: fake.Password(true, true, true, true, false, 16),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	updateForm := domain.UserUpdateInput{
		Username: fake.Username(),
		FullName: fake.Name(),
		PictureID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
	}

	suite.Transactioner.EXPECT().GetUserById(gomock.Any(), gomock.Eq(gotUserRow.ID)).Return(gotUserRow, nil)
	suite.Transactioner.EXPECT().InsertOrUpdateUser(gomock.Any(), UserUpdateParamWithMatcher{
		db.InsertOrUpdateUserParams{
			ID:           gotUserRow.ID,
			Username:     updateForm.Username,
			Email:        gotUserRow.Email,
			FullName:     updateForm.FullName,
			Role:         string(gotUserRow.Role),
			PasswordHash: gotUserRow.PasswordHash,
			PictureID:    updateForm.PictureID,
			CreatedAt:    gotUserRow.CreatedAt,
			UpdatedAt:    gotUserRow.UpdatedAt,
		},
	})

	suite.Transactioner.EXPECT().Commit(gomock.Any())
	suite.Transactioner.EXPECT().Rollback(gomock.Any())
	suite.DB.EXPECT().BeginTx(gomock.Any(), gomock.Any()).Return(suite.Transactioner, nil)

	userSvc := suite.GetService()

	err := userSvc.Update(context.Background(), gotUserRow.ID, updateForm)
	require.NoError(t, err)
}
