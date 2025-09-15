package usecase

import (
	"errors"
	"regexp"
	"user_service/internal/domain"
)

func ValidateUser(u domain.User) error {
	emailValidator := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	var errs []error
	if len(u.Password) < 6 {
		errs = append(errs, domain.ErrPasswordTooShort)
	}
	if !emailValidator.MatchString(u.Email) {
		errs = append(errs, domain.ErrInvalidEmail)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
