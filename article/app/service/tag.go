package service

import (
	"victorzhou123/vicblog/article/app/dto"
	"victorzhou123/vicblog/article/domain/tag/entity"
	"victorzhou123/vicblog/article/domain/tag/repository"
	tagdmsvc "victorzhou123/vicblog/article/domain/tag/service"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type TagAppService interface {
	ListAllTag() ([]dto.TagDto, error)
	ListTag(*dto.ListTagCmd) (dto.TagListDto, error)

	AddTags([]entity.TagName) error

	Delete(tagId cmprimitive.Id) error
}

type tagAppService struct {
	tag tagdmsvc.TagService
}

func NewTagAppService(tag tagdmsvc.TagService) TagAppService {
	return &tagAppService{
		tag: tag,
	}
}

func (s *tagAppService) ListAllTag() ([]dto.TagDto, error) {

	tags, err := s.tag.ListAllTag()
	if err != nil {
		return nil, err
	}

	tagDtos := make([]dto.TagDto, len(tags))
	for i := range tags {
		tagDtos[i] = dto.ToTagDto(tags[i])
	}

	return tagDtos, nil
}

func (s *tagAppService) ListTag(
	cmd *dto.ListTagCmd,
) (dto.TagListDto, error) {

	tagListDto, err := s.tag.ListTagByPagination(cmd.ToPagination())
	if err != nil {
		return dto.TagListDto{}, err
	}

	return dto.ToTagListDto(
		tagListDto.PaginationStatus, tagListDto.Tags,
	), nil
}

func (s *tagAppService) AddTags(tagNames []entity.TagName) error {
	return s.tag.AddTags(repository.TagNames{Names: tagNames})
}

func (s *tagAppService) Delete(tagId cmprimitive.Id) error {
	return s.tag.Delete(tagId)
}
