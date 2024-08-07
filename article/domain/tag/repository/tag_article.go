package repository

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type TagArticle interface {
	AddRelateWithArticle(articleId, tagId cmprimitive.Id) error
}
