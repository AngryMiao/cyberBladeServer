package time

import (
	"time"
)

var Day = time.Hour * time.Duration(24)

func IntToDuration(duration int) time.Duration {
	return time.Duration(duration)
}

func CalcTime(baseTime time.Time, duration int, durationUnit time.Duration) time.Time {
	var delta = time.Duration(duration) * durationUnit
	return baseTime.Add(delta)
}
