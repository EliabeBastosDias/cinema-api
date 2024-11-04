package utils

import (
	"fmt"
	"time"
)

func ParseDate(date string) (time.Time, error) {
	var dateParsed, err = time.Parse(time.DateOnly, date)

	if err != nil {
		return time.Time{}, fmt.Errorf("date %v is not in the correct format. Please use YYYY-MM-DD", date)
	}

	return dateParsed, nil
}
