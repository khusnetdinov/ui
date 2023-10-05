package bot

import (
	"encoding/json"
	"net/http"
	"log"
	"log/slog"
	"os"
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
	logger *slog.Logger
}

func New(config *Config) (*Bot, error) {
	logger := NewSlog(os.Stdout, config.Debug)

	api := telegram.NewApi(config.Token, config.HttpClientTimeout, config.Production, logger)
	logger.Info(
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
	logger.Info(
		"Authenticated:",
		"User", user,
	)

	return &Bot{
		api: api,
		config: *config,
		logger: logger,
	}, nil
}

func (bot *Bot) ListenPoolingUpdates(callback func(updates <-chan telegram.Update)) {
	bot.logger.Info(
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
				bot.logger.Info("GetUpdates", "err", err)

				time.Sleep(bot.config.HttpClientTimeout)
			}

			for _, update := range updates {
				if update.UpdateId >= bot.config.UpdateOffset {
					bot.config.UpdateOffset = update.UpdateId + 1
					updatesChan <- update

					bot.logger.Info(
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
	bot.logger.Info(
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

		bot.logger.Info(
			"WebHookUpdate",
			"update", update,
		)
	})
	go http.ListenAndServe(bot.config.WebHookListenOnPort, nil)

	callback(updatesChan)
}

func (bot *Bot) ListenHttpsWebHookUpdates(callback func(updates <-chan telegram.Update)) {
	bot.logger.Info(
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

		bot.logger.Info(
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
