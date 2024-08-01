package domain

import (
	"victorzhou123/vicblog/common/util"
	"victorzhou123/vicblog/common/validator"
)

type Text interface {
	Text() string
}

type text string

func NewTitle(v string) (Text, error) {
	v = util.XssEscape(v)

	if err := validator.IsTitle(v); err != nil {
		return nil, err
	}

	return text(v), nil
}

func (t text) Text() string {
	return string(t)
}
