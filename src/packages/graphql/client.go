package graphql

import (
	"context"

	"github.com/hungdoo/bot/src/packages/dotenv"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

var client *Client

type Client struct {
	*graphql.Client
}

func GetClient() *Client {
	if client != nil {
		return client
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: dotenv.BotEnvs["BOT_GRAPHQL_TOKEN"]},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client = &Client{graphql.NewClient(dotenv.BotEnvs["BOT_VENUS_SUBGRAPH"], httpClient)}
	return client
}
