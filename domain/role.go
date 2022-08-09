package domain

import (
	"errors"
)

/*
	1. DO NOT CHANGE THE VALUES OF THE ROLE CONSTANTS !! THIS WILL CAUSE REDUNDANT ROLES -
		TO BE SEEDED TO THE DATABASE.

	2. ANY NEW ROLES ADDED TO THE CONSTANTS MUST BE ADDED TO THE 'Roles' SLICE WITH A SUITABLE NAME.
*/

const (
	RoleAdmin   Role = "admin"
	RoleManager Role = "manager"
	RoleEditor  Role = "editor"
	RoleClient  Role = "client"
)

type RoleInfo []RoleDetails

var Roles RoleInfo = RoleInfo{
	{
		ID:   RoleAdmin,
		Name: "Administrator",
	},
	{
		ID:   RoleManager,
		Name: "Manager",
	},
	{
		ID:   RoleEditor,
		Name: "Editor",
	},
	{
		ID:   RoleClient,
		Name: "Client",
	},
}

// Init Function.
func init() {
	ids := make(map[Role]struct{})
	names := make(map[string]struct{})

	for _, v := range Roles {
		_, idExists := ids[v.ID]
		if idExists {
			panic("duplicate role id found")
		}
		ids[v.ID] = struct{}{}

		_, nameExists := names[v.Name]
		if nameExists {
			panic("duplicate role name found")
		}
		names[v.Name] = struct{}{}
	}
}

type Role string

type RoleDetails struct {
	ID   Role
	Name string
}

// Get string value of role.
func (r Role) String() string {
	return string(r)
}

// Validate default roles excluding internal roles.
func (r Role) ValidateRole() error {
	for _, rd := range Roles {
		if r == rd.ID {
			return nil
		}
	}

	return errors.New("invalid role")
}
