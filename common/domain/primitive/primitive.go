package primitive

import (
	"errors"
	"strconv"
	"time"

	"github.com/victorzhou123/vicblog/common/validator"
)

const (
	timeFormatYearToSecond = "2006-01-02 15:04:05"
	timeFormatMonthDayOnly = "01-02"
)

// id
type Id interface {
	Id() string
	IdNum() uint
}

type id string

func NewId(v string) Id {
	return id(v)
}

func NewIdByUint(v uint) Id {
	return NewId(strconv.FormatUint(uint64(v), 10))
}

func (r id) Id() string {
	return string(r)
}

// IdNum return an uint type ID.
// return 0 if it is unable to convert
func (r id) IdNum() uint {
	if num, err := strconv.ParseUint(r.Id(), 10, 64); err == nil {
		return uint(num)
	}

	return 0
}

// time
type Timex interface {
	TimeUnix() int64
	TimeYearToSecond() string
	TimeMonthDayOnly() string
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

func (t timex) TimeMonthDayOnly() string {
	return time.Unix(t.TimeUnix(), 0).Format(timeFormatMonthDayOnly)
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

// amount
type Amount interface {
	Amount() int
}

type amount int

func NewAmount(v int) (Amount, error) {

	if err := validator.IsAmount(v); err != nil {
		return nil, err
	}

	return amount(v), nil
}

func NewAmountByString(v string) (Amount, error) {

	amount, err := strconv.Atoi(v)
	if err != nil {
		return nil, errors.New("input amount is not a number")
	}

	return NewAmount(amount)
}

func (r amount) Amount() int {
	return int(r)
}
