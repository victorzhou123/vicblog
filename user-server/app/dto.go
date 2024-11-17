package app

import (
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/user-server/domain"
	"github.com/victorzhou123/vicblog/user-server/domain/repository"
)

type AccountCmd struct {
	Username cmprimitive.Username
	Password domain.Password
}

func (cmd *AccountCmd) toAccount() repository.Account {
	return repository.Account{
		Username: cmd.Username,
		Password: cmd.Password,
	}
}

type UserAndTokenDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
