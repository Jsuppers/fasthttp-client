package main

import (
	"fasthttp-client/client"
	"fmt"
	"log"
	"time"
)

const (
	server  = "http://localhost:8080"
	billion = 1000 * 1000 * 1000
	retry   = 1 * time.Second
)

func main() {
	fmt.Println("Sending one billion messages to ", server)
	for i := 1; i <= billion; {
		err := client.Send(server, i)
		if err != nil {
			log.Println("Error when sending request: ", err)
			log.Println("Retrying in: ", retry)
			time.Sleep(retry)
			continue
		}
		i++
	}
}
