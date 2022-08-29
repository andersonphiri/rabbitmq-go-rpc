package common

import "encoding/json"

func DeserializeTo[T any](data []byte) (*T, error) {
	var result T 
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}

func Serialize(data any) ([]byte, error) {
	return json.Marshal(data)
}