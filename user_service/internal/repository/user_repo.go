package repository

import (
	"context"
	"database/sql"
	"user_service/internal/domain"
	"user_service/internal/infrastructure/db/sqlc"
)

type userRepo struct {
	q  *sqlc.Queries
	db *sql.DB
}

func NewUserRepo(db *sql.DB, q *sqlc.Queries) *userRepo {
	return &userRepo{db: db, q: q}
}

func (r *userRepo) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	u, err := r.q.CreateUser(ctx, r.db, sqlc.CreateUserParams{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:    u.ID,
		Email: u.Email,
	}, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.q.GetUserByEmail(ctx, r.db, email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}, nil
}

func (r *userRepo) CreateAdminUser(ctx context.Context, u domain.User) (domain.User, error) {
	admin, err := r.q.CreateAdminUser(ctx, r.db, sqlc.CreateAdminUserParams{
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{ID: admin.ID, Email: admin.Email, Role: u.Role}, nil
}
