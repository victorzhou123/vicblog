package domain

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type Article struct {
	Id        cmprimitive.Id
	Title     Text
	Cover     cmprimitive.Urlx
	IsPublish bool
	IsTop     bool
	UpdatedAt cmprimitive.Timex
	CreatedAt cmprimitive.Timex
}
