package service

import (
	"victorzhou123/vicblog/article/domain/tag/entity"
	cment "victorzhou123/vicblog/common/domain/entity"
)

type TagListDto struct {
	cment.PaginationStatus

	Tags []entity.Tag
}
