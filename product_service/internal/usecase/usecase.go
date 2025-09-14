package usecase

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"product_service/internal/domain"
)

type ProductUsecase interface {
	GetFilteredProducts(ctx context.Context, min, max float64) ([]domain.Product, error)
	GetProductByID(ctx context.Context, id int) (domain.Product, error)
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
	GetAllProductsForAdmin(ctx context.Context) ([]domain.Product, error)
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
func (u *productUsecase) GetFilteredProducts(ctx context.Context, min, max float64) ([]domain.Product, error) {
	if min >= max {
		min = math.MinInt
		max = math.MaxInt
	}
	return u.repo.GetFilteredProducts(ctx, min, max)
}

func (u *productUsecase) GetProductByID(ctx context.Context, id int) (domain.Product, error) {
	p, err := u.repo.GetProductByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Product{}, domain.ErrProductDoesNotExist
	}
	return p, nil
}

func (u *productUsecase) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	return u.repo.GetAllProducts(ctx)
}

func (u *productUsecase) GetAllProductsForAdmin(ctx context.Context) ([]domain.Product, error) {
	role := ctx.Value("role").(string)

	if role != domain.RoleAdmin {
		return nil, domain.ErrUserNotAuthorized
	}
	return u.repo.GetAllProductsForAdmin(ctx)
}

func (u *productUsecase) CreateProduct(ctx context.Context, p domain.Product) (domain.Product, error) {
	role := ctx.Value("role").(string)
	if role != domain.RoleAdmin {
		return domain.Product{}, domain.ErrUserNotAuthorized
	}

	return u.repo.CreateProduct(ctx, domain.Product{Name: p.Name, Price: p.Price, Stock: p.Stock})

}

func (u *productUsecase) UpdateProductByID(ctx context.Context, p domain.Product) (domain.Product, error) {
	role := ctx.Value("role").(string)
	if role != domain.RoleAdmin {
		return domain.Product{}, domain.ErrUserNotAuthorized
	}

	p, err := u.repo.UpdateProduct(ctx, domain.Product{ID: p.ID, Name: p.Name, Price: p.Price, Stock: p.Stock})
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Product{}, domain.ErrProductDoesNotExist
	}
	return p, nil
}

func (u *productUsecase) DeleteProductByID(ctx context.Context, id int) error {
	role := ctx.Value("role").(string)
	if role != domain.RoleAdmin {
		return domain.ErrUserDoesNotExist
	}
	err := u.repo.DeleteProductByID(ctx, id)
	if err == sql.ErrNoRows {
		return domain.ErrProductDoesNotExist
	}
	return nil
}
