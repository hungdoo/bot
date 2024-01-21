package tombplus

import (
	"crypto/ecdsa"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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
	contractAddress := common.HexToAddress("0x45b7304fe1542b440b4c61ea04958cefbcec1089") // MasonryPlus: broadcast/TombPlusDeployment.s.sol/250/run-latest.json

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

func TestFlipmultiple(t *testing.T) {
	result, err := tombplusCli.Flipmultiple(pk, 8, true)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf(result) // forge debug <txhash>
}

func TestClaim(t *testing.T) {
	result, err := tombplusCli.Claim(pk)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf(result) // forge debug <txhash>
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

}
