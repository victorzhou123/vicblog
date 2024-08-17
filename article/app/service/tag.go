package service

import (
	"github.com/victorzhou123/vicblog/article/app/dto"
	"github.com/victorzhou123/vicblog/article/domain/tag/entity"
	"github.com/victorzhou123/vicblog/article/domain/tag/repository"
	tagdmsvc "github.com/victorzhou123/vicblog/article/domain/tag/service"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type TagAppService interface {
	ListTags(cmprimitive.Amount) ([]dto.TagWithRelatedArticleAmountDto, error)
	ListTagByPagination(*dto.ListTagCmd) (dto.TagListDto, error)

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

func (s *tagAppService) ListTags(amount cmprimitive.Amount,
) ([]dto.TagWithRelatedArticleAmountDto, error) {

	tags, err := s.tag.ListTags(amount)
	if err != nil {
		return nil, err
	}

	tagDtos := make([]dto.TagWithRelatedArticleAmountDto, len(tags))
	for i := range tags {
		tagDtos[i] = dto.ToTagWithRelatedArticleAmountDto(tags[i])
	}

	return tagDtos, nil
}

func (s *tagAppService) ListTagByPagination(
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
