package app

import (
	"log"

	"github.com/blinkinglight/thingsv2/ms/first"
	"github.com/nats-io/nats.go"
)

var (
	nc *nats.Conn
)

func Run(n *nats.Conn) {
	nc = n
	output, err := Call[first.Output]("first", first.Input{
		Value: "hello",
	})
	log.Printf("err(%v) data: %+v", err, output)
}
