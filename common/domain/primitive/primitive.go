package primitive

import (
	"time"
	"victorzhou123/vicblog/common/validator"
)

const timeFormatYearToSecond = "2006-01-02 15:04:05"

// id
type Id interface {
	Id() string
}

type id string

func NewId(v string) Id {
	return id(v)
}

func (r id) Id() string {
	return string(r)
}

// time
type Timex interface {
	TimeUnix() int64
	TimeYearToSecond() string
}

type timex int64

func NewTimeXWithUnix(v int64) Timex {
	return timex(v)
}

func NewTimeXNow() Timex {
	return timex(time.Now().Unix())
}

func (t timex) TimeUnix() int64 {
	return int64(t)
}

func (t timex) TimeYearToSecond() string {
	return time.Unix(t.TimeUnix(), 0).Format(timeFormatYearToSecond)
}

// url
type Urlx interface {
	Urlx() string
}

type urlx string

func NewUrlx(v string) (Urlx, error) {
	if err := validator.IsUrl(v); err != nil {
		return nil, err
	}

	return urlx(v), nil
}

func (u urlx) Urlx() string {
	return string(u)
}
