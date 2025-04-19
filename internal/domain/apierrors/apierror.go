package apierrors

type APIError struct {
	Domain     string `json:"domain"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Key        string `json:"key"`
}

func (e APIError) Error() string {
	return e.Message
}

var ErrInternalServerError = APIError{
	Domain:     "error",
	StatusCode: 500,
	Message:    "Internal server error",
	Key:        "InternalServerError",
}
