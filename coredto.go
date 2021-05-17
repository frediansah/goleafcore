package goleafcore

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/shopspring/decimal"
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

func (d Dto) GetInt(key string, defaultValue int) int {
	val, exist := d[key]
	if exist {
		if valStr, okCasting := val.(int); okCasting {
			return valStr
		} else {
			val, err := strconv.Atoi(fmt.Sprintf("%v", val))
			if err == nil {
				return int(val)
			}
		}
	}

	return defaultValue
}

func (d Dto) GetBool(key string, defaultValue bool) bool {
	val, exist := d[key]
	if exist {
		if valStr, okCasting := val.(bool); okCasting {
			return valStr
		} else {
			valBool, err := strconv.ParseBool(fmt.Sprintf("%v", val))
			if err == nil {
				return valBool
			}
		}
	}

	return defaultValue
}

func (d Dto) GetDecimal(key string, defaultValue decimal.Decimal) decimal.Decimal {
	val, exist := d[key]
	if exist {
		if valStr, okCasting := val.(decimal.Decimal); okCasting {
			return valStr
		} else {
			valAs, err := decimal.NewFromString(fmt.Sprintf("%v", val))
			if err == nil {
				return valAs
			}
		}
	}

	return defaultValue
}

func (d Dto) GetSlice(key string, defaultValue []interface{}) []interface{} {
	val, exist := d[key]
	if exist {
		if valStr, okCasting := val.([]interface{}); okCasting {
			return valStr
		}
	}

	return defaultValue
}

func (d Dto) GetSliceDto(key string, defaultValue []Dto) []Dto {
	val, exist := d[key]
	if exist {
		taip := reflect.TypeOf(val)
		if taip.Kind() == reflect.Slice {
			if valStr, okCasting := val.([]Dto); okCasting {
				return valStr
			}

			if valStr, okCasting := val.([]interface{}); okCasting {
				return convertToSliceOfDto(valStr)
			}
		}
	}

	return defaultValue
}

func (d Dto) GetDto(key string, defaultValue Dto) Dto {
	val, exist := d[key]
	if exist {
		if valStr, okCasting := val.(Dto); okCasting {
			return valStr
		} else {
			dto, err := NewDto(val)
			if err == nil {
				return dto
			}
		}
	}

	return defaultValue
}

func (d Dto) Put(key string, value interface{}) Dto {
	d[key] = value
	return d
}

func (d Dto) ContainKeys(key string) bool {
	_, exists := d[key]
	return exists
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

func NewOrEmpty(source interface{}) Dto {
	dto, err := NewDto(source)
	if err != nil {
		return Dto{}
	}

	return dto
}

func NewOrDefault(source interface{}, def Dto) Dto {
	dto, err := NewDto(source)
	if err != nil {
		return def
	}

	return dto
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

func convertToSliceOfDto(sliceMap []interface{}) []Dto {
	sliceDto := []Dto{}
	for _, val := range sliceMap {
		dto, err := NewDto(val)
		if err == nil {
			sliceDto = append(sliceDto, dto)
		}
	}

	return sliceDto
}
