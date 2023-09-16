package command

import (
	"fmt"
	"time"

	"github.com/hungdoo/bot/src/packages/interfaces"
)

type Command struct {
	interfaces.ICommand `bson:"-"`
	Name                string        `bson:"name"`
	Type                CommandType   `bson:"type"`
	Data                []string      `bson:"data"`
	ExecutedTime        time.Time     `bson:"executedtime"`
	IdleTime            time.Duration `bson:"idletime"`
	Enabled             bool          `bson:"enabled"`
}

// Setters
func (c *Command) SetEnabled(newValue bool) {
	c.Enabled = newValue
}
func (c *Command) SetExecutedTime(newValue time.Time) {
	c.ExecutedTime = newValue
}
func (c *Command) SetType(name string) error {
	if t := IsType(name); t != Unknown {
		c.Type = t
		return nil
	}
	return fmt.Errorf("unsupported command type for %s", name)
}

// Getters
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
