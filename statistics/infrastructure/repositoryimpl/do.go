package infrastructure

import (
	"time"

	"gorm.io/gorm"

	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/statistics/domain/entity"
)

const (
	tableNameArticleDailyVisits = "article_daily_visits"

	fieldNameTotal = "total"
	fieldNameDate  = "date"
)

type ArticleDailyVisitsDO struct {
	gorm.Model

	Total int       `gorm:"column:total"`
	Date  time.Time `gorm:"colum:date;unique"`
}

func (do *ArticleDailyVisitsDO) TableName() string {
	return tableNameArticleDailyVisits
}

func (do *ArticleDailyVisitsDO) toArticleVisits() (visits entity.ArticleDailyVisits, err error) {

	if visits.Total, err = cmprimitive.NewAmount(do.Total); err != nil {
		return
	}

	visits.Date = cmprimitive.NewTimeXWithUnix(do.Date.Unix())

	return
}
