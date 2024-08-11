package repository

type Transaction interface {
	Begin() error
	Commit() error
}
