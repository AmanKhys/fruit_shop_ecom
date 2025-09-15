package usecase

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"user_service/internal/domain"
)

func TestEnsureAdminExists(t *testing.T) {
	type testCase struct {
		name        string
		repo        UserRepository
		email       string
		password    string
		expectError error
	}

	tests := []testCase{
		{
			name:        "invalid email",
			repo:        &fakeRepo{},
			email:       "adminroot.com",
			password:    "pass@123",
			expectError: domain.ErrInvalidEmail,
		},
		{
			name:        "invalid password",
			repo:        &fakeRepo{},
			email:       "pass@root.com",
			password:    "wron",
			expectError: domain.ErrInvalidPassword,
		},
		{
			name:        "happy path",
			repo:        &fakeRepo{},
			email:       "admin@root.com",
			password:    "pass@123",
			expectError: nil,
		},
		{
			name: "user already exists",
			repo: &fakeRepo{
				user: func() domain.User {
					hash, _ := bcrypt.GenerateFromPassword([]byte("pass@123"), bcrypt.DefaultCost)
					return domain.User{Email: "admin@root.com", Password: string(hash)}
				}(),
			},
			email:       "admin@mail.com",
			password:    "holyman",
			expectError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := NewUserUsecase(tc.repo)

			_, err := u.EnsureAdminExists(context.Background(), tc.email, tc.password)
			if !errors.Is(err, tc.expectError) {
				t.Errorf("got err = %v, want %v", err, tc.expectError)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	type testCase struct {
		name        string
		repo        UserRepository
		email       string
		password    string
		expectError error
	}

	tests := []testCase{
		{
			name:        "invalid email",
			repo:        &fakeRepo{},
			email:       "nouser@mailcom",
			password:    "secret",
			expectError: domain.ErrInvalidEmail,
		},
		{
			name:        "invalid password",
			repo:        &fakeRepo{},
			email:       "test@example.com",
			password:    "wron",
			expectError: domain.ErrInvalidPassword,
		},
		{
			name:        "happy path",
			repo:        &fakeRepo{},
			email:       "test@example.com",
			password:    "validPass",
			expectError: nil,
		},
		{
			name: "user already exists",
			repo: &fakeRepo{
				user: func() domain.User {
					hash, _ := bcrypt.GenerateFromPassword([]byte("holyman"), bcrypt.DefaultCost)
					return domain.User{Email: "vivek@vs.com", Password: string(hash)}
				}(),
			},
			email:       "vivek@vs.com",
			password:    "holyman",
			expectError: domain.ErrUserAlreadyExist,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := NewUserUsecase(tc.repo)

			_, err := u.Register(context.Background(), tc.email, tc.password)
			if !errors.Is(err, tc.expectError) {
				t.Errorf("got err = %v, want %v", err, tc.expectError)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	type testCase struct {
		name        string
		repo        UserRepository
		email       string
		password    string
		expectError error
	}

	tests := []testCase{
		{
			name:        "user does not exist",
			repo:        &fakeRepo{err: sql.ErrNoRows},
			email:       "nouser@mail.com",
			password:    "secret",
			expectError: domain.ErrUserDoesNotExist,
		},
		{
			name: "invalid password",
			repo: &fakeRepo{
				user: func() domain.User {
					hash, _ := bcrypt.GenerateFromPassword([]byte("validPass"), bcrypt.DefaultCost)
					return domain.User{Email: "test@example.com", Password: string(hash)}
				}(),
			},
			email:       "test@example.com",
			password:    "wrongpass",
			expectError: domain.ErrInvalidCredentials,
		},
		{
			name: "happy path",
			repo: &fakeRepo{
				user: func() domain.User {
					hash, _ := bcrypt.GenerateFromPassword([]byte("validPass"), bcrypt.DefaultCost)
					return domain.User{Email: "test@example.com", Password: string(hash)}
				}(),
			},
			email:       "test@example.com",
			password:    "validPass",
			expectError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := NewUserUsecase(tc.repo)

			_, err := u.Login(context.Background(), tc.email, tc.password)
			if !errors.Is(err, tc.expectError) {
				t.Errorf("got err = %v, want %v", err, tc.expectError)
			}
		})
	}
}

type fakeRepo struct {
	user domain.User
	err  error
}

func (f *fakeRepo) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if f.err != nil {
		return domain.User{}, f.err
	}
	f.user = user
	return f.user, nil
}

func (f *fakeRepo) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	if f.err != nil {
		return domain.User{}, f.err
	}
	if f.user.Email == "" {
		return domain.User{}, sql.ErrNoRows
	}
	return f.user, nil
}

func (f *fakeRepo) CreateAdminUser(ctx context.Context, user domain.User) (domain.User, error) {
	return f.CreateUser(ctx, user)
}
