package domain

import (
	"errors"
)

type Role string

const (
	RoleAdministrator Role = "admin"
	RoleManager       Role = "manager"
	RoleEditor        Role = "editor"
	RoleClient        Role = "client"
)

var AllRoles []Role = []Role{
	RoleAdministrator,
	RoleManager,
	RoleEditor,
	RoleClient,
}

// Get string value of role.
func (r Role) String() string {
	return string(r)
}

// Validate default roles excluding internal roles.
func (r Role) ValidateRole() error {
	switch r {
	case RoleAdministrator, RoleManager, RoleEditor, RoleClient:
		return nil
	}

	return errors.New("invalid role")
}
