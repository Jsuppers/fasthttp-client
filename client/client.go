package client

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/valyala/fasthttp"
)

const defaultRetryDuration = 1 * time.Second

type Client interface {
	Send(amount int)
}

type client struct {
	client        *fasthttp.Client
	address       string
	maxClientID   int
	retryDuration time.Duration
}

type request struct {
	Text      string `json:"text"`
	ContentID int    `json:"content_id"`
	ClientID  int    `json:"client_id"`
	Timestamp int64  `json:"timestamp"`
}

func New(address string, maxClientID int) Client {
	c := &client{}
	c.client = &fasthttp.Client{}
	c.address = address
	c.maxClientID = maxClientID
	c.retryDuration = defaultRetryDuration
	return c
}

func (c *client) Send(amount int) {
	log.Printf("Sending %d messages to %s", amount, c.address)
	for i := 1; i <= amount; {
		err := c.send(i)
		if err != nil {
			log.Println("Error when sending request: ", err)
			log.Println("Retrying in: ", c.retryDuration)
			time.Sleep(c.retryDuration)
			continue
		}
		i++
	}
}

func (c *client) send(contentID int) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	body, err := json.Marshal(c.makeRequest(contentID))
	if err != nil {
		return err
	}

	req.SetBody(body)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetRequestURI(c.address)

	if err := c.client.Do(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *client) makeRequest(contentID int) request {
	var req request
	req.Text = "hello world"
	req.ContentID = contentID
	req.ClientID = c.getClientID()
	req.Timestamp = getMillisecondTimestamp()
	return req
}

// returns a random number between 1 and maxClientID
func (c *client) getClientID() int {
	return rand.Intn(c.maxClientID) + 1
}

// returns current time in a millisecond precision timestamp
func getMillisecondTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
