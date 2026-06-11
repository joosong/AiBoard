package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"aiboard/go/aiboard/model"
)

type DeepSeekClient struct {
	baseURL string
	http    *http.Client
}

func NewDeepSeekClient(baseURL string, client *http.Client) *DeepSeekClient {
	return &DeepSeekClient{baseURL: strings.TrimRight(baseURL, "/"), http: httpClientOrDefault(client)}
}

func (c *DeepSeekClient) Test(ctx context.Context, apiKey string) error {
	_, err := c.FetchSnapshot(ctx, model.AccountConfig{Provider: model.ProviderDeepSeek}, apiKey)
	return err
}

func (c *DeepSeekClient) FetchSnapshot(ctx context.Context, cfg model.AccountConfig, apiKey string) (model.UsageSnapshot, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/user/balance", nil)
	if err != nil {
		return errorSnapshot(cfg, err), err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return errorSnapshot(cfg, err), err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err := fmt.Errorf("deepseek balance returned HTTP %d", resp.StatusCode)
		return errorSnapshot(cfg, err), err
	}

	var parsed deepSeekBalanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return errorSnapshot(cfg, err), err
	}

	snapshot := baseSnapshot(cfg, model.StatusOK, time.Now())
	if len(parsed.BalanceInfos) > 0 {
		first := parsed.BalanceInfos[0]
		snapshot.Balance = &model.BalanceInfo{
			Currency:        first.Currency,
			Total:           first.TotalBalance,
			GrantedBalance:  first.GrantedBalance,
			ToppedUpBalance: first.ToppedUpBalance,
			Available:       parsed.Available,
		}
	}
	return snapshot, nil
}

type deepSeekBalanceResponse struct {
	Available    bool                  `json:"is_available"`
	BalanceInfos []deepSeekBalanceInfo `json:"balance_infos"`
}

type deepSeekBalanceInfo struct {
	Currency        string `json:"currency"`
	TotalBalance    string `json:"total_balance"`
	GrantedBalance  string `json:"granted_balance"`
	ToppedUpBalance string `json:"topped_up_balance"`
}
