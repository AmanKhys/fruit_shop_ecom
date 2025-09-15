package usecase

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"product_service/internal/domain"
)

type ProductUsecase interface {
	GetProductByID(ctx context.Context, id int) (domain.Product, error)
	GetProducts(ctx context.Context, min, max float64) ([]domain.Product, error)
	CreateProduct(ctx context.Context, p domain.Product) (domain.Product, error)
	UpdateProductByID(ctx context.Context, p domain.Product) (domain.Product, error)
	DeleteProductByID(ctx context.Context, id int) error
}

type productUsecase struct {
	repo ProductRepo
}

// this function returns a new &productUsecase{} that satisfies the ProductUsecase interface
func NewProductUsecase(repo ProductRepo) ProductUsecase {
	return &productUsecase{repo: repo}
}

// a fault tolerant filtering usecase to get filtered products
func (u *productUsecase) GetProducts(ctx context.Context, min, max float64) ([]domain.Product, error) {
	if min >= max {
		min = math.MinInt
		max = math.MaxInt
	}
	role, ok := ctx.Value(domain.RoleKey).(domain.ContextKey)
	if !ok {
		return u.repo.GetFilteredProducts(ctx, min, max)
	}
	if role == domain.RoleAdmin {
		return u.repo.GetAllProductsForAdmin(ctx, min, max)
	}
	return u.repo.GetFilteredProducts(ctx, min, max)

}

func (u *productUsecase) GetProductByID(ctx context.Context, id int) (domain.Product, error) {
	role, ok := ctx.Value(domain.RoleKey).(domain.ContextKey)
	if !ok || role != domain.RoleAdmin {
		p, err := u.repo.GetProductByID(ctx, id)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, domain.ErrProductDoesNotExist
		}
		return p, nil
	}
	p, err := u.repo.GetProductByIDForAdmin(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Product{}, domain.ErrProductDoesNotExist
	}
	return p, nil
}

func (u *productUsecase) CreateProduct(ctx context.Context, p domain.Product) (domain.Product, error) {
	role, ok := ctx.Value(domain.RoleKey).(domain.ContextKey)
	if !ok || role != domain.RoleAdmin {
		return domain.Product{}, domain.ErrUserNotAuthorized
	}

	if err := ValidateProduct(p); err != nil {
		return domain.Product{}, err
	}

	return u.repo.CreateProduct(ctx, domain.Product{Name: p.Name, Price: p.Price, Stock: p.Stock})

}

func (u *productUsecase) UpdateProductByID(ctx context.Context, p domain.Product) (domain.Product, error) {
	role, ok := ctx.Value(domain.RoleKey).(domain.ContextKey)
	if !ok || role != domain.RoleAdmin {
		return domain.Product{}, domain.ErrUserNotAuthorized
	}

	if err := ValidateProduct(p); err != nil {
		return domain.Product{}, err
	}

	p, err := u.repo.UpdateProduct(ctx, domain.Product{ID: p.ID, Name: p.Name, Price: p.Price, Stock: p.Stock})

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Product{}, domain.ErrProductDoesNotExist
	}
	return p, nil
}

func (u *productUsecase) DeleteProductByID(ctx context.Context, id int) error {
	role, ok := ctx.Value(domain.RoleKey).(domain.ContextKey)
	if !ok || role != domain.RoleAdmin {
		return domain.ErrUserDoesNotExist
	}
	err := u.repo.DeleteProductByID(ctx, id)
	if err == sql.ErrNoRows {
		return domain.ErrProductDoesNotExist
	}
	return nil
}
