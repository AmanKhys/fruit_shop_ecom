package http

import (
	"net/http"
	"user_service/handlers"
)

func RegisterRoutes(h *handlers.UserHandler) {
	http.HandleFunc("POST /register", h.Register)
	http.HandleFunc("POST /login", h.Login)
}
