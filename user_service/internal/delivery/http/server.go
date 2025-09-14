package http

import (
	"net/http"
)

func RegisterRoutes(h *UserHandler) {
	http.HandleFunc("POST /register", h.Register)
	http.HandleFunc("POST /login", h.Login)
}
