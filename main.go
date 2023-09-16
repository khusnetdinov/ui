package main

import (
	"fmt"
	"log"
	"time"
)

const token = "6591790550:AAE5s6Mmhs8QsGPDxmTxEB23kvKKg3KrI_w"

func main() {
	// ApiConfiguration{
	// 	token: token,
	// 	httpClientTimeout: time.Duration(1)*time.Second,
	// }

	api, err := NewApi(token, time.Duration(1)*time.Second)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(JsonPrint(api.User))
}
