package service

import (
	"victorzhou123/vicblog/article/domain/article/entity"
	"victorzhou123/vicblog/article/domain/article/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

const msgCannotFoundTheArticle = "can not found the article"

type ArticleService interface {
	GetArticleList(*ArticleListCmd) (ArticleListDto, error)
	Delete(cmprimitive.Username, cmprimitive.Id) error
	AddArticle(cmd *ArticleCmd) (articleId uint, err error)
}

type articleService struct {
	repo repository.Article
}

func NewArticleService(repo repository.Article) ArticleService {
	return &articleService{
		repo: repo,
	}
}

func (s *articleService) GetArticleList(cmd *ArticleListCmd) (ArticleListDto, error) {

	articles, total, err := s.repo.ListArticles(cmd.User, cmd.toPageListOpt())
	if err != nil {
		return ArticleListDto{}, cmdmerror.New(
			cmdmerror.ErrorCodeResourceNotFound, msgCannotFoundTheArticle,
		)
	}

	return toArticleListDto(articles, cmd, total), nil
}

func (s *articleService) Delete(user cmprimitive.Username, id cmprimitive.Id) error {
	if err := s.repo.Delete(user, id); err != nil {
		return cmdmerror.NewNoPermission("")
	}

	return nil
}

func (s *articleService) AddArticle(cmd *ArticleCmd) (uint, error) {
	return s.repo.AddArticle(&entity.ArticleInfo{
		Owner:   cmd.Owner,
		Title:   cmd.Title,
		Summary: cmd.Summary,
		Content: cmd.Content,
		Cover:   cmd.Cover,
	})
}
