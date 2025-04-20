package usershandler

import (
	"encoding/json"
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
	usersService users.Service
	jwtService   jsonwebtoken.Service
	logger       *slog.Logger
}

func NewHandler(usersService users.Service, jwtService jsonwebtoken.Service, logger *slog.Logger) *Handler {
	return &Handler{
		usersService: usersService,
		jwtService:   jwtService,
		logger:       logger,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/users", h.createUserHandler)
	mux.HandleFunc("GET /api/v1/users/{id}", middleware.AuthMiddleware(h.getUserHandler, h.jwtService, h.usersService))
	mux.HandleFunc("GET /api/v1/users", middleware.AuthMiddleware(h.getUsersHandler, h.jwtService, h.usersService))
}

func (h *Handler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type input struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	var data input
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		h.logger.Error(err.Error())
		responses.ErrorResponse(w, r, ErrInvalidUserData, h.logger)
		return
	}

	if !validators.ValidateNotEmpty(data.FirstName) || !validators.ValidateNotEmpty(data.LastName) || !validators.ValidateEmail(data.Email) || !validators.ValidatePassword(data.Password) {
		responses.ErrorResponse(w, r, ErrInvalidUserData, h.logger)
		return
	}

	user := &users.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
	}

	err = h.usersService.Create(r.Context(), user, data.Password)
	if err != nil {
		responses.ErrorResponse(w, r, err, h.logger)
		return
	}

	resp := responses.Envelope{
		"user": user,
	}
	responses.JsonResponse(w, r, http.StatusCreated, resp, h.logger)
}

// getUsersHandler If email query param is present, get user by email
// getUsersHandler If email query param is not present, get all usershandler
func (h *Handler) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email != "" && validators.ValidateEmail(email) {
		user, err := h.usersService.GetByEmail(r.Context(), email)
		if err != nil {
			responses.ErrorResponse(w, r, err, h.logger)
			return
		}
		resp := responses.Envelope{
			"data": []users.User{*user},
		}
		responses.JsonResponse(w, r, http.StatusOK, resp, h.logger)
		return
	}

	users, err := h.usersService.GetUsers(r.Context())
	if err != nil {
		responses.ErrorResponse(w, r, err, h.logger)
		return
	}

	resp := responses.Envelope{
		"data": users,
	}
	responses.JsonResponse(w, r, http.StatusOK, resp, h.logger)
}

func (h *Handler) getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Error(err.Error())
		responses.ErrorResponse(w, r, ErrInvalidUserID, h.logger)
		return
	}

	user, err := h.usersService.GetByID(r.Context(), idInt)
	if err != nil {
		responses.ErrorResponse(w, r, err, h.logger)
		return
	}

	resp := responses.Envelope{
		"data": user,
	}
	responses.JsonResponse(w, r, http.StatusOK, resp, h.logger)
}
