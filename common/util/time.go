package util

import "time"

const (
	timeFormatYearMonthDayHourMinSecond = "2006-01-02 15:04:05.000"
)

func TimeNowBaseSecond() string {
	return timeNow(timeFormatYearMonthDayHourMinSecond)
}

func timeNow(layout string) string {
	return time.Now().Format(layout)
}
