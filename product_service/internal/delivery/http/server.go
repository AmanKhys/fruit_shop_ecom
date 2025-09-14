package http

import (
	"net/http"
	"product_service/internal/domain"
)

func RegisterRoutes(h *ProductHandler) {
	Auth := AuthMiddleware([]byte(domain.AuthSecret))

	http.HandleFunc("GET /products", h.GetFilteredProductsHandler)
	http.HandleFunc("GET /admin/products", h.GetAllProductsForAdmin)
	http.HandleFunc("GET /product", h.GetProductByID)
	http.Handle("POST /product/create", Auth(h.CreateProduct))
	http.Handle("POST /product/update", Auth(h.UpdateProductByID))
	http.Handle("DELETE /product", Auth(h.DeleteProductByID))
}
