package security

import "errors"

var ErrUnauthorized = errors.New("you do not have sufficient permissions to perform this action")

var ErrNotLoggedIn = errors.New("you must be logged in to perform this action")

var ErrMustBeLoggedOut = errors.New("you must log out from the current account to perform this action")
