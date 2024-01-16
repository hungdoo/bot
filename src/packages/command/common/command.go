package command

import (
	"fmt"
	"time"

	"github.com/hungdoo/bot/src/packages/interfaces"
	"github.com/shopspring/decimal"
)

type Command struct {
	interfaces.ICommand `json:"-" bson:"-"`
	Name                string        `bson:"name"`
	Type                CommandType   `json:"-" bson:"type"`
	Data                []string      `bson:"data"`
	ExecutedTime        time.Time     `bson:"executedtime"`
	IdleTime            time.Duration `bson:"idletime"`
	Enabled             bool          `bson:"enabled"`
	Prev                string        `bson:"prev"`
	Error               string        `bson:"error"`
}

// Setters
func (c *Command) SetPrev(newValue decimal.Decimal) {
	c.Prev = newValue.String()
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
func (c *Command) SetError(err error) {
	c.Error = err.Error()
}

// Getters
func (c *Command) GetPrev() (decimal.Decimal, error) {
	d := decimal.Zero
	err := d.Scan(c.Prev)
	if err != nil {
		return decimal.Zero, err
	}
	return d, nil
}
func (c *Command) GetData() []string {
	return c.Data
}
func (c *Command) GetError() string {
	return c.Error
}
func (c *Command) GetOverview() string {
	lastErr := c.GetError()

	if len(lastErr) != 0 {
		return fmt.Sprintf("%v - %v - %.2fmin ago\nlastErr: %s", c.GetName(), c.Prev, time.Since(c.ExecutedTime).Minutes(), lastErr)
	}
	return fmt.Sprintf("%v - %v - %.2fmin ago", c.GetName(), c.Prev, time.Since(c.ExecutedTime).Minutes())

}
func (c *Command) GetName() string {
	return fmt.Sprintf("%v_%v", c.Name, c.Type.String())
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
