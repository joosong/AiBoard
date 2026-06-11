package mobile

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"aiboard/go/aiboard/model"
	"aiboard/go/aiboard/provider"
)

type Bridge struct{}

func NewBridge() *Bridge {
	return &Bridge{}
}

func (b *Bridge) RefreshSnapshot(configJSON string, apiKey string) (string, error) {
	var cfg model.AccountConfig
	if err := json.Unmarshal([]byte(configJSON), &cfg); err != nil {
		return "", err
	}

	client, err := clientForProvider(cfg.Provider)
	if err != nil {
		return "", err
	}

	snapshot, err := client.FetchSnapshot(context.Background(), cfg, apiKey)
	payload, marshalErr := json.Marshal(snapshot)
	if marshalErr != nil {
		return "", marshalErr
	}
	return string(payload), err
}

func (b *Bridge) DemoSnapshots() string {
	now := time.Now()
	snapshots := []model.UsageSnapshot{
		{
			AccountID:     "demo-openai",
			Provider:      model.ProviderOpenAI,
			Remark:        "OpenAI demo",
			Status:        model.StatusOK,
			LastUpdatedAt: now.Format(time.RFC3339),
			Usage5h:       &model.UsageWindow{InputTokens: 12800, OutputTokens: 8200, Requests: 42, ResetAt: now.Add(2 * time.Hour)},
			ResetAt5h:     now.Add(2 * time.Hour).Format(time.RFC3339),
		},
		{
			AccountID:     "demo-minimax",
			Provider:      model.ProviderMiniMax,
			Remark:        "MiniMax local",
			Status:        model.StatusLocalOnly,
			LastUpdatedAt: now.Format(time.RFC3339),
			ResetAt5h:     now.Add(time.Hour).Format(time.RFC3339),
		},
		{
			AccountID:     "demo-deepseek",
			Provider:      model.ProviderDeepSeek,
			Remark:        "DeepSeek balance",
			Status:        model.StatusOK,
			LastUpdatedAt: now.Format(time.RFC3339),
			Balance:       &model.BalanceInfo{Currency: "CNY", Total: "12.50", GrantedBalance: "2.50", ToppedUpBalance: "10.00", Available: true},
		},
	}

	payload, _ := json.Marshal(snapshots)
	return string(payload)
}

func clientForProvider(kind model.Provider) (provider.Client, error) {
	switch kind {
	case model.ProviderOpenAI:
		return provider.NewOpenAIClient("https://api.openai.com/v1", nil, time.Now), nil
	case model.ProviderMiniMax:
		return provider.NewMiniMaxLocalClient(time.Now), nil
	case model.ProviderDeepSeek:
		return provider.NewDeepSeekClient("https://api.deepseek.com", nil), nil
	default:
		return nil, errors.New("unsupported provider")
	}
}
