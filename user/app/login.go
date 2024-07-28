package app

import (
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	"victorzhou123/vicblog/user/domain/auth"
	"victorzhou123/vicblog/user/domain/repository"
)

type LoginService interface {
	Login(*AccountCmd) (UserAndTokenDto, error)
}

func NewLoginService(
	userrepo repository.User,
	auth auth.Auth,
) LoginService {
	return &loginService{
		userrepo: userrepo,
		auth:     auth,
	}
}

type loginService struct {
	userrepo repository.User
	auth     auth.Auth
}

// Login return token if login success
func (s *loginService) Login(cmd *AccountCmd) (UserAndTokenDto, error) {
	account, err := s.userrepo.GetAccountInfo(cmd.Username)
	if err != nil {
		return UserAndTokenDto{}, cmdmerror.New(
			cmdmerror.ErrorCodeAccessCertificateInvalid, msgUserNameOrPassWordError,
		)
	}

	token := s.auth.GenToken(&auth.JWTPayload{UserName: account.Username})

	return UserAndTokenDto{
		Username: account.Username.Username(),
		Email:    account.Email.Email(),
		Token:    token,
	}, nil
}
