package app

import (
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/user/domain"
)

type AccountCmd struct {
	Username cmprimitive.Username
	Password domain.Password
}

type UserAndTokenDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
