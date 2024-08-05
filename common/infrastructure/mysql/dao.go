package mysql

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"victorzhou123/vicblog/common/domain/repository"
)

type Impl interface {
	DB() *gorm.DB
	TableName() string

	// Query
	GetRecord(filter, result interface{}) error
	GetRecordByPagination(filter, result interface{}, opt PaginationOpt) (int, error)
	GetByPrimaryKey(row interface{}) error

	// Add
	Add(value interface{}) error

	// Update
	Update(filter, values interface{}) error

	// Delete
	Delete(model, filter interface{}) error
	DeleteByPrimaryKey(row interface{}) error

	// util interface
	EqualQuery(field string) string
	NotEqualQuery(field string) string
	OrderByDesc(field string) string
	InFilter(field string) string
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

func (dao *daoImpl) GetRecordByPagination(filter, result interface{}, opt PaginationOpt) (int, error) {
	var total int64

	err := dao.DB().Where(filter).Count(&total).Offset((opt.CurPage - 1) * opt.PageSize).Limit(opt.PageSize).Find(result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return int(total), err
}

func (dao *daoImpl) GetByPrimaryKey(row interface{}) error {
	err := dao.DB().First(row).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (dao *daoImpl) Add(value interface{}) error {
	err := dao.DB().Create(value).Error
	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		return repository.NewErrorConstraintViolated(err)
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return repository.NewErrorDuplicateCreating(err)
	}

	return err
}

func (dao *daoImpl) Update(filter, values interface{}) error {
	err := dao.DB().Where(filter).Updates(values).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		return repository.NewErrorConstraintViolated(err)
	}

	return err
}

func (dao *daoImpl) Delete(model, filter interface{}) error {
	err := dao.DB().Unscoped().Delete(model, filter).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (dao *daoImpl) DeleteByPrimaryKey(row interface{}) error {
	err := dao.DB().Unscoped().Delete(row).Error

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
