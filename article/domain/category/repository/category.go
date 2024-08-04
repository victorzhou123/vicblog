package repository

import (
	"victorzhou123/vicblog/article/domain/category/entity"
	"victorzhou123/vicblog/common/domain/repository"
)

type Category interface {
	AddCategory(entity.CategoryName) error
	GetCategoryList(repository.PageListOpt) ([]entity.Category, int, error)
}
