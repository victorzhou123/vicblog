package repositoryimpl

import (
	"fmt"

	"github.com/victorzhou123/vicblog/article/domain/tag/repository"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewTagArticleRepo(db mysql.Impl, tx mysql.Transaction) repository.TagArticle {

	if err := mysql.AutoMigrate(&TagArticleDO{}); err != nil {
		return nil
	}

	return &tagArticleImpl{db, tx}
}

type tagArticleImpl struct {
	db mysql.Impl
	tx mysql.Transaction
}

func (impl *tagArticleImpl) GetRelationWithArticle(articleId cmprimitive.Id) ([]cmprimitive.Id, error) {

	filterDo := TagArticleDO{}
	filterDo.ArticleId = articleId.IdNum()

	dos := []TagArticleDO{}

	if err := impl.db.GetRecords(&TagArticleDO{}, &filterDo, &dos); err != nil {
		return nil, err
	}

	tagIds := make([]cmprimitive.Id, len(dos))
	for i := range dos {
		tagIds[i] = cmprimitive.NewIdByUint(dos[i].TagId)
	}

	return tagIds, nil
}

func (impl *tagArticleImpl) GetRelatedArticleAmount(tagIds []cmprimitive.Id) (map[uint]cmprimitive.Amount, error) {

	ids := make([]uint, len(tagIds))
	for i := range tagIds {
		ids[i] = tagIds[i].IdNum()
	}

	type Result struct {
		TagId uint  `gorm:"column:tag_id"`
		Count int64 `gorm:"column:count"`
	}
	dos := []Result{}

	err := impl.db.Model(&TagArticleDO{}).
		Select(fmt.Sprintf("%s, COUNT(*) as count", fieldNameTagId)).
		Group(fieldNameTagId).Having(impl.db.InFilter(fieldNameTagId), ids).Find(&dos).Error
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

		m[dos[i].TagId] = amount
	}

	return m, nil
}

func (impl *tagArticleImpl) BuildRelationWithArticle(
	articleId cmprimitive.Id, tagIds []cmprimitive.Id,
) error {
	dos := make([]TagArticleDO, len(tagIds))

	for i := range tagIds {
		dos[i] = TagArticleDO{
			TagId:     tagIds[i].IdNum(),
			ArticleId: articleId.IdNum(),
		}
	}

	return impl.tx.Insert(&TagArticleDO{}, &dos)
}

func (impl *tagArticleImpl) RemoveAllRowsByArticleId(articleId cmprimitive.Id) error {
	do := TagArticleDO{
		ArticleId: articleId.IdNum(),
	}

	return impl.tx.Delete(&TagArticleDO{}, &do)
}
