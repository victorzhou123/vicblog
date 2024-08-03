package app

import (
	"victorzhou123/vicblog/article/domain/category/entity"
	"victorzhou123/vicblog/article/domain/category/repository"
)

type CategoryService interface {
	AddCategory(entity.CategoryName) error
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
