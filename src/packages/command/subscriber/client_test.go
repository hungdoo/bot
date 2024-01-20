package subscriber

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestSubscriber(t *testing.T) {
	wss := "wss://ethereum.publicnode.com"
	contractAddress := common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD") // uniswap universal router
	// hash := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
	methodSelector := common.HexToHash("0x3593564c").Bytes()[:4] // execute(bytes commands,bytes[] inputs,uint256 deadline)

	s, err := GetSubscriber(wss, contractAddress, string(methodSelector))
	if err != nil {
		t.Fatal(err)
	}
	go s.Listen()
	s.Logs <- types.Log{Address: common.MaxAddress}
}
