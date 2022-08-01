package domain

import (
	"context"
	"time"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type UserService interface {
	// Sends an email to the user with a link that contains a JWE token containing information about the user.
	InitRegistration(ctx context.Context, input InitRegistrationInput) error

	// Parses the user information in the registration token and creates the new user.
	CompleteRegistration(ctx context.Context, input CompleteRegistrationInput) (*User, error)

	// Update a user.
	Update(ctx context.Context, userID uuid.UUID, input UserUpdateInput) error

	// Change user role.
	ChangeRole(ctx context.Context, input RoleChangeInput) error

	// Delete a user by id.
	Delete(ctx context.Context, id uuid.UUID) error

	// Find a user by id.
	One(ctx context.Context, id uuid.UUID) (*User, error)

	// Find many users with specific filters.
	Many(ctx context.Context, filter UserFilterInput) (*ListData[User], error)

	// Find all users in a set ids.
	AllByIDS(ctx context.Context, ids ...uuid.UUID) ([]*User, error)

	// Verify provided credentials and return the user.
	VerifyCredentials(ctx context.Context, input VerifyCredentialsInput) (*User, error)
}

type User struct {
	ID           uuid.UUID     `json:"id"`
	Email        string        `json:"email"`
	Username     string        `json:"username"`
	FullName     string        `json:"fullName"`
	PasswordHash string        `json:"-"`
	Role         Role          `json:"role"`
	PictureID    uuid.NullUUID `json:"-"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	DeletedAt    null.Time     `json:"deletedAt"`
}

//*----------- USER FORMS ---------------

// Registration initialization input.
type InitRegistrationInput struct {
	Email                  string `json:"email"`
	FullName               string `json:"fullName"`
	Role                   Role   `json:"role"`
	ClientRegistrationLink string `json:"clientRegistrationLink"`
}

func (f InitRegistrationInput) Validate() error {
	return vd.ValidateStruct(&f,
		vd.Field(&f.Email, vd.Required, is.Email),
		vd.Field(&f.FullName, vd.Required),
		vd.Field(&f.Role),
		vd.Field(&f.ClientRegistrationLink, vd.Required, is.URL),
	)
}

// Registration Completion input.
type CompleteRegistrationInput struct {
	Token     string        `json:"token"`
	Username  string        `json:"username"`
	PictureID uuid.NullUUID `json:"pictureId"`
	Password  string        `json:"password"`
}

func (f CompleteRegistrationInput) Validate() error {
	return vd.ValidateStruct(&f,
		vd.Field(&f.Token, vd.Required),
		vd.Field(&f.Username, vd.Required),
		vd.Field(&f.Password, vd.Required),
	)
}

// Update user input.
type UserUpdateInput struct {
	Username  string        `json:"username"`
	FullName  string        `json:"fullName"`
	PictureID uuid.NullUUID `json:"pictureId"`
}

func (f UserUpdateInput) Validate() error {
	return vd.ValidateStruct(&f,
		vd.Field(&f.Username, vd.Required),
		vd.Field(&f.FullName, vd.Required),
	)
}

// User role change input.
type RoleChangeInput struct {
	UserID uuid.UUID `json:"userId"`
	Role   Role      `json:"role"`
}

func (f RoleChangeInput) Validate() error {
	return vd.ValidateStruct(&f,
		vd.Field(&f.UserID, is.UUIDv4),
		vd.Field(&f.Role, vd.By(func(_ interface{}) error {
			return f.Role.ValidateRole()
		})),
	)
}

type UserFilterInput struct {
	Query       null.String `schema:"search"`
	Role        null.String `schema:"role"`
	ShowDeleted null.Bool   `schema:"showDeleted"`
	Limit       null.Int    `schema:"limit" `
	Cursor      null.String `schema:"cursor"`
}

func (f UserFilterInput) Validate() error {
	return vd.ValidateStruct(&f,
		vd.Field(&f.Role, vd.When(f.Role.Valid, vd.By(func(_ interface{}) error {
			role := Role(f.Role.ValueOrZero())
			return role.ValidateRole()
		}))),
	)
}

type VerifyCredentialsInput struct {
	Email    string `json:"string"`
	Password string `json:"password"`
}

func (i VerifyCredentialsInput) Validate() error {
	return vd.ValidateStruct(&i,
		vd.Field(&i.Email, vd.Required, is.Email),
		vd.Field(&i.Password, vd.Required),
	)
}
