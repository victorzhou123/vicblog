package dto

import (
	"victorzhou123/vicblog/article/domain/article/entity"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type AddArticleCmd struct {
	Owner    cmprimitive.Username
	Title    cmprimitive.Text
	Summary  entity.ArticleSummary
	Content  cmprimitive.Text
	Cover    cmprimitive.Urlx
	Category cmprimitive.Id
	Tag      cmprimitive.Id
}
