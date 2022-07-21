package service

import "errors"

var ErrNotFound = errors.New("unable to find the target resource")
var ErrNoChange = errors.New("redundant operation, no change from existing value")
