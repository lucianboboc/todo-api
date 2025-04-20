package authhandler

import (
	"github.com/lucianboboc/todo-api/internal/domain/auth"
	"log/slog"
	"net/http"
)

type Handler struct {
	authService auth.Service
	logger      *slog.Logger
}

func NewHandler(authService auth.Service, logger *slog.Logger) *Handler {
	return &Handler{
		authService: authService,
		logger:      logger,
	}
}

// TODO: Implement handlers with data validation
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/auth/login", h.LoginHandler)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
}
