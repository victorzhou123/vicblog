package mysql

import (
	"errors"

	"gorm.io/gorm"

	"victorzhou123/vicblog/common/domain/repository"
	"victorzhou123/vicblog/common/util"
)

type Transaction interface {
	Begin() error
	Commit() error

	// Insert
	Insert(model, value any) error

	// Delete
	Delete(model, filter any) error
	SoftDelete(model, filter any) error

	// Update
	Update(model, filter, values any) error
}

type transaction struct {
	tx map[int64]*gorm.DB
}

// e.g:
//
// transactionImpl := mysql.NewTransaction()
//
// articlerepo := articleRepo(transactionImpl)
//
// tagrepo := tagRepo(transactionImpl)
func NewTransaction() Transaction {
	return &transaction{
		tx: map[int64]*gorm.DB{}, // goroutine local storage
	}
}

func (t *transaction) Begin() error {

	db := DB().Begin()

	if db.Error != nil {
		return db.Error
	}

	t.tx[util.GetGoroutineId()] = db

	return nil
}

func (t *transaction) Commit() error {
	err := t.tx[util.GetGoroutineId()].Commit().Error
	t.tx[util.GetGoroutineId()] = nil // clear the data

	return err
}

func (t *transaction) Insert(model, value any) error {

	err := t.txNow().Model(model).Create(value).Error

	if err != nil {
		t.rollback()
	}

	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		return repository.NewErrorConstraintViolated(err)
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return repository.NewErrorDuplicateCreating(err)
	}

	return err
}

func (t *transaction) Delete(model, filter any) error {

	err := t.txNow().Model(model).Unscoped().Where(filter).Delete(model).Error

	if err != nil {
		t.rollback()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (t *transaction) SoftDelete(model, filter any) error {

	err := t.txNow().Model(model).Where(filter).Delete(model).Error

	if err != nil {
		t.rollback()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	return err
}

func (t *transaction) Update(model, filter, values any) error {
	err := t.txNow().Model(model).Where(filter).Updates(values).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NewErrorResourceNotExists(errors.New("not found"))
	}

	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		return repository.NewErrorConstraintViolated(err)
	}

	return err
}

func (t *transaction) rollback() {
	t.txNow().Rollback()
	t.tx[util.GetGoroutineId()] = nil // clear the data
}

func (t *transaction) txNow() *gorm.DB {
	return t.tx[util.GetGoroutineId()]
}
