package service

import (
	articlesvc "github.com/victorzhou123/vicblog/article/domain/article/service"
	categorysvc "github.com/victorzhou123/vicblog/article/domain/category/service"
	tagsvc "github.com/victorzhou123/vicblog/article/domain/tag/service"
	"github.com/victorzhou123/vicblog/statistics/app/dto"
	dmsvc "github.com/victorzhou123/vicblog/statistics/domain/service"
)

type DashboardAppService interface {
	GetDashboardData() (dto.DashboardDataDto, error)
}

type dashboardAppService struct {
	article  articlesvc.ArticleService
	tag      tagsvc.TagService
	category categorysvc.CategoryService

	articleVisits dmsvc.ArticleVisitsService
}

func NewDashboardAppService(
	article articlesvc.ArticleService,
	tag tagsvc.TagService,
	category categorysvc.CategoryService,

	articleVisits dmsvc.ArticleVisitsService,
) DashboardAppService {
	return &dashboardAppService{
		article:       article,
		tag:           tag,
		category:      category,
		articleVisits: articleVisits,
	}
}

func (s *dashboardAppService) GetDashboardData() (dto.DashboardDataDto, error) {

	// total visits of article
	articleVisits, err := s.articleVisits.GetTotalVisits()
	if err != nil {
		return dto.DashboardDataDto{}, err
	}

	// total number of articles
	countOfArticles, err := s.article.GetTotalNumberOfArticles()
	if err != nil {
		return dto.DashboardDataDto{}, err
	}

	// total number of tag
	countOfTags, err := s.tag.GetTotalNumberOfTags()
	if err != nil {
		return dto.DashboardDataDto{}, err
	}

	// total number of category
	countOfCategories, err := s.category.GetTotalNumberOfCategories()
	if err != nil {
		return dto.DashboardDataDto{}, err
	}

	return dto.DashboardDataDto{
		ArticleVisitsCounts: articleVisits.Amount(),
		ArticleCounts:       countOfArticles.Amount(),
		TagCounts:           countOfTags.Amount(),
		CategoryCounts:      countOfCategories.Amount(),
	}, nil
}
