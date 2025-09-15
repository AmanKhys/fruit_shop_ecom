package domain

import (
	"errors"
	"testing"
)

func TestValidateUser(t *testing.T) {
	type testCase struct {
		name     string
		u        User
		wantErrs []error
	}

	tests := []testCase{
		{
			name:     "invalid email",
			u:        User{Email: "asdff.com", Password: "1132rs1"},
			wantErrs: []error{ErrInvalidEmail},
		},
		{
			name:     "invalid password",
			u:        User{Email: "asdf@asdf.com", Password: "222"},
			wantErrs: []error{ErrInvalidPassword},
		},
		{
			name:     "invalid password and email",
			u:        User{Email: "asdf@asdfcom", Password: "222"},
			wantErrs: []error{ErrInvalidPassword, ErrInvalidEmail},
		},
		{
			name:     "happy path",
			u:        User{Email: "asdf@asdf.com", Password: "2221124dfs"},
			wantErrs: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateUser(tc.u)
			if len(tc.wantErrs) == 0 {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				return
			}

			for _, expected := range tc.wantErrs {
				if !errors.Is(err, expected) {
					t.Errorf("expected error %v, got %v", expected, err)
				}
			}

		})
	}
}
