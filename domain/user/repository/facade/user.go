package facade

import (
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/user/domain"
)

type Account struct {
	Username cmprimitive.Username
	Password domain.Password
}

type User interface {
	GetUserInfo(*Account) (domain.User, error)
}
