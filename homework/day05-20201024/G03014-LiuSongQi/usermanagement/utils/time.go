package utils

import (
	"time"
)

func TimeConversion(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	// fmt.Println(t, err)
	return t
}
