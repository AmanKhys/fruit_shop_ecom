package http

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"product_service/internal/domain"
)

func RegisterRoutes(h *ProductHandler) {
	authSecret := os.Getenv(domain.AuthSecret)
	if authSecret == "" {
		log.Fatal("jwt Secret not in env")
	}
	Auth := AuthMiddleware([]byte(authSecret))

	http.Handle("GET /products", Auth(h.GetProducts))
	http.HandleFunc("GET /product", h.GetProductByID)
	http.Handle("POST /product/create", Auth(h.CreateProduct))
	http.Handle("POST /product/update", Auth(h.UpdateProductByID))
	http.Handle("DELETE /product", Auth(h.DeleteProductByID))
}
