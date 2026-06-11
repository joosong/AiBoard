package quota

import "time"

const fiveHours = 5 * time.Hour

func FixedFiveHourWindow(now time.Time) Window {
	localMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	elapsed := now.Sub(localMidnight)
	segment := int(elapsed / fiveHours)
	start := localMidnight.Add(time.Duration(segment) * fiveHours)
	end := start.Add(fiveHours)

	return Window{Start: start, End: end, ResetAt: end}
}

func WeekWindow(now time.Time, startDay time.Weekday) Window {
	dayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	delta := (int(dayStart.Weekday()) - int(startDay) + 7) % 7
	start := dayStart.AddDate(0, 0, -delta)
	end := start.AddDate(0, 0, 7)

	return Window{Start: start, End: end, ResetAt: end}
}

type Window struct {
	Start   time.Time
	End     time.Time
	ResetAt time.Time
}
