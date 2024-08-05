package app

import (
	"victorzhou123/vicblog/article/domain/tag/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
)

type TagService interface {
	AddTags(repository.TagNames) error
}

type tagService struct {
	repo repository.Tag
}

func NewTagService(repo repository.Tag) TagService {
	return &tagService{
		repo: repo,
	}
}

func (s *tagService) AddTags(names repository.TagNames) error {
	if !names.NoDuplication() {
		return cmdmerror.NewInvalidParam("input params contain duplicate tags")
	}

	return s.repo.AddBatches(names)
}
