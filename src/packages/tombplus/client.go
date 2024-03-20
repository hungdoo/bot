package tombplus

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hungdoo/bot/src/common"
	"github.com/hungdoo/bot/src/packages/utils"
	"github.com/shopspring/decimal"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type TombplusClient struct {
	ec   *ethclient.Client
	Tomb *Tombplus
}

var _client *TombplusClient

func GetClient(rpcEndpoint string, contractAddress ethCommon.Address) (*TombplusClient, error) {
	if _client != nil {
		return _client, nil
	}
	rpcclient, err := rpc.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	ec := ethclient.NewClient(rpcclient)
	tomeplusContract, err := GetTombplusContract(ec, contractAddress)
	if err != nil {
		return nil, err
	}
	_client = &TombplusClient{
		ec:   ec,
		Tomb: tomeplusContract,
	}
	return _client, nil
}

func (c *TombplusClient) MaxAllowedFutureFlips() int64 {
	val, err := c.Tomb.MaxAllowedFutureFlips(&bind.CallOpts{})
	if err != nil {
		return -1
	}
	return val.Int64()
}

func (c *TombplusClient) GameStarted() bool {
	val, err := c.Tomb.GameStarted(&bind.CallOpts{})
	if err != nil {
		return false
	}
	return val
}

func (c *TombplusClient) GetRewards(user ethCommon.Address) decimal.Decimal {
	reward, err := c.Tomb.RewardBalance(&bind.CallOpts{}, user)
	if err != nil {
		return decimal.Zero
	}
	return utils.DivDecimals(decimal.NewFromBigInt(reward, 0), 18)
}

func (c *TombplusClient) GetUserLastedVoteEpochId(user ethCommon.Address) (*big.Int, error) {
	flips, err := c.Tomb.GetUserFlips(&bind.CallOpts{}, user)
	if err != nil {
		return nil, err
	}
	latestEpoch := big.NewInt(-1)
	for _, v := range flips {
		if v.EpochId.Cmp(latestEpoch) >= 1 {
			latestEpoch = v.EpochId
		}
	}

	return latestEpoch, nil
}

func (c *TombplusClient) CanFlipForCurrentEpoch() (bool, error) {
	ok, err := c.Tomb.CanFlipForCurrentEpoch(&bind.CallOpts{})
	if err != nil {
		return false, err
	}
	return ok, nil
}
func (c *TombplusClient) CurrentEpoch() int64 {
	epochNum, err := c.Tomb.CurrentEpochId(&bind.CallOpts{})
	if err != nil {
		return -1
	}
	return epochNum.Int64()
}

func (c *TombplusClient) GetPauseGameAtEpoch() int64 {
	val, err := c.Tomb.PauseGameAtEpoch(&bind.CallOpts{})
	if err != nil {
		return 0
	}
	return val.Int64()
}

func (c *TombplusClient) GetUpgradedMasonry() ethCommon.Address {
	val, err := c.Tomb.UpgradedMasonry(&bind.CallOpts{})
	if err != nil {
		return ethCommon.Address{}
	}
	return val
}

func (c *TombplusClient) IsVotedAtEpoch(user ethCommon.Address, epoch int64) (bool, error) {
	val, err := c.Tomb.GetUserFlipIdByEpochId(&bind.CallOpts{}, user, big.NewInt(epoch))
	if err != nil {
		return false, err
	}
	return val.Found, nil
}

func (c *TombplusClient) Claim(privateKey *ecdsa.PrivateKey, maxGas *big.Int) (*types.Transaction, *common.ErrorWithSeverity) {
	noSendOpts, err := NewAuthorizedTransactor(c.ec, privateKey, 0, maxGas, big.NewInt(0))
	if err != nil {
		return nil, common.NewErrorWithSeverity(common.Error, err.Error())
	}

	signedTx, err := c.Tomb.Claim(noSendOpts)
	if err != nil {
		return nil, common.NewErrorWithSeverity(common.Error, err.Error())
	}

	return c.dryrunAndSend(noSendOpts.From, signedTx)
}

func (c *TombplusClient) Flip(privateKey *ecdsa.PrivateKey, maxGas *big.Int, up bool) (*types.Transaction, *common.ErrorWithSeverity) {
	noSendOpts, err := NewAuthorizedTransactor(c.ec, privateKey, 0, maxGas, big.NewInt(0))
	if err != nil {
		return nil, common.NewErrorWithSeverity(common.Critical, err.Error())
	}
	signedTx, err := c.Tomb.Flip(noSendOpts, up)
	if err != nil {
		return nil, common.NewErrorWithSeverity(common.Critical, err.Error())
	}
	return c.dryrunAndSend(noSendOpts.From, signedTx)
}

func (c *TombplusClient) Flipmultiple(privateKey *ecdsa.PrivateKey, maxGas *big.Int, epochs int64, up bool) (*types.Transaction, *common.ErrorWithSeverity) {
	noSendOpts, err := NewAuthorizedTransactor(c.ec, privateKey, 0, maxGas, big.NewInt(0))
	if err != nil {
		return nil, common.NewErrorWithSeverity(common.Error, err.Error())
	}

	ups := make([]bool, epochs)
	for i := int64(0); i < epochs; i++ {
		ups[i] = up
	}

	signedTx, err := c.Tomb.FlipMultiple(noSendOpts, ups)
	if err != nil {
		return nil, common.NewErrorWithSeverity(common.Error, err.Error())
	}
	return c.dryrunAndSend(noSendOpts.From, signedTx)
}

func (c *TombplusClient) CheckResult(txHash string) *common.ErrorWithSeverity {
	// Wait for the transaction to be mined
	// receipt, err := bind.WaitMined(context.Background(), c.ec, tx)
	receipt, err := c.ec.TransactionReceipt(context.Background(), ethCommon.HexToHash(txHash))
	if err != nil {
		// rpc error
		return common.NewErrorWithSeverity(common.Error, err.Error())
	}

	if receipt.Status == 0 {
		reasons := []string{}
		for _, l := range receipt.Logs {
			reasons = append(reasons, string(l.Data))
		}

		// onchain tx reverted, raise to Critical
		return common.NewErrorWithSeverity(common.Critical, fmt.Sprintf("transaction %s was reverted with reason: %s", receipt.TxHash.Hex(), strings.Join(reasons, "\n")))
	}

	return nil
}

func (c *TombplusClient) dryrunAndSend(fromAddress ethCommon.Address, signedTx *types.Transaction) (*types.Transaction, *common.ErrorWithSeverity) {
	// Dryrun first to save gas in case of revert
	_, err := c.ec.CallContract(context.Background(), ethereum.CallMsg{
		To:       signedTx.To(),
		From:     fromAddress,
		Gas:      signedTx.Gas(),
		GasPrice: signedTx.GasPrice(),
		Value:    big.NewInt(0),
		Data:     signedTx.Data(),
	}, nil)
	if err != nil {
		return nil, common.NewErrorWithSeverity(common.Critical, err.Error())
	}

	err = c.ec.SendTransaction(context.Background(), signedTx)
	if err != nil {
		// rpc error
		return nil, common.NewErrorWithSeverity(common.Critical, err.Error())
	}
	return signedTx, nil
}
