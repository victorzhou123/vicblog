package mysql

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/victorzhou123/vicblog/common/domain/repository"
)

type Impl interface {
	Model(model any) *gorm.DB

	// Get
	GetRecord(model, filter, result any) error
	GetRecords(model, filter, result any) error
	GetLimitRecords(model, filter, result any, amount int) error
	GetRecordsByPagination(model, filter, result any, opt PaginationOpt, filterArgs ...any) (total int, err error)
	GetByPrimaryKey(model, row any) error

	// Add
	Add(model, value any) error

	// Update
	Update(model, filter, values any) error
	Increase(model, filter any, column string, increaseNum int) error

	// Delete
	Delete(model, filter any) error
	DeleteByPrimaryKey(model, row any) error

	// Helper
	EqualQuery(field string) string
	NotEqualQuery(field string) string
	GreaterQuery(field string) string
	LessQuery(field string) string
	OrderByDesc(field string) string
	InFilter(field string) string
}

func DAO() *daoImpl {
	return &daoImpl{}
}

type daoImpl struct{}

// Each operation must generate a new gorm.DB instance.
// If using the same gorm.DB instance by different operations, they will share the same error.
func (dao *daoImpl) Model(model any) *gorm.DB {
	return db.Model(model)
}

func (dao *daoImpl) GetRecord(model, filter, result any) error {
	err := dao.Model(model).Where(filter).First(result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (dao *daoImpl) GetRecords(model, filter, result any) error {
	err := dao.Model(model).Where(filter).Find(result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (dao *daoImpl) GetLimitRecords(model, filter, result any, amount int) error {

	err := dao.Model(model).Where(filter).Limit(amount).Find(result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (dao *daoImpl) GetRecordsByPagination(model, filter, result any, opt PaginationOpt, filterArgs ...any) (int, error) {
	var total int64

	err := dao.Model(model).Where(filter, filterArgs...).Count(&total).Offset((opt.CurPage - 1) * opt.PageSize).Limit(opt.PageSize).Find(result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return int(total), err
}

func (dao *daoImpl) GetByPrimaryKey(model, row any) error {
	err := dao.Model(model).First(row).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (dao *daoImpl) Add(model, value any) error {

	err := dao.Model(model).Create(value).Error

	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		return repository.NewErrorConstraintViolated(err)
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return repository.NewErrorDuplicateCreating(err)
	}

	return err
}

func (dao *daoImpl) Update(model, filter, values any) error {
	err := dao.Model(model).Where(filter).Updates(values).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		return repository.NewErrorConstraintViolated(err)
	}

	return err
}

func (dao *daoImpl) Increase(model, filter any, column string, increaseNum int) error {

	err := dao.Model(model).Where(filter).
		UpdateColumn(column, gorm.Expr(fmt.Sprintf("%s + ?", column), increaseNum)).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		return repository.NewErrorConstraintViolated(err)
	}

	return err
}

func (dao *daoImpl) Delete(model, filter any) error {
	err := dao.Model(model).Unscoped().Delete(model, filter).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (dao *daoImpl) DeleteByPrimaryKey(model, row any) error {
	err := dao.Model(model).Unscoped().Delete(row).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

// helper
func (dao *daoImpl) EqualQuery(field string) string {
	return fmt.Sprintf(`%s = ?`, field)
}

func (dao *daoImpl) NotEqualQuery(field string) string {
	return fmt.Sprintf(`%s <> ?`, field)
}

func (dao *daoImpl) GreaterQuery(field string) string {
	return fmt.Sprintf(`%s > ?`, field)
}

func (dao *daoImpl) LessQuery(field string) string {
	return fmt.Sprintf(`%s < ?`, field)
}

func (dao *daoImpl) OrderByDesc(field string) string {
	return field + " desc"
}

func (dao *daoImpl) InFilter(field string) string {
	return fmt.Sprintf(`%s IN ?`, field)
}
