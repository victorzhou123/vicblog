package entity

import cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"

type Category struct {
	Id        cmprimitive.Id
	Name      CategoryName
	CreatedAt cmprimitive.Timex
}

type CategoryWithRelatedArticleAmount struct {
	Category

	RelatedArticleAmount cmprimitive.Amount
}
