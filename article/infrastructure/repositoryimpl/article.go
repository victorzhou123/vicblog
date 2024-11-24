package repositoryimpl

import (
	"errors"

	"gorm.io/gorm"

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

func (impl *articleRepoImpl) ListArticleCards(
	articleIds []cmprimitive.Id, opt cment.Pagination,
) ([]entity.ArticleCard, int, error) {

	filter := impl.db.InFilter(fieldNamePrimaryKeyId)

	option := mysql.PaginationOpt{
		CurPage:  opt.CurPage.CurPage(),
		PageSize: opt.PageSize.PageSize(),
	}

	ids := make([]uint, len(articleIds))
	for i := range articleIds {
		ids[i] = articleIds[i].IdNum()
	}

	dos := []ArticleCardDO{}

	total, err := impl.db.GetRecordsByPagination(ArticleDO{}, filter, &dos, option, ids)
	if err != nil {
		return nil, 0, err
	}

	articleCards := make([]entity.ArticleCard, len(dos))
	for i := range dos {
		if articleCards[i], err = dos[i].toArticleCard(); err != nil {
			return nil, 0, err
		}
	}

	return articleCards, total, nil
}

func (impl *articleRepoImpl) ListArticlesByPagination(opt cment.Pagination) ([]entity.ArticleCard, int, error) {

	option := mysql.PaginationOpt{
		CurPage:  opt.CurPage.CurPage(),
		PageSize: opt.PageSize.PageSize(),
	}

	dos := []ArticleCardDO{}

	total, err := impl.db.GetRecordsByPagination(ArticleDO{}, &ArticleCardDO{}, &dos, option)
	if err != nil {
		return nil, 0, err
	}

	articleCards := make([]entity.ArticleCard, len(dos))
	for i := range dos {
		if articleCards[i], err = dos[i].toArticleCard(); err != nil {
			return nil, 0, err
		}
	}

	return articleCards, total, nil
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

func (impl *articleRepoImpl) GetPreAndNextArticle(articleId cmprimitive.Id) (articleArr [2]*entity.ArticleIdTitle, err error) {

	preArticleDo := ArticleIdTitleDO{}
	nextArticleDo := ArticleIdTitleDO{}

	var preNotFound, nextNotFound bool

	// find next article
	err = impl.db.Model(&ArticleDO{}).Where(impl.db.GreaterQuery(fieldNamePrimaryKeyId), articleId.IdNum()).
		First(&nextArticleDo).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}

		nextNotFound = true
	}

	if !nextNotFound {
		if articleArr[0], err = nextArticleDo.toArticleIdTitle(); err != nil {
			return
		}
	}

	// find prev article
	err = impl.db.Model(&ArticleDO{}).Where(impl.db.LessQuery(fieldNamePrimaryKeyId), articleId.IdNum()).
		Last(&preArticleDo).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}

		preNotFound = true
	}

	if !preNotFound {
		if articleArr[1], err = preArticleDo.toArticleIdTitle(); err != nil {
			return
		}
	}

	return articleArr, nil
}

func (impl *articleRepoImpl) GetTotalNumberOfArticle() (cmprimitive.Amount, error) {
	var count int64
	if err := impl.db.GetTotalNumber(&ArticleDO{}, &ArticleDO{}, &count); err != nil {
		return nil, err
	}

	return cmprimitive.NewAmount(int(count))
}

func (impl *articleRepoImpl) GetRecentArticleCards(startDate cmprimitive.Timex) ([]entity.ArticleCard, error) {

	dos := []ArticleCardDO{}

	err := impl.db.GetRecords(&ArticleDO{}, impl.db.GreaterQuery(fieldCreatedAt), &dos, startDate.Time())
	if err != nil {
		if cmdmerror.IsNotFound(err) {
			return nil, nil
		}

		return nil, err
	}

	cards := make([]entity.ArticleCard, len(dos))
	for i := range dos {
		if cards[i], err = dos[i].toArticleCard(); err != nil {
			return nil, err
		}
	}

	return cards, nil
}

func (impl *articleRepoImpl) SearchArticle(word cmprimitive.Text, opt cment.Pagination) ([]entity.ArticleCardWithSummary, int, error) {

	dos := []ArticleCardWithSummaryDO{}

	filter := impl.db.LikeQuery(fieldTitle)
	option := mysql.PaginationOpt{
		CurPage:  opt.CurPage.CurPage(),
		PageSize: opt.PageSize.PageSize(),
	}

	total, err := impl.db.GetRecordsByPagination(&ArticleDO{}, filter, &dos, option, "%"+word.Text()+"%")
	if err != nil {
		if cmdmerror.IsNotFound(err) {
			return nil, 0, nil
		}

		return nil, 0, err
	}

	articlesSearch := make([]entity.ArticleCardWithSummary, len(dos))
	for i := range articlesSearch {
		if articlesSearch[i], err = dos[i].toArticleCardWithSummary(); err != nil {
			return nil, 0, err
		}
	}

	return articlesSearch, total, nil
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
		Cover:   articleInfo.Cover.Urlx(),
	}

	return impl.tx.Update(&ArticleDO{}, &filterDo, &do)
}
