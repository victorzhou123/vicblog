package entity

import (
	"github.com/victorzhou123/vicblog/common/util"
	"github.com/victorzhou123/vicblog/common/validator"
)

type TagName interface {
	TagName() string
	Equal(TagName) bool
}

type tagName string

func NewTagName(v string) (TagName, error) {
	if err := validator.IsTagName(v); err != nil {
		return nil, err
	}

	return tagName(util.XssEscape(v)), nil
}

func (t tagName) TagName() string {
	return string(t)
}

func (t tagName) Equal(name TagName) bool {
	return t.TagName() == name.TagName()
}
