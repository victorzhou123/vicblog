package service

import (
	"github.com/victorzhou123/vicblog/statistics/app/dto"
	dmsvc "github.com/victorzhou123/vicblog/statistics/domain/service"
)

type ArticleVisitsAppService interface {
	GetArticleVisitsOfAWeek() (dto.VisitsOfAWeekDto, error)
}

type articleVisitsAppService struct {
	articleVisits dmsvc.ArticleVisitsService
}

func NewArticleVisitsAppService(articleVisits dmsvc.ArticleVisitsService) ArticleVisitsAppService {
	return &articleVisitsAppService{articleVisits}
}

func (s *articleVisitsAppService) GetArticleVisitsOfAWeek() (dto.VisitsOfAWeekDto, error) {

	visits, err := s.articleVisits.GetAscendVisitsOfAWeek()
	if err != nil {
		return dto.VisitsOfAWeekDto{}, err
	}

	return dto.ToVisitsOfAWeekDto(visits), nil
}
