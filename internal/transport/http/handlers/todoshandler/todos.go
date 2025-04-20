package todoshandler

import (
	"github.com/lucianboboc/todo-api/internal/domain/todos"
	"github.com/lucianboboc/todo-api/internal/domain/users"
	"github.com/lucianboboc/todo-api/internal/intrastructure/jsonwebtoken"
	"github.com/lucianboboc/todo-api/internal/transport/http/middleware"
	"log/slog"
	"net/http"
)

type Handler struct {
	todoService  todos.Service
	usersService users.Service
	jwtService   jsonwebtoken.Service
	logger       *slog.Logger
}

func NewHandler(todoService todos.Service, usersService users.Service, jwtService jsonwebtoken.Service, logger *slog.Logger) *Handler {
	return &Handler{
		todoService:  todoService,
		usersService: usersService,
		jwtService:   jwtService,
		logger:       logger,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/todos", middleware.AuthMiddleware(h.createTodoHandler, h.jwtService, h.usersService))
	mux.HandleFunc("GET /api/v1/todos", middleware.AuthMiddleware(h.getTodosHandler, h.jwtService, h.usersService))
	mux.HandleFunc("GET /api/v1/todos/{id}", middleware.AuthMiddleware(h.getTodoHandler, h.jwtService, h.usersService))
}

// TODO: Implement handlers with data validation
func (h *Handler) createTodoHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) getTodosHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) getTodoHandler(w http.ResponseWriter, r *http.Request) {
}
