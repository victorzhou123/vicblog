package service

import (
	categoryett "victorzhou123/vicblog/article/domain/category/entity"
	"victorzhou123/vicblog/common/domain/repository"
	dmservice "victorzhou123/vicblog/common/domain/service"
)

type CategoryListCmd struct {
	dmservice.ListCmd
}

func (cmd *CategoryListCmd) Validate() error {
	return cmd.ListCmd.Validate()
}

func (cmd *CategoryListCmd) toPageListOpt() repository.PageListOpt {
	return cmd.ListCmd.ToPageListOpt()
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
