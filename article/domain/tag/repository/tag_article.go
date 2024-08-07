package repository

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type TagArticle interface {
	AddRelateWithArticle(articleId cmprimitive.Id, tagIds []cmprimitive.Id) error
}
