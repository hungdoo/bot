package balance

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/shopspring/decimal"
)

func TestBalance(t *testing.T) {
	rpc := "https://mainnet.gateway.tenderly.co"

	cmd := &BalanceCommand{
		Command: command.Command{
			Name:     "test",
			Enabled:  true,
			Type:     command.Balance,
			IdleTime: time.Second * 30,
		},
		Rpc: rpc,
		Wallets: []Wallet{
			{
				Address:    common.HexToAddress("0x3b245f18149c1cb082cc3ddf05a483e8a79ae677"),
				PreBalance: decimal.Zero,
			},
		},
	}

	res, err := cmd.Execute(true, "")

	t.Log(res)
	t.Log(err)
}
