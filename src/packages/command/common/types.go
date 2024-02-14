package command

type CommandType int32

const (
	Unknown CommandType = iota
	ContractCall
	Tomb
	Debank
	Balance
	BybitIdo
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
	case BybitIdo:
		return "BybitIdo"
	default:
		return "unknown"
	}
}
