package command

import (
	"time"

	"github.com/hungdoo/bot/src/common"
	"github.com/shopspring/decimal"
)

type ICommand interface {
	SetPrev(newValue decimal.Decimal)
	SetData(newValue []string) error
	SetExecutedTime(newValue time.Time)
	SetIdleTime(newValue time.Duration)
	SetEnabled(newValue bool)
	SetError(err string)
	SetDisplayMsg(msg string)

	GetPrev() decimal.Decimal
	GetDisplayMsg() string
	GetName() string
	GetType() CommandType
	GetError() string
	GetOverview() string
	GetData() []string
	GetUnderlying() interface{}
	IsEnabled() bool
	IsIdle() bool

	Validate(data []string) error
	Execute(mustReport bool, cmd string) (string, *common.ErrorWithSeverity)
}
