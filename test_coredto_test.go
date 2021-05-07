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

	dtoFromAnotherDto.Put("decimal", decimalValue)
	log.Println("Dto print decimal : ", dtoFromAnotherDto.ToJsonString())

}
