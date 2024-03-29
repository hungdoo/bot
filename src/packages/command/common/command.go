package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type Command struct {
	ICommand     `json:"-" bson:"-"`
	Name         string        `json:"name" bson:"name,omitempty"`
	Type         CommandType   `json:"type" bson:"type"`
	Data         []string      `json:"data" bson:"data"`
	ExecutedTime time.Time     `json:"-" bson:"executed_time"`
	IdleTime     time.Duration `json:"idletime" bson:"idletime"`
	Enabled      bool          `json:"-" bson:"enabled"`
	Prev         string        `json:"-" bson:"prev"`
	DisplayMsg   string        `json:"-" bson:"display_msg"`
	Error        string        `json:"-" bson:"error"`
}

type CustomCommand struct {
	Command
	Id string `bson:"_id,unique"`
}

// Setters
func (c *Command) SetDisplayMsg(msg string) {
	c.DisplayMsg = msg
}
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
func (c *Command) SetError(err string) {
	c.Error = err
}

// Getters
func (c *Command) GetPrev() decimal.Decimal {
	d := decimal.Zero
	if len(strings.TrimSpace(c.Prev)) == 0 {
		return decimal.Zero
	}
	err := d.Scan(c.Prev)
	if err != nil {
		return decimal.Zero
	}
	return d
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
		return fmt.Sprintf("[%s:%v]\n%v - %.2fm\nlastErr: %s", c.Type.String(), c.Name, c.GetDisplayMsg(), time.Since(c.ExecutedTime).Minutes(), lastErr)
	}
	return fmt.Sprintf("[%s:%v]\n%v - %.2fm", c.Type.String(), c.Name, c.GetDisplayMsg(), time.Since(c.ExecutedTime).Minutes())

}
func (c *Command) GetType() CommandType {
	return c.Type
}

func (c *Command) GetDisplayMsg() string {
	return c.DisplayMsg
}

// @dev don't change. need consistency for db access
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
