package domain

import "errors"

var (
	ErrUserAlreadyExist          = errors.New("user already exists")
	ErrUserDoesNotExist          = errors.New("user does not exist")
	ErrInternalErrorFetchingUser = errors.New("internal error fetching user")
	ErrRegisteringUser           = errors.New("error registering user")

	ErrInvalidPassword    = errors.New("password too short")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrInvalidCredentials = errors.New("invalid credentials")

	JwtSecret = "JWT_SECRET"
)
