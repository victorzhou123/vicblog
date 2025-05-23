package repository

import (
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/tag-server/domain/tag/entity"
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
	GetTagList(cmprimitive.Amount) ([]entity.Tag, error)
	GetTotalNumberOfTag() (cmprimitive.Amount, error)
	Delete(cmprimitive.Id) error
}
