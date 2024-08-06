package service

import (
	"victorzhou123/vicblog/article/domain/tag/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type TagService interface {
	AddTags(repository.TagNames) error
	GetTagList(*TagListCmd) (TagListDto, error)
	Delete(cmprimitive.Id) error
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

func (s *tagService) GetTagList(cmd *TagListCmd) (TagListDto, error) {

	tags, total, err := s.repo.GetTagList(cmd.ToPageListOpt())
	if err != nil {
		if cmdmerror.IsNotFound(err) {
			return TagListDto{}, nil
		}

		return TagListDto{}, err
	}

	return toTagListDto(tags, cmd, total), nil
}

func (s *tagService) Delete(id cmprimitive.Id) error {
	return s.repo.Delete(id)
}
