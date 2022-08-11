package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/db"
	"github.com/thanishsid/goserver/infrastructure/mailer"
	"github.com/thanishsid/goserver/infrastructure/tokenizer"
	"github.com/thanishsid/goserver/mock"
	"gopkg.in/guregu/null.v4"
)

func TestUser_InitRegistration(t *testing.T) {
	t.Parallel()

	type fields struct {
		Tokens         tokenizer.Tokenizer
		Mail           mailer.Mailer
		DB             db.DB
		SessionService domain.SessionService
	}

	type args struct {
		ctx   context.Context
		input domain.InitRegistrationInput
	}

	validInput := domain.InitRegistrationInput{
		Email:                  "john@gmail.com",
		FullName:               fake.Name(),
		Role:                   getRandRole(),
		ClientRegistrationLink: fake.URL(),
	}

	invalidEmailInput := validInput
	invalidEmailInput.Email = "john@@gmail.com"

	invalidLinkInput := validInput
	invalidLinkInput.ClientRegistrationLink = "blabla://craplink. bla"

	fakeToken := "abcd"

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Invalid Email",
			args: args{
				ctx:   context.Background(),
				input: invalidEmailInput,
			},
			wantErr: true,
		},
		{
			name: "Invalid Client Registration Link",
			args: args{
				ctx:   context.Background(),
				input: invalidLinkInput,
			},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				Tokens: &mock.TokenizerMock{
					CreateTokenFunc: func(ctx context.Context, claims tokenizer.Validateable) (string, error) {
						tClaim, ok := claims.(tokenizer.RegistrationClaims)
						require.True(t, ok)
						require.Equal(t, tClaim, tokenizer.RegistrationClaims{
							Email:    validInput.Email,
							FullName: validInput.FullName,
							Role:     validInput.Role,
							Expiry:   tClaim.Expiry,
						})
						require.True(t, tClaim.Expiry.After(time.Now()))
						return fakeToken, nil
					},
				},
				Mail: &mock.MailerMock{
					SendLinkMailFunc: func(ctx context.Context, data mailer.LinkMailData) error {
						require.Equal(t, data.To, validInput.Email)
						require.Equal(t, data.Link, validInput.ClientRegistrationLink+"/"+fakeToken)
						require.NotZero(t, data.Subject)
						require.NotZero(t, data.Title)
						require.NotZero(t, data.PrimaryMessage)
						require.NotZero(t, data.SecondaryMessage)
						return nil
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				input: validInput,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Tokens:         tt.fields.Tokens,
				Mail:           tt.fields.Mail,
				DB:             tt.fields.DB,
				SessionService: tt.fields.SessionService,
			}
			if err := u.InitRegistration(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("User.InitRegistration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_CompleteRegistration(t *testing.T) {
	t.Parallel()

	type fields struct {
		Tokens         tokenizer.Tokenizer
		Mail           mailer.Mailer
		DB             db.DB
		SessionService domain.SessionService
	}

	type args struct {
		ctx   context.Context
		input domain.CompleteRegistrationInput
	}

	input := domain.CompleteRegistrationInput{
		Token:    "abcd",
		Username: fake.Name(),
		PictureID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		Password: fake.Password(true, true, true, true, false, 16),
	}

	tokenClaims := tokenizer.RegistrationClaims{
		Email:    "john@gmail.com",
		FullName: fake.Name(),
		Role:     getRandRole(),
		Expiry:   time.Now().Add(config.REGISTRATION_TOKEN_TTL),
	}

	var user domain.User

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Tokens: &mock.TokenizerMock{
					GetClaimsFunc: func(ctx context.Context, secureToken string, claims tokenizer.Validateable) error {
						require.Equal(t, secureToken, input.Token)
						tClaims, ok := claims.(*tokenizer.RegistrationClaims)
						require.True(t, ok)
						*tClaims = tokenClaims
						return nil
					},
				},
				DB: &mock.DBMock{
					InsertOrUpdateUserFunc: func(ctx context.Context, arg db.InsertOrUpdateUserParams) error {
						require.Equal(t, arg.Email, tokenClaims.Email)
						require.Equal(t, arg.FullName, tokenClaims.FullName)
						require.Equal(t, arg.Username, input.Username)
						require.Equal(t, arg.Role, tokenClaims.Role.String())
						require.Equal(t, arg.PictureID.UUID, input.PictureID.UUID)
						require.Equal(t, arg.PictureID.Valid, input.PictureID.Valid)
						require.NotZero(t, arg.CreatedAt)
						require.NotZero(t, arg.UpdatedAt)
						require.NotZero(t, arg.PasswordHash)

						user = domain.User{
							ID:           arg.ID,
							Email:        arg.Email,
							Username:     arg.Username,
							FullName:     arg.FullName,
							PasswordHash: arg.PasswordHash,
							Role:         domain.Role(arg.Role),
							PictureID:    arg.PictureID,
							CreatedAt:    arg.CreatedAt,
							UpdatedAt:    arg.UpdatedAt,
						}

						return nil
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				input: input,
			},
			want:    &user,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Tokens:         tt.fields.Tokens,
				Mail:           tt.fields.Mail,
				DB:             tt.fields.DB,
				SessionService: tt.fields.SessionService,
			}

			got, err := u.CompleteRegistration(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.CompleteRegistration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.CompleteRegistration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	t.Parallel()

	type fields struct {
		Tokens         tokenizer.Tokenizer
		Mail           mailer.Mailer
		DB             db.DB
		SessionService domain.SessionService
	}

	type args struct {
		ctx    context.Context
		userID uuid.UUID
		input  domain.UserUpdateInput
	}

	userRow := db.GetUserByIdRow{
		ID:           uuid.New(),
		Email:        fake.Email(),
		Username:     fake.Username(),
		FullName:     fake.Name(),
		Role:         string(domain.RoleAdmin),
		PasswordHash: null.StringFrom(fake.Password(true, true, true, true, false, 16)),
		CreatedAt:    time.Now().Add(-time.Hour),
		UpdatedAt:    time.Now().Add(-time.Hour),
	}

	input := domain.UserUpdateInput{
		Username: fake.Username(),
		FullName: fake.Name(),
		PictureID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				DB: &mock.DBMock{
					BeginTxFunc: func(ctx context.Context, txOpts pgx.TxOptions) (db.Transactioner, error) {
						return &mock.TransactionerMock{
							RollbackFunc: func(ctx context.Context) error {
								return nil
							},
							GetUserByIdFunc: func(ctx context.Context, userID uuid.UUID) (db.GetUserByIdRow, error) {
								return userRow, nil
							},
							InsertOrUpdateUserFunc: func(ctx context.Context, arg db.InsertOrUpdateUserParams) error {
								require.Equal(t, arg, db.InsertOrUpdateUserParams{
									ID:           userRow.ID,
									Username:     input.Username,
									Email:        userRow.Email,
									FullName:     input.FullName,
									Role:         userRow.Role,
									PasswordHash: userRow.PasswordHash,
									PictureID:    input.PictureID,
									CreatedAt:    userRow.CreatedAt,
									UpdatedAt:    arg.UpdatedAt,
								})
								require.True(t, arg.UpdatedAt.After(userRow.UpdatedAt))
								return nil
							},
							CommitFunc: func(ctx context.Context) error {
								return nil
							},
						}, nil
					},
				},
			},
			args: args{
				ctx:    context.Background(),
				userID: userRow.ID,
				input:  input,
			},
			wantErr: false,
		},
		{
			name: "Transaction Start Failed",
			fields: fields{
				DB: &mock.DBMock{
					BeginTxFunc: func(ctx context.Context, txOpts pgx.TxOptions) (db.Transactioner, error) {
						return nil, fmt.Errorf("transaction not started")
					},
				},
			},
			args: args{
				ctx:    context.Background(),
				userID: uuid.New(),
				input:  input,
			},
			wantErr: true,
		},
		{
			name: "User not found",
			fields: fields{
				DB: &mock.DBMock{
					BeginTxFunc: func(ctx context.Context, txOpts pgx.TxOptions) (db.Transactioner, error) {
						return &mock.TransactionerMock{
							RollbackFunc: func(ctx context.Context) error {
								return nil
							},
							GetUserByIdFunc: func(ctx context.Context, userID uuid.UUID) (db.GetUserByIdRow, error) {
								return db.GetUserByIdRow{}, pgx.ErrNoRows
							},
						}, nil
					},
				},
			},
			args: args{
				ctx:    context.Background(),
				userID: uuid.New(),
				input:  input,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Tokens:         tt.fields.Tokens,
				Mail:           tt.fields.Mail,
				DB:             tt.fields.DB,
				SessionService: tt.fields.SessionService,
			}
			if err := u.Update(tt.args.ctx, tt.args.userID, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("User.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_ChangeRole(t *testing.T) {
	t.Parallel()

	type fields struct {
		Tokens         tokenizer.Tokenizer
		Mail           mailer.Mailer
		DB             db.DB
		SessionService domain.SessionService
	}

	type args struct {
		ctx   context.Context
		input domain.RoleChangeInput
	}

	testUser := getFakeUser(t)

	input := domain.RoleChangeInput{
		UserID: testUser.User.ID,
		Role:   getRandRole(),
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				DB: &mock.DBMock{
					BeginTxFunc: func(ctx context.Context, txOpts pgx.TxOptions) (db.Transactioner, error) {
						return &mock.TransactionerMock{
							RollbackFunc: func(ctx context.Context) error {
								return nil
							},
							GetUserByIdFunc: func(ctx context.Context, userID uuid.UUID) (db.GetUserByIdRow, error) {
								return db.GetUserByIdRow{
									ID:           testUser.User.ID,
									Username:     testUser.User.Username,
									Email:        testUser.User.Email,
									FullName:     testUser.User.FullName,
									Role:         testUser.User.Role.String(),
									PasswordHash: testUser.User.PasswordHash,
									PictureID:    testUser.User.PictureID,
									CreatedAt:    testUser.User.CreatedAt,
									UpdatedAt:    testUser.User.UpdatedAt,
									DeletedAt:    testUser.User.DeletedAt,
								}, nil
							},
							InsertOrUpdateUserFunc: func(ctx context.Context, arg db.InsertOrUpdateUserParams) error {
								require.Equal(t, arg, db.InsertOrUpdateUserParams{
									ID:           testUser.User.ID,
									Username:     testUser.User.Username,
									Email:        testUser.User.Email,
									FullName:     testUser.User.FullName,
									Role:         input.Role.String(),
									PasswordHash: testUser.User.PasswordHash,
									PictureID:    testUser.User.PictureID,
									CreatedAt:    testUser.User.CreatedAt,
									UpdatedAt:    arg.UpdatedAt,
								})

								require.True(t, arg.UpdatedAt.After(testUser.User.UpdatedAt))

								return nil
							},
							CommitFunc: func(ctx context.Context) error {
								return nil
							},
						}, nil
					},
				},
				SessionService: &mock.SessionServiceMock{
					UpdateRoleByUserIDFunc: func(ctx context.Context, userID uuid.UUID, role domain.Role) error {
						require.Equal(t, userID, testUser.User.ID)
						require.Equal(t, role, input.Role)
						return nil
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				input: input,
			},
			wantErr: false,
		},
		{
			name: "User Not Found",
			fields: fields{
				DB: &mock.DBMock{
					BeginTxFunc: func(ctx context.Context, txOpts pgx.TxOptions) (db.Transactioner, error) {
						return &mock.TransactionerMock{
							RollbackFunc: func(ctx context.Context) error {
								return nil
							},
							GetUserByIdFunc: func(ctx context.Context, userID uuid.UUID) (db.GetUserByIdRow, error) {
								return db.GetUserByIdRow{}, pgx.ErrNoRows
							},
						}, nil
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				input: input,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Tokens:         tt.fields.Tokens,
				Mail:           tt.fields.Mail,
				DB:             tt.fields.DB,
				SessionService: tt.fields.SessionService,
			}
			if err := u.ChangeRole(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("User.ChangeRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Delete(t *testing.T) {
	t.Parallel()

	type fields struct {
		Tokens         tokenizer.Tokenizer
		Mail           mailer.Mailer
		DB             db.DB
		SessionService domain.SessionService
	}

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}

	testUser := getFakeUser(t)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				DB: &mock.DBMock{
					GetUserByIdFunc: func(ctx context.Context, userID uuid.UUID) (db.GetUserByIdRow, error) {
						require.Equal(t, userID, testUser.User.ID)
						return db.GetUserByIdRow{
							ID:           userID,
							Username:     testUser.User.Username,
							Email:        testUser.User.Email,
							FullName:     testUser.User.FullName,
							Role:         testUser.User.Role.String(),
							PasswordHash: testUser.User.PasswordHash,
							PictureID:    testUser.User.PictureID,
							CreatedAt:    testUser.User.CreatedAt,
							UpdatedAt:    testUser.User.UpdatedAt,
							DeletedAt:    testUser.User.DeletedAt,
						}, nil
					},
					SoftDeleteUserFunc: func(ctx context.Context, userID uuid.UUID) error {
						require.Equal(t, userID, testUser.User.ID)
						return nil
					},
				},
				SessionService: &mock.SessionServiceMock{
					DeleteAllByUserIDFunc: func(ctx context.Context, userID uuid.UUID) error {
						require.Equal(t, userID, testUser.User.ID)
						return nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				id:  testUser.User.ID,
			},
			wantErr: false,
		},
		{
			name: "User not found",
			fields: fields{
				DB: &mock.DBMock{
					GetUserByIdFunc: func(ctx context.Context, userID uuid.UUID) (db.GetUserByIdRow, error) {
						return db.GetUserByIdRow{}, pgx.ErrNoRows
					},
				},
			},
			args: args{
				ctx: context.Background(),
				id:  testUser.User.ID,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Tokens:         tt.fields.Tokens,
				Mail:           tt.fields.Mail,
				DB:             tt.fields.DB,
				SessionService: tt.fields.SessionService,
			}
			if err := u.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("User.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_One(t *testing.T) {
	t.Parallel()

	type fields struct {
		Tokens         tokenizer.Tokenizer
		Mail           mailer.Mailer
		DB             db.DB
		SessionService domain.SessionService
	}

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}

	testUser := getFakeUser(t)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				DB: &mock.DBMock{
					GetUserByIdFunc: func(ctx context.Context, userID uuid.UUID) (db.GetUserByIdRow, error) {
						require.Equal(t, userID, testUser.User.ID)
						return db.GetUserByIdRow{
							ID:           testUser.User.ID,
							Username:     testUser.User.Username,
							Email:        testUser.User.Email,
							FullName:     testUser.User.FullName,
							Role:         testUser.User.Role.String(),
							PasswordHash: testUser.User.PasswordHash,
							PictureID:    testUser.User.PictureID,
							CreatedAt:    testUser.User.CreatedAt,
							UpdatedAt:    testUser.User.UpdatedAt,
							DeletedAt:    testUser.User.DeletedAt,
						}, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				id:  testUser.User.ID,
			},
			want:    testUser.User,
			wantErr: false,
		},
		{
			name: "User not found",
			fields: fields{
				DB: &mock.DBMock{
					GetUserByIdFunc: func(ctx context.Context, userID uuid.UUID) (db.GetUserByIdRow, error) {
						return db.GetUserByIdRow{}, pgx.ErrNoRows
					},
				},
			},
			args: args{
				ctx: context.Background(),
				id:  testUser.User.ID,
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Tokens:         tt.fields.Tokens,
				Mail:           tt.fields.Mail,
				DB:             tt.fields.DB,
				SessionService: tt.fields.SessionService,
			}
			got, err := u.One(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.One() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.One() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Many(t *testing.T) {
	t.Parallel()

	var getManyUserRows = func(count int) []db.GetManyUsersRow {
		testUsers := getManyFakeUsers(t, count)
		userRows := make([]db.GetManyUsersRow, count)
		for i := range userRows {
			userRows[i] = db.GetManyUsersRow{
				ID:        testUsers.Users[i].ID,
				Username:  testUsers.Users[i].Username,
				Email:     testUsers.Users[i].Email,
				FullName:  testUsers.Users[i].FullName,
				Role:      testUsers.Users[i].Role.String(),
				PictureID: testUsers.Users[i].PictureID,
				CreatedAt: testUsers.Users[i].CreatedAt,
				UpdatedAt: testUsers.Users[i].UpdatedAt,
				DeletedAt: testUsers.Users[i].DeletedAt,
			}
		}
		return userRows
	}

	type fields struct {
		Tokens         tokenizer.Tokenizer
		Mail           mailer.Mailer
		DB             db.DB
		SessionService domain.SessionService
	}

	type args struct {
		ctx    context.Context
		filter domain.UserFilterInput
	}

	input := domain.UserFilterInput{
		Query:       null.StringFrom("This is a test query"),
		Role:        null.StringFrom(domain.RoleEditor.String()),
		ShowDeleted: null.BoolFrom(true),
		Limit:       null.IntFrom(20),
	}

	cursorObj := db.GetManyUsersParams{
		Search:        null.StringFrom("This is another test query"),
		Role:          null.StringFrom(getRandRole().String()),
		UpdatedBefore: null.TimeFrom(fake.Date()),
		ShowDeleted:   true,
		UsersLimit:    30,
	}

	cursor, err := encodeCursor(cursorObj)
	require.NoError(t, err)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Check Return Error On Invalid Role Parameter",
			args: args{
				ctx: context.Background(),
				filter: domain.UserFilterInput{
					Role: null.StringFrom("invalidRole"),
				},
			},
			wantErr: true,
		},
		{
			name: "Check All Params are Passed to DB Query Except Cursor",
			fields: fields{
				DB: &mock.DBMock{
					GetManyUsersFunc: func(ctx context.Context, arg db.GetManyUsersParams) ([]db.GetManyUsersRow, error) {
						require.Equal(t, arg, db.GetManyUsersParams{
							Search:      input.Query,
							Role:        input.Role,
							UsersLimit:  input.Limit.ValueOrZero() + 1,
							ShowDeleted: input.ShowDeleted.Bool,
						})

						userRows := getManyUserRows(int(arg.UsersLimit))

						return userRows, nil
					},
				},
			},
			args: args{
				ctx:    context.Background(),
				filter: input,
			},
			wantErr: false,
		},
		{
			name: "Check if Default Users Limit is used when limit is null",
			fields: fields{
				DB: &mock.DBMock{
					GetManyUsersFunc: func(ctx context.Context, arg db.GetManyUsersParams) ([]db.GetManyUsersRow, error) {
						require.Equal(t, arg.UsersLimit, int64(config.DEFAULT_USERS_LIST_LIMIT+1))
						userRows := getManyUserRows(int(arg.UsersLimit))
						return userRows, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "Test Passing a encoded cursor of db.GetManyUsers Param",
			fields: fields{
				DB: &mock.DBMock{
					GetManyUsersFunc: func(ctx context.Context, arg db.GetManyUsersParams) ([]db.GetManyUsersRow, error) {
						require.Equal(t, arg, db.GetManyUsersParams{
							Search:        cursorObj.Search,
							Role:          cursorObj.Role,
							UpdatedBefore: cursorObj.UpdatedBefore,
							ShowDeleted:   cursorObj.ShowDeleted,
							UsersLimit:    cursorObj.UsersLimit + 1,
						})
						userRows := getManyUserRows(int(arg.UsersLimit))
						return userRows, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				filter: domain.UserFilterInput{
					Cursor: null.StringFrom(cursor),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Tokens:         tt.fields.Tokens,
				Mail:           tt.fields.Mail,
				DB:             tt.fields.DB,
				SessionService: tt.fields.SessionService,
			}

			got, err := u.Many(tt.args.ctx, tt.args.filter)

			if (err != nil) != tt.wantErr {
				t.Errorf("User.Many() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			nodeCount := len(got.Nodes)
			cursorCount := len(got.Cursors)

			if cursorCount != nodeCount {
				t.Errorf("Number of User nodes and cursors must be equal but got %d nodes and %d cursors",
					nodeCount, cursorCount)
				return
			}

			if !tt.args.filter.Cursor.Valid {
				if tt.args.filter.Limit.Valid && nodeCount != int(tt.args.filter.Limit.ValueOrZero()) {
					t.Errorf("Number of nodes returned %d is not equal to the limit requested %d",
						nodeCount, tt.args.filter.Limit.ValueOrZero())
					return
				}

				if tt.args.filter.Limit.Valid {
					limit := tt.args.filter.Limit.ValueOrZero()

					if nodeCount < int(limit) && got.HasMore {
						t.Errorf("Cannot have next page when node count is less than the requested limit. "+
							"Got %d user nodes for the limit of %d", nodeCount, limit)
					}
					return
				}
			}

			if cursorCount > 0 {
				if got.StartCursor.ValueOrZero() != got.Cursors[0] {
					t.Errorf("StartCursor should be equal to the first cursor in the cursors array, "+
						"but StartCursor = %s and cursors[0] = %s",
						got.StartCursor.ValueOrZero(), got.Cursors[0],
					)
					return
				}

				if got.EndCursor.ValueOrZero() != got.Cursors[cursorCount-1] {
					t.Errorf("EndCursor should be equal to the last cursor in the cursors array, "+
						"but EndCursor = %s and cursors[%d] = %s",
						got.StartCursor.ValueOrZero(), cursorCount-1, got.Cursors[cursorCount-1],
					)
					return
				}
			}

		})
	}
}

func TestUser_AllByIDS(t *testing.T) {
	t.Parallel()

	type fields struct {
		Tokens         tokenizer.Tokenizer
		Mail           mailer.Mailer
		DB             db.DB
		SessionService domain.SessionService
	}

	type args struct {
		ctx context.Context
		ids []uuid.UUID
	}

	testUsers := getManyFakeUsers(t, 100)

	for _, user := range testUsers.Users {
		user.PasswordHash = null.String{}
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.User
		wantErr bool
	}{
		{
			name: "Test Get all users with the given IDS",
			fields: fields{
				DB: &mock.DBMock{
					GetAllUsersInIDSFunc: func(ctx context.Context, userIds []uuid.UUID) ([]db.GetAllUsersInIDSRow, error) {
						require.Equal(t, userIds, testUsers.UserIDS)
						userRows := make([]db.GetAllUsersInIDSRow, len(testUsers.Users))

						for i, user := range testUsers.Users {
							userRows[i] = db.GetAllUsersInIDSRow{
								ID:        user.ID,
								Username:  user.Username,
								Email:     user.Email,
								FullName:  user.FullName,
								Role:      user.Role.String(),
								PictureID: user.PictureID,
								CreatedAt: user.CreatedAt,
								UpdatedAt: user.UpdatedAt,
								DeletedAt: user.DeletedAt,
							}
						}

						return userRows, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				ids: testUsers.UserIDS,
			},
			want:    testUsers.Users,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Tokens:         tt.fields.Tokens,
				Mail:           tt.fields.Mail,
				DB:             tt.fields.DB,
				SessionService: tt.fields.SessionService,
			}
			got, err := u.AllByIDS(tt.args.ctx, tt.args.ids...)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.AllByIDS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.AllByIDS() = %v, want %v", got, tt.want)
			}
		})
	}
}
