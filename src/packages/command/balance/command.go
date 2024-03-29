package balance

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hungdoo/bot/src/common"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/log"
	"github.com/hungdoo/bot/src/packages/utils"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
)

type Wallet struct {
	Address    ethCommon.Address `bson:"address"`
	PreBalance decimal.Decimal   `bson:"prebalance"`
}

func (w *Wallet) MarshalBSON() ([]byte, error) {
	return bson.Marshal(&struct {
		Address    string
		PreBalance string
	}{
		Address:    w.Address.String(),
		PreBalance: w.PreBalance.String(),
	})
}
func (w *Wallet) UnmarshalBSON(data []byte) error {
	var d bson.D
	err := bson.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	for _, v := range d {
		if v.Key == "address" {
			w.Address = ethCommon.HexToAddress(v.Value.(string))
		} else if v.Key == "prebalance" {
			w.PreBalance, err = decimal.NewFromString(v.Value.(string))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type BalanceCommand struct {
	command.Command
	Id      string   `bson:"_id,unique"`
	Rpc     string   `bson:"rpc"`
	Wallets []Wallet `bson:"wallets"`
}

func (c BalanceCommand) MarshalJSON() ([]byte, error) {
	addresses := make([]string, len(c.Wallets))
	for i, w := range c.Wallets {
		addresses[i] = w.Address.String()
	}
	return json.Marshal(&struct {
		Name string `json:"name"`
		Type string `json:"type"`
		// Data     []string `json:"data"`
		IdleTime string   `json:"idletime"`
		Rpc      string   `json:"rpc"`
		Wallets  []Wallet `json:"wallets"`
		Command  string   `json:"command"`
	}{
		Name:     c.Name,
		Type:     c.Type.String(),
		IdleTime: c.IdleTime.String(),
		Rpc:      c.Rpc,
		Wallets:  c.Wallets,
		Command:  fmt.Sprintf("add balance %s %s %s", c.Name, c.Rpc, strings.Join(addresses, " ")),
	})
}

func (c *BalanceCommand) Validate(data []string) error {
	if len(data) < 2 {
		return fmt.Errorf("invalid params: rpc, ...wallets")
	}
	return nil
}

func (c *BalanceCommand) SetData(newValue []string) (err error) {
	if err = c.Validate(newValue); err != nil {
		return err
	}
	// c.Data = newValue
	c.Rpc = newValue[0]
	for _, v := range newValue[1:] {
		c.Wallets = append(
			c.Wallets,
			Wallet{
				Address:    ethCommon.HexToAddress(v),
				PreBalance: decimal.Zero,
			},
		)
	}
	return nil
}

func (c *BalanceCommand) Execute(mustReport bool, subcommand string) (string, *common.ErrorWithSeverity) {
	rpcCli, err := rpc.Dial(c.Rpc)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	// Create a batch of RPC calls to get the balance of each address
	batch := make([]rpc.BatchElem, len(c.Wallets))
	for i, wallet := range c.Wallets {
		batch[i] = rpc.BatchElem{
			Method: "eth_getBalance",
			Args: []interface{}{
				wallet.Address,
				"latest",
			},
			Result: new(string),
		}
	}

	// Execute the batch call
	err = rpcCli.BatchCall(batch)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	// Process the results
	results := []string{}
	var errors []string
	for i, call := range batch {
		wallet := c.Wallets[i]
		if call.Error != nil {
			errors = append(errors, fmt.Sprintf("address %s: err-%v", wallet.Address, call.Error))
			continue
		}

		// Assert the type to string
		resultString, ok := call.Result.(*string)
		if !ok {
			errors = append(errors, fmt.Sprintf("unexpected result type for address %s", wallet.Address))
			continue
		}

		decimalValue, success := new(big.Int).SetString(*resultString, 0)
		if !success {
			errors = append(errors, fmt.Sprintf("error converting balance for address %s: %v", wallet.Address, resultString))
			continue
		}

		balance := decimal.NewFromBigInt(decimalValue, 0)
		log.GeneralLogger.Printf("%s: %.5f", wallet.Address, utils.DivDecimals(balance, 18).InexactFloat64())
		if mustReport || !mustReport && !balance.Equal(wallet.PreBalance) {
			c.Wallets[i].PreBalance = balance
			results = append(results, fmt.Sprintf("%s\n=> [P:%.5f | V:%.5f]\n", wallet.Address,
				utils.DivDecimals(wallet.PreBalance, 18).InexactFloat64(),
				utils.DivDecimals(balance, 18).InexactFloat64()))
		}
	}

	if len(errors) != 0 {
		log.GeneralLogger.Printf("%s", strings.Join(errors, "\n"))
		c.SetError(strings.Join(errors, "\n"))
	}
	return strings.Join(results, ""), nil
}
