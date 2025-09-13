package usecase

import (
	"context"
	"user_service/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	CreateAdminUser(ctx context.Context, user domain.User) (domain.User, error)
}
