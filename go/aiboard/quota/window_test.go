package quota

import (
	"testing"
	"time"
)

func TestFixedFiveHourWindowReturnsCurrentSegmentAndReset(t *testing.T) {
	loc := time.FixedZone("CST", 8*60*60)
	now := time.Date(2026, 6, 11, 13, 20, 0, 0, loc)

	window := FixedFiveHourWindow(now)

	if want := time.Date(2026, 6, 11, 10, 0, 0, 0, loc); !window.Start.Equal(want) {
		t.Fatalf("start = %s, want %s", window.Start, want)
	}
	if want := time.Date(2026, 6, 11, 15, 0, 0, 0, loc); !window.End.Equal(want) {
		t.Fatalf("end = %s, want %s", window.End, want)
	}
	if !window.ResetAt.Equal(window.End) {
		t.Fatalf("reset = %s, want end %s", window.ResetAt, window.End)
	}
}

func TestFixedFiveHourWindowCrossesMidnight(t *testing.T) {
	loc := time.FixedZone("CST", 8*60*60)
	now := time.Date(2026, 6, 12, 0, 30, 0, 0, loc)

	window := FixedFiveHourWindow(now)

	if want := time.Date(2026, 6, 12, 0, 0, 0, 0, loc); !window.Start.Equal(want) {
		t.Fatalf("start = %s, want %s", window.Start, want)
	}
	if want := time.Date(2026, 6, 12, 5, 0, 0, 0, loc); !window.End.Equal(want) {
		t.Fatalf("end = %s, want %s", window.End, want)
	}
}

func TestWeeklyWindowStartsMondayAndResetsNextMonday(t *testing.T) {
	loc := time.FixedZone("CST", 8*60*60)
	now := time.Date(2026, 6, 11, 13, 20, 0, 0, loc)

	window := WeekWindow(now, time.Monday)

	if want := time.Date(2026, 6, 8, 0, 0, 0, 0, loc); !window.Start.Equal(want) {
		t.Fatalf("start = %s, want %s", window.Start, want)
	}
	if want := time.Date(2026, 6, 15, 0, 0, 0, 0, loc); !window.End.Equal(want) {
		t.Fatalf("end = %s, want %s", window.End, want)
	}
}
