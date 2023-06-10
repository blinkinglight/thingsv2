package shared

import "encoding/json"

type SystemMessage struct {
	Payload string `json:"payload"`
}

func (s *SystemMessage) Marshal() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SystemMessage) Unmarshal(payload []byte) error {
	return json.Unmarshal(payload, s)
}
