package goleafcore

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Dto map[string]interface{}

func (d Dto) ToJsonString() string {
	jsonByte, err := json.Marshal(d)
	if err != nil {
		return "{}"
	}

	return string(jsonByte)
}

func (d Dto) Get(key string, defaultValue interface{}) interface{} {
	val, exist := d[key]
	if exist {
		return val
	}

	return defaultValue
}

func (d Dto) GetString(key, defaultValue string) string {
	val, exist := d[key]
	if exist {
		if valStr, okCasting := val.(string); okCasting {
			return valStr
		} else {
			return fmt.Sprintf("%v", val)
		}
	}

	return defaultValue
}

func (d Dto) GetInt64(key string, defaultValue int64) int64 {
	val, exist := d[key]
	if exist {
		if valStr, okCasting := val.(int64); okCasting {
			return valStr
		} else {
			val, err := strconv.Atoi(fmt.Sprintf("%v", val))
			if err == nil {
				return int64(val)
			}
		}
	}

	return defaultValue
}

func (d Dto) Put(key string, value interface{}) Dto {
	d[key] = value
	return d
}

func NewDto(source interface{}) (Dto, error) {
	typ := reflect.TypeOf(source)
	val := source
	if typ.Kind() == reflect.Ptr {
		val = reflect.ValueOf(source).Elem().Interface()
	}

	if valStr, okStr := val.(string); okStr {
		return newFromString(valStr)
	}

	if valSByte, okSByte := val.([]byte); okSByte {
		return newFromSByte(valSByte)
	}

	return newFromStruct(val)
}

func newFromString(str string) (Dto, error) {
	return newFromSByte([]byte(str))
}

func newFromSByte(data []byte) (Dto, error) {
	var dto Dto
	err := json.Unmarshal(data, &dto)

	return dto, err
}

func newFromStruct(obj interface{}) (Dto, error) {
	var dto Dto
	jsonByte, err := json.Marshal(obj)
	if err == nil {
		err = json.Unmarshal(jsonByte, &dto)
	}

	return dto, err
}
