package balance

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/log"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	Address    common.Address
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
				Address:    common.HexToAddress(v),
				PreBalance: decimal.Zero,
			},
		)
	}
	return nil
}

func (c *BalanceCommand) Execute(mustReport bool, subcommand string) (string, error) {
	rpcCli, err := rpc.Dial(c.Rpc)
	if err != nil {
		return "", err
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
		return "", err
	}

	// Process the results
	results := []string{c.GetName()}
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
			errors = append(errors, fmt.Sprintf("Unexpected result type for address %s", wallet.Address))
			continue
		}

		decimalValue, success := new(big.Int).SetString(*resultString, 0)
		if !success {
			errors = append(errors, fmt.Sprintf("Error converting balance for address %s: %v", wallet.Address, resultString))
			continue
		}

		balance := decimal.NewFromBigInt(decimalValue, 0)
		log.GeneralLogger.Printf("%s: %.5f", wallet.Address, balance.Div(decimal.NewFromBigInt(common.Big1, 18)).InexactFloat64())
		if mustReport || !mustReport && !balance.Equal(wallet.PreBalance) {
			c.Wallets[i].PreBalance = balance
			results = append(results, fmt.Sprintf("%s[%.5f]\n", wallet.Address, balance.Div(decimal.NewFromBigInt(common.Big1, 18)).InexactFloat64()))
		}
	}

	if len(errors) != 0 {
		log.GeneralLogger.Printf("%s", strings.Join(errors, "\n"))
		c.SetError(fmt.Errorf("%s", strings.Join(errors, "\n")))
	}
	return strings.Join(results, "\n"), nil
}
