package repository

import (
	"victorzhou123/vicblog/article/domain/tag/entity"
	cment "victorzhou123/vicblog/common/domain/entity"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type TagNames struct {
	Names []entity.TagName
}

func (t *TagNames) NoDuplication() bool {
	for i := range t.Names {
		for j := i + 1; j < len(t.Names); j++ {
			if t.Names[i].Equal(t.Names[j]) {
				return false
			}
		}
	}

	return true
}

type Tag interface {
	AddBatches(TagNames) error
	GetBatchTags(tagIds []cmprimitive.Id) ([]entity.Tag, error)
	GetTagListByPagination(cment.Pagination) ([]entity.Tag, int, error)
	GetAllTagList() ([]entity.Tag, error)
	Delete(cmprimitive.Id) error
}
