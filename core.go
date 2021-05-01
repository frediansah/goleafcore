package goleafcore

import js "encoding/json"

type Dto map[string]interface{}

func (d Dto) ToJsonString() string {
	jsonByte, err := js.Marshal(d)
	if err != nil {
		return "{}"
	}

	return string(jsonByte)
}
