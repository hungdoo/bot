package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hungdoo/bot/src/packages/graphql"
)

func TestGraphQlClient(it *testing.T) {
	it.Run("Query Venus subgraph", func(t *testing.T) {
		client := graphql.GetClient()
		err := client.Query(context.Background(), &graphql.MarketQuery, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(graphql.MarketQuery)
	})
}
