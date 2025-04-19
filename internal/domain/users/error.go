package users

import (
	"github.com/lucianboboc/todo-api/internal/domain/apierrors"
	"net/http"
)

const errorDomain = "users"

var ErrUserAlreadyExists = apierrors.APIError{
	Domain:     errorDomain,
	StatusCode: http.StatusConflict,
	Message:    "User already exists",
	Key:        "UserAlreadyExists",
}

var ErrUserNotFound = apierrors.APIError{
	Domain:     errorDomain,
	StatusCode: http.StatusNotFound,
	Message:    "User not found",
	Key:        "UserNotFound",
}
