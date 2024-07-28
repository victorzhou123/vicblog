package domain

import (
	"victorzhou123/vicblog/common/validator"
)

// password
type password string

type Password interface {
	Password() string
}

func NewPassword(v string) (Password, error) {
	if err := validator.IsPassword(v); err != nil {
		return nil, err
	}

	return password(v), nil
}

func (e password) Password() string {
	return string(e)
}

// email
type email string

type Email interface {
	Email() string
}

func NewEmail(v string) (Email, error) {
	if err := validator.IsEmail(v); err != nil {
		return nil, err
	}

	return email(v), nil
}

func (e email) Email() string {
	return string(e)
}
