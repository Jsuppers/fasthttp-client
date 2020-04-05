package main

import (
	"fasthttp-client/client"
)

const (
	server      = "http://localhost:8080"
	billion     = 1000 * 1000 * 1000
	maxClientID = 10
)

func main() {
	sender := client.New(server, maxClientID)
	sender.Send(billion)
}
