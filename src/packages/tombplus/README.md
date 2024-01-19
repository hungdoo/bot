### Build go code from sol & abi

cd /Users/hungdong/go/pkg/mod/github.com/ethereum/go-ethereum@v1.13.10/
make
make devtools

abigen --abi src/packages/tombplus/abi/tombplus.json  --pkg tombplus --out src/packages/tombplus/TombPlus.go

```go
  // Instantiate the contract with the deployment address and the client
  contract, err := tombplus.NewTombPlus(common.HexToAddress("contract address"), client)
  if err != nil {
      log.Fatalf("Failed to instantiate a NewTombPlus contract: %v", err)
  }
```