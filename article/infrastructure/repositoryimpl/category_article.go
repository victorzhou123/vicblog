package repositoryimpl

import (
	"fmt"

	"github.com/victorzhou123/vicblog/article/domain/category/repository"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
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

func (impl *categoryArticleImpl) GetRelatedArticleAmount(categoryIds []cmprimitive.Id) (map[uint]cmprimitive.Amount, error) {

	ids := make([]uint, len(categoryIds))
	for i := range categoryIds {
		ids[i] = categoryIds[i].IdNum()
	}

	type Result struct {
		CategoryId uint  `gorm:"column:category_id"`
		Count      int64 `gorm:"column:count"`
	}
	dos := []Result{}

	err := impl.db.Model(&CategoryArticleDO{}).
		Select(fmt.Sprintf("%s, COUNT(*) as count", fieldNameCategoryId)).
		Group(fieldNameCategoryId).Having(impl.db.InFilter(fieldNameCategoryId), ids).Find(&dos).Error
	if err != nil {
		return nil, err
	}

	// to map
	m := make(map[uint]cmprimitive.Amount, 0)
	for i := range dos {

		amount, err := cmprimitive.NewAmount(int(dos[i].Count))
		if err != nil {
			return nil, err
		}

		m[dos[i].CategoryId] = amount
	}

	return m, nil
}

func (impl *categoryArticleImpl) BuildRelationWithArticle(articleId, cateId cmprimitive.Id) error {
	do := CategoryArticleDO{
		CategoryId: cateId.IdNum(),
		ArticleId:  articleId.IdNum(),
	}

	return impl.tx.Insert(&CategoryArticleDO{}, &do)
}

func (impl *categoryArticleImpl) RemoveAllRowsByArticleId(articleId cmprimitive.Id) error {
	do := CategoryArticleDO{
		ArticleId: articleId.IdNum(),
	}

	return impl.tx.Delete(&CategoryArticleDO{}, &do)
}
