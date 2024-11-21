package pkg

import "github.com/shopspring/decimal"

func TurnYuanToCentInt64(amount float64) int64 {
	amountDecimal := decimal.NewFromFloat(amount)
	amountDecimal = amountDecimal.Mul(decimal.NewFromInt(100))

	return amountDecimal.IntPart()
}

func TurnCentToYuanFloat64(amount int64) (float64, bool) {
	decimal.DivisionPrecision = 2
	amountDecimal := decimal.New(amount, -2)
	return amountDecimal.Float64()
}
