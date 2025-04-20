package auth

import (
	"github.com/lucianboboc/todo-api/internal/domain/apierrors"
	"net/http"
)

const errorDomain = "auth"

var ErrUserUnauthorized = apierrors.APIError{
	Domain:     errorDomain,
	StatusCode: http.StatusUnauthorized,
	Message:    "User unauthorized",
	Key:        "UserUnauthorized",
}
