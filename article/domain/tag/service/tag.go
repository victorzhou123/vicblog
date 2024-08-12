package service

import (
	"victorzhou123/vicblog/article/domain/tag/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/log"
)

type TagService interface {
	AddTags(repository.TagNames) error
	GetArticleTag(articleId cmprimitive.Id) ([]TagDto, error)
	GetTagList(*TagListCmd) (TagListDto, error)
	ListAllTag() ([]TagDto, error)
	Delete(cmprimitive.Id) error

	GetRelationWithArticle(articleId cmprimitive.Id) (tags []cmprimitive.Id, err error)
	BuildRelationWithArticle(articleId cmprimitive.Id, tagIds []cmprimitive.Id) error
	RemoveRelationWithArticle(articleId cmprimitive.Id) error
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

		err := cmdmerror.NewInvalidParam("input params contain duplicate tags")

		log.Errorf("add tags failed, err: %s", err.Error())

		return err
	}

	if err := s.repo.AddBatches(names); err != nil {

		log.Errorf("add batches tags failed, err: %s", err.Error())

		return err
	}

	return nil
}

func (s *tagService) GetArticleTag(articleId cmprimitive.Id) ([]TagDto, error) {

	// get article relate tags
	tagIds, err := s.tagArticleRepo.GetRelationWithArticle(articleId)
	if err != nil {
		return nil, err
	}

	// get tags information
	tags, err := s.repo.GetBatchTags(tagIds)
	if err != nil {
		return nil, err
	}

	dtos := make([]TagDto, len(tags))
	for i := range tags {
		dtos[i] = toTagDto(tags[i])
	}

	return dtos, nil
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

	if err := s.repo.Delete(id); err != nil {

		log.Errorf("tag %s deleted failed, err: %s", id.Id(), err.Error())

		return err
	}

	return nil
}

func (s *tagService) GetRelationWithArticle(articleId cmprimitive.Id) ([]cmprimitive.Id, error) {
	return s.tagArticleRepo.GetRelationWithArticle(articleId)
}

func (s *tagService) BuildRelationWithArticle(
	articleId cmprimitive.Id, tagIds []cmprimitive.Id,
) error {

	if err := s.tagArticleRepo.BuildRelationWithArticle(articleId, tagIds); err != nil {

		log.Errorf("article %s build relation with tags %s failed, err: %s",
			articleId.Id(), tagIds, err.Error())

		return err
	}

	return nil
}

func (s *tagService) RemoveRelationWithArticle(articleId cmprimitive.Id) error {

	if err := s.tagArticleRepo.RemoveAllRowsByArticleId(articleId); err != nil {

		log.Errorf("remove all tags relation with article %s failed, err: %s",
			articleId.Id(), err.Error())

		return err
	}

	return nil
}
