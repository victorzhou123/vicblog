package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/tag/repository"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/infrastructure/mysql"
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
