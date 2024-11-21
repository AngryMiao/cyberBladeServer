package pkg

import (
	"github.com/shopspring/decimal"
)

func Ratio(a, b int) float64 {
	if b == 0 {
		return 0
	}

	value := decimal.NewFromInt(int64(a))
	value = value.Div(decimal.NewFromInt(int64(b)))
	value = value.Mul(decimal.NewFromInt(100))
	result, _ := value.RoundBank(2).Float64()

	return result
}

func MulFloat64(a, b float64) float64 {
	if b == 0 {
		return 0
	}

	value := decimal.NewFromFloat(a)
	value = value.Mul(decimal.NewFromFloat(b))
	result, _ := value.RoundBank(2).Float64()

	return result
}

func AddFloat64(a, b float64) float64 {
	value := decimal.NewFromFloat(a)
	value = value.Add(decimal.NewFromFloat(b))
	result, _ := value.RoundBank(2).Float64()

	return result
}
