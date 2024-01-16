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
	fromAddress, err := AddressFromPriKey(privateKeyECDSA)
	if err != nil {
		return nil, err
	}

	nonce, err := ec.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	chainID, err := ec.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chainId: %w", err)
	}
	gasPrice, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gasPrice: %w", err)
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
