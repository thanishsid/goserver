package security

import "errors"

type Role int32

const (
	Administrator Role = iota + 1
	Manager
	Editor
	Affiliate
	User
	Guest
)

var roleNames = map[Role]string{
	Administrator: "Administrator",
	Manager:       "Manager",
	Editor:        "Editor",
	Affiliate:     "Affiliate",
	User:          "User",
	Guest:         "Guest",
}

func (r Role) ValidateRole() error {
	switch r {
	case Administrator, Manager, Editor, Affiliate, User, Guest:
		return nil
	}

	return errors.New("invalid role")
}

func (r Role) GetName() string {
	return roleNames[r]
}
