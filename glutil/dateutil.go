package glutil

import (
	"time"

	"gitlab.com/stsgoleaf/goleafcore/glconstant"
)

func DateTimeNow() string {
	now := time.Now()
	return now.Format(glconstant.DATETIME_FORMAT)
}

func DateNow() string {
	now := time.Now()
	return now.Format(glconstant.DATE_FORMAT)
}
