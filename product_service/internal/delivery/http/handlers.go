package http

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"product_service/internal/domain"
	"product_service/internal/dto"
	"product_service/internal/usecase"
	"strconv"
)

type ProductHandler struct {
	u usecase.ProductUsecase
}

func NewProductHandler(u usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{u: u}
}

func (h ProductHandler) GetFilteredProductsHandler(w http.ResponseWriter, r *http.Request) {
	minStr := r.URL.Query().Get("min")
	maxStr := r.URL.Query().Get("max")
	var respMsg []string
	min, err := strconv.ParseFloat(minStr, 64)
	if err != nil && minStr != "" {
		respMsg = append(respMsg, "min value is not valid")
	}
	max, err := strconv.ParseFloat(maxStr, 64)
	if err != nil && maxStr != "" {
		respMsg = append(respMsg, "max value is not valid")
	}

	products, err := h.u.GetFilteredProducts(r.Context(), min, max)
	if err != nil {
		http.Error(w, "fetching products failed", http.StatusInternalServerError)
		return
	}

	var resp struct {
		Products []dto.ProductUserResponse `json:"products"`
		Messages []string                  `json:"messages"`
	}
	var respProducts []dto.ProductUserResponse
	for _, p := range products {
		respProducts = append(respProducts, dto.ProductUserResponse{
			ID:    p.ID,
			Name:  p.Name,
			Price: p.Price,
			Stock: p.Stock,
		})
	}
	resp.Products = respProducts
	resp.Messages = respMsg
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Fatal(err)
	}
}

func (h ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(domain.ID)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, domain.ErrProductDoesNotExistResponse, http.StatusBadRequest)
		return
	}
	p, err := h.u.GetProductByID(r.Context(), id)
	if err == domain.ErrProductDoesNotExist {
		http.Error(w, domain.ErrProductDoesNotExistResponse, http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func (h ProductHandler) GetAllProductsForAdmin(w http.ResponseWriter, r *http.Request) {
	products, err := h.u.GetAllProductsForAdmin(r.Context())
	if err == domain.ErrUserNotAuthorized {
		http.Error(w, domain.ErrUserNotAuthorized.Error(), http.StatusForbidden)
		return
	}
	if err != nil {
		http.Error(w, domain.ErrProductsFetchingFailed, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Fatal(err)
	}
}

func (h ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var reqProduct domain.Product
	err := json.NewDecoder(r.Body).Decode(&reqProduct)
	if err != nil {
		http.Error(w, domain.ErrPoorlyFormedRequest, http.StatusBadRequest)
		return
	}
	p, err := h.u.CreateProduct(r.Context(), reqProduct)
	if err == domain.ErrUserNotAuthorized {
		http.Error(w, domain.ErrUserNotAuthorized.Error(), http.StatusForbidden)
		return
	}
	var resp struct {
		Msg     string         `json:"message"`
		Product domain.Product `json:"product"`
	}
	resp.Msg = "product successfully created"
	resp.Product = p
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Fatal(err)
	}
}

func (h ProductHandler) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	var reqProduct domain.Product
	err := json.NewDecoder(r.Body).Decode(&reqProduct)
	if err != nil {
		http.Error(w, domain.ErrPoorlyFormedRequest, http.StatusBadRequest)
		return
	}
	p, err := h.u.UpdateProductByID(r.Context(), reqProduct)
	if err == domain.ErrUserNotAuthorized {
		http.Error(w, domain.ErrUserNotAuthorized.Error(), http.StatusForbidden)
		return
	}
	if err == domain.ErrProductDoesNotExist {
		http.Error(w, domain.ErrProductDoesNotExistResponse, http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, domain.ErrProductFetchingFailed, http.StatusInternalServerError)
		return
	}

	var resp struct {
		Msg     string         `json:"message"`
		Product domain.Product `json:"product"`
	}
	resp.Msg = "product successfully created"
	resp.Product = p
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Fatal(err)
	}
}

func (h ProductHandler) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(domain.ID)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, domain.ErrProductDoesNotExistResponse, http.StatusBadRequest)
		return
	}
	err = h.u.DeleteProductByID(r.Context(), id)
	if err == domain.ErrUserNotAuthorized {
		http.Error(w, domain.ErrUserNotAuthorized.Error(), http.StatusForbidden)
		return
	}
	if err == domain.ErrProductDoesNotExist {
		http.Error(w, domain.ErrProductDoesNotExistResponse, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	msg := fmt.Sprintf("successfully deleted product with id:%d", id)
	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}

}
