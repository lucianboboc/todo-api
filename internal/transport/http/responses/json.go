package responses

import (
	"encoding/json"
	"errors"
	"github.com/lucianboboc/todo-api/internal/common/apierrors"
	"log/slog"
	"net/http"
)

type Envelope map[string]any

func JsonResponse(w http.ResponseWriter, r *http.Request, statusCode int, data any, logger *slog.Logger) {
	respJson, err := json.Marshal(data)
	if err != nil {
		logError(err, r, logger)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(respJson)
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, err error, logger *slog.Logger) {
	var apierror apierrors.APIError
	if errors.As(err, &apierror) {
		e := Envelope{
			"error": apierror,
		}
		JsonResponse(w, r, apierror.StatusCode, e, logger)
	} else {
		logError(err, r, logger)
		JsonResponse(w, r, apierrors.ErrInternalServerError.StatusCode, apierrors.ErrInternalServerError, logger)
	}
}

func logError(err error, r *http.Request, logger *slog.Logger) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
}
