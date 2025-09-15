package usecase

import (
	"context"
	"product_service/internal/domain"
)

type ProductRepo interface {
	GetFilteredProducts(ctx context.Context, min, max float64) ([]domain.Product, error)
	GetProductByID(ctx context.Context, id int) (domain.Product, error)
	GetProductByIDForAdmin(ctx context.Context, id int) (domain.Product, error)
	GetAllProductsForAdmin(ctx context.Context, min, max float64) ([]domain.Product, error)
	CreateProduct(ctx context.Context, p domain.Product) (domain.Product, error)
	UpdateProduct(ctx context.Context, p domain.Product) (domain.Product, error)
	DeleteProductByID(ctx context.Context, id int) error
}
