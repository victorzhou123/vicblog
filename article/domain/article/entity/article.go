package entity

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type Article struct {
	Id        cmprimitive.Id
	Owner     cmprimitive.Username
	Title     cmprimitive.Text
	Summary   ArticleSummary
	Content   cmprimitive.Text
	Cover     cmprimitive.Urlx
	ReadTimes int
	IsPublish bool
	IsTop     bool
	UpdatedAt cmprimitive.Timex
	CreatedAt cmprimitive.Timex
}
