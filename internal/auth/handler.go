package auth

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"gamira/common"
	"gamira/internal/divar"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	redis *redis.Client
	repo  Repository
}

func NewHandler(redis *redis.Client, repo Repository) *Handler {
	return &Handler{redis: redis, repo: repo}
}

func (h *Handler) Init(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	sessionId := r.Header.Get("session_id")
	if sessionId == "" {
		return &common.HTTPError{Code: http.StatusBadRequest, Message: "Session Id is empty"}
	}

	// validate session and its existence
	sessionKey := fmt.Sprintf("session:%s", sessionId)
	val, err := h.redis.Get(ctx, sessionKey).Result()
	if err != nil {
		return err
	}
	var session divar.Session
	if err := json.NewDecoder(strings.NewReader(val)).Decode(&session); err != nil {
		return err
	}

	src := rand.NewSource(time.Now().UnixNano())
	rnd := strconv.Itoa(rand.New(src).Int())

	// store state
	stateId := sha256.Sum256([]byte(rnd))
	stateKey := fmt.Sprintf("state:%s", stateId)
	var state = &State{PostToken: session.PostToken, CompletionUrl: session.CompletionUrl}
	err = h.redis.Set(ctx, stateKey, state, time.Hour).Err()
	if err != nil {
		return err
	}

	// construct auth uri

	authUri := os.Getenv("DIVAR_CONNECT_AUTH")
	redirect := os.Getenv("AUTH_REDIRECT_URL")
	clientId := os.Getenv("AUTH_CLIENT_ID")
	scp := os.Getenv("AUTH_SCOPE")

	uri, _ := url.Parse(authUri)
	q := uri.Query()
	q.Add("response_type", "code")
	q.Add("redirect_uri", redirect)
	q.Add("client_id", clientId)
	q.Add("scope", scp)
	q.Add("state", fmt.Sprintf("%s", state))
	uri.RawQuery = q.Encode()

	err = common.WriteJSON(w, http.StatusOK, &InitResponse{URL: uri.String()})
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	// check state existence/validation

	stateId := r.Header.Get("state")

	stateKey := fmt.Sprintf("state:%s", stateId)
	stateCache, err := h.redis.Get(ctx, stateKey).Result()
	if err != nil {
		return err
	}

	if stateCache == "" {
		return &common.HTTPError{Code: http.StatusForbidden, Message: "State not found"}
	}

	var state State
	if err := json.NewDecoder(strings.NewReader(stateCache)).Decode(&state); err != nil {
		return err
	}

	// send connect token request
	code := r.Header.Get("code")
	accessToken, expiresIn, err := connectToken(ctx, code)

	// retrieve user phone
	phoneNumber, err := getPhoneNumber(ctx, accessToken)
	if err != nil {
		return err
	}

	// create user

	res, err := h.repo.GetByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return err
	}
	if res == nil {
		id, err := h.repo.Create(ctx, phoneNumber)
		if err != nil {
			return err
		}
		res, err = h.repo.GetById(ctx, id)
	}

	// store user data
	userData := &UserData{
		PostToken:        state.PostToken,
		CompletionUrl:    state.CompletionUrl,
		DivarAccessToken: accessToken,
		PhoneNumber:      phoneNumber,
	}
	key := fmt.Sprintf("user:%s", res)
	err = h.redis.Set(ctx, key, userData, time.Duration(expiresIn)*time.Second).Err()
	if err != nil {
		return err
	}

	// make jwt and return
	claims := jwt.MapClaims{
		"sub": res.ID,
		"exp": time.Now().Add(time.Duration(expiresIn) * time.Second).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("DIVAR_ACCESS_TOKEN")))
	if err != nil {
		return err
	}

	err = common.WriteJSON(w, http.StatusOK, &CallbackResponse{AccessToken: tokenString})
	if err != nil {
		return err
	}

	return nil
}

func connectToken(ctx context.Context, code string) (string, int, error) {
	tokenUri := os.Getenv("DIVAR_CONNECT_TOKEN")
	redirect := os.Getenv("AUTH_REDIRECT_URL")
	clientId := os.Getenv("AUTH_CLIENT_ID")
	clientSecret := os.Getenv("AUTH_CLIENT_SECRET")

	form := url.Values{}
	form.Set("code", code)
	form.Set("redirect_uri", redirect)
	form.Set("client_id", clientId)
	form.Set("client_secret", clientSecret)
	form.Set("grant_type", "authorization_code")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenUri, strings.NewReader(form.Encode()))
	if err != nil {
		return "", 0, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", 0, err
	}

	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		return "", 0, err
	}

	accessToken := data["access_token"]
	expiresIn := data["expires_in"]

	return accessToken.(string), int(expiresIn.(float64)), nil
}

func getPhoneNumber(ctx context.Context, accessToken string) (string, error) {
	uri := os.Getenv("DIVAR_USER_INFO_URL")
	apiKey := os.Getenv("DIVAR_API_KEY")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	phoneNumber := data["phone_number"]
	if phoneNumber != nil {
		return phoneNumber.(string), nil
	}

	return "", &common.HTTPError{Code: http.StatusInternalServerError, Message: "retrieved invalid phone number"}
}
