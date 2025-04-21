package todoshandler

import (
	"github.com/lucianboboc/todo-api/internal/common/apierrors"
	"net/http"
)

const errorDomain = "todos"

var ErrTodoCannotBeCreated = apierrors.APIError{
	Domain:     errorDomain,
	StatusCode: http.StatusUnprocessableEntity,
	Message:    "Todo cannot be created",
	Key:        "TodoCannotBeCreated",
}

var ErrInvalidTodoData = apierrors.APIError{
	Domain:     errorDomain,
	StatusCode: http.StatusBadRequest,
	Message:    "Invalid todo data",
	Key:        "InvalidTodoData",
}

var ErrInvalidTodoID = apierrors.APIError{
	Domain:     errorDomain,
	StatusCode: http.StatusBadRequest,
	Message:    "Invalid todo ID",
	Key:        "InvalidTodoID",
}

var ErrTodoNotFound = apierrors.APIError{
	Domain:     errorDomain,
	StatusCode: http.StatusNotFound,
	Message:    "Todo not found",
	Key:        "TodoNotFound",
}
