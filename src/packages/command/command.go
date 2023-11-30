package command

import (
	"fmt"
	"time"

	"github.com/hungdoo/bot/src/packages/interfaces"
	"github.com/shopspring/decimal"
)

type Command struct {
	interfaces.ICommand `json:"-" bson:"-"`
	Name                string          `bson:"name"`
	Type                CommandType     `json:"-" bson:"type"`
	Data                []string        `bson:"data"`
	ExecutedTime        time.Time       `bson:"executedtime"`
	IdleTime            time.Duration   `bson:"idletime"`
	Enabled             bool            `bson:"enabled"`
	Prev                decimal.Decimal `bson:"Prev"`
}

// Setters
func (c *Command) SetPrev(newValue decimal.Decimal) {
	c.Prev = newValue
}
func (c *Command) SetEnabled(newValue bool) {
	c.Enabled = newValue
}
func (c *Command) SetExecutedTime(newValue time.Time) {
	c.ExecutedTime = newValue
}
func (c *Command) SetIdleTime(newValue time.Duration) {
	c.IdleTime = newValue
}
func (c *Command) SetType(name string) error {
	if t := IsType(name); t != Unknown {
		c.Type = t
		return nil
	}
	return fmt.Errorf("unsupported command type for %s", name)
}

// Getters
func (c *Command) GetPrev() decimal.Decimal {
	return c.Prev
}
func (c *Command) GetData() []string {
	return c.Data
}
func (c *Command) GetName() string {
	return c.Name
}
func (c *Command) IsEnabled() bool {
	return c.Enabled
}
func (c *Command) IsIdle() bool {
	return time.Since(c.ExecutedTime) < c.IdleTime
}
func (c *Command) GetUnderlying() interface{} {
	return *c
}
