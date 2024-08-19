package service

import (
	"errors"

	"github.com/victorzhou123/vicblog/article/domain/category/entity"
	"github.com/victorzhou123/vicblog/article/domain/category/repository"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmdmerror "github.com/victorzhou123/vicblog/common/domain/error"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/common/log"
)

type CategoryService interface {
	AddCategory(entity.CategoryName) error
	ListCategoryByPagination(*cment.Pagination) (CategoryListDto, error)
	ListCategories(amount cmprimitive.Amount) ([]entity.CategoryWithRelatedArticleAmount, error)
	GetArticleCategory(articleId cmprimitive.Id) (entity.Category, error)
	DelCategory(cmprimitive.Id) error

	GetRelationWithArticle(articleId cmprimitive.Id) (cateId cmprimitive.Id, err error)
	GetRelatedArticleIdsThroughCateId(cateId cmprimitive.Id) (articleIds []cmprimitive.Id, err error)
	BuildRelationWithArticle(articleId, cateId cmprimitive.Id) error
	RemoveRelationWithArticle(articleId cmprimitive.Id) error
}

type categoryService struct {
	repo                repository.Category
	categoryArticleRepo repository.CategoryArticle
}

func NewCategoryService(
	repo repository.Category,
	categoryArticleRepo repository.CategoryArticle,
) CategoryService {
	return &categoryService{
		repo:                repo,
		categoryArticleRepo: categoryArticleRepo,
	}
}

func (s *categoryService) AddCategory(category entity.CategoryName) error {

	if err := s.repo.AddCategory(category); err != nil {

		log.Errorf("add category %s failed, err: %s", category.CategoryName(), err.Error())

		return err
	}

	return nil
}

func (s *categoryService) ListCategoryByPagination(pagination *cment.Pagination) (CategoryListDto, error) {

	cates, total, err := s.repo.GetCategoryListByPagination(*pagination)
	if err != nil {
		if cmdmerror.IsNotFound(err) {
			return CategoryListDto{}, nil
		}

		return CategoryListDto{}, err
	}

	categoryIds := make([]cmprimitive.Id, len(cates))
	for i := range cates {
		categoryIds[i] = cates[i].Id
	}

	// get amount of category related article
	amountMap, err := s.categoryArticleRepo.GetRelatedArticleAmount(categoryIds)
	if err != nil {
		return CategoryListDto{}, err
	}

	cateWithAmounts := make([]entity.CategoryWithRelatedArticleAmount, len(cates))
	for i := range cates {

		v, ok := amountMap[cates[i].Id.IdNum()]
		if !ok {
			v, _ = cmprimitive.NewAmount(0) // set a default amount
		}

		cateWithAmounts[i] = entity.CategoryWithRelatedArticleAmount{
			Category:             cates[i],
			RelatedArticleAmount: v,
		}
	}

	return CategoryListDto{
		PaginationStatus: pagination.ToPaginationStatus(total),
		Categories:       cateWithAmounts,
	}, nil
}

func (s *categoryService) ListCategories(amount cmprimitive.Amount) ([]entity.CategoryWithRelatedArticleAmount, error) {

	cates, err := s.repo.GetCategoryList(amount)
	if err != nil {
		return nil, err
	}

	categoryIds := make([]cmprimitive.Id, len(cates))
	for i := range cates {
		categoryIds[i] = cates[i].Id
	}

	amountMap, err := s.categoryArticleRepo.GetRelatedArticleAmount(categoryIds)
	if err != nil {
		return nil, err
	}

	cateWithAmounts := make([]entity.CategoryWithRelatedArticleAmount, len(cates))
	for i := range cates {

		v, ok := amountMap[cates[i].Id.IdNum()]
		if !ok {
			v, _ = cmprimitive.NewAmount(0) // set a default amount
		}

		cateWithAmounts[i] = entity.CategoryWithRelatedArticleAmount{
			Category:             cates[i],
			RelatedArticleAmount: v,
		}
	}

	return cateWithAmounts, nil
}

func (s *categoryService) GetArticleCategory(articleId cmprimitive.Id) (entity.Category, error) {

	// get article relate category
	cateId, err := s.categoryArticleRepo.GetRelationWithArticle(articleId)
	if err != nil {
		return entity.Category{}, err
	}

	return s.repo.GetCategory(cateId)
}

func (s *categoryService) DelCategory(id cmprimitive.Id) error {

	if err := s.repo.DelCategory(id); err != nil {

		log.Errorf("delete category %s failed, err: %s", id.Id(), err.Error())

		return err
	}

	return nil
}

func (s *categoryService) GetRelationWithArticle(
	articleId cmprimitive.Id,
) (cmprimitive.Id, error) {

	cateId, err := s.categoryArticleRepo.GetRelationWithArticle(articleId)
	if err != nil {

		log.Errorf("get all category related to article %s failed, err: %s",
			articleId.Id(), err.Error())

		return nil, err
	}

	return cateId, nil
}

func (s *categoryService) GetRelatedArticleIdsThroughCateId(
	cateId cmprimitive.Id) ([]cmprimitive.Id, error) {

	if cateId == nil {
		return nil, errors.New("category id must be provided")
	}

	return s.categoryArticleRepo.GetRelatedArticleIdsThoroughCateId(cateId)
}

func (s *categoryService) BuildRelationWithArticle(articleId, cateId cmprimitive.Id) error {

	if err := s.categoryArticleRepo.BuildRelationWithArticle(articleId, cateId); err != nil {

		log.Errorf("article %s build relation with category %s failed, err: %s",
			articleId.Id(), cateId.Id(), err.Error())

		return err
	}

	return nil
}

func (s *categoryService) RemoveRelationWithArticle(articleId cmprimitive.Id) error {

	if err := s.categoryArticleRepo.RemoveAllRowsByArticleId(articleId); err != nil {

		log.Errorf("remove all category relation with article %s failed, err: %s",
			articleId.Id(), err.Error())

		return err
	}

	return nil
}
