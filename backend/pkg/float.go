package pkg

import (
	"github.com/shopspring/decimal"
)

func FloatEquals(a, b float64) bool {
	return decimal.NewFromFloat(a).Equal(decimal.NewFromFloat(b))
}

func FloatGreaterThanOrEqual(a, b float64) bool {
	return decimal.NewFromFloat(a).GreaterThanOrEqual(decimal.NewFromFloat(b))
}

func FloatGreaterThan(a, b float64) bool {
	return decimal.NewFromFloat(a).GreaterThan(decimal.NewFromFloat(b))
}

func BoolToPointer(v bool) *bool {
	return &v
}
