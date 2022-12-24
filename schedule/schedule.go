package schedule

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Schedule struct {
	Name          string
	Data          map[string]any
	runsEvery     time.Duration
	runsAfter     time.Time
	Enabled       bool
	LastRun       time.Time
	dontRunBefore time.Duration
	dontRunAfter  time.Duration
	runCount      uint64
}

func NewSchedule(name string) *Schedule {
	return &Schedule{
		Name:      name,
		Data:      make(map[string]any),
		runsEvery: time.Duration(0),
		runsAfter: time.Time{},
		Enabled:   true,
		LastRun:   time.Time{},
	}
}

func (s *Schedule) Every(duration time.Duration) *Schedule {
	s.runsEvery = duration
	return s
}

func (s *Schedule) After(t time.Time) *Schedule {
	s.runsAfter = t
	return s
}

func (s *Schedule) DontRunBefore(duration time.Duration) *Schedule {
	s.dontRunBefore = duration
	return s
}

func (s *Schedule) DontRunAfter(duration time.Duration) *Schedule {
	s.dontRunAfter = duration
	return s
}

func (s *Schedule) NextRun() time.Time {
	if !s.Enabled {
		return time.Time{}
	}
	if s.runsEvery > 0 {
		return s.getNextRunWindow(s.LastRun.Add(s.runsEvery))
	}
	if s.runsAfter.After(time.Now()) {
		return s.getNextRunWindow(s.runsAfter)
	}
	return time.Time{}
}

func (s *Schedule) getNextRunWindow(nextRun time.Time) time.Time {
	if s.dontRunBefore > 0 {
		dontRunBefore := Today().Add(s.dontRunBefore)
		if time.Now().Before(dontRunBefore) {
			return Tomorrow().Add(s.dontRunBefore)
		}
	}

	if s.dontRunAfter > 0 {
		dontRunAfter := Today().Add(s.dontRunAfter)
		if time.Now().After(dontRunAfter) {
			return Tomorrow().Add(s.dontRunAfter)
		}
	}

	return nextRun
}

func (s *Schedule) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Schedule('%s'", s.Name))
	if s.runsEvery > 0 {
		sb.WriteString(fmt.Sprintf(", runs every %v", s.runsEvery))
	}
	if !s.runsAfter.IsZero() {
		sb.WriteString(fmt.Sprintf(", runs after %v", s.runsAfter))
	}
	if !s.Enabled {
		sb.WriteString(", disabled")
	}
	if !s.LastRun.IsZero() {
		sb.WriteString(fmt.Sprintf(", last run %v", s.LastRun))
	}
	if !s.NextRun().IsZero() {
		sb.WriteString(fmt.Sprintf(", next run %v", s.NextRun()))
	}
	sb.WriteString(")")
	return sb.String()
}

func (s *Schedule) RunCount() uint64 {
	return s.runCount
}

func (s *Schedule) Run(callback ScheduleCallBack, logger *log.Logger) {
	callback(s)
	s.LastRun = time.Now()
	s.runCount++
	if logger!=nil {
		logger.Printf("Running schedule: %v", s)
	}
}
