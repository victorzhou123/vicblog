package controller

import (
	"strconv"

	articleent "victorzhou123/vicblog/article/domain/article/entity"
	"victorzhou123/vicblog/article/domain/category/entity"
	categorysvc "victorzhou123/vicblog/article/domain/category/service"
	tagent "victorzhou123/vicblog/article/domain/tag/entity"
	"victorzhou123/vicblog/article/domain/tag/repository"
	tagsvc "victorzhou123/vicblog/article/domain/tag/service"
	articlesvc "victorzhou123/vicblog/article/domain/article/service"
	"victorzhou123/vicblog/common/domain/primitive"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
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

func (req *reqCategoryList) emptyValue() bool {
	return req.CurPage == "" && req.PageSize == ""
}

func (req *reqCategoryList) toCmd() (cmd categorysvc.CategoryListCmd, err error) {
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

func (req *reqTagList) emptyValue() bool {
	return req.CurPage == "" && req.PageSize == ""
}

func (req *reqTagList) toCmd() (cmd tagsvc.TagListCmd, err error) {
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

type reqArticle struct {
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Content  string `json:"content"`
	Cover    string `json:"cover"`
	Category uint   `json:"categoryId"`
	Tags     []uint `json:"tags"`
}

func (req *reqArticle) toCmd(user primitive.Username) (cmd articlesvc.AddArticleCmd, err error) {

	if cmd.Title, err = cmprimitive.NewTitle(req.Title); err != nil {
		return
	}

	if cmd.Summary, err = articleent.NewArticleSummary(req.Summary); err != nil {
		return
	}

	if cmd.Content, err = cmprimitive.NewArticleContent(req.Content); err != nil {
		return
	}

	if cmd.Cover, err = cmprimitive.NewUrlx(req.Cover); err != nil {
		return
	}

	cmd.Owner = user

	cmd.Category = cmprimitive.NewIdByUint(req.Category)

	cmd.Tags = make([]cmprimitive.Id, len(req.Tags))
	for i := range req.Tags {
		cmd.Tags[i] = cmprimitive.NewIdByUint(req.Tags[i])
	}

	return
}
