package helper

import "encoding/json"

func UnmarshalJSON(data []byte, target interface{}) error {
	return json.Unmarshal(data, target)
}
