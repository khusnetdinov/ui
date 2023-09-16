package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const telegramApiUrl = "https://api.telegram.org/bot%s/"

type Api struct {
	client *http.Client
	url    string
	User   User
}

type ApiResponse struct {
	Ok          bool                   `json:"ok"`
	Result      json.RawMessage        `json:"result,omitempty"`
	ErrorCode   int                    `json:"error_code,omitempty"`
	Description string                 `json:"description,omitempty"`
	Parameters  *ApiResponseParameters `json:"parameters,omitempty"`
}

type ApiResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}

type ApiError struct {
	Code    int
	Message string
	ApiResponseParameters
}

func (e ApiError) Error() string {
	return e.Message
}

func NewApi(token string, timeout time.Duration) (*Api, error) {
	api := &Api{
		client: &http.Client{Timeout: timeout},
		url:    fmt.Sprintf(telegramApiUrl, token),
	}

	response, err := api.GetMe()
	if err != nil {
		return &Api{}, err
	}

	var user User
	if err := json.Unmarshal(response.Result, &user); err != nil {
		return &Api{}, err
	}
	api.User = user

	return api, nil
}
