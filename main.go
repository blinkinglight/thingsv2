package main

import (
	"os"
	"time"

	"github.com/blinkinglight/thingsv2/app"
	"github.com/blinkinglight/thingsv2/ms/first"
	"github.com/blinkinglight/thingsv2/shared"
	"github.com/nats-io/nats.go"
)

var (
	reg = map[string]shared.MSI{}
)

func main() {

	if os.Getenv("NATS_URL") == "" {
		panic("NATS_URL is not set")
	}

	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	reg["first"] = shared.MSI{
		Fn:  first.Call,
		In:  first.Input{},
		Out: first.Output{},
	}

	for topic, fn := range reg {
		nc.Subscribe(topic, func(m *nats.Msg) {
			var message shared.SystemMessage
			err := message.Unmarshal(m.Data)
			if err != nil {
				panic(err)
			}
			output, err := fn.Fn.Call([]byte(message.Payload))

			if err != nil {
				panic(err)
			}

			message.Payload = string(output)

			b, err := message.Marshal()
			if err != nil {
				panic(err)
			}
			nc.Publish(m.Reply, b)
		})
	}

	time.Sleep(1 * time.Second)

	app.Run(nc)

	select {}
}
