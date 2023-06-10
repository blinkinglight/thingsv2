package first

import "encoding/json"

type Input struct {
	Value string `json:"value"`
}

func (i *Input) Marshal() ([]byte, error) {
	return json.Marshal(i)
}

func (i *Input) Unmarshal(payload []byte) error {
	return json.Unmarshal(payload, i)
}

type Output struct {
	Value string `json:"value"`
}
