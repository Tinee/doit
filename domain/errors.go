package domain

import (
	"errors"
)

// General errors
var (
	ErrInternal = errors.New("Internal error")
)

// User database errors
var (
	ErrUserExists        = errors.New("User already exists")
	ErrorUserNotFound    = errors.New("User not found in collection")
	ErrPasswordMissmatch = errors.New("Password is not matching")
)
