package app

import "victorzhou123/vicblog/user/domain/repository"

type UserService interface {
	// return token if login success
	Login() (TokenDTO, error)
}

func NewUserService(
	userRepo repository.User,
) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

type userService struct {
	userRepo repository.User
}

func (s *userService) Login() (TokenDTO, error) {
	return TokenDTO{}, nil
}
