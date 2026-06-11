package model

import "time"

type Provider string

const (
	ProviderOpenAI   Provider = "openai"
	ProviderMiniMax  Provider = "minimax"
	ProviderDeepSeek Provider = "deepseek"
)

type Status string

const (
	StatusOK        Status = "ok"
	StatusLocalOnly Status = "local_only"
	StatusError     Status = "error"
)

type AccountConfig struct {
	ID       string   `json:"id"`
	Provider Provider `json:"provider"`
	Remark   string   `json:"remark"`
	Enabled  bool     `json:"enabled"`
}

type UsageWindow struct {
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	ResetAt      time.Time `json:"reset_at"`
	InputTokens  int64     `json:"input_tokens"`
	OutputTokens int64     `json:"output_tokens"`
	Requests     int64     `json:"requests"`
	QuotaLabel   string    `json:"quota_label,omitempty"`
}

type BalanceInfo struct {
	Currency        string `json:"currency"`
	Total           string `json:"total"`
	GrantedBalance  string `json:"granted_balance"`
	ToppedUpBalance string `json:"topped_up_balance"`
	Available       bool   `json:"available"`
}

type CostInfo struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type UsageSnapshot struct {
	AccountID     string       `json:"account_id"`
	Provider      Provider     `json:"provider"`
	Remark        string       `json:"remark"`
	Status        Status       `json:"status"`
	Balance       *BalanceInfo `json:"balance,omitempty"`
	Usage5h       *UsageWindow `json:"usage_5h,omitempty"`
	UsageWeek     *UsageWindow `json:"usage_week,omitempty"`
	Cost          *CostInfo    `json:"cost,omitempty"`
	ResetAt5h     string       `json:"reset_at_5h,omitempty"`
	ResetAtWeek   string       `json:"reset_at_week,omitempty"`
	LastUpdatedAt string       `json:"last_updated_at"`
	ErrorMessage  string       `json:"error_message,omitempty"`
}
