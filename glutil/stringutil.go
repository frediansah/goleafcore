package glutil

import (
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
