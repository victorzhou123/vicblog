package mysql

import (
	"errors"

	"gorm.io/gorm"

	"victorzhou123/vicblog/common/domain/repository"
	"victorzhou123/vicblog/common/util"
)

type Transaction interface {
	Begin()
	Commit()

	// Insert
	Insert(model, value any) error
}

type transaction struct {
	tx map[int64]*gorm.DB
}

// e.g:
//
// addArticleTx := mysql.NewTransaction()
//
// articlerepo := articleRepo(addArticleTx)
//
// tagrepo := tagRepo(addArticleTx)
func NewTransaction() Transaction {
	return &transaction{
		tx: map[int64]*gorm.DB{}, // goroutine local storage
	}
}

func (t *transaction) Begin() {
	t.tx[util.GetGoroutineId()] = DB().Begin()
}

func (t *transaction) Commit() {
	t.tx[util.GetGoroutineId()].Commit()
	t.tx[util.GetGoroutineId()] = nil // clear the data
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

func (t *transaction) rollback() {
	t.txNow().Rollback()
	t.tx[util.GetGoroutineId()] = nil // clear the data
}

func (t *transaction) txNow() *gorm.DB {
	return t.tx[util.GetGoroutineId()]
}
