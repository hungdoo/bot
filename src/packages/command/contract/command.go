package contract

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/alethio/web3-go/ethrpc"
	"github.com/hungdoo/bot/src/common"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/log"
	"github.com/hungdoo/bot/src/packages/math"
	"github.com/shopspring/decimal"
)

const DECIMAL_POINTS = 4

type Command struct {
	command.Command
}

func (c *Command) Validate(data []string) error {
	if len(data) < 7 {
		return fmt.Errorf("invalid params: rpc, contract address, method (getUserAccountData(address)(uint256), params(pr1;pr2), value index, margin(1%%), precision")
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

func (c *Command) Execute(mustReport bool, _ string) (string, *common.ErrorWithSeverity) {
	rpc, contractAddr, method, params, valueIdx, marginStr, precisionStr := c.Data[0], c.Data[1], c.Data[2], c.Data[3], c.Data[4], c.Data[5], c.Data[6]
	log.GeneralLogger.Printf("[%s] Execute: %v", c.GetName(), c.Data)

	eth, err := GetETH(rpc)
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	args := []interface{}{}
	for _, p := range strings.Split(params, ";") {
		args = append(args, p)
	}
	vc := NewViewCall(
		method,
		nil,
	)
	argTypes := vc.ArgumentTypes()
	if len(argTypes) != 0 {
		vc.arguments = args
	}

	packed, err := vc.CallData()
	if err != nil {
		log.ErrorLogger.Printf("CallData Err: %v-%v", vc, err)
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	res, err := eth.CallContractFunction("0x"+hex.EncodeToString(packed), contractAddr, ethrpc.DefaultCallGas)
	if err != nil {
		log.ErrorLogger.Printf("CallContractFunction Err: %v-%v", vc, err)
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	bytes, err := hex.DecodeString(strings.TrimPrefix(res, "0x"))
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}
	values, err := vc.Decode(bytes)
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	valueIndex, err := strconv.ParseInt(valueIdx, 10, 0)
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}
	margin, err := decimal.NewFromString(marginStr)
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}
	precision, err := strconv.ParseInt(precisionStr, 10, 0)
	if err != nil {
		log.ErrorLogger.Printf("Err: %v", err)
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	if value, ok := values[valueIndex].(*big.Int); ok {
		if !margin.IsPositive() {
			margin = decimal.NewFromInt(1)
		}
		valueDecimal := decimal.NewFromBigInt(value, 0)
		_prev := c.GetPrev()
		_high := _prev.Mul(decimal.NewFromInt(100).Add(margin)).Div(decimal.NewFromInt(100))
		_low := _prev.Mul(decimal.NewFromInt(100).Sub(margin)).Div(decimal.NewFromInt(100))
		if mustReport {
			return fmt.Sprintf("%v\nV:%v | Pre: %v", c.Name, math.ShortenDecimal(valueDecimal, int32(precision), DECIMAL_POINTS), math.ShortenDecimal(_prev, int32(precision), DECIMAL_POINTS)), nil
		} else if valueDecimal.GreaterThan(_high) || valueDecimal.LessThan(_low) {
			c.SetPrev(valueDecimal)

			newHigh := _prev.Mul(decimal.NewFromInt(100).Add(margin)).Div(decimal.NewFromInt(100))
			newLow := _prev.Mul(decimal.NewFromInt(100).Sub(margin)).Div(decimal.NewFromInt(100))
			return fmt.Sprintf("%v\nV:%v | Pre: %v | L:%v | H:%v", c.Name, math.ShortenDecimal(valueDecimal, int32(precision), DECIMAL_POINTS), math.ShortenDecimal(_prev, int32(precision), DECIMAL_POINTS), math.ShortenDecimal(newLow, int32(precision), DECIMAL_POINTS), math.ShortenDecimal(newHigh, int32(precision), DECIMAL_POINTS)), nil
		}
		return "", common.NewErrorWithSeverity(common.Info, fmt.Sprintf("no report for result [%v]", value))
	} else {
		return "", common.NewErrorWithSeverity(common.Error, fmt.Sprintf("cannot parse value [%v]", values...))
	}
}
