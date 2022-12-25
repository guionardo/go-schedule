package schedule

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

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

func ParseDuration(str string) (duration time.Duration, err error) {
	duration, err = time.ParseDuration(str)
	if err == nil {
		return
	}

	w := strings.Split(str, ":")
	if len(w) < 2 {
		err = errors.New("invalid duration")
		return
	}
	var h, m, s int
	if h, err = strconv.Atoi(w[0]); err != nil {
		return
	}
	if m, err = strconv.Atoi(w[1]); err != nil {
		return
	}
	if len(w) > 2 {
		if s, err = strconv.Atoi(w[2]); err != nil {
			return
		}
	}
	duration = time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second
	return
}
