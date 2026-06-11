package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"aiboard/go/aiboard/model"
	"aiboard/go/aiboard/quota"
)

type OpenAIClient struct {
	baseURL string
	http    *http.Client
	now     func() time.Time
}

func NewOpenAIClient(baseURL string, client *http.Client, now func() time.Time) *OpenAIClient {
	if now == nil {
		now = time.Now
	}
	return &OpenAIClient{baseURL: strings.TrimRight(baseURL, "/"), http: httpClientOrDefault(client), now: now}
}

func (c *OpenAIClient) Test(ctx context.Context, apiKey string) error {
	_, err := c.FetchSnapshot(ctx, model.AccountConfig{Provider: model.ProviderOpenAI}, apiKey)
	return err
}

func (c *OpenAIClient) FetchSnapshot(ctx context.Context, cfg model.AccountConfig, apiKey string) (model.UsageSnapshot, error) {
	now := c.now()
	five := quota.FixedFiveHourWindow(now)
	week := quota.WeekWindow(now, time.Monday)

	endpoint, err := url.Parse(c.baseURL + "/organization/usage/completions")
	if err != nil {
		return errorSnapshot(cfg, err), err
	}
	q := endpoint.Query()
	q.Set("start_time", strconv.FormatInt(five.Start.Unix(), 10))
	q.Set("end_time", strconv.FormatInt(five.End.Unix(), 10))
	q.Set("bucket_width", "1h")
	endpoint.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
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
		err := fmt.Errorf("openai usage returned HTTP %d", resp.StatusCode)
		return errorSnapshot(cfg, err), err
	}

	var parsed openAIUsageResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return errorSnapshot(cfg, err), err
	}

	snapshot := baseSnapshot(cfg, model.StatusOK, now)
	snapshot.Usage5h = usageFromWindow(five, "OpenAI five-hour usage")
	snapshot.UsageWeek = usageFromWindow(week, "OpenAI weekly reset")
	snapshot.ResetAt5h = five.ResetAt.Format(time.RFC3339)
	snapshot.ResetAtWeek = week.ResetAt.Format(time.RFC3339)

	for _, bucket := range parsed.Data {
		for _, result := range bucket.Results {
			snapshot.Usage5h.InputTokens += result.InputTokens
			snapshot.Usage5h.OutputTokens += result.OutputTokens
			snapshot.Usage5h.Requests += result.Requests
		}
	}

	return snapshot, nil
}

type openAIUsageResponse struct {
	Data []openAIUsageBucket `json:"data"`
}

type openAIUsageBucket struct {
	Results []openAIUsageResult `json:"results"`
}

type openAIUsageResult struct {
	InputTokens  int64 `json:"input_tokens"`
	OutputTokens int64 `json:"output_tokens"`
	Requests     int64 `json:"num_model_requests"`
}
