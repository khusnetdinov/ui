package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/telegrapha/ui/api"
)

func main() {
	config := api.Config{
		Token:             "6591790550:AAE5s6Mmhs8QsGPDxmTxEB23kvKKg3KrI_w",
		HttpClientTimeout: time.Duration(5) * time.Second,
		UpdateOffset:      0,
		UpdateLimit:       1,
		UpdateTimeout:     1,
	}
	ui, err := api.New(&config)
	if err != nil {
		log.Panic(err)
	}

	requestParams := api.GetWebhookInfoParams{}
	response, err := ui.GetWebhookInfo(requestParams)
	if err != nil {
		log.Panic(err)
	}

	var webhookInfo api.WebhookInfo
	if err := json.Unmarshal(response.Result, &webhookInfo); err != nil {
		log.Println(err)
	}

	log.Println(JsonPrint(webhookInfo))

	// ui.ListenPoolingUpdates(func(updates <-chan api.Update) {
	// 	for update := range updates {
	// 		if update.Message != nil {
	// 			message := api.SendMessageParams{
	// 				ChatId:           update.Message.Chat.Id,
	// 				ReplyToMessageId: update.Message.MessageId,
	// 				Text:             update.Message.Text,
	// 			}

	// 			ui.SendMessage(message)
	// 		}
	// 	}
	// })
}
