package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/article/entity"
	"victorzhou123/vicblog/article/domain/article/repository"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewArticleRepo(db mysql.Impl, tx mysql.Transaction) repository.Article {

	if err := mysql.AutoMigrate(
		&ArticleDO{},
		&CategoryArticleDO{},
		&TagArticleDO{},
	); err != nil {
		return nil
	}

	return &articleRepoImpl{db, tx}
}

type articleRepoImpl struct {
	db mysql.Impl
	tx mysql.Transaction
}

func (impl *articleRepoImpl) GetArticles(owner cmprimitive.Username) ([]entity.Article, error) {
	articleDo := &ArticleDO{}
	articleDo.Owner = owner.Username()

	articlesDo := []ArticleDO{}

	if err := impl.db.GetRecords(&ArticleDO{}, &articleDo, &articlesDo); err != nil {
		return nil, err
	}

	// convert []ArticleDO to []domain.Article
	var err error
	dmArticles := make([]entity.Article, len(articlesDo))
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

	// transaction begin
	impl.tx.Begin()

	return impl.tx.Delete(&ArticleDO{}, &articleDo)
}

func (impl *articleRepoImpl) AddArticle(info *entity.ArticleInfo) (uint, error) {
	do := ArticleDO{
		Owner:   info.Owner.Username(),
		Title:   info.Title.Text(),
		Summary: info.Summary.ArticleSummary(),
		Content: info.Content.Text(),
		Cover:   info.Cover.Urlx(),
	}

	// transaction begin
	impl.tx.Begin()

	if err := impl.tx.Insert(&ArticleDO{}, &do); err != nil {
		return 0, err
	}

	return do.ID, nil
}
