package pkg

import "time"

func UTC8Now() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc)
}

func CurrentDate(format string) string {
	currentTime := time.Now()
	return currentTime.Format(format)
}