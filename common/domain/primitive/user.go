package primitive

import "github.com/victorzhou123/vicblog/common/validator"

// username
type username string

type Username interface {
	Username() string
}

func NewUsername(v string) (Username, error) {
	if err := validator.IsUsername(v); err != nil {
		return nil, err
	}

	return username(v), nil
}

func (e username) Username() string {
	return string(e)
}
