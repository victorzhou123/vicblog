package repository

import (
	"victorzhou123/vicblog/article/domain/article/entity"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	cmrepo "victorzhou123/vicblog/common/domain/repository"
)

type Article interface {
	GetArticle(user cmprimitive.Username, articleId cmprimitive.Id) (entity.Article, error)
	ListArticles(cmprimitive.Username, cmrepo.PageListOpt) ([]entity.Article, int, error)
	ListAllArticles(cmrepo.PageListOpt) ([]entity.Article, int, error)

	Delete(cmprimitive.Username, cmprimitive.Id) error

	AddArticle(*entity.ArticleInfo) (articleId uint, err error)

	Update(*entity.ArticleUpdate) error
}
