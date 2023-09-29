package contract

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/alethio/web3-go/ethrpc"
	"github.com/hungdoo/bot/src/packages/command"
	"github.com/hungdoo/bot/src/packages/log"
	"github.com/hungdoo/bot/src/packages/math"
	"github.com/shopspring/decimal"
)

type Command struct {
	command.Command
	prev decimal.Decimal
}

func (c *Command) Validate(data []string) error {
	if len(data) < 7 {
		return fmt.Errorf("contract command needs at least 4 params: rpc, contract address, method (getUserAccountData(address)(uint256), params(pr1;pr2), value index, margin(1%%), precision")
	}
	return nil
}
func (c *Command) SetData(newValue []string) error {
	if err := c.Validate(newValue); err != nil {
		return err
	}
	c.Data = newValue
	return nil
}

func (c *Command) Execute(noCondition bool) (string, error) {
	rpc, contractAddr, method, params, valueIdx, marginStr, precisionStr := c.Data[0], c.Data[1], c.Data[2], c.Data[3], c.Data[4], c.Data[5], c.Data[6]
	log.GeneralLogger.Printf("[%s] Execute: %v", c.GetName(), c.Data)

	eth, err := GetETH(rpc)
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", err
	}

	args := []interface{}{}
	for _, p := range strings.Split(params, ";") {
		args = append(args, p)
	}
	vc := NewViewCall(
		method,
		args,
	)

	packed, err := vc.CallData()
	if err != nil {
		log.ErrorLogger.Printf("CallData Err: %v-%v", vc, err)
		return "", err
	}

	res, err := eth.CallContractFunction("0x"+hex.EncodeToString(packed), contractAddr, ethrpc.DefaultCallGas)
	if err != nil {
		log.ErrorLogger.Printf("CallContractFunction Err: %v-%v", vc, err)
		return "", err
	}

	bytes, err := hex.DecodeString(strings.TrimPrefix(res, "0x"))
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", err
	}
	values, err := vc.Decode(bytes)
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", err
	}

	valueIndex, err := strconv.ParseInt(valueIdx, 10, 0)
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", err
	}
	margin, err := decimal.NewFromString(marginStr)
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", err
	}
	precision, err := strconv.ParseInt(precisionStr, 10, 0)
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", err
	}

	if value, ok := values[valueIndex].(*big.Int); ok {
		log.GeneralLogger.Printf("[%s] execution result: [%s]", c.GetName(), value)
		if !margin.IsPositive() {
			margin = decimal.NewFromInt(1)
		}
		valueDecimal := decimal.NewFromBigInt(value, 0)
		_prev := c.prev
		_high := c.prev.Mul(decimal.NewFromInt(100).Add(margin)).Div(decimal.NewFromInt(100))
		_low := c.prev.Mul(decimal.NewFromInt(100).Sub(margin)).Div(decimal.NewFromInt(100))
		if noCondition {
			return fmt.Sprintf("%v\nV:%v | Pre: %v", c.Name, math.ShortenDecimal(valueDecimal, int32(precision), 2), math.ShortenDecimal(_prev, int32(precision), 2)), nil
		} else if valueDecimal.GreaterThan(_high) || valueDecimal.LessThan(_low) {
			c.prev = valueDecimal
			newHigh := c.prev.Mul(decimal.NewFromInt(100).Add(margin)).Div(decimal.NewFromInt(100))
			newLow := c.prev.Mul(decimal.NewFromInt(100).Sub(margin)).Div(decimal.NewFromInt(100))
			return fmt.Sprintf("%v\nV:%v | Pre: %v | L:%v | H:%v", c.Name, math.ShortenDecimal(valueDecimal, int32(precision), 2), math.ShortenDecimal(_prev, int32(precision), 2), math.ShortenDecimal(newLow, int32(precision), 2), math.ShortenDecimal(newHigh, int32(precision), 2)), nil
		}
	} else {
		return "", fmt.Errorf("cannot parse value [%v]", values...)
	}
	return "", nil
}
