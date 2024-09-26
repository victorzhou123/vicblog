package repository

import (
	"github.com/victorzhou123/vicblog/comment/domain/comment/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type Comment interface {
	Add(entity.Comment) error
	GetCommentsByArticleId(articleId cmprimitive.Id) ([]entity.Comment, error)
}
