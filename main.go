package main

import (
	"log"
	"time"

	"github.com/telegrapha/ui/api"
)

func main() {

	config := api.Config{
		Debug:               true,
		Production:          true,
		Token:               "6591790550:AAE5s6Mmhs8QsGPDxmTxEB23kvKKg3KrI_w",
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
	ui, err := api.New(&config)
	if err != nil {
		log.Fatal(err)
	}

	ui.ListenPoolingUpdates(func(updates <-chan api.Update) {
		for update := range updates {
			if update.Message != nil {
				var message api.Message
				requestParams := api.SendMessageParams{
					ChatId:           update.Message.Chat.Id,
					ReplyToMessageId: update.Message.MessageId,
					Text:             update.Message.Text,
				}
				ui.SendMessage(requestParams, &message)
			}
		}
	})
}
