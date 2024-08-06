package service

import (
	"victorzhou123/vicblog/article/domain/category/entity"
	"victorzhou123/vicblog/article/domain/category/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type CategoryService interface {
	AddCategory(entity.CategoryName) error
	ListCategory(*CategoryListCmd) (CategoryListDto, error)
	ListAllCategory() ([]CategoryDto, error)
	DelCategory(cmprimitive.Id) error
}

type categoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) AddCategory(category entity.CategoryName) error {
	return s.repo.AddCategory(category)
}

func (s *categoryService) ListCategory(cmd *CategoryListCmd) (CategoryListDto, error) {

	cates, total, err := s.repo.GetCategoryList(cmd.toPageListOpt())
	if err != nil {
		if cmdmerror.IsNotFound(err) {
			return CategoryListDto{}, nil
		}

		return CategoryListDto{}, err
	}

	return toCategoryListDto(cates, cmd, total), nil
}

func (s *categoryService) ListAllCategory() ([]CategoryDto, error) {

	cates, err := s.repo.GetAllCategoryList()
	if err != nil {
		return nil, err
	}

	dtos := make([]CategoryDto, len(cates))
	for i := range cates {
		dtos[i] = toCategoryDto(cates[i])
	}

	return dtos, nil
}

func (s *categoryService) DelCategory(id cmprimitive.Id) error {
	return s.repo.DelCategory(id)
}
