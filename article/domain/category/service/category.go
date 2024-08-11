package service

import (
	"victorzhou123/vicblog/article/domain/category/entity"
	"victorzhou123/vicblog/article/domain/category/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/log"
)

type CategoryService interface {
	AddCategory(entity.CategoryName) error
	ListCategory(*CategoryListCmd) (CategoryListDto, error)
	ListAllCategory() ([]CategoryDto, error)
	DelCategory(cmprimitive.Id) error

	GetRelationWithArticle(articleId cmprimitive.Id) (cateId cmprimitive.Id, err error)
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

func (s *categoryService) ListCategory(cmd *CategoryListCmd) (CategoryListDto, error) {

	cates, total, err := s.repo.GetCategoryList(cmd.toPageListOpt())
	if err != nil {
		if cmdmerror.IsNotFound(err) {
			return CategoryListDto{}, nil
		}

		return CategoryListDto{}, err
	}

	return toCategoryListDto(cates, cmd, total), nil
}

func (s *categoryService) ListAllCategory() ([]CategoryDto, error) {

	cates, err := s.repo.GetAllCategoryList()
	if err != nil {
		return nil, err
	}

	dtos := make([]CategoryDto, len(cates))
	for i := range cates {
		dtos[i] = toCategoryDto(cates[i])
	}

	return dtos, nil
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
