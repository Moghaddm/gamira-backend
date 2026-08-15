package common

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

type Handler func(http.ResponseWriter, *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		var httpError *HTTPError
		if errors.As(err, &httpError) {
			_ = WriteJSON(w, httpError.Code, map[string]string{"error": httpError.Message})
			return
		}
		slog.Error(err.Error())
		_ = WriteJSON(w, httpError.Code, map[string]string{"error": httpError.Message})
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
