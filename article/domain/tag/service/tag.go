package service

import (
	"victorzhou123/vicblog/article/domain/tag/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type TagService interface {
	AddTags(repository.TagNames) error
	GetTagList(*TagListCmd) (TagListDto, error)
	ListAllTag() ([]TagDto, error)
	Delete(cmprimitive.Id) error

	BuildRelationWithArticle(articleId cmprimitive.Id, tagIds []cmprimitive.Id) error
}

type tagService struct {
	repo           repository.Tag
	tagArticleRepo repository.TagArticle
}

func NewTagService(
	repo repository.Tag, tagArticleRepo repository.TagArticle,
) TagService {
	return &tagService{
		repo:           repo,
		tagArticleRepo: tagArticleRepo,
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

func (s *tagService) ListAllTag() ([]TagDto, error) {

	tags, err := s.repo.GetAllTagList()
	if err != nil {
		return nil, err
	}

	dtos := make([]TagDto, len(tags))
	for i := range tags {
		dtos[i] = toTagDto(tags[i])
	}

	return dtos, nil
}

func (s *tagService) Delete(id cmprimitive.Id) error {
	return s.repo.Delete(id)
}

func (s *tagService) BuildRelationWithArticle(
	articleId cmprimitive.Id, tagIds []cmprimitive.Id,
) error {
	return s.tagArticleRepo.AddRelateWithArticle(articleId, tagIds)
}