package divar

import (
	"encoding/json"
	"fmt"
	"gamira/common"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	redis *redis.Client
}

func NewHandler(redis *redis.Client) *Handler {
	return &Handler{redis: redis}
}

func (h *Handler) StartFlow(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	// validate content-type
	c := r.Header.Get("Content-Type")
	if c != "application/json" {
		return &common.HTTPError{Code: http.StatusUnsupportedMediaType, Message: http.StatusText(http.StatusUnsupportedMediaType)}
	}

	// retrieve identification key
	a := r.Header.Get("Authorization")
	if a == "" || !strings.HasPrefix(a, "Bearer ") {
		return &common.HTTPError{Code: http.StatusUnauthorized, Message: "authorization header missing"}
	}

	// check identification key
	identification := os.Getenv("DIVAR_IDENTIFICATION_KEY")
	if identification == "" {
		return &common.HTTPError{Code: http.StatusInternalServerError, Message: "DIVAR_IDENTIFICATION_KEY required"}
	}
	if identification != a {
		return &common.HTTPError{Code: http.StatusForbidden, Message: "authorization header mismatch!"}
	}

	// read body
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error(err.Error())
		}
	}(r.Body)

	// decode body
	var createSession struct {
		PostToken     string `json:"post_token"`
		CompletionUrl string `json:"completion_url"`
	}
	if err := json.Unmarshal(b, &createSession); err != nil {
		return err
	}

	// generate session id
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src).Int31()

	// set session
	session := fmt.Sprintf("session:%s", strconv.Itoa(int(rnd)))
	err = h.redis.Set(ctx, session, createSession, time.Hour).Err()
	if err != nil {
		return err
	}

	// construct/build-up final response

	client := os.Getenv("CLIENT_URI")
	if client == "" {
		return &common.HTTPError{Code: http.StatusInternalServerError, Message: "CLIENT_URI required"}
	}

	uri, _ := url.Parse(client)
	uri.Path = "/start-flow"

	q := uri.Query()
	q.Add("session_id", session)
	uri.RawQuery = q.Encode()

	err = common.WriteJSON(w, http.StatusOK, &StartFlowResponse{URL: uri.String()})
	if err != nil {
		return err
	}

	return nil
}
