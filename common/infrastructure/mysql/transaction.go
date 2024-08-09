package mysql

import (
	"errors"

	"gorm.io/gorm"

	"victorzhou123/vicblog/common/domain/repository"
)

type Transaction interface {
	Begin()
	Commit()

	// Insert
	Insert(model, value any) error
}

type transaction struct {
	tx *gorm.DB
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
		tx: DB(),
	}
}

func (t *transaction) Begin() {
	t.tx = t.tx.Begin()
}

func (t *transaction) Commit() {
	t.tx = t.tx.Commit()
}

func (t *transaction) Insert(model, value any) error {

	err := t.tx.Model(model).Create(value).Error

	if err != nil {
		t.tx.Rollback()
	}

	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		return repository.NewErrorConstraintViolated(err)
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return repository.NewErrorDuplicateCreating(err)
	}

	return err
}
