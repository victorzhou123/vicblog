package service

import (
	"fmt"

	categoryett "victorzhou123/vicblog/article/domain/category/entity"
	"victorzhou123/vicblog/common/domain/repository"
)

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
