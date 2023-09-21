package main

import (
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

	var resultSet bool
	requestSetParams := api.SetWebHookParams{}
	if err := ui.SetWebHook(requestSetParams, &resultSet); err != nil {
		log.Panic(err)
	}
	log.Println("SetWebHook: %s", resultSet)

	var resultDelete bool
	requestDeleteParams := api.DeleteWebHookParams{
		DropPendingUpdates: false,
	}
	if err := ui.DeleteWebHook(requestDeleteParams, &resultDelete); err != nil {
		log.Panic(err)
	}
	log.Println("DeleteWebHook: %s", resultDelete)

	// ui.ListenPoolingUpdates(func(updates <-chan api.Update) {
	// 	for update := range updates {
	// 		if update.Message != nil {
	// 			var message api.Message
	// 			requestParams := api.SendMessageParams{
	// 				ChatId:           update.Message.Chat.Id,
	// 				ReplyToMessageId: update.Message.MessageId,
	// 				Text:             update.Message.Text,
	// 			}
	// 			ui.SendMessage(requestParams, &message)

	// 			log.Println(JsonPrint(message))
	// 		}
	// 	}
	// })
}
