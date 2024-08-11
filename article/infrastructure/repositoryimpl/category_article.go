package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/category/repository"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewCategoryArticleRepo(db mysql.Impl, tx mysql.Transaction) repository.CategoryArticle {

	if err := mysql.AutoMigrate(&CategoryArticleDO{}); err != nil {
		return nil
	}

	return &categoryArticleImpl{db, tx}
}

type categoryArticleImpl struct {
	db mysql.Impl
	tx mysql.Transaction
}

func (impl *categoryArticleImpl) GetRelationWithArticle(articleId cmprimitive.Id) (cmprimitive.Id, error) {

	do := CategoryArticleDO{
		ArticleId: articleId.IdNum(),
	}

	if err := impl.db.GetByPrimaryKey(&CategoryArticleDO{}, &do); err != nil {
		return nil, err
	}

	return cmprimitive.NewIdByUint(do.ID), nil
}

func (impl *categoryArticleImpl) BuildRelationWithArticle(articleId, cateId cmprimitive.Id) error {
	do := CategoryArticleDO{
		CategoryId: cateId.IdNum(),
		ArticleId:  articleId.IdNum(),
	}

	if err := impl.tx.Insert(&CategoryArticleDO{}, &do); err != nil {
		return err
	}

	return nil
}

func (impl *categoryArticleImpl) RemoveAllRowsByArticleId(articleId cmprimitive.Id) error {
	do := CategoryArticleDO{
		ArticleId: articleId.IdNum(),
	}

	if err := impl.tx.Delete(&CategoryArticleDO{}, &do); err != nil {
		return err
	}

	return nil
}
