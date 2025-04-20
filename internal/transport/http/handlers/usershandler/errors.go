package usershandler

import (
	"github.com/lucianboboc/todo-api/internal/domain/apierrors"
	"net/http"
)

const errorDomain = "users"

var ErrInvalidUserData = apierrors.APIError{
	Domain:     errorDomain,
	StatusCode: http.StatusBadRequest,
	Message:    "Invalid user data",
	Key:        "InvalidUserData",
}

var ErrInvalidUserID = apierrors.APIError{
	Domain:     errorDomain,
	StatusCode: http.StatusBadRequest,
	Message:    "Invalid user ID",
	Key:        "InvalidUserID",
}
