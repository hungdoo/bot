package command

type CommandType int

const (
	Unknown CommandType = iota
	ContractCall
	Tomb
	Debank
	Balance
)

func (s CommandType) String() string {
	switch s {
	case ContractCall:
		return "callcontract"
	case Tomb:
		return "tomb"
	case Debank:
		return "debank"
	case Balance:
		return "balance"
	default:
		return "unknown"
	}
}
