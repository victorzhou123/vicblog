package app

import "victorzhou123/vicblog/user/domain/repository"

type UserService interface {
	Login() (TokenDTO, error)
}

func NewUserService(
	userrepo repository.User,
) UserService {
	return &userService{
		userrepo: userrepo,
	}
}

type userService struct {
	userrepo repository.User
}

// Login return token if login success
func (s *userService) Login() (TokenDTO, error) {
	return TokenDTO{}, nil
}
