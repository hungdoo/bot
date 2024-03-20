package tombplus

import (
	"crypto/ecdsa"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

var pk *ecdsa.PrivateKey
var tombplusCli *TombplusClient

func TestMain(m *testing.M) {
	// Set up your fixture before running tests
	setup()

	// Run all the tests
	exitCode := m.Run()

	// Teardown your fixture after running tests
	teardown()

	// Exit with the code from the tests
	os.Exit(exitCode)
}

func setup() {
	// Set up your fixture data
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" // get from anvil pub 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
	// rpcEndpoint := "http://localhost:8545" // anvil
	rpcEndpoint := "https://rpc.ftm.tools"
	contractAddress := common.HexToAddress("0xA979F47480b4B598bf6a8bFA73aC0B6aEccBa505") // MasonryPlus: broadcast/TombPlusDeployment.s.sol/250/run-latest.json

	var err error
	pk, err = PrivateKeyFromHex(privateKeyHex)
	if err != nil {
		panic(err)
	}

	tombplusCli, err = GetClient(rpcEndpoint, contractAddress)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	// Clean up your fixture data
}

func TestNewAuthorizedTransactor(t *testing.T) {
	opts, err := NewAuthorizedTransactor(tombplusCli.ec, pk, 0, nil, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(opts.GasPrice)

	opts, err = NewAuthorizedTransactor(tombplusCli.ec, pk, 0, new(big.Int).Mul(common.Big1, big.NewInt(params.GWei)), big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(opts.GasPrice)
}
func TestFlipmultiple(t *testing.T) {
	tx, err := tombplusCli.Flipmultiple(pk, nil, 8, true)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tx) // forge debug <txhash>
}

func TestClaim(t *testing.T) {
	tx, err := tombplusCli.Claim(pk, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tx) // forge debug <txhash>
}

func TestViewCalls_VoteData(t *testing.T) {

	data, err := tombplusCli.Tomb.UpcomingEpochData(&bind.CallOpts{}, big.NewInt(1))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("data: %+v", data)
}

func TestViewCalls(t *testing.T) {
	currentEpoch := tombplusCli.CurrentEpoch()
	t.Logf("GameStarted: %v", tombplusCli.GameStarted())
	t.Logf("currentEpoch: %v", currentEpoch)

	t.Logf("MaxAllowedFutureFlips: %v", tombplusCli.MaxAllowedFutureFlips())

	latestEpoch, err := tombplusCli.GetUserLastedVoteEpochId(common.HexToAddress("")) // kimchi,kimbap
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetUserLastedVoteEpochId: %v", latestEpoch)

	canFlip, err := tombplusCli.CanFlipForCurrentEpoch()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("canFlip: %v", canFlip)
}
