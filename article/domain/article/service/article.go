package service

import (
	"github.com/victorzhou123/vicblog/article/domain/article/entity"
	"github.com/victorzhou123/vicblog/article/domain/article/repository"
	cmentt "github.com/victorzhou123/vicblog/common/domain/entity"
	cmdmerror "github.com/victorzhou123/vicblog/common/domain/error"
	cmdmmd2html "github.com/victorzhou123/vicblog/common/domain/md2html"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/common/log"
)

const msgCannotFoundTheArticle = "can not found the article"

type ArticleService interface {
	GetArticleByIdWithContentParsed(articleId cmprimitive.Id) (entity.Article, error)
	GetArticle(*GetArticleCmd) (entity.Article, error)
	GetArticleList(*ArticleListCmd) (ArticleListDto, error)
	GetArticleCardList(*ArticleCardsCmd) (ArticleCardsDto, error)
	PaginationListArticle(*cmentt.Pagination) (ArticleListDto, error)
	ListArticlesClassifiedByMonth(*cmentt.Pagination) (ArticleListClassifyByMonthDto, error)
	GetPrevAndNextArticle(articleId cmprimitive.Id) (ArticlePrevAndNextDto, error)
	GetTotalNumberOfArticles() (cmprimitive.Amount, error)

	Delete(cmprimitive.Username, cmprimitive.Id) error

	AddArticle(*entity.ArticleInfo) (articleId cmprimitive.Id, err error)
	AddArticleReadTimes(articleId cmprimitive.Id, plus cmprimitive.Amount) error

	UpdateArticle(articleId cmprimitive.Id, articleInfo *entity.ArticleInfo) error
}

type articleService struct {
	repo repository.Article
	m2h  cmdmmd2html.Md2Html
}

func NewArticleService(repo repository.Article, m2h cmdmmd2html.Md2Html) ArticleService {
	return &articleService{
		repo: repo,
		m2h:  m2h,
	}
}

func (s *articleService) GetArticleByIdWithContentParsed(articleId cmprimitive.Id) (entity.Article, error) {

	article, err := s.repo.GetArticleById(articleId)
	if err != nil {
		return entity.Article{}, err
	}

	// parse md to html
	article.Content = s.m2h.Render(article.Content)

	return article, nil
}

func (s *articleService) GetArticle(cmd *GetArticleCmd) (entity.Article, error) {

	// get article
	article, err := s.repo.GetArticle(cmd.User, cmd.ArticleId)
	if err != nil {

		log.Errorf("user %s get article %s details failed, err: %s",
			cmd.User.Username(), cmd.ArticleId.Id(), err.Error())

		return entity.Article{}, err
	}

	return article, nil
}

func (s *articleService) GetArticleList(cmd *ArticleListCmd) (ArticleListDto, error) {

	articles, total, err := s.repo.ListArticles(cmd.User, cmd.Pagination)
	if err != nil {
		return ArticleListDto{}, cmdmerror.New(
			cmdmerror.ErrorCodeResourceNotFound, msgCannotFoundTheArticle,
		)
	}

	return ArticleListDto{
		PaginationStatus: cmd.ToPaginationStatus(total),
		Articles:         articles,
	}, nil
}

func (s *articleService) GetArticleCardList(cmd *ArticleCardsCmd) (ArticleCardsDto, error) {

	if err := cmd.validate(); err != nil {
		return ArticleCardsDto{}, err
	}

	articleCards, total, err := s.repo.ListArticleCards(cmd.ArticleIds, cmd.Pagination)
	if err != nil {
		return ArticleCardsDto{}, err
	}

	return ArticleCardsDto{
		PaginationStatus: cmd.ToPaginationStatus(total),
		ArticleCards:     articleCards,
	}, nil
}

func (s *articleService) PaginationListArticle(pagination *cmentt.Pagination) (ArticleListDto, error) {

	articles, total, err := s.repo.ListAllArticles(*pagination)
	if err != nil {
		if cmdmerror.IsNotFound(err) {
			return ArticleListDto{}, cmdmerror.New(
				cmdmerror.ErrorCodeResourceNotFound, msgCannotFoundTheArticle,
			)
		}

		return ArticleListDto{}, err
	}

	return ArticleListDto{
		PaginationStatus: pagination.ToPaginationStatus(total),
		Articles:         articles,
	}, nil
}

func (s *articleService) ListArticlesClassifiedByMonth(pagination *cmentt.Pagination) (ArticleListClassifyByMonthDto, error) {

	articleCards, total, err := s.repo.ListArticlesByPagination(*pagination)
	if err != nil {
		return ArticleListClassifyByMonthDto{}, err
	}

	return toArticleListClassifyByMonthDto(articleCards, pagination, total), nil
}

func (s *articleService) GetPrevAndNextArticle(articleId cmprimitive.Id) (ArticlePrevAndNextDto, error) {

	articles, err := s.repo.GetPreAndNextArticle(articleId)
	if err != nil {
		return ArticlePrevAndNextDto{}, err
	}

	return ArticlePrevAndNextDto{
		Prev: articles[0],
		Next: articles[1],
	}, nil
}

func (s *articleService) GetTotalNumberOfArticles() (cmprimitive.Amount, error) {
	return s.repo.GetTotalNumberOfArticle()
}

func (s *articleService) Delete(user cmprimitive.Username, id cmprimitive.Id) error {
	if err := s.repo.Delete(user, id); err != nil {

		log.Errorf("user %s delete article %s failed, err: %s", user.Username(), id.Id(), err.Error())

		return cmdmerror.NewNoPermission("")
	}

	return nil
}

func (s *articleService) AddArticle(articleInfo *entity.ArticleInfo) (cmprimitive.Id, error) {

	articleId, err := s.repo.AddArticle(&entity.ArticleInfo{
		Owner:   articleInfo.Owner,
		Title:   articleInfo.Title,
		Summary: articleInfo.Summary,
		Content: articleInfo.Content,
		Cover:   articleInfo.Cover,
	})
	if err != nil {

		log.Errorf("user %s add article failed, err: %s", articleInfo.Owner.Username(), err.Error())

		return nil, err
	}

	return articleId, nil
}

func (s *articleService) AddArticleReadTimes(articleId cmprimitive.Id, plus cmprimitive.Amount) error {
	return s.repo.AddArticleReadTimes(articleId, plus)
}

func (s *articleService) UpdateArticle(articleId cmprimitive.Id, articleInfo *entity.ArticleInfo) error {

	if err := s.repo.Update(articleId, &entity.ArticleInfo{
		Owner:   articleInfo.Owner,
		Title:   articleInfo.Title,
		Content: articleInfo.Content,
		Summary: articleInfo.Summary,
		Cover:   articleInfo.Cover,
	}); err != nil {

		log.Errorf("user %s update article %s failed, err: %s",
			articleInfo.Owner.Username(), articleId.Id(), err.Error())

		return err
	}

	return nil
}
