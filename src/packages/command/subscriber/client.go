package subscriber

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Subscriber struct {
	WssClient    *ethclient.Client
	Logs         chan types.Log
	Subscription ethereum.Subscription
}

var _sub *Subscriber

func GetSubscriber(wss string, contract common.Address, method string) (*Subscriber, error) {
	if _sub != nil {
		return _sub, nil
	}
	_client, err := ethclient.Dial(wss)
	if err != nil {
		return nil, err
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(rpc.LatestBlockNumber.Int64()),
		ToBlock:   nil,
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{common.HexToHash(method)}},
	}
	_sub = &Subscriber{
		WssClient: _client,
		Logs:      make(chan types.Log),
	}

	subscription, err := _sub.WssClient.SubscribeFilterLogs(context.TODO(), query, _sub.Logs)
	if err != nil {
		return nil, err
	}
	_sub.Subscription = subscription

	return _sub, nil
}

func (s *Subscriber) Listen() {
	errChan := s.Subscription.Err()
	defer s.Subscription.Unsubscribe()
	defer close(s.Logs)

	// fmt.Printf("Listen to %+v\n", s.Subscription)
	for {
		select {
		case err := <-errChan:
			fmt.Println("Error", err)
			return
		case l := <-s.Logs:
			fmt.Println("New log", l)
		}
	}
}
