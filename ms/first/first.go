package first

import "encoding/json"

func Call(payload []byte) ([]byte, error) {
	var input Input
	err := input.Unmarshal(payload)
	if err != nil {
		return nil, err
	}
	output := Output{Value: input.Value}
	return json.Marshal(output)
}
