package repository

import "github.com/victorzhou123/vicblog/article/domain/picture/entity"

type Picture interface {
	Upload(username, pictureName string, file entity.Picture) (url string, err error)
}
