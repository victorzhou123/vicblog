package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/tag/repository"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewTagArticleRepo(db mysql.Impl) repository.TagArticle {
	tableNameTagArticle = db.TableName()

	if err := mysql.AutoMigrate(&TagArticleDO{}); err != nil {
		return nil
	}

	return &tagArticleImpl{db}
}

type tagArticleImpl struct {
	mysql.Impl
}

func (impl *tagArticleImpl) AddRelateWithArticle(articleId, tagId cmprimitive.Id) error {
	do := TagArticleDO{
		TagId:     tagId.IdNum(),
		ArticleId: articleId.IdNum(),
	}

	return impl.Impl.Add(&do)
}
