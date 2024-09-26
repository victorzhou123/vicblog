package entity

import (
	"github.com/victorzhou123/vicblog/common/validator"
)

// qq number
type QQNumber interface {
	QQNumber() string
}

type qqNumber string

func NewQQNumber(v string) (QQNumber, error) {

	if err := validator.IsQQNumber(v); err != nil {
		return nil, err
	}

	return qqNumber(v), nil
}

func (r qqNumber) QQNumber() string {
	return string(r)
}
