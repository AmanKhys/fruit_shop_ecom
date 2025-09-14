package http

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"user_service/internal/domain"
	"user_service/internal/usecase"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var reqUser domain.User

	err := json.NewDecoder(r.Body).Decode(&reqUser)
	if err != nil {
		http.Error(w, "request body not in format", http.StatusBadRequest)
		return
	}
	user, err := h.usecase.Register(r.Context(), reqUser.Email, reqUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	msg := fmt.Sprintf("user: %s registered successfully with userID: %d", user.Email, user.ID)
	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var reqUser domain.User

	err := json.NewDecoder(r.Body).Decode(&reqUser)
	if err != nil {
		http.Error(w, "request body not in format", http.StatusBadRequest)
		return
	}

	user, err := h.usecase.Login(r.Context(), reqUser.Email, reqUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}
