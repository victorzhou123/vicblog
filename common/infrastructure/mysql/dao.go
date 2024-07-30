package mysql

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"victorzhou123/vicblog/common/domain/repository"
)

type Impl interface {
	GetRecord(filter, result interface{}) error
	GetByPrimaryKey(row interface{}) error
	DeleteByPrimaryKey(row interface{}) error
	EqualQuery(field string) string
	NotEqualQuery(field string) string
	OrderByDesc(field string) string
	InFilter(field string) string
	DB() *gorm.DB
	TableName() string
}

func DAO(table string) *daoImpl {
	return &daoImpl{
		table: table,
	}
}

type daoImpl struct {
	table string
}

// Each operation must generate a new gorm.DB instance.
// If using the same gorm.DB instance by different operations, they will share the same error.
func (dao *daoImpl) DB() *gorm.DB {
	return db.Table(dao.table)
}

func (dao *daoImpl) GetRecord(filter, result interface{}) error {
	err := dao.DB().Where(filter).First(result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (dao *daoImpl) GetByPrimaryKey(row interface{}) error {
	err := dao.DB().First(row).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (dao *daoImpl) DeleteByPrimaryKey(row interface{}) error {
	err := dao.DB().Delete(row).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (dao *daoImpl) EqualQuery(field string) string {
	return fmt.Sprintf(`%s = ?`, field)
}

func (dao *daoImpl) NotEqualQuery(field string) string {
	return fmt.Sprintf(`%s <> ?`, field)
}

func (dao *daoImpl) OrderByDesc(field string) string {
	return field + " desc"
}

func (dao *daoImpl) InFilter(field string) string {
	return fmt.Sprintf(`%s IN ?`, field)
}

func (dao *daoImpl) TableName() string {
	return dao.table
}
