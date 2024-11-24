package service

import (
	"errors"

	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmdmerror "github.com/victorzhou123/vicblog/common/domain/error"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/common/log"
	"github.com/victorzhou123/vicblog/tag-server/domain/tag/entity"
	"github.com/victorzhou123/vicblog/tag-server/domain/tag/repository"
)

type TagService interface {
	AddTags(repository.TagNames) error
	GetArticleTag(articleId cmprimitive.Id) ([]entity.Tag, error)
	ListTagByPagination(*cment.Pagination) (TagListDto, error)
	ListTags(cmprimitive.Amount) ([]entity.TagWithRelatedArticleAmount, error)
	GetTotalNumberOfTags() (cmprimitive.Amount, error)
	Delete(cmprimitive.Id) error

	GetRelationWithArticle(articleId cmprimitive.Id) (tags []cmprimitive.Id, err error)
	GetRelatedArticleIdsThroughTagId(tagId cmprimitive.Id) (tagIds []cmprimitive.Id, err error)
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

func (s *tagService) GetArticleTag(articleId cmprimitive.Id) ([]entity.Tag, error) {

	// get article relate tags
	tagIds, err := s.tagArticleRepo.GetRelationWithArticle(articleId)
	if err != nil {
		return nil, err
	}

	return s.repo.GetBatchTags(tagIds)
}

func (s *tagService) ListTagByPagination(pagination *cment.Pagination) (TagListDto, error) {

	tags, total, err := s.repo.GetTagListByPagination(*pagination)
	if err != nil {
		if cmdmerror.IsNotFound(err) {
			return TagListDto{}, nil
		}

		return TagListDto{}, err
	}

	// get amount of tag related article
	tagIds := make([]cmprimitive.Id, len(tags))
	for i := range tags {
		tagIds[i] = tags[i].Id
	}

	amountMap, err := s.tagArticleRepo.GetRelatedArticleAmount(tagIds)
	if err != nil {
		return TagListDto{}, err
	}

	tagWithAmounts := make([]entity.TagWithRelatedArticleAmount, len(tags))
	for i := range tags {

		v, ok := amountMap[tags[i].Id.IdNum()]
		if !ok {
			v, _ = cmprimitive.NewAmount(0) // set a default amount
		}

		tagWithAmounts[i] = entity.TagWithRelatedArticleAmount{
			Tag:                  tags[i],
			RelatedArticleAmount: v,
		}
	}

	return TagListDto{
		PaginationStatus: pagination.ToPaginationStatus(total),
		Tags:             tagWithAmounts,
	}, nil
}

func (s *tagService) ListTags(amount cmprimitive.Amount,
) ([]entity.TagWithRelatedArticleAmount, error) {

	tags, err := s.repo.GetTagList(amount)
	if err != nil {
		return nil, err
	}

	tagIds := make([]cmprimitive.Id, len(tags))
	for i := range tags {
		tagIds[i] = tags[i].Id
	}

	amountMap, err := s.tagArticleRepo.GetRelatedArticleAmount(tagIds)
	if err != nil {
		return nil, err
	}

	tagWithAmounts := make([]entity.TagWithRelatedArticleAmount, len(tags))
	for i := range tags {

		amount, ok := amountMap[tags[i].Id.IdNum()]
		if !ok {
			amount, _ = cmprimitive.NewAmount(0)
		}

		tagWithAmounts[i] = entity.TagWithRelatedArticleAmount{
			Tag:                  tags[i],
			RelatedArticleAmount: amount,
		}
	}

	return tagWithAmounts, nil
}

func (s *tagService) GetTotalNumberOfTags() (cmprimitive.Amount, error) {
	return s.repo.GetTotalNumberOfTag()
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

func (s *tagService) GetRelatedArticleIdsThroughTagId(
	tagId cmprimitive.Id) (tagIds []cmprimitive.Id, err error) {

	if tagId == nil {
		return nil, errors.New("tag id must be provided")
	}

	return s.tagArticleRepo.GetRelatedArticleIdsThoroughTagId(tagId)

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
