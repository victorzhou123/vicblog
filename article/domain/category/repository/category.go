package repository

import (
	"github.com/victorzhou123/vicblog/article/domain/category/entity"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type Category interface {
	AddCategory(entity.CategoryName) error

	GetCategory(cateId cmprimitive.Id) (entity.Category, error)
	GetCategoryListByPagination(cment.Pagination) ([]entity.Category, int, error)
	GetCategoryList(amount cmprimitive.Amount) ([]entity.Category, error) // nil means get all
	GetTotalNumberOfCategories() (cmprimitive.Amount, error)

	DelCategory(cmprimitive.Id) error
}
