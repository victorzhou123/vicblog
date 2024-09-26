package qqinfo

import (
	commentent "github.com/victorzhou123/vicblog/comment/domain/comment/entity"
	"github.com/victorzhou123/vicblog/comment/domain/qqinfo/entity"
)

type QQInfo interface {
	GetQQInfo(entity.QQNumber) (commentent.CommentUserInfo, error)
}
