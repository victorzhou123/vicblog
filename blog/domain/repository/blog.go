package repository

import "github.com/victorzhou123/vicblog/blog/domain/entity"

type Blog interface {
	GetBlogInfo() (entity.Blog, error)
}
