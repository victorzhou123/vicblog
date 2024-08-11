package repository

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type TagArticle interface {
	GetRelationWithArticle(articleId cmprimitive.Id) (tags []cmprimitive.Id, err error)

	BuildRelationWithArticle(articleId cmprimitive.Id, tagIds []cmprimitive.Id) error

	RemoveAllRowsByArticleId(articleId cmprimitive.Id) error
}
