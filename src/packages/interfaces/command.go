package interfaces

import (
	"time"

	"github.com/shopspring/decimal"
)

type ICommand interface {
	SetPrev(newValue decimal.Decimal)
	SetData(newValue []string) error
	SetExecutedTime(newValue time.Time)
	SetIdleTime(newValue time.Duration)
	SetEnabled(newValue bool)
	SetError(err error)

	GetPrev() (decimal.Decimal, error)
	GetName() string
	GetError() string
	GetOverview() string
	GetData() []string
	GetUnderlying() interface{}
	IsEnabled() bool
	IsIdle() bool

	Validate(data []string) error
	Execute(noCondition bool, cmd string) (string, error)
}
