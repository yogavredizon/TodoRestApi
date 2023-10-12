package helper

import "time"

func GetHour(hour string) time.Time {
	getHour, _ := time.Parse(time.Kitchen, hour)

	return getHour

}
