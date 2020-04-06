package main

import (
	"fasthttp-client/client"
	"fmt"
	"os"
	"time"
)

const (
	messagesToSend       = 1000 * 1000 * 1000 // one billion
	measureDuration      = 10000              // ten thousand
	maxClientID          = 10
	defaultServerIP      = "http://localhost:8080"
	defaultRetryDuration = 1 * time.Second
)

func main() {
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		fmt.Println("SERVER_ADDRESS not set, using the default server: ", defaultServerIP)
		address = defaultServerIP
	}

	sender := client.New(address, maxClientID, defaultRetryDuration, measureDuration)
	sender.SendMessages(messagesToSend)
}
