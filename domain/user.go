package domain

import (
	"context"
	"errors"
	"time"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/infrastructure/security"
	"github.com/thanishsid/goserver/input"
)

type UserService interface {
	// Sends an email to the user with a link that contains a JWE token containing information about the user.
	InitRegistration(ctx context.Context, input input.InitRegistration) error

	// Parses the user information in the registration token and creates the new user.
	CompleteRegistration(ctx context.Context, input input.CompleteRegistration) (*User, error)

	// Update a user.
	Update(ctx context.Context, userID uuid.UUID, input input.UserUpdate) error

	// Change user role.
	ChangeRole(ctx context.Context, input input.RoleChange) error

	// Delete a user by id.
	Delete(ctx context.Context, id uuid.UUID) error

	// Find a user by id.
	User(ctx context.Context, id uuid.UUID) (*User, error)

	// Find users with specific filters.
	Users(ctx context.Context, filter input.UserFilter) (*ListWithCursor[User], error)
}

type UserRepository interface {
	// Upsert user.
	SaveOrUpdate(ctx context.Context, user *User) error

	// Soft delete user and remove user from search index.
	Delete(ctx context.Context, id uuid.UUID) error

	// Get user by id.
	OneByID(ctx context.Context, id uuid.UUID) (*User, error)

	// Get user by email.
	OneByEmail(ctx context.Context, email string) (*User, error)

	// Get many users.
	Many(ctx context.Context, params input.UserFilterBase) ([]User, error)
}

type User struct {
	ID           uuid.UUID     `json:"id"`
	Email        string        `json:"email"`
	Username     string        `json:"username"`
	FullName     string        `json:"fullName"`
	PasswordHash string        `json:"-"`
	Role         security.Role `json:"role"`
	PictureID    uuid.NullUUID `json:"-"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	DeletedAt    null.Time     `json:"deletedAt"`
}

// validate the user object.
func (u User) Validate() error {
	return vd.ValidateStruct(&u,
		vd.Field(&u.ID, is.UUIDv4),
		vd.Field(&u.Email, vd.Required),
		vd.Field(&u.Username, vd.Required),
		vd.Field(&u.FullName, vd.Required),
		vd.Field(&u.PasswordHash, vd.Required),
		vd.Field(&u.Role, vd.By(func(_ interface{}) error {
			return u.Role.ValidateRole()
		})),
		vd.Field(&u.CreatedAt, vd.By(func(_ interface{}) error {
			if u.CreatedAt.IsZero() {
				return errors.New("invalid createdAt time")
			}
			return nil
		})),
		vd.Field(&u.UpdatedAt, vd.By(func(_ interface{}) error {
			if u.UpdatedAt.IsZero() {
				return errors.New("invalid updatedAt time")
			}
			return nil
		})),
	)
}

// Check if the user object is equal to the provided user object.
func (u User) IsEqual(c *User) bool {
	return u.ID == c.ID &&
		u.Email == c.Email &&
		u.Username == c.Username &&
		u.FullName == c.FullName &&
		u.PasswordHash == c.PasswordHash &&
		u.Role == c.Role &&
		u.PictureID == c.PictureID &&
		u.CreatedAt.Equal(c.CreatedAt) &&
		u.UpdatedAt.Equal(c.UpdatedAt) &&
		u.DeletedAt.Equal(c.DeletedAt)
}

// Hash the provided password and store it in the user object.
func (u *User) CreatePasswordHash(password string) error {

	if u.PasswordHash == "" {
		return errors.New("password hash is empty")
	}

	pswdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(pswdHash)

	return nil
}

// Compare the provided password against the passwordHash in the user object.
func (u User) ComparePassword(password string) error {

	if u.PasswordHash == "" {
		return errors.New("password hash is empty")
	}

	if password == "" {
		return errors.New("password is empty")
	}

	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}
