package debank

import (
	"fmt"

	"github.com/hungdoo/bot/src/packages/command"
	"github.com/hungdoo/bot/src/packages/log"
	"github.com/hungdoo/bot/src/packages/math"
	"github.com/shopspring/decimal"
)

type Command struct {
	command.Command
	prev decimal.Decimal
}

func (c *Command) Validate(data []string) error {
	if len(data) < 2 {
		return fmt.Errorf("contract command needs at least 3 params: address, offset(+50000)")
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
	userAddr, offsetStr := c.Data[0], c.Data[1]
	log.GeneralLogger.Printf("[%s] Execute: [%v]", c.GetName(), c.Data)

	debt, err := GetDebt(userAddr)
	if err != nil {
		return "", err
	}

	offset, err := decimal.NewFromString(offsetStr)
	if err != nil {
		return "", err
	}

	if debt.IsPositive() {
		log.GeneralLogger.Printf("[%s] execution result: [%s]", c.GetName(), debt)
		if !offset.IsPositive() {
			offset = decimal.NewFromInt(50000)
		}
		_prev := c.prev
		if noCondition {
			return fmt.Sprintf("%v\nV:%v | Pre: %v", c.Name, math.ShortenDecimal(debt, 0, 2), math.ShortenDecimal(_prev, 0, 2)), nil
		} else if debt.GreaterThan(_prev.Add(offset)) || debt.LessThan(_prev.Sub(offset)) {
			c.prev = debt
			return fmt.Sprintf("%v\nV:%v | Pre: %v | L:%v | H:%v", c.Name, math.ShortenDecimal(debt, 0, 2), math.ShortenDecimal(_prev, 0, 2), math.ShortenDecimal(c.prev.Sub(offset), 0, 2), math.ShortenDecimal(c.prev.Add(offset), 0, 2)), nil
		}
	}
	return "", nil
}
