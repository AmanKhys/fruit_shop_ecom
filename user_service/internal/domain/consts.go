package domain

import "errors"

var (
	ErrUserDoesNotExist          = errors.New("user does not exist")
	ErrInternalErrorFetchingUser = errors.New("internal error fetching user")
)
