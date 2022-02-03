package types

import "time"

type ICommand interface {
	SetData(newValue []string) error
	SetExecutedTime(newValue time.Time)
	SetEnabled(newValue bool)

	GetName() string
	GetData() []string
	IsEnabled() bool
	IsIdle() bool

	Validate(data []string) error
	Execute(noCondition bool) (string, error)
}

type Command struct {
	Name         string
	Data         []string
	ExecutedTime time.Time
	IdleTime     time.Duration
	Enabled      bool
}

func (c *Command) SetEnabled(newValue bool) {
	c.Enabled = newValue
}
func (c *Command) SetExecutedTime(newValue time.Time) {
	c.ExecutedTime = newValue
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
