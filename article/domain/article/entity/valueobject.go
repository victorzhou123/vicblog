package entity

import (
	"victorzhou123/vicblog/common/util"
	"victorzhou123/vicblog/common/validator"
)

type ArticleSummary interface {
	ArticleSummary() string
}

type articleSummary string

func NewArticleSummary(v string) (ArticleSummary, error) {
	if err := validator.IsArticleSummary(v); err != nil {
		return nil, err
	}

	return articleSummary(util.XssEscape(v)), nil
}

func (r articleSummary) ArticleSummary() string {
	return string(r)
}
