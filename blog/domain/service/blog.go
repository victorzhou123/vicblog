package service

import (
	"victorzhou123/vicblog/blog/domain/entity"
	"victorzhou123/vicblog/blog/domain/repository"
)

type BlogService interface {
	GetBlogInformation() (entity.Blog, error)
}

type blogService struct {
	repo repository.Blog
}

func NewBlogService(repo repository.Blog) BlogService {
	return &blogService{
		repo: repo,
	}
}

func (s *blogService) GetBlogInformation() (entity.Blog, error) {
	return s.repo.GetBlogInfo()
}
