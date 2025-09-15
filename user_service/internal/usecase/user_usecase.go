package usecase

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user_service/internal/domain"
)

type UserUsecase interface {
	Register(ctx context.Context, email, pasword string) (domain.User, error)
	Login(ctx context.Context, email, password string) (domain.User, error)
	EnsureAdminExists(ctx context.Context, email, password string) (domain.User, error)
}

type userUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Register(ctx context.Context, email, password string) (domain.User, error) {
	if err := domain.ValidateUser(domain.User{Email: email, Password: password}); err != nil {
		return domain.User{}, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, domain.ErrRegisteringUser
	}
	_, err = u.repo.GetUserByEmail(ctx, email)
	if err == nil {
		return domain.User{}, domain.ErrUserAlreadyExist
	}

	user, err := u.repo.CreateUser(ctx, domain.User{Email: email, Password: string(hashedPassword)})
	if err != nil {
		return domain.User{}, domain.ErrInternalErrorFetchingUser
	}
	return user, nil
}

func (u *userUsecase) Login(ctx context.Context, email, password string) (domain.User, error) {
	if err := domain.ValidateUser(domain.User{Email: email, Password: password}); err != nil {
		return domain.User{}, err
	}
	user, err := u.repo.GetUserByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, domain.ErrUserDoesNotExist
	}
	if err != nil {
		return domain.User{}, domain.ErrInternalErrorFetchingUser
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return domain.User{}, domain.ErrInvalidCredentials
	}
	return user, nil
}

func (u *userUsecase) EnsureAdminExists(ctx context.Context, email, password string) (domain.User, error) {
	if err := domain.ValidateUser(domain.User{Email: email, Password: password}); err != nil {
		return domain.User{}, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}

	admin, err := u.repo.CreateAdminUser(ctx, domain.User{Email: email, Password: string(hashedPassword)})
	if errors.Is(err, sql.ErrNoRows) {
		return u.repo.GetUserByEmail(ctx, email)
	} else if err != nil {
		return domain.User{}, err
	}
	return admin, err
}
