package service

import (
	"victorzhou123/vicblog/article/domain/category/entity"
	cment "victorzhou123/vicblog/common/domain/entity"
)

type CategoryListDto struct {
	cment.PaginationStatus

	Categories []entity.Category
}
