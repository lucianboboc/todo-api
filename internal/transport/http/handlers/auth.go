package handlers

import (
	"github.com/lucianboboc/todo-api/internal/domain/auth"
	"log/slog"
	"net/http"
)

type AuthHandler struct {
	authService auth.Service
	logger      *slog.Logger
}

func NewAuthHandler(authService auth.Service, logger *slog.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

// TODO: Implement handlers with data validation
func (h *AuthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/auth/login", h.LoginHandler)
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
}
