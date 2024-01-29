package contract

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/alethio/web3-go/ethrpc"
	"github.com/hungdoo/bot/src/common"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/log"
	"github.com/hungdoo/bot/src/packages/utils"
	"github.com/shopspring/decimal"
)

type ContractCommand struct {
	command.Command
	Id         string   `bson:"_id,unique"`
	Rpc        string   `json:"rpc" bson:"rpc"`
	Contract   string   `json:"contract" bson:"contract"`
	Method     string   `json:"method" bson:"method"`
	Params     []string `json:"params" bson:"params"`
	ValueIndex int64    `json:"index" bson:"index"`
	Margin     int64    `json:"margin" bson:"margin"`
	Precision  int64    `json:"precision" bson:"precision"`
}

func (c ContractCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name string `json:"name"`
		Type string `json:"type"`
		// Data     []string `json:"data"`
		IdleTime   string `json:"idletime"`
		Rpc        string `json:"rpc"`
		Contract   string `json:"contract"`
		Method     string `json:"method" bson:"method"`
		Params     string `json:"params" bson:"params"`
		ValueIndex int64  `json:"index" bson:"index"`
		Margin     int64  `json:"margin" bson:"margin"`
		Precision  int64  `json:"precision" bson:"precision"`
		Command    string `json:"command"`
	}{
		Name: c.Name,
		Type: c.Type.String(),
		// Data:     c.Data,
		IdleTime:   c.IdleTime.String(),
		Rpc:        c.Rpc,
		Contract:   c.Contract,
		Method:     c.Method,
		Params:     strings.Join(c.Params, ";"),
		ValueIndex: c.ValueIndex,
		Margin:     c.Margin,
		Precision:  c.Precision,
		Command:    fmt.Sprintf("add call %s %s %s %s %s %d %d %d", c.Name, c.Rpc, c.Contract, c.Method, strings.Join(c.Params, ";"), c.ValueIndex, c.Margin, c.Precision),
	})
}

func (c *ContractCommand) Validate(data []string) error {
	if len(data) < 7 {
		return fmt.Errorf("invalid params: rpc, contract address, method (getUserAccountData(address)(uint256), params(pr1;pr2) or \"\", value index, margin(1%%), precision")
	}
	return nil
}
func (c *ContractCommand) SetData(newValue []string) (err error) {
	if err := c.Validate(newValue); err != nil {
		return err
	}
	// c.Data = newValue
	c.Rpc = newValue[0]
	c.Contract = newValue[1]
	c.Method = newValue[2]
	if newValue[3] != "\"\"" {
		c.Params = strings.Split(newValue[3], ";")
	}
	c.ValueIndex, err = strconv.ParseInt(newValue[4], 10, 0)
	if err != nil {
		return err
	}
	c.Margin, err = strconv.ParseInt(newValue[5], 10, 0)
	if err != nil {
		return err
	}
	c.Precision, err = strconv.ParseInt(newValue[6], 10, 0)
	if err != nil {
		return err
	}
	return nil
}

func (c *ContractCommand) Execute(mustReport bool, _ string) (string, *common.ErrorWithSeverity) {
	// Backward compatiple: try to fill struct with old c.Data if empty prop
	if len(c.Rpc) == 0 {
		err := c.SetData(c.GetData())
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Error, err.Error())
		}
	}
	eth, err := GetETH(c.Rpc)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	args := []interface{}{}
	for _, p := range c.Params {
		args = append(args, p)
	}
	vc := NewViewCall(
		c.Method,
		nil,
	)
	argTypes := vc.ArgumentTypes()
	if len(argTypes) != 0 {
		vc.arguments = args
	}

	packed, err := vc.CallData()
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	res, err := eth.CallContractFunction("0x"+hex.EncodeToString(packed), c.Contract, ethrpc.DefaultCallGas)
	if err != nil {
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

	margin := decimal.NewFromInt(c.Margin)
	if value, ok := values[c.ValueIndex].(*big.Int); ok {
		if !margin.IsPositive() {
			margin = decimal.NewFromInt(1)
		}
		valueDecimal := decimal.NewFromBigInt(value, 0)
		_prev := c.GetPrev()
		_high := _prev.Mul(decimal.NewFromInt(100).Add(margin)).Div(decimal.NewFromInt(100))
		_low := _prev.Mul(decimal.NewFromInt(100).Sub(margin)).Div(decimal.NewFromInt(100))

		if mustReport {
			return fmt.Sprintf(
				"V:%.4f | Pre: %.4f",
				utils.DivDecimals(valueDecimal, int32(c.Precision)).InexactFloat64(),
				utils.DivDecimals(_prev, int32(c.Precision)).InexactFloat64(),
			), nil
		} else if valueDecimal.GreaterThan(_high) || valueDecimal.LessThan(_low) {
			c.SetPrev(valueDecimal)

			newHigh := _prev.Mul(decimal.NewFromInt(100).Add(margin)).Div(decimal.NewFromInt(100))
			newLow := _prev.Mul(decimal.NewFromInt(100).Sub(margin)).Div(decimal.NewFromInt(100))
			return fmt.Sprintf(
				"V:%.4f | Pre: %.4f | L:%.4f | H:%.4f",
				utils.DivDecimals(valueDecimal, int32(c.Precision)).InexactFloat64(),
				utils.DivDecimals(_prev, int32(c.Precision)).InexactFloat64(),
				utils.DivDecimals(newLow, int32(c.Precision)).InexactFloat64(),
				utils.DivDecimals(newHigh, int32(c.Precision)).InexactFloat64(),
			), nil
		}
		return "", common.NewErrorWithSeverity(common.Info, fmt.Sprintf("no report for result [%v]", value))
	} else {
		return "", common.NewErrorWithSeverity(common.Error, fmt.Sprintf("cannot parse value [%v]", values...))
	}
}
