package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"aiboard/go/aiboard/model"
)

func TestDeepSeekParsesBalanceResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Fatalf("authorization header = %q", r.Header.Get("Authorization"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"is_available":true,"balance_infos":[{"currency":"CNY","total_balance":"12.50","granted_balance":"2.50","topped_up_balance":"10.00"}]}`))
	}))
	defer server.Close()

	client := NewDeepSeekClient(server.URL, server.Client())
	snapshot, err := client.FetchSnapshot(context.Background(), model.AccountConfig{ID: "deepseek-1", Provider: model.ProviderDeepSeek, Remark: "personal"}, "test-key")
	if err != nil {
		t.Fatalf("FetchSnapshot error = %v", err)
	}

	if snapshot.Status != model.StatusOK {
		t.Fatalf("status = %s, want ok", snapshot.Status)
	}
	if snapshot.Balance == nil || snapshot.Balance.Total != "12.50" || snapshot.Balance.Currency != "CNY" {
		t.Fatalf("balance = %+v", snapshot.Balance)
	}
}

func TestMiniMaxLocalFallbackUsesConfiguredWindows(t *testing.T) {
	now := time.Date(2026, 6, 11, 13, 20, 0, 0, time.FixedZone("CST", 8*60*60))
	client := NewMiniMaxLocalClient(func() time.Time { return now })

	snapshot, err := client.FetchSnapshot(context.Background(), model.AccountConfig{ID: "mini-1", Provider: model.ProviderMiniMax, Remark: "team"}, "ignored")
	if err != nil {
		t.Fatalf("FetchSnapshot error = %v", err)
	}

	if snapshot.Usage5h == nil || snapshot.UsageWeek == nil {
		t.Fatalf("expected both usage windows: %+v", snapshot)
	}
	if snapshot.ResetAt5h != now.Location().String() && snapshot.Usage5h.ResetAt.IsZero() {
		t.Fatalf("expected five-hour reset to be populated")
	}
	if snapshot.Status != model.StatusLocalOnly {
		t.Fatalf("status = %s, want local_only", snapshot.Status)
	}
}

func TestOpenAIAggregatesUsageBuckets(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/organization/usage/completions" {
			t.Fatalf("path = %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":[{"results":[{"input_tokens":10,"output_tokens":15,"num_model_requests":2}]},{"results":[{"input_tokens":7,"output_tokens":8,"num_model_requests":1}]}]}`))
	}))
	defer server.Close()

	client := NewOpenAIClient(server.URL, server.Client(), func() time.Time {
		return time.Date(2026, 6, 11, 13, 20, 0, 0, time.UTC)
	})
	snapshot, err := client.FetchSnapshot(context.Background(), model.AccountConfig{ID: "openai-1", Provider: model.ProviderOpenAI, Remark: "org"}, "admin-key")
	if err != nil {
		t.Fatalf("FetchSnapshot error = %v", err)
	}

	if snapshot.Usage5h == nil {
		t.Fatalf("expected usage window")
	}
	if snapshot.Usage5h.InputTokens != 17 || snapshot.Usage5h.OutputTokens != 23 || snapshot.Usage5h.Requests != 3 {
		t.Fatalf("usage = %+v", snapshot.Usage5h)
	}
}
