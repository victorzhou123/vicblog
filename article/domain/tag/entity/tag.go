package entity

import cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"

type Tag struct {
	Id        cmprimitive.Id
	Name      TagName
	CreatedAt cmprimitive.Timex
}

type TagWithRelatedArticleAmount struct {
	Tag

	RelatedArticleAmount cmprimitive.Amount
}
