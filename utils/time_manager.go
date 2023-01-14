package utils

import (
	"errors"
	"time"
)

func GetCurrentDateTimeString(format *string) (*string, error) {
	dt := time.Now()
	var dateTime string
	if format == nil {
		dateTime = dt.Format("2006-01-02T15:04:05-0700")
	} else {
		dateTime = dt.Format(*format)
	}
	if len(dateTime) == 0 {
		return nil, errors.New("could not format current datetime with given format string")
	}
	return &dateTime, nil
}
