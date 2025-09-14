package domain

import "errors"

var (
	ID = "id"
	// authentication key
	AuthSecret = "AUTH_SECRET"

	RoleAdmin = "admin"
	RoleUser  = "user"

	// internal errors
	ErrUserDoesNotExist    = errors.New("user does not exist")
	ErrUserNotAuthorized   = errors.New("user not authorized")
	ErrProductDoesNotExist = errors.New("product does not exist")

	// response errors
	ErrProductDoesNotExistResponse = "accessing product does not exist"
	ErrProductFetchingFailed       = "fetching product failed"
	ErrProductsFetchingFailed      = "fetching products failed"
	ErrPoorlyFormedRequest         = "request doesn't contain all the necessary data"
)

type ContextKey string

var (
	// request context and query keys
	UserIDKey ContextKey = "userId"
	RoleKey   ContextKey = "role"
)
