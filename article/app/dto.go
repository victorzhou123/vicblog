package app

import (
	"fmt"
	"victorzhou123/vicblog/article/domain/article/entity"
	categoryett "victorzhou123/vicblog/article/domain/category/entity"
	tagett "victorzhou123/vicblog/article/domain/tag/entity"
	cmapp "victorzhou123/vicblog/common/app"
	"victorzhou123/vicblog/common/domain/repository"
)

type ArticleListDto struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	IsPublish bool   `json:"isPublish"`
	IsTop     bool   `json:"isTop"`
	CreatedAt string `json:"createTime"`
}

func toArticleListDto(v entity.Article) ArticleListDto {
	return ArticleListDto{
		Id:        v.Id.IdNum(),
		Title:     v.Title.Text(),
		Cover:     v.Cover.Urlx(),
		IsPublish: v.IsPublish,
		IsTop:     v.IsTop,
		CreatedAt: v.CreatedAt.TimeYearToSecond(),
	}
}

type CategoryListCmd struct {
	CurPage  int
	PageSize int
}

func (cmd *CategoryListCmd) Validate() error {
	if cmd.CurPage <= 0 {
		return fmt.Errorf("current page must > 0")
	}

	if cmd.PageSize <= 0 {
		return fmt.Errorf("page size must > 0")
	}

	return nil
}

func (cmd *CategoryListCmd) toPageListOpt() repository.PageListOpt {
	return repository.PageListOpt{
		CurPage:  cmd.CurPage,
		PageSize: cmd.PageSize,
	}
}

type CategoryListDto struct {
	Total     int           `json:"total"`
	PageCount int           `json:"pages"`
	PageSize  int           `json:"size"`
	CurPage   int           `json:"current"`
	Category  []CategoryDto `json:"category"`
}

func toCategoryListDto(cates []categoryett.Category, cmd *CategoryListCmd, total int) CategoryListDto {

	pageCount := total/cmd.PageSize + 1

	categoryDos := make([]CategoryDto, len(cates))
	for i := range cates {
		categoryDos[i] = toCategoryDto(cates[i])
	}

	return CategoryListDto{
		Total:     total,
		PageCount: pageCount,
		PageSize:  cmd.PageSize,
		CurPage:   cmd.CurPage,
		Category:  categoryDos,
	}
}

type CategoryDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createTime"`
}

func toCategoryDto(cate categoryett.Category) CategoryDto {
	return CategoryDto{
		Id:        cate.Id.Id(),
		Name:      cate.Name.CategoryName(),
		CreatedAt: cate.CreatedAt.TimeYearToSecond(),
	}
}

type TagListCmd struct {
	cmapp.ListCmd
}

type TagListDto struct {
	Total     int      `json:"total"`
	PageCount int      `json:"pages"`
	PageSize  int      `json:"size"`
	CurPage   int      `json:"current"`
	Tag       []TagDto `json:"tag"`
}

func toTagListDto(tags []tagett.Tag, cmd *TagListCmd, total int) TagListDto {

	pageCount := total/cmd.PageSize + 1

	dtos := make([]TagDto, len(tags))
	for i := range tags {
		dtos[i] = toTagDto(tags[i])
	}

	return TagListDto{
		Total:     total,
		PageCount: pageCount,
		PageSize:  cmd.PageSize,
		CurPage:   cmd.CurPage,
		Tag:       dtos,
	}
}

type TagDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createTime"`
}

func toTagDto(tag tagett.Tag) TagDto {
	return TagDto{
		Id:        tag.Id.Id(),
		Name:      tag.Name.TagName(),
		CreatedAt: tag.CreatedAt.TimeYearToSecond(),
	}
}
