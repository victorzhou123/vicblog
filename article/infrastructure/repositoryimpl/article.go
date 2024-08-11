package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/article/entity"
	"victorzhou123/vicblog/article/domain/article/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	cmrepo "victorzhou123/vicblog/common/domain/repository"
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

func (impl *articleRepoImpl) GetArticle(
	user cmprimitive.Username, articleId cmprimitive.Id,
) (entity.Article, error) {

	do := ArticleDO{}
	do.ID = articleId.IdNum()
	do.Owner = user.Username()

	if err := impl.db.GetByPrimaryKey(&ArticleDO{}, &do); err != nil {

		if cmdmerror.IsNotFound(err) {
			return entity.Article{}, cmdmerror.NewNoPermission("")
		}

		return entity.Article{}, err
	}

	return do.toArticle()
}

func (impl *articleRepoImpl) ListArticles(
	user cmprimitive.Username, opt cmrepo.PageListOpt,
) ([]entity.Article, int, error) {
	articleDo := &ArticleDO{}
	articleDo.Owner = user.Username()

	articlesDo := []ArticleDO{}

	option := mysql.PaginationOpt{
		CurPage:  opt.CurPage,
		PageSize: opt.PageSize,
	}

	total, err := impl.db.GetRecordsByPagination(&ArticleDO{}, &articleDo, &articlesDo, option)
	if err != nil {
		return nil, 0, err
	}

	// convert []ArticleDO to []domain.Article
	dmArticles := make([]entity.Article, len(articlesDo))
	for i := range articlesDo {

		if dmArticles[i], err = articlesDo[i].toArticle(); err != nil {
			return nil, 0, err
		}

	}

	return dmArticles, total, nil
}

func (impl *articleRepoImpl) Delete(user cmprimitive.Username, id cmprimitive.Id) error {
	articleDo := &ArticleDO{}
	articleDo.Owner = user.Username()
	articleDo.ID = id.IdNum()

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

	if err := impl.tx.Insert(&ArticleDO{}, &do); err != nil {
		return 0, err
	}

	return do.ID, nil
}

func (impl *articleRepoImpl) Update(article *entity.ArticleUpdate) error {

	filterDo := ArticleDO{}
	filterDo.ID = article.Id.IdNum()
	filterDo.Owner = article.Owner.Username()

	do := ArticleDO{
		Title:   article.Title.Text(),
		Content: article.Content.Text(),
		Summary: article.Summary.ArticleSummary(),
	}

	return impl.tx.Update(&ArticleDO{}, &filterDo, &do)
}
