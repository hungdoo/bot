package math

import "github.com/shopspring/decimal"

func ShortenDecimal(value decimal.Decimal, precision int32, round int32) decimal.Decimal {
	return value.DivRound(decimal.New(1, precision), round)
}
