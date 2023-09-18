package main

import (
	"fmt"
	"log"
	"time"

	"github.com/telegrapha/ui/api"
)

func main() {
	config := api.Config{
		Token:             "6591790550:AAE5s6Mmhs8QsGPDxmTxEB23kvKKg3KrI_w",
		HttpClientTimeout: time.Duration(1) * time.Second,
		UpdateOffset:      100,
		UpdateLimit:       0,
		UpdateTimeout:     0,
	}

	ui, err := api.NewApi(config)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(JsonPrint(ui.User))
}
