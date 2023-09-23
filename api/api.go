package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const telegramApiUrl = "https://api.telegram.org/bot%s/"

const (
	UpdatesViaPooling      = "polling"
	UpdatesViaNgrokWebhook = "ngrok"
	UpdatesViaLocalServer  = "local"
	UpdatesViaTlsServer    = "tlc"
)

type Config struct {
	Token               string
	HttpClientTimeout   time.Duration
	UpdatesVia          string
	UpdateBuffer        int64
	UpdateOffset        int64
	UpdateLimit         int64
	UpdateTimeout       int64
	WebHookUrl          string
	WebHookListenOnPort string
	WebHookListenOnPath string
	WebHookTlsCertPem   string
	WebHookTlsKeyPem    string
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

func New(config *Config) (*Api, error) {
	api := &Api{
		config: *config,
		client: &http.Client{Timeout: config.HttpClientTimeout},
		url:    fmt.Sprintf(telegramApiUrl, config.Token),
	}

	var user User
	requestParams := GetMeParams{}
	if err := api.GetMe(requestParams, &user); err != nil {
		return &Api{}, err
	}
	api.User = user

	return api, nil
}

func (api *Api) ListenPoolingUpdates(callback func(updates <-chan Update)) {
	updatesChan := make(chan Update, api.config.UpdateBuffer)

	go func() {
		for {
			var updates []Update
			requestParams := GetUpdatesParams{
				Offset:  api.config.UpdateOffset,
				Limit:   api.config.UpdateLimit,
				Timeout: api.config.UpdateTimeout,
			}
			if err := api.GetUpdates(requestParams, &updates); err != nil {
				log.Panic(err)

				time.Sleep(api.config.HttpClientTimeout)
			}

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

func (api *Api) ListenHttpWebHookUpdates(callback func(updates <-chan Update)) {
	updatesChan := make(chan Update, api.config.UpdateBuffer)

	http.HandleFunc(api.config.WebHookListenOnPath, func(writer http.ResponseWriter, request *http.Request) {
		var update Update
		if err := json.NewDecoder(request.Body).Decode(&update); err != nil {
			log.Panic(err)
		}

		updatesChan <- update
	})
	go http.ListenAndServe(api.config.WebHookListenOnPort, nil)

	callback(updatesChan)
}

func (api Api) ListenHttpsWebHookUpdates(callback func(updates <-chan Update)) {
	updatesChan := make(chan Update, api.config.UpdateBuffer)

	http.HandleFunc(api.config.WebHookListenOnPath, func(writer http.ResponseWriter, request *http.Request) {
		var update Update
		if err := json.NewDecoder(request.Body).Decode(&update); err != nil {
			log.Panic(err)
		}

		updatesChan <- update
	})
	go http.ListenAndServeTLS(
		api.config.WebHookListenOnPort,
		api.config.WebHookTlsCertPem,
		api.config.WebHookTlsKeyPem,
		nil,
	)

	callback(updatesChan)
}
