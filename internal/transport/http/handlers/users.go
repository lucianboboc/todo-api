package handlers

import (
	"github.com/lucianboboc/todo-api/internal/domain/users"
	"github.com/lucianboboc/todo-api/internal/intrastructure/jsonwebtoken"
	"github.com/lucianboboc/todo-api/internal/transport/http/middleware"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	usersService users.Service
	jwtService   jsonwebtoken.Service
	logger       *slog.Logger
}

func NewUserHandler(usersService users.Service, jwtService jsonwebtoken.Service, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		usersService: usersService,
		jwtService:   jwtService,
		logger:       logger,
	}
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/users", h.createUserHandler)
	mux.HandleFunc("GET /api/v1/users", middleware.AuthMiddleware(h.getUsersHandler, h.jwtService, h.usersService))
	mux.HandleFunc("GET /api/v1/users/{id}", middleware.AuthMiddleware(h.getUserHandler, h.jwtService, h.usersService))
}

// TODO: Implement handlers with data validation
func (h *UserHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
}

// getUsersHandler If email query param is present, get user by email
// getUsersHandler If email query param is not present, get all users
func (h *UserHandler) getUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) getUserHandler(w http.ResponseWriter, r *http.Request) {
}
