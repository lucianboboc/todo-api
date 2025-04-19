package todos

import (
	"github.com/lucianboboc/todo-api/internal/domain/apierrors"
	"net/http"
)

const errorDomain = "todos"

var ErrTodoNotFound = apierrors.APIError{
	Domain:     errorDomain,
	StatusCode: http.StatusNotFound,
	Message:    "Todo not found",
	Key:        "TodoNotFound",
}
