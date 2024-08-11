package repository

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type CategoryArticle interface {
	GetRelationWithArticle(articleId cmprimitive.Id) (cateId cmprimitive.Id, err error)

	BuildRelationWithArticle(articleId, cateId cmprimitive.Id) error

	RemoveAllRowsByArticleId(articleId cmprimitive.Id) error
}
