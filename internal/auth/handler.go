package auth

import (
	"fmt"
	"gamira/common"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	redis *redis.Client
}

func NewHandler(redis *redis.Client) *Handler {
	return &Handler{redis: redis}
}

func (h *Handler) Init(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	session := r.Header.Get("session_id")
	if session == "" {
		return &common.HTTPError{Code: http.StatusBadRequest, Message: "Session Id is empty"}
	}

	key := fmt.Sprintf("session:%s", session)
	if h.redis.Exists(ctx, key).Val() == 0 {
		return &common.HTTPError{Code: http.StatusForbidden, Message: "Session does not exist"}
	}
	// todo
	return nil
}
