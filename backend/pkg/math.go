package pkg

import "math"

func AvgCeil(a, b int) float64 {
	if b == 0 {
		return 0
	}

	return math.Ceil(float64(a) / float64(b))
}
