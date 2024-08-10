package repository

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type CategoryArticle interface {
	BuildRelationWithArticle(articleId, cateId cmprimitive.Id) error
	RemoveAllRowsByArticleId(articleId cmprimitive.Id) error
}
