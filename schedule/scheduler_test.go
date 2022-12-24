package schedule

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
)

func TestNewScheduler(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		logger := log.New(os.Stdout, "", log.LstdFlags)
		scheduler := NewScheduler().SetInterval(1 * time.Second).SetLogger(logger)

		s1 := NewSchedule("test").Every(3 * time.Second)
		s2 := NewSchedule("test2").Every(5 * time.Second)
		scheduler.AddSchedule(s1)
		scheduler.AddSchedule(s2)

		ns := scheduler.GetNextSchedule()
		if ns.Name != "test" {
			t.Errorf("GetNextSchedule() = %v, want %v", ns.Name, "test")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		scheduler.Run(ctx, func(schedule *Schedule) {
			t.Logf("Running schedule: %v", schedule)
		})

		if s1.RunCount() != 4 {
			t.Errorf("s1.Runs() = %v, want %v", s1.RunCount(), 4)
		}
		if s2.RunCount() != 2 {
			t.Errorf("s2.Runs() = %v, want %v", s2.RunCount(), 2)
		}

	})

}
