package service

import (
	commentent "github.com/victorzhou123/vicblog/comment/domain/comment/entity"
	"github.com/victorzhou123/vicblog/comment/domain/qqinfo/entity"
	"github.com/victorzhou123/vicblog/comment/domain/qqinfo/qqinfo"
)

type QQInfoService interface {
	GetQQInfo(entity.QQNumber) (commentent.CommentUserInfo, error)
}

type qqInfoService struct {
	qqInfo qqinfo.QQInfo
}

func NewQQInfoService(qqInfo qqinfo.QQInfo) QQInfoService {
	return &qqInfoService{
		qqInfo: qqInfo,
	}
}

func (s *qqInfoService) GetQQInfo(num entity.QQNumber) (commentent.CommentUserInfo, error) {
	return s.qqInfo.GetQQInfo(num)
}
