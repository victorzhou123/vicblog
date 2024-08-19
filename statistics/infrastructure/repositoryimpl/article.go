package infrastructure

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	cmrepo "github.com/victorzhou123/vicblog/common/domain/repository"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	"github.com/victorzhou123/vicblog/statistics/domain/entity"
	"github.com/victorzhou123/vicblog/statistics/domain/repository"
)

const daysOfAWeek = 7

type timeCreator interface {
	FirstTimeOfTodayBaseDay() time.Time
}

func NewArticleVisitsRepo(db mysql.Impl, t timeCreator) repository.ArticleVisits {

	if err := mysql.AutoMigrate(
		&ArticleDailyVisitsDO{},
	); err != nil {
		return nil
	}

	return &articleVisitsRepoImpl{t, db}
}

type articleVisitsRepoImpl struct {
	timeCreator timeCreator

	db mysql.Impl
}

func (impl *articleVisitsRepoImpl) GetTotalVisits() (cmprimitive.Amount, error) {

	var total int
	err := impl.db.Model(&ArticleDailyVisitsDO{}).
		Select(fmt.Sprintf("SUM(%s) as total", fieldNameTotal)).Find(&total).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cmprimitive.NewAmount(0)
		}

		return nil, err
	}

	return cmprimitive.NewAmount(total)
}

func (impl *articleVisitsRepoImpl) IncreaseTodayVisits(amount cmprimitive.Amount) error {

	filterDo := ArticleDailyVisitsDO{
		Date: impl.timeCreator.FirstTimeOfTodayBaseDay(),
	}

	if err := impl.db.Increase(&ArticleDailyVisitsDO{}, &filterDo, fieldNameTotal, amount.Amount()); err != nil {
		if !cmrepo.IsErrorNotAffected(err) {
			return err
		}

		do := ArticleDailyVisitsDO{
			Total: 1,
			Date:  filterDo.Date,
		}

		// add a record of visits today, if function increase not affect any rows
		if err := impl.db.Add(&ArticleDailyVisitsDO{}, &do); err != nil {
			return err
		}
	}

	return nil
}

func (impl *articleVisitsRepoImpl) GetDailyVisitsOfWeek() ([]entity.ArticleDailyVisits, error) {

	dos := []ArticleDailyVisitsDO{}

	err := impl.db.Model(&ArticleDailyVisitsDO{}).Where(impl.db.LessQuery(fieldNameDate), time.Now()).
		Order(impl.db.OrderByDesc(fieldNameDate)).Limit(daysOfAWeek).Find(&dos).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	visits := make([]entity.ArticleDailyVisits, len(dos))
	for i := range dos {
		if visits[i], err = dos[i].toArticleVisits(); err != nil {
			return nil, err
		}
	}

	return visits, nil
}
