package provider

import (
	"time"

	"aiboard/go/aiboard/model"
	"aiboard/go/aiboard/quota"
)

func baseSnapshot(cfg model.AccountConfig, status model.Status, now time.Time) model.UsageSnapshot {
	return model.UsageSnapshot{
		AccountID:     cfg.ID,
		Provider:      cfg.Provider,
		Remark:        cfg.Remark,
		Status:        status,
		LastUpdatedAt: now.Format(time.RFC3339),
	}
}

func errorSnapshot(cfg model.AccountConfig, err error) model.UsageSnapshot {
	snapshot := baseSnapshot(cfg, model.StatusError, time.Now())
	snapshot.ErrorMessage = err.Error()
	return snapshot
}

func usageFromWindow(window quota.Window, label string) *model.UsageWindow {
	return &model.UsageWindow{
		Start:      window.Start,
		End:        window.End,
		ResetAt:    window.ResetAt,
		QuotaLabel: label,
	}
}
