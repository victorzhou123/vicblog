package service

import (
	"github.com/victorzhou123/vicblog/article/domain/tag/entity"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
)

type TagListDto struct {
	cment.PaginationStatus

	Tags []entity.TagWithRelatedArticleAmount
}
