package service

import (
	"github.com/victorzhou123/vicblog/comment/app/dto"
	"github.com/victorzhou123/vicblog/comment/domain/qqinfo/entity"
	"github.com/victorzhou123/vicblog/comment/domain/qqinfo/service"
)

type QQInfoAppService interface {
	GetQQInfo(qqNum entity.QQNumber) (dto.QQInfoDto, error)
}

type qqInfoAppService struct {
	qqInfo service.QQInfoService
}

func NewQQInfoAppService(qqInfo service.QQInfoService) QQInfoAppService {
	return &qqInfoAppService{
		qqInfo: qqInfo,
	}
}

func (s *qqInfoAppService) GetQQInfo(qqNum entity.QQNumber) (dto.QQInfoDto, error) {

	userInfo, err := s.qqInfo.GetQQInfo(qqNum)
	if err != nil {
		return dto.QQInfoDto{}, err
	}

	return dto.QQInfoDto{
		Avatar:   userInfo.Avatar.Urlx(),
		NickName: userInfo.NickName,
		Email:    userInfo.Email.Email(),
	}, nil
}
