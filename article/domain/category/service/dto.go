package service

import (
	categoryett "victorzhou123/vicblog/article/domain/category/entity"
	"victorzhou123/vicblog/common/domain/repository"
	dmservice "victorzhou123/vicblog/common/domain/service"
)

type CategoryListCmd struct {
	dmservice.PaginationCmd
}

func (cmd *CategoryListCmd) Validate() error {
	return cmd.PaginationCmd.Validate()
}

func (cmd *CategoryListCmd) toPageListOpt() repository.PageListOpt {
	return cmd.PaginationCmd.ToPageListOpt()
}

type CategoryListDto struct {
	dmservice.PaginationDto

	Category []CategoryDto `json:"category"`
}

func toCategoryListDto(cates []categoryett.Category, cmd *CategoryListCmd, total int) CategoryListDto {

	categoryDos := make([]CategoryDto, len(cates))
	for i := range cates {
		categoryDos[i] = toCategoryDto(cates[i])
	}

	return CategoryListDto{
		PaginationDto: cmd.ToPaginationDto(total),
		Category:      categoryDos,
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
