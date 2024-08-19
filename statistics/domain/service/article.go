package service

import (
	"sort"

	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/statistics/domain/entity"
	"github.com/victorzhou123/vicblog/statistics/domain/repository"
)

type ArticleVisitsService interface {
	GetTotalVisits() (cmprimitive.Amount, error)
	GetAscendVisitsOfAWeek() ([]entity.ArticleDailyVisits, error)
	IncreaseOneVisitsOfToday() error
}

type articleVisitsService struct {
	repo repository.ArticleVisits
}

func NewArticleVisitsService(repo repository.ArticleVisits) ArticleVisitsService {
	return &articleVisitsService{
		repo: repo,
	}
}

func (s *articleVisitsService) GetTotalVisits() (cmprimitive.Amount, error) {
	return s.repo.GetTotalVisits()
}

func (s *articleVisitsService) GetAscendVisitsOfAWeek() ([]entity.ArticleDailyVisits, error) {

	visits, err := s.repo.GetDailyVisitsOfWeek()
	if err != nil {
		return nil, err
	}

	sort.Slice(visits, func(i, j int) bool {
		return visits[i].Date.TimeUnix() < visits[j].Date.TimeUnix()
	})

	return visits, nil
}

func (s *articleVisitsService) IncreaseOneVisitsOfToday() error {

	one, err := cmprimitive.NewAmount(1)
	if err != nil {
		return err
	}

	return s.repo.IncreaseTodayVisits(one)
}
