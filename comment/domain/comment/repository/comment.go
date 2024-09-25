package repository

import "github.com/victorzhou123/vicblog/comment/domain/comment/entity"

type Comment interface {
	Add(entity.Comment) error
}
