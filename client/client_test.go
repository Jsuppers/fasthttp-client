package client

import (
	"net"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestSend(t *testing.T) {
	received := false
	ln := fasthttputil.NewInmemoryListener()
	s := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			received = true
		},
	}
	go s.Serve(ln) //nolint:errcheck
	client = &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}
	defer func() {
		client = &fasthttp.Client{}
	}()

	Send("http://make.fasthttp.great?again", 1000)

	// give some time for the server to receive the request
	time.Sleep(100 * time.Millisecond)

	if !received {
		t.Error("request was not sent to server")
	}
}
