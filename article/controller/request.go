package controller

import (
	"strconv"
	"victorzhou123/vicblog/article/app"
	"victorzhou123/vicblog/article/domain/category/entity"
	tagent "victorzhou123/vicblog/article/domain/tag/entity"
	"victorzhou123/vicblog/article/domain/tag/repository"
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

type reqTag struct {
	Names []string `json:"names"`
}

func (req *reqTag) toTagNames() (repository.TagNames, error) {
	tagNames := make([]tagent.TagName, len(req.Names))

	var err error
	for i := range req.Names {
		tagNames[i], err = tagent.NewTagName(req.Names[i])
		if err != nil {
			return repository.TagNames{}, err
		}
	}

	return repository.TagNames{Names: tagNames}, nil
}

type reqTagList struct {
	CurPage  string `json:"current"`
	PageSize string `json:"size"`
}

func (req *reqTagList) toCmd() (cmd app.TagListCmd, err error) {
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
