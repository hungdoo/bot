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

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type TombplusClient struct {
	ec   *ethclient.Client
	tomb *Tombplus
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
		tomb: tomeplusContract,
	}
	return _client, nil
}

func (c *TombplusClient) MaxAllowedFutureFlips() int64 {
	val, err := c.tomb.MaxAllowedFutureFlips(&bind.CallOpts{})
	if err != nil {
		return -1
	}
	return val.Int64()
}

func (c *TombplusClient) GameStarted() bool {
	val, err := c.tomb.GameStarted(&bind.CallOpts{})
	if err != nil {
		return false
	}
	return val
}

func (c *TombplusClient) CurrentEpoch() int64 {
	epochNum, err := c.tomb.EpochNumber(&bind.CallOpts{})
	if err != nil {
		return -1
	}
	return epochNum.Int64()
}

func (c *TombplusClient) IsVotedAtEpoch(user ethCommon.Address, epoch int64) (bool, error) {
	val, err := c.tomb.GetUserFlipIdByEpochId(&bind.CallOpts{}, user, big.NewInt(epoch))
	if err != nil {
		return false, err
	}
	return val.Found, nil
}

func (c *TombplusClient) Claim(privateKey *ecdsa.PrivateKey) (string, *common.ErrorWithSeverity) {
	noSendOpts, err := NewAuthorizedTransactor(c.ec, privateKey, 0, big.NewInt(0))
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	signedTx, err := c.tomb.Claim(noSendOpts)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	return c.dryrunAndSend(noSendOpts.From, signedTx)
}

func (c *TombplusClient) Flipmultiple(privateKey *ecdsa.PrivateKey, fromEpoch int64, epochs int64, up bool) (string, *common.ErrorWithSeverity) {
	noSendOpts, err := NewAuthorizedTransactor(c.ec, privateKey, 0, big.NewInt(0))
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	epochIds := make([]*big.Int, epochs)
	ups := make([]bool, epochs)
	for i := int64(0); i < epochs; i++ {
		epochIds[i] = big.NewInt(fromEpoch + i)
		ups[i] = up
	}

	signedTx, err := c.tomb.FlipMultiple(noSendOpts, epochIds, ups)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}
	return c.dryrunAndSend(noSendOpts.From, signedTx)
}

func (c *TombplusClient) checkResult(tx *types.Transaction) (string, *common.ErrorWithSeverity) {
	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), c.ec, tx)
	if err != nil {
		// rpc error
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	if receipt.Status == 0 {
		reasons := []string{}
		for _, l := range receipt.Logs {
			reasons = append(reasons, string(l.Data))
		}

		// onchain tx reverted, raise to Critical
		return "", common.NewErrorWithSeverity(common.Critical, fmt.Sprintf("transaction %s was reverted with reason: %s", receipt.TxHash.Hex(), strings.Join(reasons, "\n")))
	}

	return receipt.TxHash.Hex(), nil
}

func (c *TombplusClient) dryrunAndSend(fromAddress ethCommon.Address, signedTx *types.Transaction) (string, *common.ErrorWithSeverity) {
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
		// tx will revert, no need to send
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	err = c.ec.SendTransaction(context.Background(), signedTx)
	if err != nil {
		// rpc error
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}
	return c.checkResult(signedTx)
}
