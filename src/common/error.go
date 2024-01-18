package common

import (
	"errors"
	"fmt"
)

type SeverityLevel int

const (
	Info SeverityLevel = iota
	Warning
	Error
	Critical
)

func (s SeverityLevel) String() string {
	switch s {
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case Critical:
		return "Critical"
	default:
		return "Unknown"
	}
}

type ErrorWithSeverity struct {
	Level         SeverityLevel
	OriginalError error
}

func NewErrorWithSeverity(level SeverityLevel, msg string) *ErrorWithSeverity {
	return &ErrorWithSeverity{
		Level:         level,
		OriginalError: errors.New(msg),
	}
}

func (e *ErrorWithSeverity) Error() string {
	return fmt.Sprintf("[%s] %s", e.Level.String(), e.OriginalError.Error())
}
