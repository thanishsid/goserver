package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/mailer"
	"github.com/thanishsid/goserver/infrastructure/tokenizer"
	"github.com/thanishsid/goserver/input"
	"github.com/thanishsid/goserver/repository"
)

type userService struct {
	*ServiceDeps
}

var _ domain.UserService = (*userService)(nil)

// Sends an email to the user with a link that contains a JWE token containing information about the user.
func (u *userService) InitRegistration(ctx context.Context, input input.InitRegistration) error {

	if err := input.Validate(); err != nil {
		return err
	}

	regToken, err := u.Tokens.CreateToken(ctx, tokenizer.RegistrationClaims{
		Username:  input.Username,
		Email:     input.Email,
		FullName:  input.FullName,
		Role:      input.Role,
		PictureID: input.PictureID,
		Expiry:    time.Now().Add(config.REGISTRATION_TOKEN_TTL),
	})
	if err != nil {
		return fmt.Errorf("InitRegistration.CreateToken: %w", err)
	}

	regLink, err := url.Parse(input.ClientRegistrationLink)
	if err != nil {
		return fmt.Errorf("InitRegistration.ParseRegistrationLink: %w", err)
	}

	regLink.Path = path.Join(regLink.Path, regToken)

	if err = u.Mail.SendLinkMail(ctx, mailer.LinkMailTemplateData{
		Title:            "Account Registration",
		PrimaryMessage:   fmt.Sprintf("Hi %s thank you for registering with us.", input.FullName),
		SecondaryMessage: "Please click the link below to complete your registration.",
		Link:             regLink.String(),
	}); err != nil {
		return fmt.Errorf("InitRegistration.SendLinkMail: %w", err)
	}

	return nil
}

// Parses the user information in the registration token and creates the new user.
func (u *userService) CompleteRegistration(ctx context.Context, input input.CompleteRegistration) (*domain.User, error) {

	if err := input.Validate(); err != nil {
		return nil, err
	}

	var claims tokenizer.RegistrationClaims

	if err := u.Tokens.GetClaims(ctx, input.RegistrationToken, &claims); err != nil {
		return nil, fmt.Errorf("CompleteRegistration.GetClaims: %w", err)
	}

	now := time.Now()

	user := domain.User{
		ID:        uuid.New(),
		Email:     claims.Email,
		Username:  claims.Username,
		FullName:  claims.FullName,
		Role:      claims.Role,
		PictureID: claims.PictureID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := user.CreatePasswordHash(input.Password); err != nil {
		return nil, fmt.Errorf("CompleteRegistration.CreatePasswordHash: %w", err)
	}

	if err := u.Repo.UserRepository().SaveOrUpdate(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Update a user.
func (u *userService) Update(ctx context.Context, userID uuid.UUID, input input.UserUpdate) error {

	if err := input.Validate(); err != nil {
		return err
	}

	return u.Repo.ExecTx(ctx, pgx.TxOptions{}, func(ctx context.Context, repo repository.TxRepository) error {

		r := repo.UserRepository()

		targetUser, err := r.OneByID(ctx, userID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return ErrNotFound
			}
		}

		updatedUser := *targetUser
		updatedUser.Username = input.Username
		updatedUser.FullName = input.FullName
		updatedUser.PictureID = input.PictureID

		if updatedUser.IsEqual(targetUser) {
			return ErrNoChange
		}

		updatedUser.UpdatedAt = time.Now()
		return r.SaveOrUpdate(ctx, &updatedUser)
	})
}

// Change user role.
func (u *userService) ChangeRole(ctx context.Context, input input.RoleChange) error {

	if err := input.Validate(); err != nil {
		return err
	}

	return u.Repo.ExecTx(ctx, pgx.TxOptions{}, func(ctx context.Context, repo repository.TxRepository) error {

		r := repo.UserRepository()

		targetUser, err := r.OneByID(ctx, input.UserID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return ErrNotFound
			}
		}

		updatedUser := *targetUser
		updatedUser.Role = input.Role

		if targetUser.IsEqual(&updatedUser) {
			return ErrNoChange
		}

		return r.SaveOrUpdate(ctx, &updatedUser)
	})
}

// Delete a user by id.
func (u *userService) Delete(ctx context.Context, id uuid.UUID) error {
	return u.Repo.ExecTx(ctx, pgx.TxOptions{}, func(ctx context.Context, repo repository.TxRepository) error {

		r := repo.UserRepository()

		targetUser, err := r.OneByID(ctx, id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return ErrNotFound
			}
		}

		return r.Delete(ctx, targetUser.ID)
	})
}

// Find a user by id.
func (u *userService) User(ctx context.Context, id uuid.UUID) (*domain.User, error) {

	r := u.Repo.UserRepository()

	user, err := r.OneByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

// Find users with specific filters and returns a cursor for pagination.
func (u *userService) Users(ctx context.Context, filter input.UserFilter) (*domain.ListWithCursor[domain.User], error) {
	var baseFilter input.UserFilterBase
	var err error

	if filter.Cursor.Valid {
		baseFilter, err = filter.GetFilterBaseFromCursor()
		if err != nil {
			return nil, err
		}
	} else {
		baseFilter = filter.UserFilterBase
	}

	if baseFilter.Limit.ValueOrZero() == 0 {
		baseFilter.Limit = null.IntFrom(config.DEFAULT_USERS_LIST_LIMIT)
	}

	r := u.Repo.UserRepository()

	// Limit incremented by 1 to find if next page exists based on
	// whether the returned array size is equal to the speculation limit.
	speculationLimit := baseFilter.Limit.ValueOrZero() + 1

	users, err := r.Many(ctx, input.UserFilterBase{
		Query:        baseFilter.Query,
		Role:         baseFilter.Role,
		ShowDeleted:  baseFilter.ShowDeleted,
		Limit:        null.IntFrom(speculationLimit),
		UpdatedAfter: baseFilter.UpdatedAfter,
	})
	if err != nil {
		return nil, err
	}

	if len(users) < int(speculationLimit) {
		return &domain.ListWithCursor[domain.User]{
			Data: users,
		}, nil
	}

	nextCursor, err := baseFilter.CreateCursor()
	if err != nil {
		return nil, err
	}

	return &domain.ListWithCursor[domain.User]{
		Data:       users[:baseFilter.Limit.ValueOrZero()],
		NextCursor: null.StringFrom(nextCursor),
	}, nil
}
