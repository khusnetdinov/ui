package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const telegramApiUrl = "https://api.telegram.org/bot%s/"

type Config struct {
	Token             string
	HttpClientTimeout time.Duration
	UpdateOffset      int64
	UpdateLimit       int64
	UpdateTimeout     int64
}

type Api struct {
	client *http.Client
	config Config
	url    string
	User   User
}

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

func (api Api) ListenLongPoolingUpdates(callback func(updates <-chan Update)) {
	updatesChan := make(chan Update, 100)

	go func() {
		for {
			requestParams := GetUpdatesParams{
				Offset:  api.config.UpdateOffset,
				Limit:   api.config.UpdateLimit,
				Timeout: api.config.UpdateTimeout,
			}

			// []Update: begin
			response, err := api.GetUpdates(requestParams)
			if err != nil {
				log.Println(err)
			}

			var updates []Update
			if err := json.Unmarshal(response.Result, &updates); err != nil {
				log.Println(err)
			}
			// []Update: end

			for _, update := range updates {
				if update.UpdateId >= api.config.UpdateOffset {
					api.config.UpdateOffset = update.UpdateId + 1
					updatesChan <- update
				}

			}
		}
	}()

	callback(updatesChan)
}

func (api Api) ListenWebHookUpdates(callback func()) {
	callback()
}
