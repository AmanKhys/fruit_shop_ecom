package usecase

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user_service/internal/domain"
)

type UserUsecase interface {
	Register(ctx context.Context, email, pasword string) (domain.User, error)
	Login(ctx context.Context, email, password string) (domain.User, error)
}

type userUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Register(ctx context.Context, email, password string) (domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return u.repo.CreateUser(ctx, user)
}

func (u *userUsecase) Login(ctx context.Context, email, password string) (domain.User, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return domain.User{}, errors.New("invalid credentials")
	}
	return user, nil
}
