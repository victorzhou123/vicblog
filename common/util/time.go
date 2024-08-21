package util

import "time"

const (
	timeFormatYearMonthDayHourMinSecond = "2006-01-02 15:04:05.000"
	timeFormatYearMonthDay              = "2006-01-02"
)

type timeCreator struct{}

func NewTimerCreator() timeCreator {
	return timeCreator{}
}

// AddUnix add time based on current time and return an unix format time
func (t *timeCreator) AddUnix(add time.Duration) int64 {
	return time.Now().Add(add).Unix()
}

func (t *timeCreator) FirstTimeOfTodayBaseDay() time.Time {
	now := time.Now()

	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func (t *timeCreator) GetPastYearUnixTime() int64 {
	return t.AddUnix(-365 * 24 * time.Hour)
}

func TimeNowBaseSecond() string {
	return timeNow(timeFormatYearMonthDayHourMinSecond)
}

func timeNow(layout string) string {
	return time.Now().Format(layout)
}
