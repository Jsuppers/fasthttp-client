package client

import (
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestNew(t *testing.T) {
	mockClient := &client{
		client:          &fasthttp.Client{},
		address:         "address",
		maxClientID:     1,
		retryDuration:   1 * time.Second,
		measureMessages: 1,
	}

	type args struct {
		address         string
		maxClientID     int
		retryDuration   time.Duration
		measureMessages int
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		{"creates a new client", args{address: "address", maxClientID: 1, retryDuration: 1 * time.Second, measureMessages: 1}, mockClient},
	}
	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			got := New(test.args.address, test.args.maxClientID, test.args.retryDuration, test.args.measureMessages)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("New() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSendMessages(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(listener *fasthttputil.InmemoryListener, retried *bool) Client
		wantRetried bool
	}{
		{"sends message", func(listener *fasthttputil.InmemoryListener, retried *bool) Client {
			memClient := &fasthttp.Client{
				Dial: func(addr string) (net.Conn, error) {
					return listener.Dial()
				},
			}
			return &client{
				client:          memClient,
				maxClientID:     1,
				address:         "http://make.fasthttp.great?again",
				measureMessages: 1,
			}
		}, false},
		{"retries sending if error with the listener", func(listener *fasthttputil.InmemoryListener, retried *bool) Client {
			memClient := &fasthttp.Client{
				Dial: func(addr string) (net.Conn, error) {
					if *retried == false {
						*retried = true
						return nil, fmt.Errorf("cannot connect")
					}
					return listener.Dial()
				},
			}
			return &client{
				client:          memClient,
				maxClientID:     1,
				address:         "http://make.fasthttp.great?again",
				measureMessages: 1,
			}
		}, true},
		{"retries sending if error with json marshaller", func(listener *fasthttputil.InmemoryListener, retried *bool) Client {
			memClient := &fasthttp.Client{
				Dial: func(addr string) (net.Conn, error) {
					return listener.Dial()
				},
			}
			jsonMarshal = func(v interface{}) (bytes []byte, err error) {
				if *retried == false {
					*retried = true
					return nil, fmt.Errorf("error")
				}
				return json.Marshal(v)
			}
			return &client{
				client:          memClient,
				maxClientID:     1,
				address:         "http://make.fasthttp.great?again",
				measureMessages: 1,
			}
		}, true},
	}
	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				jsonMarshal = json.Marshal
			}()

			received := false
			ln := fasthttputil.NewInmemoryListener()
			s := &fasthttp.Server{
				Handler: func(ctx *fasthttp.RequestCtx) {
					received = true
				},
			}

			go s.Serve(ln) //nolint:errcheck

			retried := false
			c := test.setup(ln, &retried)

			c.SendMessages(1)

			if !received {
				t.Errorf("server did not receive request")
			}

			if retried != test.wantRetried {
				t.Errorf("wanted %v, but got %v", test.wantRetried, retried)
			}
		})
	}
}
