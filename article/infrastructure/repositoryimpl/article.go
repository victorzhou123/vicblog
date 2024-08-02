package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain"
	"victorzhou123/vicblog/article/domain/repository"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewArticleRepo(db mysql.Impl) repository.Article {
	tableNameArticle = db.TableName()

	if err := mysql.AutoMigrate(&ArticleDO{}); err != nil {
		return nil
	}

	return &articleRepoImpl{db}
}

type articleRepoImpl struct {
	mysql.Impl
}

func (impl *articleRepoImpl) GetArticles(owner cmprimitive.Username) ([]domain.Article, error) {
	articleDo := &ArticleDO{}
	articleDo.Owner = owner.Username()

	articlesDo := []ArticleDO{}

	if err := impl.GetRecord(&articleDo, &articlesDo); err != nil {
		return nil, err
	}

	// convert []ArticleDO to []domain.Article
	var err error
	dmArticles := make([]domain.Article, len(articlesDo))
	for i := range articlesDo {

		if dmArticles[i], err = articlesDo[i].toArticle(); err != nil {
			return nil, err
		}

	}

	return dmArticles, nil
}

func (impl *articleRepoImpl) Delete(user cmprimitive.Username, id cmprimitive.Id) error {
	articleDo := &ArticleDO{}
	articleDo.Owner = user.Username()
	articleDo.ID = id.IdNum()

	return impl.Impl.Delete(&ArticleDO{}, &articleDo)
}
