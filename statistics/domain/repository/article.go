package repository

import (
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/statistics/domain/entity"
)

type ArticleVisits interface {
	IncreaseTodayVisits(cmprimitive.Amount) error

	GetDailyVisitsOfWeek() ([]entity.ArticleDailyVisits, error)
	GetTotalVisits() (cmprimitive.Amount, error)
}
