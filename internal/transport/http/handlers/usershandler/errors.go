package usershandler

import (
	"github.com/lucianboboc/todo-api/internal/common/apierrors"
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
