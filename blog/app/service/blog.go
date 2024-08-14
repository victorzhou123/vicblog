package service

import (
	"github.com/victorzhou123/vicblog/blog/app/dto"
	"github.com/victorzhou123/vicblog/blog/domain/service"
)

type BlogAppService interface {
	GetBlogInformation() (dto.BlogInformationDto, error)
}

type blogAppService struct {
	blog service.BlogService
}

func NewBlogAppService(blog service.BlogService) BlogAppService {
	return &blogAppService{
		blog: blog,
	}
}

func (s *blogAppService) GetBlogInformation() (dto.BlogInformationDto, error) {

	blog, err := s.blog.GetBlogInformation()
	if err != nil {
		return dto.BlogInformationDto{}, err
	}

	return dto.ToBlogInformationDto(blog), nil
}
