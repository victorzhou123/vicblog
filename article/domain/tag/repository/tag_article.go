package repository

import cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"

type TagArticle interface {
	GetRelationWithArticle(articleId cmprimitive.Id) (tags []cmprimitive.Id, err error)
	GetRelatedArticleAmount(tagIds []cmprimitive.Id) (map[uint]cmprimitive.Amount, error)
	GetRelatedArticleIdsThoroughTagId(tagId cmprimitive.Id) ([]cmprimitive.Id, error)

	BuildRelationWithArticle(articleId cmprimitive.Id, tagIds []cmprimitive.Id) error

	RemoveAllRowsByArticleId(articleId cmprimitive.Id) error
}
