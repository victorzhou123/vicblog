package service

import (
	"victorzhou123/vicblog/article/domain/article/entity"
	"victorzhou123/vicblog/article/domain/article/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/log"
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

		log.Errorf("user %s delete article %s failed, err: %s", user.Username(), id.Id(), err.Error())

		return cmdmerror.NewNoPermission("")
	}

	return nil
}

func (s *articleService) AddArticle(cmd *ArticleCmd) (uint, error) {

	articleId, err := s.repo.AddArticle(&entity.ArticleInfo{
		Owner:   cmd.Owner,
		Title:   cmd.Title,
		Summary: cmd.Summary,
		Content: cmd.Content,
		Cover:   cmd.Cover,
	})
	if err != nil {

		log.Errorf("user %s add article failed, err: %s", cmd.Owner.Username(), err.Error())

		return 0, err
	}

	return articleId, nil
}
