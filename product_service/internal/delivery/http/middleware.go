package http

import (
	"context"
	"net/http"
	"product_service/internal/domain"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func AuthMiddleware(secret []byte) func(http.HandlerFunc) http.Handler {
	return func(next http.HandlerFunc) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "missing token", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			// attach user info to context
			ctx := context.WithValue(r.Context(), domain.UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, domain.RoleKey, domain.ContextKey(claims.Role))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
