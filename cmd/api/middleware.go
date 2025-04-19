package main

import (
	"context"
	"github.com/lucianboboc/todo-api/internal/domain/users"
	"github.com/lucianboboc/todo-api/internal/intrastructure/jsonwebtoken"
	"net/http"
	"strings"
)

type userKey string

const userCtx userKey = "user"

func AuthMiddleware(
	next http.HandlerFunc,
	jwtService jsonwebtoken.Service,
	usersService users.Service,
) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authToken := parts[1]
		userID, err := jwtService.ValidateToken(authToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, err := usersService.GetByID(r.Context(), userID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
