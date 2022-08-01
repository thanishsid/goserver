package service

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/db"
	"github.com/thanishsid/goserver/infrastructure/mailer"
	"github.com/thanishsid/goserver/infrastructure/tokenizer"
)

func NewUserService(token tokenizer.Tokenizer, mail mailer.Mailer, dbs db.DB) domain.UserService {
	return &userService{
		Tokens: token,
		Mail:   mail,
		DB:     dbs,
	}
}

type userService struct {
	Tokens tokenizer.Tokenizer
	Mail   mailer.Mailer
	DB     db.DB
}

var _ domain.UserService = (*userService)(nil)

// Sends an email to the user with a link that contains a JWE token containing information about the user.
func (u *userService) InitRegistration(ctx context.Context, input domain.InitRegistrationInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	regToken, err := u.Tokens.CreateToken(ctx, tokenizer.RegistrationClaims{
		Email:    input.Email,
		FullName: input.FullName,
		Role:     input.Role,
		Expiry:   time.Now().Add(config.REGISTRATION_TOKEN_TTL),
	})
	if err != nil {
		return fmt.Errorf("InitRegistration.CreateToken: %w", err)
	}

	regLink, err := url.Parse(input.ClientRegistrationLink)
	if err != nil {
		return fmt.Errorf("InitRegistration.ParseRegistrationLink: %w", err)
	}

	regLink.Path = path.Join(regLink.Path, regToken)

	if err = u.Mail.SendLinkMail(ctx, mailer.LinkMailData{
		To:               input.Email,
		Subject:          "Account Registration",
		Title:            "New Account Registration",
		PrimaryMessage:   fmt.Sprintf("Hi %s thank you for registering with us.", input.FullName),
		SecondaryMessage: "Please click the link below to complete your registration.",
		Link:             regLink.String(),
	}); err != nil {
		return fmt.Errorf("InitRegistration.SendLinkMail: %w", err)
	}

	return nil
}

// Parses the user information in the registration token and creates the new user.
func (u *userService) CompleteRegistration(ctx context.Context, input domain.CompleteRegistrationInput) (*domain.User, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	var claims tokenizer.RegistrationClaims

	if err := u.Tokens.GetClaims(ctx, input.Token, &claims); err != nil {
		return nil, fmt.Errorf("CompleteRegistration.GetClaims: %w", err)
	}

	userID := uuid.New()
	now := time.Now()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if err := u.DB.InsertOrUpdateUser(ctx, db.InsertOrUpdateUserParams{
		ID:           userID,
		Username:     input.Username,
		Email:        claims.Email,
		FullName:     claims.FullName,
		Role:         string(claims.Role),
		PasswordHash: string(passwordHash),
		PictureID:    input.PictureID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}); err != nil {
		return nil, err
	}

	user := domain.User{
		ID:           userID,
		Email:        claims.Email,
		Username:     input.Username,
		FullName:     claims.FullName,
		Role:         claims.Role,
		PasswordHash: string(passwordHash),
		PictureID:    input.PictureID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	return &user, nil
}

// Update a user.
func (u *userService) Update(ctx context.Context, userID uuid.UUID, input domain.UserUpdateInput) error {

	if err := input.Validate(); err != nil {
		return err
	}

	tx, err := u.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	user, err := tx.GetUserById(ctx, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNotFound
		}

		return err
	}

	if err := tx.InsertOrUpdateUser(ctx, db.InsertOrUpdateUserParams{
		ID:           userID,
		Email:        user.Email,
		Role:         user.Role,
		PasswordHash: user.PasswordHash,
		Username:     input.Username,
		FullName:     input.FullName,
		PictureID:    input.PictureID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    time.Now(),
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Change user role.
func (u *userService) ChangeRole(ctx context.Context, input domain.RoleChangeInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	tx, err := u.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	user, err := tx.GetUserById(ctx, input.UserID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNotFound
		}

		return err
	}

	if err := tx.InsertOrUpdateUser(ctx, db.InsertOrUpdateUserParams{
		ID:           input.UserID,
		Username:     user.Username,
		Email:        user.Email,
		FullName:     user.FullName,
		Role:         string(input.Role),
		PasswordHash: user.PasswordHash,
		PictureID:    user.PictureID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    time.Now(),
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete a user by id.
func (u *userService) Delete(ctx context.Context, id uuid.UUID) error {
	user, err := u.DB.GetUserById(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	return u.DB.SoftDeleteUser(ctx, user.ID)
}

// Find a user by id.
func (u *userService) One(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	userRow, err := u.DB.GetUserById(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	user := &domain.User{
		ID:           userRow.ID,
		Email:        userRow.Email,
		Username:     userRow.Username,
		FullName:     userRow.FullName,
		PasswordHash: userRow.PasswordHash,
		Role:         domain.Role(userRow.Role),
		PictureID:    userRow.PictureID,
		CreatedAt:    userRow.CreatedAt,
		UpdatedAt:    userRow.UpdatedAt,
		DeletedAt:    userRow.DeletedAt,
	}

	return user, nil
}

// Find users with specific filters and returns a cursor for pagination.
func (u *userService) Many(ctx context.Context, filter domain.UserFilterInput) (*domain.ListData[domain.User], error) {
	var err error

	if err = filter.Validate(); err != nil {
		return nil, err
	}

	if filter.Limit.ValueOrZero() == 0 {
		filter.Limit = null.IntFrom(config.DEFAULT_USERS_LIST_LIMIT)
	}

	var dbParams db.GetManyUsersParams

	if filter.Cursor.Valid {
		dbParams, err = DecodeCursor[db.GetManyUsersParams](filter.Cursor.ValueOrZero())
		if err != nil {
			return nil, err
		}
	} else {
		dbParams = db.GetManyUsersParams{
			Search:      filter.Query,
			Role:        filter.Role,
			ShowDeleted: filter.ShowDeleted.ValueOrZero(),
			UsersLimit:  filter.Limit.ValueOrZero(),
		}
	}

	// save original limit and increment db param limit by 1 to check if next page exists.
	originalLimit := dbParams.UsersLimit
	dbParams.UsersLimit++

	// fetch the users from the database.
	userRows, err := u.DB.GetManyUsers(ctx, dbParams)
	if err != nil {
		return nil, fmt.Errorf("DB.GetManyUsers: %w", err)
	}

	listData := new(domain.ListData[domain.User])

	// if number records returned from the database is greater than the original limit
	// before incrementing by 1 then set hasNext page bool to true in the list data object.
	if len(userRows) > int(originalLimit) {
		listData.HasMore = true
	}

	for i, row := range userRows {
		if i == int(originalLimit) {
			break
		}

		user := &domain.User{
			ID:        row.ID,
			Email:     row.Email,
			Username:  row.Username,
			FullName:  row.FullName,
			Role:      domain.Role(row.Role),
			PictureID: row.PictureID,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
			DeletedAt: row.DeletedAt,
		}

		// create a s
		cursorDbParams := dbParams
		cursorDbParams.UpdatedBefore = null.TimeFrom(user.UpdatedAt)
		cursor, err := EncodeCursor(cursorDbParams)
		if err != nil {
			return nil, err
		}

		listData.Nodes = append(listData.Nodes, user)
		listData.Cursors = append(listData.Cursors, cursor)

		if i == 0 {
			listData.StartCursor = null.StringFrom(cursor)
		}

		if i == int(originalLimit)-1 {
			listData.EndCursor = null.StringFrom(cursor)
		}
	}

	return listData, nil
}

// Find all users in a set ids.
func (u *userService) AllByIDS(ctx context.Context, ids ...uuid.UUID) ([]*domain.User, error) {

	userRows, err := u.DB.GetAllUsersInIDS(ctx, ids)
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(userRows))

	for i, row := range userRows {

		user := domain.User{
			ID:        row.ID,
			Email:     row.Email,
			Username:  row.Username,
			FullName:  row.FullName,
			Role:      domain.Role(row.Role),
			PictureID: row.PictureID,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
			DeletedAt: row.DeletedAt,
		}

		users[i] = &user
	}

	return users, nil
}

// Verify provided credentials and return the user.
func (u *userService) VerifyCredentials(ctx context.Context, input domain.VerifyCredentialsInput) (*domain.User, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	user, err := u.DB.GetUserByEmail(ctx, input.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return &domain.User{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		FullName:     user.FullName,
		PasswordHash: user.PasswordHash,
		Role:         domain.Role(user.Role),
		PictureID:    user.PictureID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
	}, nil
}
