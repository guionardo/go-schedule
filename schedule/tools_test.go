package schedule

import (
	"testing"
	"time"
)

func TestToday(t *testing.T) {

	t.Run("Default", func(t *testing.T) {
		got := Today()
		if got.Hour() != 0 || got.Minute() != 0 || got.Second() != 0 || got.Nanosecond() != 0 {
			t.Errorf("Today() = %v, want %v", got, time.Time{})
		}
	})

}

func TestWallClock(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		if got := WallClock(); got.Seconds() == 0 {
			t.Errorf("WallClock() = %v, want positive", got)
		}
	})

}

func TestTomorrow(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		if got := Tomorrow(); got.Day() == time.Now().Day() {
			t.Errorf("Tomorrow() = %v", got)
		}
	})

}
