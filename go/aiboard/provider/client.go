package provider

import (
	"context"
	"net/http"

	"aiboard/go/aiboard/model"
)

type Client interface {
	FetchSnapshot(ctx context.Context, cfg model.AccountConfig, apiKey string) (model.UsageSnapshot, error)
	Test(ctx context.Context, apiKey string) error
}

func httpClientOrDefault(client *http.Client) *http.Client {
	if client != nil {
		return client
	}
	return http.DefaultClient
}
