package bot

import (
	"encoding/json"
	"net/http"
	"log"
	"time"

	telegram "github.com/telegrapha/ui/api"
)

type Config struct {
	Debug               bool
	Token               string
	Production          bool
	HttpClientTimeout   time.Duration
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

type Bot struct {
	api    *telegram.Api
	config Config
}

func New(config *Config) (*Bot, error) {
	api := telegram.NewApi(config.Token, config.HttpClientTimeout, config.Production, config.Debug)
	api.Logger.Info(
		"Configuration:",
		"debug", config.Debug,
		"production", config.Production,
	)

	var user telegram.User
	requestParams := telegram.RequestParamsGetMe{}
	if err := api.GetMe(requestParams, &user); err != nil {
		return &Bot{}, err
	}
	api.User = user
	api.Logger.Info(
		"Authenticated:",
		"User", user,
	)

	return &Bot{
		api: api,
		config: *config,
	}, nil
}

func (bot *Bot) ListenPoolingUpdates(callback func(updates <-chan telegram.Update)) {
	bot.api.Logger.Info(
		"ListenPoolingUpdates:",
		"offset", bot.config.UpdateOffset,
		"limit", bot.config.UpdateLimit,
		"timeout", bot.config.UpdateTimeout,
	)

	updatesChan := make(chan telegram.Update, bot.config.UpdateBuffer)

	go func() {
		for {
			var updates []telegram.Update
			requestParams := telegram.RequestParamsGetUpdates{
				Offset:  bot.config.UpdateOffset,
				Limit:   bot.config.UpdateLimit,
				Timeout: bot.config.UpdateTimeout,
			}
			if err := bot.api.GetUpdates(requestParams, &updates); err != nil {
				bot.api.Logger.Info("GetUpdates", "err", err)

				time.Sleep(bot.config.HttpClientTimeout)
			}

			for _, update := range updates {
				if update.UpdateId >= bot.config.UpdateOffset {
					bot.config.UpdateOffset = update.UpdateId + 1
					updatesChan <- update

					bot.api.Logger.Info(
						"PoolingUpdate",
						"update", update,
						"offset", bot.config.UpdateOffset,
					)
				}
			}
		}
	}()

	callback(updatesChan)
}

func (bot *Bot) ListenHttpWebHookUpdates(callback func(updates <-chan telegram.Update)) {
	bot.api.Logger.Info(
		"ListenHttpWebHookUpdates:",
		"webhook", bot.config.WebHookListenOnPath,
		"port", bot.config.WebHookListenOnPort,
		"cert", nil,
		"key", nil,
		"tsl", false,
	)

	updatesChan := make(chan telegram.Update, bot.config.UpdateBuffer)

	http.HandleFunc(bot.config.WebHookListenOnPath, func(writer http.ResponseWriter, request *http.Request) {
		var update telegram.Update
		if err := json.NewDecoder(request.Body).Decode(&update); err != nil {
			log.Panic(err)
		}

		updatesChan <- update

		bot.api.Logger.Info(
			"WebHookUpdate",
			"update", update,
		)
	})
	go http.ListenAndServe(bot.config.WebHookListenOnPort, nil)

	callback(updatesChan)
}

func (bot *Bot) ListenHttpsWebHookUpdates(callback func(updates <-chan telegram.Update)) {
	bot.api.Logger.Info(
		"ListenHttpWebHookUpdates:",
		"webhook", bot.config.WebHookListenOnPath,
		"port", bot.config.WebHookListenOnPort,
		"cert", bot.config.WebHookTlsCertPem,
		"key", bot.config.WebHookTlsKeyPem,
		"tsl", true,
	)

	updatesChan := make(chan telegram.Update, bot.config.UpdateBuffer)

	http.HandleFunc(bot.config.WebHookListenOnPath, func(writer http.ResponseWriter, request *http.Request) {
		var update telegram.Update
		if err := json.NewDecoder(request.Body).Decode(&update); err != nil {
			log.Panic(err)
		}

		updatesChan <- update

		bot.api.Logger.Info(
			"WebHookUpdate",
			"update", update,
		)
	})
	go http.ListenAndServeTLS(
		bot.config.WebHookListenOnPort,
		bot.config.WebHookTlsCertPem,
		bot.config.WebHookTlsKeyPem,
		nil,
	)

	callback(updatesChan)
}
