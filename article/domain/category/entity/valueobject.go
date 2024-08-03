package entity

import (
	"victorzhou123/vicblog/common/util"
	"victorzhou123/vicblog/common/validator"
)

type CategoryName interface {
	CategoryName() string
}

type categoryName string

func NewCategoryName(v string) (CategoryName, error) {
	if err := validator.IsCategoryName(v); err != nil {
		return nil, err
	}

	return categoryName(util.XssEscape(v)), nil
}

func (c categoryName) CategoryName() string {
	return string(c)
}
