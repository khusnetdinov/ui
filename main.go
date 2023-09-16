package main

import (
	"fmt"
	"log"
	"time"

	"github.com/telegrapha/ui/api"
)

const token = "6591790550:AAE5s6Mmhs8QsGPDxmTxEB23kvKKg3KrI_w"

func main() {
	// ApiConfiguration{
	// 	token: token,
	// 	httpClientTimeout: time.Duration(1)*time.Second,
	// }

	bot, err := NewApi(token, time.Duration(1)*time.Second)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(JsonPrint(bot.User))
	api.HelloApi()
}
