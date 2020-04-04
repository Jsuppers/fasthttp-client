package client

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	maxClientID = 10
	client      = &fasthttp.Client{}
)

type Request struct {
	Text      string `json:"text"`
	ContentID int    `json:"content_id"`
	ClientID  int    `json:"client_id"`
	Timestamp int64  `json:"timestamp"`
}

func Send(address string, contentID int) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetBody(makeRequest(contentID))
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetRequestURI(address)

	if err := client.Do(req, nil); err != nil {
		panic(err)
	}
}

func makeRequest(contentID int) []byte {
	var request Request
	request.Text = "hello world"
	request.ContentID = contentID
	request.ClientID = getClientID()
	request.Timestamp = getMillisecondTimestamp()

	body, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("error when parsing request %v", err)
	}
	return body
}

// returns a random number between 1 and maxClientID
func getClientID() int {
	return rand.Intn(maxClientID) + 1
}

// returns current time in a millisecond precision timestamp
func getMillisecondTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
