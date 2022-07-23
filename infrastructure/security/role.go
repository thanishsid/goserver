package security

import (
	"errors"
)

type Role string

const (
	Administrator Role = "Administrator"
	Manager       Role = "Manager"
	Editor        Role = "Editor"
	Affiliate     Role = "Affiliate"
	User          Role = "User"
	Guest         Role = "Guest"
)

var AllRoles []Role = []Role{
	Administrator,
	Manager,
	Editor,
	Affiliate,
	User,
	Guest,
}

// Validate role.
func (r Role) ValidateRole() error {
	switch r {
	case Administrator, Manager, Editor, Affiliate, User, Guest:
		return nil
	}

	return errors.New("invalid role")
}
