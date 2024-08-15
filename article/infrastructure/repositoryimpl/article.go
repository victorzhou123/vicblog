package repositoryimpl

import (
	"github.com/victorzhou123/vicblog/article/domain/article/entity"
	"github.com/victorzhou123/vicblog/article/domain/article/repository"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmdmerror "github.com/victorzhou123/vicblog/common/domain/error"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
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

func (impl *articleRepoImpl) GetArticleById(articleId cmprimitive.Id) (entity.Article, error) {

	do := ArticleDO{}
	do.ID = articleId.IdNum()

	if err := impl.db.GetByPrimaryKey(&ArticleDO{}, &do); err != nil {
		return entity.Article{}, err
	}

	return do.toArticle()
}

func (impl *articleRepoImpl) GetArticle(
	user cmprimitive.Username, articleId cmprimitive.Id,
) (entity.Article, error) {

	do := ArticleDO{}
	do.ID = articleId.IdNum()
	do.Owner = user.Username()

	if err := impl.db.GetRecord(&ArticleDO{}, &do, &do); err != nil {

		if cmdmerror.IsNotFound(err) {
			return entity.Article{}, cmdmerror.NewNoPermission("")
		}

		return entity.Article{}, err
	}

	return do.toArticle()
}

func (impl *articleRepoImpl) ListArticles(
	user cmprimitive.Username, opt cment.Pagination,
) ([]entity.Article, int, error) {
	articleDo := &ArticleDO{}
	articleDo.Owner = user.Username()

	articlesDo := []ArticleDO{}

	option := mysql.PaginationOpt{
		CurPage:  opt.CurPage.CurPage(),
		PageSize: opt.PageSize.PageSize(),
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

// TODO ignore content while list articles
func (impl *articleRepoImpl) ListAllArticles(opt cment.Pagination) ([]entity.Article, int, error) {

	dos := []ArticleDO{}

	option := mysql.PaginationOpt{
		CurPage:  opt.CurPage.CurPage(),
		PageSize: opt.PageSize.PageSize(),
	}

	total, err := impl.db.GetRecordsByPagination(&ArticleDO{}, &ArticleDO{}, &dos, option)
	if err != nil {
		return nil, 0, err
	}

	// convert []ArticleDO to []domain.Article
	dmArticles := make([]entity.Article, len(dos))
	for i := range dos {

		if dmArticles[i], err = dos[i].toArticle(); err != nil {
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

func (impl *articleRepoImpl) AddArticle(info *entity.ArticleInfo) (cmprimitive.Id, error) {
	do := ArticleDO{
		Owner:   info.Owner.Username(),
		Title:   info.Title.Text(),
		Summary: info.Summary.ArticleSummary(),
		Content: info.Content.Text(),
		Cover:   info.Cover.Urlx(),
	}

	if err := impl.tx.Insert(&ArticleDO{}, &do); err != nil {
		return nil, err
	}

	return cmprimitive.NewIdByUint(do.ID), nil
}

func (impl *articleRepoImpl) AddArticleReadTimes(articleId cmprimitive.Id, plusNum cmprimitive.Amount) error {

	filterDo := ArticleDO{}
	filterDo.ID = articleId.IdNum()

	return impl.db.Increase(&ArticleDO{}, &filterDo, fieldNameArticleReadTimes, plusNum.Amount())
}

func (impl *articleRepoImpl) Update(
	articleId cmprimitive.Id, articleInfo *entity.ArticleInfo,
) error {

	filterDo := ArticleDO{}
	filterDo.ID = articleId.IdNum()
	filterDo.Owner = articleInfo.Owner.Username()

	do := ArticleDO{
		Title:   articleInfo.Title.Text(),
		Content: articleInfo.Content.Text(),
		Summary: articleInfo.Summary.ArticleSummary(),
	}

	return impl.tx.Update(&ArticleDO{}, &filterDo, &do)
}
