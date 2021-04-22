package glutil

import (
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
