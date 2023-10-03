package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const telegramApiUrl = "https://api.telegram.org/bot%s/"
const localApiUrl = "http://localhost:8081/bot%s/"

type Config struct {
	Debug               bool
	Token               string
	Production          bool
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
	config Config
	client *http.Client
	logger *Slog
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

func New(config *Config) (*Api, error) {
	var url string
	if config.Production {
		url = fmt.Sprintf(telegramApiUrl, config.Token)
	} else {
		url = fmt.Sprintf(localApiUrl, config.Token)
	}

	api := &Api{
		config: *config,
		logger: NewSlog(os.Stdout, config.Debug),
		client: &http.Client{Timeout: config.HttpClientTimeout},
		url:    url,
	}
	api.logger.Info(
		"Configuration:",
		"debug", api.config.Debug,
		"production", api.config.Production,
	)

	var user User
	requestParams := RequestParamsGetMe{}
	if err := api.GetMe(requestParams, &user); err != nil {
		return &Api{}, err
	}
	api.User = user
	api.logger.Info(
		"Authenticated:",
		"User", user,
	)

	return api, nil
}

func (api *Api) ListenPoolingUpdates(callback func(updates <-chan Update)) {
	api.logger.Info(
		"ListenPoolingUpdates:",
		"offset", api.config.UpdateOffset,
		"limit", api.config.UpdateLimit,
		"timeout", api.config.UpdateTimeout,
	)

	updatesChan := make(chan Update, api.config.UpdateBuffer)

	go func() {
		for {
			var updates []Update
			requestParams := RequestParamsGetUpdates{
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

					api.logger.Info(
						"PoolingUpdate",
						"update", update,
						"offset", api.config.UpdateOffset,
					)
				}
			}
		}
	}()

	callback(updatesChan)
}

func (api *Api) ListenHttpWebHookUpdates(callback func(updates <-chan Update)) {
	api.logger.Info(
		"ListenHttpWebHookUpdates:",
		"webhook", api.config.WebHookListenOnPath,
		"port", api.config.WebHookListenOnPort,
		"cert", nil,
		"key", nil,
		"tsl", false,
	)

	updatesChan := make(chan Update, api.config.UpdateBuffer)

	http.HandleFunc(api.config.WebHookListenOnPath, func(writer http.ResponseWriter, request *http.Request) {
		var update Update
		if err := json.NewDecoder(request.Body).Decode(&update); err != nil {
			log.Panic(err)
		}

		updatesChan <- update

		api.logger.Info(
			"WebHookUpdate",
			"update", update,
		)
	})
	go http.ListenAndServe(api.config.WebHookListenOnPort, nil)

	callback(updatesChan)
}

func (api Api) ListenHttpsWebHookUpdates(callback func(updates <-chan Update)) {
	api.logger.Info(
		"ListenHttpWebHookUpdates:",
		"webhook", api.config.WebHookListenOnPath,
		"port", api.config.WebHookListenOnPort,
		"cert", api.config.WebHookTlsCertPem,
		"key", api.config.WebHookTlsKeyPem,
		"tsl", true,
	)

	updatesChan := make(chan Update, api.config.UpdateBuffer)

	http.HandleFunc(api.config.WebHookListenOnPath, func(writer http.ResponseWriter, request *http.Request) {
		var update Update
		if err := json.NewDecoder(request.Body).Decode(&update); err != nil {
			log.Panic(err)
		}

		updatesChan <- update

		api.logger.Info(
			"WebHookUpdate",
			"update", update,
		)
	})
	go http.ListenAndServeTLS(
		api.config.WebHookListenOnPort,
		api.config.WebHookTlsCertPem,
		api.config.WebHookTlsKeyPem,
		nil,
	)

	callback(updatesChan)
}
