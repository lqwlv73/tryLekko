// Generated by Lekko. DO NOT EDIT.
package lekko

import (
	"context"

	client "github.com/lekkodev/go-sdk/client"
	lekkoexample "lqwlv73/tryLekko/lekko/example"
)

type LekkoClient struct {
	Example *lekkoexample.LekkoClient
	Close   client.CloseFunc
}

// Initializes the Lekko SDK client.
// For remote configs to be fetched correctly, the LEKKO_API_KEY env variable is required.
// If the env variable is missing or if there are any connection errors, the static fallbacks will be used.
func NewLekkoClient(ctx context.Context, opts ...client.ProviderOption) *LekkoClient {
	repoOwner := "lqwlv73"
	repoName := "lekko-configs"
	cli, close := client.NewClientFromEnv(ctx, repoOwner, repoName, opts...)
	return &LekkoClient{
		Example: &lekkoexample.LekkoClient{Client: cli},
		Close:   close,
	}
}
