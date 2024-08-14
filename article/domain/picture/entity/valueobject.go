package entity

import "github.com/victorzhou123/vicblog/common/validator"

type PictureName interface {
	PictureName() string
}

type pictureName string

func NewPictureName(v string) (PictureName, error) {
	if err := validator.IsPictureName(v); err != nil {
		return nil, err
	}

	return pictureName(v), nil
}

func (p pictureName) PictureName() string {
	return string(p)
}
