package domain

import "errors"

var (
	RoleAdmin = "admin"
	RoleUser  = "user"

	ErrUserDoesNotExist  = errors.New("user does not exist")
	ErrUserNotAuthorized = errors.New("user not authorized")

	ErrProductDoesNotExist = errors.New("product does not exist")
)
