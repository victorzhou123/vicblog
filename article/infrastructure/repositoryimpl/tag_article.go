package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/tag/repository"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewTagArticleRepo(db mysql.Impl) repository.TagArticle {

	if err := mysql.AutoMigrate(&TagArticleDO{}); err != nil {
		return nil
	}

	return &tagArticleImpl{db}
}

type tagArticleImpl struct {
	mysql.Impl
}

func (impl *tagArticleImpl) AddRelateWithArticle(
	articleId cmprimitive.Id, tagIds []cmprimitive.Id,
) error {
	dos := make([]TagArticleDO, len(tagIds))

	for i := range tagIds {
		dos[i] = TagArticleDO{
			TagId:     tagIds[i].IdNum(),
			ArticleId: articleId.IdNum(),
		}
	}

	return impl.Impl.Add(&TagArticleDO{}, &dos)
}
