package domain

import (
	"errors"
	"net/mail"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"gopkg.in/guregu/null.v4"
)

type ListData[T any] struct {
	Nodes       []*T
	Cursors     []string
	HasMore     bool
	StartCursor null.String
	EndCursor   null.String
}

type EmailStringType interface {
	string | null.String
}

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
