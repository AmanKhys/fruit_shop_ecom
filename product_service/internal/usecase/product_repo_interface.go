package usecase

import (
	"context"
	"product_service/internal/domain"
)

type ProductRepo interface {
	GetFilteredProducts(ctx context.Context, min, max float64) ([]domain.Product, error)
	GetProductByID(ctx context.Context, id int) (domain.Product, error)
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
	CreateProduct(ctx context.Context, p domain.Product, userID int) (domain.Product, error)
	UpdateProduct(ctx context.Context, p domain.Product, userID int) (domain.Product, error)
}
