package controller

import (
	"strconv"
	"victorzhou123/vicblog/article/app"
	"victorzhou123/vicblog/article/domain/category/entity"
)

type reqCategory struct {
	Name string `json:"name"`
}

func (req *reqCategory) toCategoryName() (entity.CategoryName, error) {
	return entity.NewCategoryName(req.Name)
}

type reqCategoryList struct {
	CurPage  string `json:"current"`
	PageSize string `json:"size"`
}

func (req *reqCategoryList) toCmd() (cmd app.CategoryListCmd, err error) {
	if cmd.CurPage, err = strconv.Atoi(req.CurPage); err != nil {
		return
	}

	if cmd.PageSize, err = strconv.Atoi(req.PageSize); err != nil {
		return
	}

	if err = cmd.Validate(); err != nil {
		return
	}

	return
}
