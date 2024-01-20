package utils

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func DivDecimals(d decimal.Decimal, decimalPoints int32) decimal.Decimal {
	return d.Div(decimal.NewFromBigInt(big.NewInt(1), decimalPoints))
}
