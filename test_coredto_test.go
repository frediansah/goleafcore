package goleafcore_test

import (
	"log"
	"testing"

	"github.com/frediansah/goleafcore"
	"github.com/shopspring/decimal"
)

func TestDto(t *testing.T) {
	dto := goleafcore.Dto{
		"name":    "Frediansah",
		"age":     400,
		"gender":  "Male",
		"cool":    true,
		"baddass": false,
	}

	log.Println("Dto getString gender : ", dto.GetString("gender", ""))
	log.Println("Dto getString age : ", dto.GetString("age", "90"))

	log.Println("Dto before : ", dto.ToJsonString())

	dto.Put("carrie", "Pro-gamer")

	log.Println("Dto after : ", dto.ToJsonString())

	dtoFromStruct, errFromDto := goleafcore.NewDto(struct {
		ProductId   int64  `json:"productId"`
		ProductName string `json:"productName"`
	}{
		ProductId:   12,
		ProductName: "Sandals",
	})

	log.Println("Error casting dto : ", errFromDto)
	log.Println("Dto from struct : ", dtoFromStruct.ToJsonString())

	dtoFromString, errFromDto := goleafcore.NewDto("{\"heroName\":\"Dark Lord Kasel\",\"attack\":36000,\"def\":353566}")
	log.Println("Error casting dto : ", errFromDto)
	log.Println("Dto from string : ", dtoFromString.ToJsonString())

	dtoFromAnotherDto, errFromDto := goleafcore.NewDto(dto)
	log.Println("Error casting dto : ", errFromDto)
	log.Println("Dto from anotherDto : ", dtoFromAnotherDto.ToJsonString())

	decimalValue, _ := decimal.NewFromString("2442.234324")
	sliceDto := []goleafcore.Dto{}
	sliceDto = append(sliceDto, goleafcore.Dto{
		"id": 1,
	})
	sliceDto = append(sliceDto, goleafcore.Dto{
		"id": 2,
	})

	dtoFromAnotherDto.Put("decimal", decimalValue)
	dtoFromAnotherDto.Put("slice", []interface{}{"aku", "adalah", "anak", "gembala"})
	dtoFromAnotherDto.Put("sliceDto", sliceDto)

	decValue := dtoFromAnotherDto.GetDecimal("decimal", decimal.Zero)
	containOraOno := dtoFromAnotherDto.ContainKeys("containOraOno")
	valueSlice := dtoFromAnotherDto.GetSlice("slice", []interface{}{})

	log.Println("sliceDto : ", dtoFromAnotherDto.GetSliceDto("sliceDto", []goleafcore.Dto{}))
	log.Println("Decimal : ", decValue)
	log.Println("containOraOno : ", containOraOno)
	log.Println("Dto print decimal : ", dtoFromAnotherDto.ToJsonString())
	log.Println("slice of interface : ", valueSlice)

	sliceDtoFromGet := dtoFromAnotherDto.GetSliceDto("sliceDto", []goleafcore.Dto{})
	log.Println("Item dto from get : ")
	for _, itemDto := range sliceDtoFromGet {
		log.Println("item ", itemDto.ToJsonString())
	}

}

func TestDtoWithSlice(t *testing.T) {
	str := `{"userList":[{"id":10, "name":"Name1"},{"id":11, "name":"Name2"}]}`

	dto := goleafcore.NewOrEmpty(str)

	userList := dto.GetSliceDto("userList", []goleafcore.Dto{})
	log.Println("User list : ")
	for _, userDto := range userList {
		log.Println("User : ", userDto.ToJsonString())
	}
}

func TestDtoGetDto(t *testing.T) {
	str := `{"payload":{"status":"OK", "id":10}}`

	dto := goleafcore.NewOrEmpty(str)

	payloadDto := dto.GetDto("payload", goleafcore.Dto{})
	log.Println("payloadDto : ", payloadDto.ToJsonString())

	for key, val := range payloadDto {
		log.Println("key : ", key, " --> val :", val)
	}
}
