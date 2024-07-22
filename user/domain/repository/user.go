package repository

import "victorzhou123/vicblog/user/domain"

type User interface {
	GetAccountInfo() (domain.Account, error)
}
