package authhandler

import (
	"github.com/lucianboboc/todo-api/internal/common/apierrors"
	"net/http"
)

const errorDomain = "auth"

var ErrInvalidCredentials = apierrors.APIError{
	Domain:     errorDomain,
	StatusCode: http.StatusBadRequest,
	Message:    "Invalid credentials",
	Key:        "InvalidCredentials",
}
