package repository

import (
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/user-server/domain"
)

type Account struct {
	Username cmprimitive.Username
	Password domain.Password
}

type User interface {
	GetUserInfo(*Account) (domain.User, error)
}
