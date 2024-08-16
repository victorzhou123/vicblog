package service

import (
	"github.com/victorzhou123/vicblog/article/domain/category/entity"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
)

type CategoryListDto struct {
	cment.PaginationStatus

	Categories []entity.CategoryWithRelatedArticleAmount
}
