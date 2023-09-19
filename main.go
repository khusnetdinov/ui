package main

import (
	"log"
	"time"

	"github.com/telegrapha/ui/api"
)

func main() {
	config := api.Config{
		Token:             "6591790550:AAE5s6Mmhs8QsGPDxmTxEB23kvKKg3KrI_w",
		HttpClientTimeout: time.Duration(1) * time.Second,
		UpdateOffset:      0,
		UpdateLimit:       100,
		UpdateTimeout:     0,
	}
	ui, err := api.New(config)
	if err != nil {
		log.Panic(err)
	}

	ui.ListenLongPoolingUpdates(func(updates <-chan api.Update) {
		for update := range updates {
			if update.Message != nil {
				log.Println(JsonPrint(update.Message))
			}
		}
	})
}
