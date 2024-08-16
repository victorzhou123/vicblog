package service

import (
	"github.com/victorzhou123/vicblog/article/app/dto"
	"github.com/victorzhou123/vicblog/article/domain/category/entity"
	categorydmsvc "github.com/victorzhou123/vicblog/article/domain/category/service"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type CategoryAppService interface {
	ListCategories(amount cmprimitive.Amount) ([]dto.CategoryWithRelatedArticleAmountDto, error)
	ListCategoryByPagination(*dto.ListCategoryCmd) (dto.CategoryListDto, error)

	AddCategory(entity.CategoryName) error

	DelCategory(categoryId cmprimitive.Id) error
}

type categoryAppService struct {
	cate categorydmsvc.CategoryService
}

func NewCategoryAppService(cate categorydmsvc.CategoryService) CategoryAppService {
	return &categoryAppService{
		cate: cate,
	}
}

func (s *categoryAppService) ListCategories(amount cmprimitive.Amount) ([]dto.CategoryWithRelatedArticleAmountDto, error) {

	cateWithAmounts, err := s.cate.ListCategories(amount)
	if err != nil {
		return nil, err
	}

	dtos := make([]dto.CategoryWithRelatedArticleAmountDto, len(cateWithAmounts))
	for i := range cateWithAmounts {
		dtos[i] = dto.ToCategoryWithRelatedArticleAmountDto(cateWithAmounts[i])
	}

	return dtos, nil
}

func (s *categoryAppService) ListCategoryByPagination(cmd *dto.ListCategoryCmd) (dto.CategoryListDto, error) {

	cateListDto, err := s.cate.ListCategoryByPagination(cmd.ToPagination())
	if err != nil {
		return dto.CategoryListDto{}, err
	}

	return dto.ToCategoryListDto(cateListDto.PaginationStatus, cateListDto.Categories), nil
}

func (s *categoryAppService) AddCategory(cateName entity.CategoryName) error {
	return s.cate.AddCategory(cateName)
}

func (s *categoryAppService) DelCategory(categoryId cmprimitive.Id) error {
	return s.cate.DelCategory(categoryId)
}
