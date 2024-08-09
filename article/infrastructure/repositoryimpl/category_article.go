package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/category/repository"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewCategoryArticleRepo(db mysql.Impl) repository.CategoryArticle {

	if err := mysql.AutoMigrate(&CategoryArticleDO{}); err != nil {
		return nil
	}

	return &categoryArticleImpl{db}
}

type categoryArticleImpl struct {
	mysql.Impl
}

func (impl *categoryArticleImpl) BindCategoryAndArticle(articleId, cateId cmprimitive.Id) error {
	do := CategoryArticleDO{
		CategoryId: cateId.IdNum(),
		ArticleId:  articleId.IdNum(),
	}

	return impl.Impl.Add(&CategoryArticleDO{}, &do)
}
