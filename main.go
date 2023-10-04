package main

import (
	"log"
	"time"

	"github.com/telegrapha/ui/api"
	"github.com/telegrapha/ui/bot"
)

func main() {
	config := bot.Config{
		Debug:      true,
		Production: true,
		// Token:               "6591790550:AAE5s6Mmhs8QsGPDxmTxEB23kvKKg3KrI_w",
		Token:               "6082780877:AAFuH3vVGM2k5JZHQncGQnCkuQiDOB0ikMI",
		HttpClientTimeout:   time.Duration(5) * time.Second,
		UpdateBuffer:        100,
		UpdateOffset:        0,
		UpdateLimit:         1,
		UpdateTimeout:       1,
		WebHookUrl:          "",
		WebHookListenOnPath: "/",
		WebHookListenOnPort: ":8081",
		WebHookTlsCertPem:   "",
		WebHookTlsKeyPem:    "",
	}

	ui, err := bot.New(&config)
	if err != nil {
		log.Fatal(err)
	}

	ui.ListenPoolingUpdates(func(updates <-chan api.Update) {
		for update := range updates {
			if update.Message != nil {
				log.Printf("Message: %s ", update.Message.Text)
			}
		}
	})
}
