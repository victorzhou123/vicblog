package util

import "time"

const (
	timeFormatYearMonthDayHourMinSecond = "2006-01-02 15:04:05.000"
)

type timeCreator struct{}

func NewTimerCreator() timeCreator {
	return timeCreator{}
}

// AddUnix add time based on current time and return an unix format time
func (t *timeCreator) AddUnix(add time.Duration) int64 {
	return time.Now().Add(add).Unix()
}

func TimeNowBaseSecond() string {
	return timeNow(timeFormatYearMonthDayHourMinSecond)
}

func timeNow(layout string) string {
	return time.Now().Format(layout)
}
