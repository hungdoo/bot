package command

import "strings"

type CommandType int

const (
	Unknown CommandType = iota
	ContractCall
	Tomb
	Debank
)

func (s CommandType) String() string {
	switch s {
	case ContractCall:
		return "callcontract"
	case Tomb:
		return "tomb"
	case Debank:
		return "debank"
	default:
		return "unknown"
	}
}

func IsType(name string) CommandType {
	for i := ContractCall; i < Debank; i++ {
		if strings.HasPrefix(strings.ToLower(name), i.String()) {
			return i
		}
	}
	return Unknown
}
