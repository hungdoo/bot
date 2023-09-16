package graphql

import "github.com/shopspring/decimal"

type Markets struct {
	Id                 string
	Name               string
	CollateralFactor   decimal.Decimal
	UnderlyingSymbol   string
	UnderlyingDecimals decimal.Decimal
}

var MarketQuery struct {
	Markets []Markets `graphql:"markets"`
}
