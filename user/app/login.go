package app

import (
	cmdmauth "victorzhou123/vicblog/common/domain/auth"
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
	account := cmd.toAccount()

	user, err := s.userrepo.GetUserInfo(&account)
	if err != nil {
		return UserAndTokenDto{}, cmdmerror.New(
			cmdmerror.ErrorCodeAccessCertificateInvalid, msgUserNameOrPassWordError,
		)
	}

	token, err := s.auth.GenToken(&cmdmauth.Payload{UserName: user.Username})
	if err != nil {
		return UserAndTokenDto{}, err
	}

	return UserAndTokenDto{
		Username: user.Username.Username(),
		Email:    user.Email.Email(),
		Token:    token,
	}, nil
}
