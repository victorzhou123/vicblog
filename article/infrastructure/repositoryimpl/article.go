package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/article/entity"
	"victorzhou123/vicblog/article/domain/article/repository"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewArticleRepo(db mysql.Impl) repository.Article {

	if err := mysql.AutoMigrate(&ArticleDO{}); err != nil {
		return nil
	}

	return &articleRepoImpl{db}
}

type articleRepoImpl struct {
	mysql.Impl
}

func (impl *articleRepoImpl) GetArticles(owner cmprimitive.Username) ([]entity.Article, error) {
	articleDo := &ArticleDO{}
	articleDo.Owner = owner.Username()

	articlesDo := []ArticleDO{}

	if err := impl.GetRecords(&ArticleDO{}, &articleDo, &articlesDo); err != nil {
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

	return impl.Impl.Delete(&ArticleDO{}, &articleDo)
}

func (impl *articleRepoImpl) AddArticle(info *entity.ArticleWithCateAndTagInfo) error {
	articleDo := ArticleDO{
		Owner:   info.Article.Owner.Username(),
		Title:   info.Article.Title.Text(),
		Summary: info.Article.Summary.ArticleSummary(),
		Content: info.Article.Content.Text(),
		Cover:   info.Article.Cover.Urlx(),
	}

	// transaction begin
	tx := impl.Impl.Begin()

	// add article
	if err := impl.Impl.TxAdd(tx, &ArticleDO{}, &articleDo); err != nil {
		return err
	}

	// bind category
	cateArticleDo := CategoryArticleDO{
		CategoryId: info.Category.IdNum(),
		ArticleId:  articleDo.ID,
	}

	if err := impl.Impl.TxAdd(tx, &CategoryArticleDO{}, &cateArticleDo); err != nil {
		return err
	}

	// bind tags
	tagArticleDos := make([]TagArticleDO, len(info.Tags))
	for i := range info.Tags {
		tagArticleDos[i] = TagArticleDO{
			TagId:     info.Tags[i].IdNum(),
			ArticleId: articleDo.ID,
		}
	}

	if err := impl.Impl.TxAdd(tx, &TagArticleDO{}, &tagArticleDos); err != nil {
		return err
	}

	// transaction commit
	tx.Commit()

	return nil
}
