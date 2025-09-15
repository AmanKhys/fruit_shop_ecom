package usecase

import (
	"errors"
	"product_service/internal/domain"
	"regexp"
	"strings"
)

var nameValidator = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9 ]{2,}$`)

func ValidateProduct(p domain.Product) error {
	var errs []error

	if p.Price <= 0 {
		errs = append(errs, domain.ErrPriceInvalid)
	}
	if p.Stock < 0 {
		errs = append(errs, domain.ErrStockInvalid)
	}
	if !nameValidator.MatchString(strings.TrimSpace(p.Name)) {
		errs = append(errs, domain.ErrNameInvalid)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
