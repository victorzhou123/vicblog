package app

import (
	cmdmauth "github.com/victorzhou123/vicblog/common/domain/auth"
	cmdmerror "github.com/victorzhou123/vicblog/common/domain/error"
	"github.com/victorzhou123/vicblog/common/log"
	"github.com/victorzhou123/vicblog/user-server/domain/auth"
	"github.com/victorzhou123/vicblog/user-server/domain/repository"
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

		log.Errorf("user %s get user info failed, err: %s", cmd.Username.Username(), err.Error())

		if cmdmerror.IsNotFound(err) {
			return UserAndTokenDto{}, cmdmerror.New(
				cmdmerror.ErrorCodeAccessCertificateInvalid, msgUserNameOrPassWordError,
			)
		}

		return UserAndTokenDto{}, err
	}

	token, err := s.auth.GenToken(&cmdmauth.Payload{UserName: user.Username})
	if err != nil {

		log.Errorf("user %s gen token failed, err: %s", cmd.Username.Username(), err.Error())

		return UserAndTokenDto{}, err
	}

	log.Infof("user %s login success", cmd.Username.Username())

	return UserAndTokenDto{
		Username: user.Username.Username(),
		Email:    user.Email.Email(),
		Token:    token,
	}, nil
}
