package repository

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type TagArticle interface {
	BuildRelationWithArticle(articleId cmprimitive.Id, tagIds []cmprimitive.Id) error
	RemoveAllRowsByArticleId(articleId cmprimitive.Id) error
}
