package main

import "fasthttp-client/client"

const (
	server  = "http://localhost:8080"
	billion = 1000 * 1000 * 1000
)

func main() {
	for i := 1; i <= billion; i++ {
		client.Send(server, i)
	}
}
