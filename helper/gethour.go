package helper

import "time"

func GetHour(hour string) time.Time {
	getHour, _ := time.Parse("15:04:05", hour)

	return getHour

}
