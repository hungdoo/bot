package graphql

import (
	"context"
	"fmt"
	"testing"
)

func TestGraphQlClient(it *testing.T) {
	it.Run("Query Venus subgraph", func(t *testing.T) {
		client := GetClient()
		err := client.Query(context.Background(), &MarketQuery, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", MarketQuery)
	})
}
