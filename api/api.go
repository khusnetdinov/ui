package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const telegramApiUrl = "https://api.telegram.org/bot%s/"

type Config struct {
	Token string
	HttpClientTimeout time.Duration
	UpdateOffset int
	UpdateLimit int
	UpdateTimeout int
}

type Api struct {
	client *http.Client
	config Config
	url    string
	User   User
}

type RequestParams []byte

type Response struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	ErrorCode   int             `json:"error_code,omitempty"`
	Description string          `json:"description,omitempty"`
	Parameters  *ResponseParams `json:"parameters,omitempty"`
}

type ResponseParams struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}

type Error struct {
	Code    int
	Message string
	ResponseParams
}

func (e Error) Error() string {
	return e.Message
}

func New(config Config) (*Api, error) {
	api := &Api{
		config: config,
		client: &http.Client{Timeout: config.HttpClientTimeout},
		url:    fmt.Sprintf(telegramApiUrl, config.Token),
	}

	// User: begin
	response, err := api.GetMe()
	if err != nil {
		return &Api{}, err
	}

	var user User
	if err := json.Unmarshal(response.Result, &user); err != nil {
		return &Api{}, err
	}
	api.User = user
	// User: end

	return api, nil
}

func (api Api) ListenLongPoolingUpdates() (string) {
	return "LongPoolUpdate"
}


func (api Api) ListenWebHookUpdates() (string) {
	return "WebHookUpdate"
}
