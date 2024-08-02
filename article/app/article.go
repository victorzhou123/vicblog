package app

import (
	"victorzhou123/vicblog/article/domain/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

const msgCannotFoundTheArticle = "can not found the article"

type ArticleService interface {
	GetArticleList(cmprimitive.Username) ([]ArticleListDto, error)
	Delete(cmprimitive.Username, cmprimitive.Id) error
}

type articleService struct {
	repo repository.Article
}

func NewArticleService(repo repository.Article) ArticleService {
	return &articleService{
		repo: repo,
	}
}

func (s *articleService) GetArticleList(user cmprimitive.Username) ([]ArticleListDto, error) {
	articles, err := s.repo.GetArticles(user)
	if err != nil {
		return []ArticleListDto{}, cmdmerror.New(
			cmdmerror.ErrorCodeResourceNotFound, msgCannotFoundTheArticle,
		)
	}

	dtos := make([]ArticleListDto, len(articles))
	for i := range articles {
		dtos[i] = toArticleListDto(articles[i])
	}

	return dtos, nil
}

func (s *articleService) Delete(user cmprimitive.Username, id cmprimitive.Id) error {
	if err := s.repo.Delete(user, id); err != nil {
		return cmdmerror.NewNoPermission("")
	}

	return nil
}
