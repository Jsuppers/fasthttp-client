package main

import (
	"fasthttp-client/client"
	"fmt"
	"os"
)

const (
	billion         = 1000 * 1000 * 1000
	maxClientID     = 10
	defaultServerIP = "http://172.17.0.1:8080"
)

func main() {
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		fmt.Println("SERVER_ADDRESS not set, using the default server: ", defaultServerIP)
		address = defaultServerIP
	}

	sender := client.New(address, maxClientID)
	sender.Send(billion)
}
