package app

import (
	"encoding/json"
	"log"
	"time"

	"github.com/blinkinglight/thingsv2/shared"
)

func Call[T any](topic string, I any) (T, error) {
	b, err := json.Marshal(I)
	if err != nil {
		panic(err)
	}
	var request shared.SystemMessage
	request.Payload = string(b)
	b, err = request.Marshal()
	if err != nil {
		panic(err)
	}
	m, err := nc.Request(topic, b, 1*time.Second)
	if err != nil {
		panic(err)
	}
	var response shared.SystemMessage
	err = response.Unmarshal(m.Data)
	if err != nil {
		panic(err)
	}
	log.Printf("%s", response.Payload)
	var output T
	json.Unmarshal([]byte(response.Payload), &output)
	return output, err
}
