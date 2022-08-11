package service

import "errors"

var ErrInvalidCredentials = errors.New("the provided credentials are invalid")

var ErrNotFound = errors.New("unable to find the target resource")
var ErrNoChange = errors.New("redundant operation, no change from existing value")

var ErrIncompleteFile = errors.New("incomplete file, the received file is broken")
var ErrFileAlreadyExists = errors.New("duplicate filename, uploaded file already exists")
