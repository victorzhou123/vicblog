package repository

import "victorzhou123/vicblog/article/domain/category/entity"

type Category interface {
	AddCategory(entity.CategoryName) error
}
