package graphql

import "github.com/shopspring/decimal"

type Markets struct {
	Id                 string
	Name               string
	CollateralFactor   decimal.NullDecimal
	UnderlyingSymbol   string
	UnderlyingDecimals decimal.NullDecimal
}

var MarketQuery struct {
	Markets []Markets `graphql:"markets"`
}
