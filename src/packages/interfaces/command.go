package interfaces

import "time"

type ICommand interface {
	SetData(newValue []string) error
	SetExecutedTime(newValue time.Time)
	SetEnabled(newValue bool)
	SetType(name string) error

	GetName() string
	GetData() []string
	GetUnderlying() interface{}
	IsEnabled() bool
	IsIdle() bool

	Validate(data []string) error
	Execute(noCondition bool) (string, error)
}
