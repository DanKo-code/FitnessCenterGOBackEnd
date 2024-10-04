package auth

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	InvalidPassword      = errors.New("invalid password")
)
