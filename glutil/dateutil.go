package glutil

import (
	"errors"
	"time"

	"github.com/frediansah/goleafcore/glconstant"
)

func DateTimeNow() string {
	now := time.Now()
	return now.Format(glconstant.DATETIME_FORMAT)
}

func DateNow() string {
	now := time.Now()
	return now.Format(glconstant.DATE_FORMAT)
}

func ParseDate(dateStr string) (*time.Time, error) {
	if len(dateStr) == len(glconstant.DATE_FORMAT) {
		val, err := time.Parse(glconstant.DATE_FORMAT, dateStr)
		return &val, err
	}

	if len(dateStr) == len(glconstant.DATETIME_FORMAT) {
		val, err := time.Parse(glconstant.DATETIME_FORMAT, dateStr)
		return &val, err
	}

	return nil, errors.New("invalid date string use date layout  '" + glconstant.DATE_FORMAT + "' or '" + glconstant.DATETIME_FORMAT + "'")
}
