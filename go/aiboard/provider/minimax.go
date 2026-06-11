package provider

import (
	"context"
	"time"

	"aiboard/go/aiboard/model"
	"aiboard/go/aiboard/quota"
)

type MiniMaxLocalClient struct {
	now func() time.Time
}

func NewMiniMaxLocalClient(now func() time.Time) *MiniMaxLocalClient {
	if now == nil {
		now = time.Now
	}
	return &MiniMaxLocalClient{now: now}
}

func (c *MiniMaxLocalClient) Test(ctx context.Context, apiKey string) error {
	return nil
}

func (c *MiniMaxLocalClient) FetchSnapshot(ctx context.Context, cfg model.AccountConfig, apiKey string) (model.UsageSnapshot, error) {
	now := c.now()
	five := quota.FixedFiveHourWindow(now)
	week := quota.WeekWindow(now, time.Monday)

	snapshot := baseSnapshot(cfg, model.StatusLocalOnly, now)
	snapshot.Usage5h = usageFromWindow(five, "MiniMax local five-hour window")
	snapshot.UsageWeek = usageFromWindow(week, "MiniMax local weekly window")
	snapshot.ResetAt5h = five.ResetAt.Format(time.RFC3339)
	snapshot.ResetAtWeek = week.ResetAt.Format(time.RFC3339)
	return snapshot, nil
}
