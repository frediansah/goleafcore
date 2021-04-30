package goleafcore

import "encoding/json"

type Dto map[string]interface{}

func (d Dto) ToJsonString() string {
	jsonByte, err := json.Marshal(d)
	if err != nil {
		return "{}"
	}

	return string(jsonByte)
}
