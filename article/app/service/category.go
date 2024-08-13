package service

import (
	"victorzhou123/vicblog/article/app/dto"
	"victorzhou123/vicblog/article/domain/category/entity"
	categorydmsvc "victorzhou123/vicblog/article/domain/category/service"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type CategoryAppService interface {
	ListCategories(amount cmprimitive.Amount) ([]dto.CategoryDto, error)
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

func (s *categoryAppService) ListCategories(amount cmprimitive.Amount) ([]dto.CategoryDto, error) {

	cates, err := s.cate.ListCategories(amount)
	if err != nil {
		return nil, err
	}

	cateDtos := make([]dto.CategoryDto, len(cates))
	for i := range cates {
		cateDtos[i] = dto.ToCategoryDto(cates[i])
	}

	return cateDtos, nil
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
