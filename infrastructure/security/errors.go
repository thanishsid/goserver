package security

import "errors"

var ErrUnauthorized = errors.New("unauthorized, you do not have sufficient permissions to perform this action")

var ErrNotLoggedIn = errors.New("unable to perform operation, need to be logged in")
