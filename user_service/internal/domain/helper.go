package domain

import (
	"errors"
	"regexp"
)

func ValidateUser(u User) error {
	emailValidator := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	var errs []error
	if len(u.Password) < 6 {
		errs = append(errs, ErrInvalidPassword)
	}
	if !emailValidator.MatchString(u.Email) {
		errs = append(errs, ErrInvalidEmail)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
