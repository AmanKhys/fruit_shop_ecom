package http

import (
	"net/http"
	"product_service/internal/usecase"
)

// GetFilteredProducts(ctx context.Context, min, max float64) ([]domain.Product, error)
// GetProductByID(ctx context.Context, id int) (domain.Product, error)
// GetAllProducts(ctx context.Context) ([]domain.Product, error)
// GetAllProductsForAdmin(ctx context.Context) ([]domain.Product, error)
// CreateProduct(ctx context.Context, p domain.Product) (domain.Product, error)
// UpdateProduct(ctx context.Context, p domain.Product) (domain.Product, error)

type ProductHandler struct {
	u usecase.ProductUsecase
}

func (h ProductHandler) GetFilteredProductsHandler(w http.ResponseWriter, r *http.Request) {
}

func (h ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
}

func (h ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
}

func (h ProductHandler) GetAllProductsForAdmin(w http.ResponseWriter, r *http.Request) {
}

func (h ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
}

func (h ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
}
