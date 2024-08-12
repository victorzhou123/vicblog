package primitive

import (
	"errors"
	"strconv"
)

// curPage
type CurPage interface {
	CurPage() int
}

type curPage int

func NewCurPageWithString(v string) (CurPage, error) {

	cur, err := strconv.Atoi(v)
	if err != nil {
		return nil, errors.New("input curPage not a number")
	}

	return NewCurPage(cur)
}

func NewCurPage(v int) (CurPage, error) {

	if v <= 0 {
		return nil, errors.New("curPage must bigger than 0")
	}

	return curPage(v), nil
}

func (r curPage) CurPage() int {
	return int(r)
}

// pageSize
type PageSize interface {
	PageSize() int
}

type pageSize int

func NewPageSizeWithString(v string) (PageSize, error) {

	cur, err := strconv.Atoi(v)
	if err != nil {
		return nil, errors.New("input pageSize not a number")
	}

	return NewPageSize(cur)
}

func NewPageSize(v int) (PageSize, error) {

	if v <= 0 {
		return nil, errors.New("pageSize must bigger than 0")
	}

	return pageSize(v), nil
}

func (r pageSize) PageSize() int {
	return int(r)
}
