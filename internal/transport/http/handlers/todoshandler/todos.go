package todoshandler

import (
	"encoding/json"
	"github.com/lucianboboc/todo-api/internal/domain/todos"
	"github.com/lucianboboc/todo-api/internal/domain/users"
	"github.com/lucianboboc/todo-api/internal/intrastructure/jsonwebtoken"
	"github.com/lucianboboc/todo-api/internal/transport/http/middleware"
	"github.com/lucianboboc/todo-api/internal/transport/http/responses"
	"github.com/lucianboboc/todo-api/internal/transport/http/validators"
	"log/slog"
	"net/http"
	"strconv"
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

func (h *Handler) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	type input struct {
		Text      string `json:"text"`
		UserID    int    `json:"user_id"`
		Completed bool   `json:"completed,omitempty"`
	}

	var data input
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		h.logger.Error(err.Error())
		responses.ErrorResponse(w, r, ErrInvalidTodoData, h.logger)
		return
	}

	if !validators.ValidateNotEmpty(data.Text) || !validators.ValidateUserId(data.UserID) {
		responses.ErrorResponse(w, r, ErrInvalidTodoData, h.logger)
		return
	}

	todo := &todos.Todo{
		Text:      data.Text,
		UserID:    data.UserID,
		Completed: data.Completed,
	}

	err = h.todoService.Create(r.Context(), todo)
	if err != nil {
		h.logger.Error(err.Error())
		responses.ErrorResponse(w, r, ErrTodoCannotBeCreated, h.logger)
		return
	}

	resp := responses.Envelope{
		"data": todo,
	}
	responses.JsonResponse(w, r, http.StatusCreated, resp, h.logger)
}

func (h *Handler) getTodosHandler(w http.ResponseWriter, r *http.Request) {
	data, err := h.todoService.GetAll(r.Context())
	if err != nil {
		responses.ErrorResponse(w, r, err, h.logger)
		return
	}

	resp := responses.Envelope{
		"data": data,
	}
	responses.JsonResponse(w, r, http.StatusOK, resp, h.logger)
}

func (h *Handler) getTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Error(err.Error())
		responses.ErrorResponse(w, r, ErrInvalidTodoID, h.logger)
		return
	}

	todo, err := h.todoService.GetById(r.Context(), idInt)
	if err != nil {
		responses.ErrorResponse(w, r, err, h.logger)
		return
	}

	resp := responses.Envelope{
		"data": todo,
	}
	responses.JsonResponse(w, r, http.StatusOK, resp, h.logger)
}
