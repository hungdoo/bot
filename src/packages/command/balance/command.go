package balance

import (
	"fmt"
	"math/big"
	"strings"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hungdoo/bot/src/common"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/log"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	Address    ethCommon.Address
	PreBalance decimal.Decimal
}
type BalanceCommand struct {
	command.Command
	Rpc     string
	Wallets []Wallet
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
	c.Data = newValue
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
		log.GeneralLogger.Printf("%s: %.5f", wallet.Address, balance.Div(decimal.NewFromBigInt(ethCommon.Big1, 18)).InexactFloat64())
		if mustReport || !mustReport && !balance.Equal(wallet.PreBalance) {
			c.Wallets[i].PreBalance = balance
			results = append(results, fmt.Sprintf("\n%s[P:%.5f|V:%.5f]", wallet.Address,
				wallet.PreBalance.Div(decimal.NewFromBigInt(ethCommon.Big1, 18)).InexactFloat64(),
				balance.Div(decimal.NewFromBigInt(ethCommon.Big1, 18)).InexactFloat64()))
		}
	}

	if len(errors) != 0 {
		log.GeneralLogger.Printf("%s", strings.Join(errors, "\n"))
		c.SetError(fmt.Errorf("%s", strings.Join(errors, "\n")))
	}
	return strings.Join(results, ""), nil
}
