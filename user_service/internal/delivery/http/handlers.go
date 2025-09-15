package http

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
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

	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		http.Error(w, "request body not in format", http.StatusBadRequest)
		return
	}

	user, err := h.usecase.Login(r.Context(), reqUser.Email, reqUser.Password)
	if err == domain.ErrUserDoesNotExist {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour * 7).Unix(), // expire in 7 days
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv(domain.JwtSecret)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "failed to sign token", http.StatusInternalServerError)
		return
	}

	// Response
	resp := struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
		Role  string `json:"role"`
		Msg   string `json:"message"`
		Token string `json:"token"`
	}{
		ID:    int(user.ID),
		Email: user.Email,
		Role:  user.Role,
		Msg:   "successfully logged in",
		Token: signedToken,
	}

	w.Header().Set("Authorization", "Bearer "+signedToken)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("error encoding response:", err)
	}
}
