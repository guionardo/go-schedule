package schedule

import (
	"testing"
	"time"
)

func TestNewSchedule_Empty_ShouldHaveNoNext(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		s := NewSchedule("test")
		if s.NextRun().IsZero() == false {
			t.Errorf("NewSchedule() = %v, want %v", s.NextRun(), time.Time{})
		}
	})
}