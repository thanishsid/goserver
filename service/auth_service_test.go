package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/mock"
)

func TestAuth_PasswordLogin(t *testing.T) {
	type fields struct {
		UserService    domain.UserService
		SessionService domain.SessionService
		GoogleConfig   *oauth2.Config
	}

	type args struct {
		ctx   context.Context
		input domain.PasswordLoginInput
	}

	testUser := getFakeUser(t)

	user := testUser.User
	password := testUser.Password

	validInput := domain.PasswordLoginInput{
		Email:     user.Email,
		Password:  password,
		UserAgent: fake.UserAgent(),
	}

	session := getFakeSession(t, user.ID, user.Role, validInput.UserAgent, nil, nil)

	invalidEmailInput := validInput
	invalidEmailInput.Email = "test@@gmail.com"

	noPasswordUser := *user
	noPasswordUser.PasswordHash = null.String{}
	noPasswordUserInput := domain.PasswordLoginInput{
		Email:     noPasswordUser.Email,
		Password:  validInput.Password,
		UserAgent: validInput.UserAgent,
	}

	incorrectPasswordInput := domain.PasswordLoginInput{
		Email:     user.Email,
		Password:  "incorrect_password",
		UserAgent: validInput.UserAgent,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Session
		wantErr bool
	}{
		{
			name: "Invalid Email Input must fail",
			args: args{
				ctx:   context.Background(),
				input: invalidEmailInput,
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "User not found must fail",
			fields: fields{
				UserService: &mock.UserServiceMock{
					OneByEmailFunc: func(ctx context.Context, email string) (*domain.User, error) {
						return nil, ErrNotFound
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				input: validInput,
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "User without a password must fail",
			fields: fields{
				UserService: &mock.UserServiceMock{
					OneByEmailFunc: func(ctx context.Context, email string) (*domain.User, error) {
						require.Equal(t, noPasswordUser.Email, email)
						return &noPasswordUser, nil
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				input: noPasswordUserInput,
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "Invalid Password must fail",
			fields: fields{
				UserService: &mock.UserServiceMock{
					OneByEmailFunc: func(ctx context.Context, email string) (*domain.User, error) {
						require.Equal(t, user.Email, email)
						return user, nil
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				input: incorrectPasswordInput,
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "Failed to create session must fail",
			fields: fields{
				UserService: &mock.UserServiceMock{
					OneByEmailFunc: func(ctx context.Context, email string) (*domain.User, error) {
						require.Equal(t, user.Email, email)
						return user, nil
					},
				},
				SessionService: &mock.SessionServiceMock{
					CreateFunc: func(ctx context.Context, input domain.CreateSessionInput) (*domain.Session, error) {
						return nil, fmt.Errorf("failed to create session")
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				input: validInput,
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "All Ok Must Pass",
			fields: fields{
				UserService: &mock.UserServiceMock{
					OneByEmailFunc: func(ctx context.Context, email string) (*domain.User, error) {
						require.Equal(t, user.Email, email)
						return user, nil
					},
				},
				SessionService: &mock.SessionServiceMock{
					CreateFunc: func(ctx context.Context, input domain.CreateSessionInput) (*domain.Session, error) {
						require.Equal(t, input.UserID, user.ID)
						require.Equal(t, input.UserRole, user.Role)
						require.Equal(t, input.UserAgent, validInput.UserAgent)
						return session, nil
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				input: validInput,
			},
			wantErr: false,
			want:    session,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Auth{
				UserService:    tt.fields.UserService,
				SessionService: tt.fields.SessionService,
				GoogleConfig:   tt.fields.GoogleConfig,
			}
			got, err := a.PasswordLogin(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth.PasswordLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth.PasswordLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
