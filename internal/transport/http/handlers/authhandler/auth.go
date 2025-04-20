package authhandler

import (
	"encoding/json"
	"github.com/lucianboboc/todo-api/internal/domain/auth"
	"github.com/lucianboboc/todo-api/internal/transport/http/responses"
	"github.com/lucianboboc/todo-api/internal/transport/http/validators"
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

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/auth/login", h.LoginHandler)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	type input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var data input
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		h.logger.Error(err.Error())
		responses.ErrorResponse(w, r, ErrInvalidCredentials, h.logger)
		return
	}

	if !validators.ValidateEmail(data.Email) || !validators.ValidatePassword(data.Password) {
		responses.ErrorResponse(w, r, ErrInvalidCredentials, h.logger)
	}

	token, err := h.authService.Login(r.Context(), data.Email, data.Password)
	if err != nil {
		responses.ErrorResponse(w, r, err, h.logger)
		return
	}

	resp := responses.Envelope{
		"token": token,
	}
	responses.JsonResponse(w, r, http.StatusOK, resp, h.logger)
}
