package tombplus

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

var _contract *Tombplus

func GetTombplusContract(ec *ethclient.Client, contractAdr common.Address) (*Tombplus, error) {
	if _contract != nil {
		return _contract, nil
	}
	_contract, err := NewTombplus(contractAdr, ec)
	if err != nil {
		return nil, err
	}
	return _contract, nil
}

func NewAuthorizedTransactor(ec *ethclient.Client, privateKeyECDSA *ecdsa.PrivateKey, gaslimit uint64, value *big.Int) (*bind.TransactOpts, error) {
	ctx := context.Background()
	fromAddress, err := AddressFromPriKey(privateKeyECDSA)
	if err != nil {
		return nil, err
	}

	nonce, err := ec.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	chainID, err := ec.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chainId: %w", err)
	}
	gasPrice, err := ec.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get gasPrice: %w", err)
	}
	multitude := new(big.Int).Div(gasPrice, big.NewInt(100*params.GWei))
	if multitude.Int64() > 1 { // ftm gas > 100 Gwei
		gasPrice = big.NewInt(80 * params.GWei)
	}
	signerFn := func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(transaction, types.LatestSignerForChainID(chainID), privateKeyECDSA)
	}
	opts := &bind.TransactOpts{
		From:     fromAddress,
		Signer:   signerFn,
		GasPrice: gasPrice,
		GasLimit: gaslimit, // 0 = estimate
		Value:    value,
		Nonce:    big.NewInt(int64(nonce)),
		NoSend:   true,
	}

	return opts, nil
}
