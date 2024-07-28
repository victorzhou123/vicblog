package repository

import (
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/user/domain"
)

type User interface {
	GetAccountInfo(user cmprimitive.Username) (domain.User, error)
}
