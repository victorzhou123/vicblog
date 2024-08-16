package repository

import (
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type CategoryArticle interface {
	GetRelationWithArticle(articleId cmprimitive.Id) (cateId cmprimitive.Id, err error)
	GetRelatedArticleAmount(categoryIds []cmprimitive.Id) (map[uint]cmprimitive.Amount, error)

	BuildRelationWithArticle(articleId, cateId cmprimitive.Id) error

	RemoveAllRowsByArticleId(articleId cmprimitive.Id) error
}
