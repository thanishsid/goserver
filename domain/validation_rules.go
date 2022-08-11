package domain

import (
	"errors"
	"fmt"
	"net/mail"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"gopkg.in/guregu/null.v4"
)

const PASSWORD_MIN_ENTROPY_BITS = 60

// Email validation rule.
var IsEmail = vd.By(func(value interface{}) error {
	switch t := value.(type) {
	case string:
		_, err := mail.ParseAddress(t)
		if err != nil {
			return errors.New("invalid email address")
		}
	case *string:
		_, err := mail.ParseAddress(null.StringFromPtr(t).ValueOrZero())
		if err != nil {
			return errors.New("invalid email address")
		}
	case null.String:
		_, err := mail.ParseAddress(t.ValueOrZero())
		if err != nil {
			return errors.New("invalid email address")
		}
	default:
		return errors.New("invalid email address data type")
	}

	return nil
})

// Role validation rule.
var IsRole = vd.By(func(value interface{}) error {
	switch t := value.(type) {
	case Role:
		return t.ValidateRole()
	case null.String:
		return Role(t.ValueOrZero()).ValidateRole()
	case *string:
		return Role(null.StringFromPtr(t).ValueOrZero()).ValidateRole()
	default:
		return errors.New("invalid role data type")
	}
})

var IsValidPassword = vd.By(func(value interface{}) error {
	var password string

	switch v := value.(type) {
	case string:
		password = v
	case *string:
		password = *v
	case null.String:
		password = v.ValueOrZero()
	default:
		return fmt.Errorf("invalid password data type")
	}

	return passwordvalidator.Validate(password, PASSWORD_MIN_ENTROPY_BITS)
})
