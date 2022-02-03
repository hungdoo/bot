package contract

import (
	"time"

	"github.com/alethio/web3-go/ethrpc"
	"github.com/alethio/web3-go/ethrpc/provider/httprpc"
)

func GetETH(url string) (ethrpc.ETHInterface, error) {
	provider, err := httprpc.New(url)
	if err != nil {
		return nil, err
	}
	provider.SetHTTPTimeout(5 * time.Second)
	return ethrpc.New(provider)
}
