package main

import (
	"fmt"
	"log"
	"time"

	"github.com/telegrapha/ui/api"
)

const token = "6591790550:AAE5s6Mmhs8QsGPDxmTxEB23kvKKg3KrI_w"

func main() {
	ui, err := api.NewApi(token, time.Duration(1) * time.Second)
	if err != nil {
		log.Panic(err)
	}

	ui.ListenLongPoolingUpdates()
	ui.ListenWebHookUpdates()

	fmt.Println(JsonPrint(ui.User))
}
