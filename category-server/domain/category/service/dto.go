package service

import (
	"strconv"

	"github.com/victorzhou123/vicblog/category-server/domain/category/entity"
	"github.com/victorzhou123/vicblog/common/controller/rpc"
	cmdto "github.com/victorzhou123/vicblog/common/domain/dto"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type CategoryListDto struct {
	cment.PaginationStatus

	Categories []entity.CategoryWithRelatedArticleAmount
}

func (dto *CategoryListDto) toProto() *rpc.CategoryList {
	if dto == nil {
		return nil
	}

	categories := make([]*rpc.CategoryWithRelatedArticleAmount, len(dto.Categories))
	for i := range categories {
		categories[i] = toProtoCategoryWithRelatedArticleAmount(dto.Categories[i])
	}

	return &rpc.CategoryList{
		PaginationStatus: &rpc.PaginationStatus{
			Pagination: &rpc.Pagination{
				CurPage:  strconv.Itoa(dto.CurPage.CurPage()),
				PageSize: strconv.Itoa(dto.PageSize.PageSize()),
			},
		},
		Categories: nil,
	}
}

func toProtoCategoryWithRelatedArticleAmount(c entity.CategoryWithRelatedArticleAmount) *rpc.CategoryWithRelatedArticleAmount {
	return &rpc.CategoryWithRelatedArticleAmount{
		Category: &rpc.Category{
			Id:        c.Id.Id(),
			Name:      c.Name.CategoryName(),
			CreatedAt: c.CreatedAt.TimeUnix(),
		},
		RelatedArticleAmount: int64(c.RelatedArticleAmount.Amount()),
	}
}

func toCategoryWithRelatedArticleAmount(in *rpc.CategoryWithRelatedArticleAmount) entity.CategoryWithRelatedArticleAmount {

	amount, err := cmprimitive.NewAmount(int(in.GetRelatedArticleAmount()))
	if err != nil {
		return entity.CategoryWithRelatedArticleAmount{}
	}

	categoryName, err := entity.NewCategoryName(in.Category.GetName())
	if err != nil {
		return entity.CategoryWithRelatedArticleAmount{}
	}

	return entity.CategoryWithRelatedArticleAmount{
		Category: entity.Category{
			Id:        cmprimitive.NewId(in.Category.GetId()),
			Name:      categoryName,
			CreatedAt: cmprimitive.NewTimeXWithUnix(in.Category.CreatedAt),
		},
		RelatedArticleAmount: amount,
	}
}

func toCategoryListDto(categoryList *rpc.CategoryList) (CategoryListDto, error) {
	categories := make([]entity.CategoryWithRelatedArticleAmount, len(categoryList.Categories))
	for i := range categories {

		amount, err := cmprimitive.NewAmount(int(categoryList.Categories[i].GetRelatedArticleAmount()))
		if err != nil {
			return CategoryListDto{}, err
		}

		categoryName, err := entity.NewCategoryName(categoryList.Categories[i].Category.GetName())
		if err != nil {
			return CategoryListDto{}, err
		}

		categories[i] = entity.CategoryWithRelatedArticleAmount{
			Category: entity.Category{
				Id:        cmprimitive.NewId(categoryList.Categories[i].Category.GetId()),
				Name:      categoryName,
				CreatedAt: cmprimitive.NewTimeXWithUnix(categoryList.Categories[i].Category.CreatedAt),
			},
			RelatedArticleAmount: amount,
		}
	}

	paginationStatus, err := cmdto.ToPaginationStatus(categoryList.GetPaginationStatus())
	if err != nil {
		return CategoryListDto{}, err
	}

	return CategoryListDto{
		PaginationStatus: paginationStatus,
		Categories:       categories,
	}, nil
}
