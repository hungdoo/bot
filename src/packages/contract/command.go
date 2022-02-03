package contract

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/alethio/web3-go/ethrpc"
	u "github.com/hungdoo/bot/src/packages/utils"
	"github.com/hungdoo/bot/src/types"
	"github.com/shopspring/decimal"
)

type ContractCommand struct {
	types.Command
	low  decimal.Decimal
	high decimal.Decimal
}

func (c *ContractCommand) Validate(data []string) error {
	if len(data) < 7 {
		return fmt.Errorf("contract command needs at least 4 params: rpc, contract address, method (getUserAccountData(address)(uint256), params(pr1;pr2), value index, margin(1%%), precision")
	}
	return nil
}
func (c *ContractCommand) SetData(newValue []string) error {
	if err := c.Validate(newValue); err != nil {
		return err
	}
	c.Data = newValue
	return nil
}

func (c *ContractCommand) Execute(noCondition bool) (string, error) {
	rpc, contractAddr, method, params, valueIdx, marginStr, precisionStr := c.Data[0], c.Data[1], c.Data[2], c.Data[3], c.Data[4], c.Data[5], c.Data[6]
	u.GeneralLogger.Printf("[%s] Execute: [%v]", c.GetName(), c.Data)

	eth, err := GetETH(rpc)
	if err != nil {
		return "", err
	}

	args := []interface{}{}
	for _, p := range strings.Split(params, ";") {
		args = append(args, p)
	}
	vc := u.NewViewCall(
		method,
		args,
	)

	packed, err := vc.CallData()
	if err != nil {
		return "", err
	}

	res, err := eth.CallContractFunction("0x"+hex.EncodeToString(packed), contractAddr, ethrpc.DefaultCallGas)
	if err != nil {
		return "", err
	}

	bytes, err := hex.DecodeString(strings.TrimPrefix(res, "0x"))
	if err != nil {
		return "", err
	}
	values, err := vc.Decode(bytes)
	if err != nil {
		return "", err
	}

	valueIndex, err := strconv.ParseInt(valueIdx, 10, 0)
	if err != nil {
		return "", err
	}
	margin, err := decimal.NewFromString(marginStr)
	if err != nil {
		return "", err
	}
	precision, err := strconv.ParseInt(precisionStr, 10, 0)
	if err != nil {
		return "", err
	}

	if value, ok := values[valueIndex].(*big.Int); ok {
		u.GeneralLogger.Printf("[%s] execution result: [%s]", c.GetName(), value)
		if !margin.IsPositive() {
			margin = decimal.NewFromInt(1)
		}
		valueDecimal := decimal.NewFromBigInt(value, 0)
		if noCondition {
			return fmt.Sprintf("<strong>V:%v</strong>", valueDecimal.DivRound(decimal.New(1, int32(precision)), 2)), nil
		} else if valueDecimal.GreaterThan(c.high) || valueDecimal.LessThan(c.low) {
			c.high = valueDecimal.Mul(decimal.NewFromInt(100).Add(margin)).Div(decimal.NewFromInt(100))
			c.low = valueDecimal.Mul(decimal.NewFromInt(100).Sub(margin)).Div(decimal.NewFromInt(100))
			return fmt.Sprintf("<strong>V:%v | L:%v | H:%v</strong>", valueDecimal.DivRound(decimal.New(1, int32(precision)), 2), c.low.DivRound(decimal.New(1, int32(precision)), 2), c.high.DivRound(decimal.New(1, int32(precision)), 2)), nil
		}
	} else {
		return "", fmt.Errorf("cannot parse value [%v]", values...)
	}
	return "", nil
}
