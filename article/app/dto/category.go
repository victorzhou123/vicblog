package dto

import (
	"github.com/victorzhou123/vicblog/category-server/domain/category/entity"
	cmappdto "github.com/victorzhou123/vicblog/common/app/dto"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
)

// list category
type ListCategoryCmd struct {
	cmappdto.PaginationCmd
}

type CategoryDto struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

func ToCategoryDto(cate entity.Category) CategoryDto {
	return CategoryDto{
		Id:        cate.Id.IdNum(),
		Name:      cate.Name.CategoryName(),
		CreatedAt: cate.CreatedAt.TimeYearToSecond(),
	}
}

type CategoryWithRelatedArticleAmountDto struct {
	CategoryDto

	RelatedArticleAmount int `json:"relatedArticleAmount"`
}

func ToCategoryWithRelatedArticleAmountDto(
	c entity.CategoryWithRelatedArticleAmount,
) CategoryWithRelatedArticleAmountDto {
	return CategoryWithRelatedArticleAmountDto{
		CategoryDto:          ToCategoryDto(c.Category),
		RelatedArticleAmount: c.RelatedArticleAmount.Amount(),
	}
}

type CategoryListDto struct {
	cmappdto.PaginationDto

	Category []CategoryWithRelatedArticleAmountDto `json:"category"`
}

func ToCategoryListDto(
	ps cment.PaginationStatus, cates []entity.CategoryWithRelatedArticleAmount,
) CategoryListDto {

	cateDtos := make([]CategoryWithRelatedArticleAmountDto, len(cates))
	for i := range cates {
		cateDtos[i] = ToCategoryWithRelatedArticleAmountDto(cates[i])
	}

	return CategoryListDto{
		PaginationDto: cmappdto.ToPaginationDto(ps),
		Category:      cateDtos,
	}
}
