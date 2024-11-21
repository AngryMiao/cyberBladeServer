package delay

import (
	"time"
)

func Delay(duration time.Duration) {
	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {

	case <-timer.C:
		break
	}

}
