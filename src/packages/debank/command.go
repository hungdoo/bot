package debank

import (
	"fmt"
	"strconv"

	u "github.com/hungdoo/bot/src/packages/utils"
	"github.com/hungdoo/bot/src/types"
	"github.com/shopspring/decimal"
)

type Command struct {
	types.Command
	low  decimal.Decimal
	high decimal.Decimal
	prev decimal.Decimal
}

func (c *Command) Validate(data []string) error {
	if len(data) < 3 {
		return fmt.Errorf("contract command needs at least 3 params: address, margin(1%%), precision")
	}
	return nil
}
func (c *Command) SetData(newValue []string) error {
	if err := c.Validate(newValue); err != nil {
		return err
	}
	c.Data = newValue
	return nil
}

func (c *Command) Execute(noCondition bool) (string, error) {
	userAddr, marginStr, precisionStr := c.Data[0], c.Data[1], c.Data[2]
	u.GeneralLogger.Printf("[%s] Execute: [%v]", c.GetName(), c.Data)

	debt, err := GetDebt(userAddr)
	if err != nil {
		return "", err
	}

	margin, err := decimal.NewFromString(marginStr)
	if err != nil {
		return "", err
	}
	precision, err := strconv.ParseInt(precisionStr, 10, 0)
	if err != nil {
		return "", err
	}

	if debt.IsPositive() {
		u.GeneralLogger.Printf("[%s] execution result: [%s]", c.GetName(), debt)
		if !margin.IsPositive() {
			margin = decimal.NewFromInt(1)
		}
		if noCondition {
			return fmt.Sprintf("%v\n<strong>V:%v | Pre: %v</strong>", c.Name, u.ShortenDecimal(debt, int32(precision), 2), u.ShortenDecimal(c.prev, int32(precision), 2)), nil
		} else if debt.GreaterThan(c.high) || debt.LessThan(c.low) {
			c.prev = debt
			c.high = debt.Mul(decimal.NewFromInt(100).Add(margin)).Div(decimal.NewFromInt(100))
			c.low = debt.Mul(decimal.NewFromInt(100).Sub(margin)).Div(decimal.NewFromInt(100))
			return fmt.Sprintf("%v\n<strong>V:%v | Pre: %v | L:%v | H:%v</strong>", c.Name, u.ShortenDecimal(debt, int32(precision), 2), u.ShortenDecimal(c.prev, int32(precision), 2), u.ShortenDecimal(c.low, int32(precision), 2), u.ShortenDecimal(c.high, int32(precision), 2)), nil
		}
	}
	return "", nil
}
