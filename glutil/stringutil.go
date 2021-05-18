package glutil

import (
	"fmt"
	"regexp"
	"strings"
)

func ToUnderedScore(str string) string {
	if len(str) > 0 {
		a := regexp.MustCompile(`[A-Z][^A-Z]*`)
		splitByUpperCase := a.FindAllString(str, -1)

		merged := ""
		for _, item := range splitByUpperCase {
			if len(merged) > 0 {
				merged = merged + "_" + strings.ToLower(item)
			} else {
				merged = strings.ToLower(item)
			}
		}

		return merged
	}

	return str
}

func ToString(value interface{}) string {
	str := fmt.Sprintf("%v", value)
	return str
}

func AppendSliceString(list []string, appender string) string {
	var result string
	for _, item := range list {
		if len(result) > 0 {
			result = result + appender
		}

		result = result + item
	}
	return result
}
