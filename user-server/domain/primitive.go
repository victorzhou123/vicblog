package domain

import (
	"github.com/victorzhou123/vicblog/common/validator"
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
