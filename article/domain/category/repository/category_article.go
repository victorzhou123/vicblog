package repository

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type CategoryArticle interface {
	BindCategoryAndArticle(articleId, cateId cmprimitive.Id) error
}
