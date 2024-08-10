package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/tag/repository"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewTagArticleRepo(tx mysql.Transaction) repository.TagArticle {

	if err := mysql.AutoMigrate(&TagArticleDO{}); err != nil {
		return nil
	}

	return &tagArticleImpl{tx}
}

type tagArticleImpl struct {
	tx mysql.Transaction
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
