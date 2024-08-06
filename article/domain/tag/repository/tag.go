package repository

import (
	"victorzhou123/vicblog/article/domain/tag/entity"
	"victorzhou123/vicblog/common/domain/primitive"
	cmrepo "victorzhou123/vicblog/common/domain/repository"
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
	GetTagList(cmrepo.PageListOpt) ([]entity.Tag, int, error)
	GetAllTagList() ([]entity.Tag, error)
	Delete(primitive.Id) error
}
