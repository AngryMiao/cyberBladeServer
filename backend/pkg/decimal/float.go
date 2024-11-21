package decimal

import "github.com/shopspring/decimal"

func Float64RatioRoundDown(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	value := decimal.NewFromFloat(a)
	value = value.Div(decimal.NewFromFloat(b))
	result, _ := value.RoundDown(2).Float64()

	return result
}

func Float64Mul(a, b float64) float64 {
	if b == 0 {
		return 0
	}

	value := decimal.NewFromFloat(a)
	value = value.Mul(decimal.NewFromFloat(b))
	result, _ := value.RoundBank(2).Float64()

	return result
}

func Float64Add(a, b float64) float64 {
	value := decimal.NewFromFloat(a)
	value = value.Add(decimal.NewFromFloat(b))
	result, _ := value.RoundBank(2).Float64()

	return result
}

func Float64Sub(a, b float64) float64 {
	value := decimal.NewFromFloat(a)
	value = value.Sub(decimal.NewFromFloat(b))
	result, _ := value.RoundBank(2).Float64()

	return result
}

func Float64LessThan(a, b float64) bool {
	decimalA := decimal.NewFromFloat(a)
	decimalB := decimal.NewFromFloat(b)
	return decimalA.LessThan(decimalB)
}

func Float64Small(a, b float64) float64 {
	decimalA := decimal.NewFromFloat(a)
	decimalB := decimal.NewFromFloat(b)

	var result float64
	if decimalA.LessThan(decimalB) {
		result, _ = decimalA.Float64()
	} else {
		result, _ = decimalB.Float64()
	}

	return result
}
