package repository

import (
	"victorzhou123/vicblog/article/domain/category/entity"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/domain/repository"
)

type Category interface {
	AddCategory(entity.CategoryName) error

	GetCategory(cateId cmprimitive.Id) (entity.Category, error)
	GetCategoryList(repository.PageListOpt) ([]entity.Category, int, error)
	GetAllCategoryList() ([]entity.Category, error)

	DelCategory(cmprimitive.Id) error
}
