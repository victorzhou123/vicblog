package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/category/repository"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewCategoryArticleRepo(tx mysql.Transaction) repository.CategoryArticle {

	if err := mysql.AutoMigrate(&CategoryArticleDO{}); err != nil {
		return nil
	}

	return &categoryArticleImpl{tx}
}

type categoryArticleImpl struct {
	tx mysql.Transaction
}

func (impl *categoryArticleImpl) BindCategoryAndArticle(articleId, cateId cmprimitive.Id) error {
	do := CategoryArticleDO{
		CategoryId: cateId.IdNum(),
		ArticleId:  articleId.IdNum(),
	}

	if err := impl.tx.Insert(&CategoryArticleDO{}, &do); err != nil {
		return err
	}

	// transaction commit
	impl.tx.Commit()

	return nil
}
