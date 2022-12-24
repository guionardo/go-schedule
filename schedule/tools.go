package schedule

import "time"

func WallClock() time.Duration {
	now := time.Now()
	return now.Sub(Today())
}

func Today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

func Tomorrow() time.Time {
	return Today().Add(24 * time.Hour)
}
