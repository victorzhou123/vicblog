package repository

import (
	"victorzhou123/vicblog/article/domain/article/entity"
	cment "victorzhou123/vicblog/common/domain/entity"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type Article interface {
	GetArticle(user cmprimitive.Username, articleId cmprimitive.Id) (entity.Article, error)
	ListArticles(cmprimitive.Username, cment.Pagination) ([]entity.Article, int, error)
	ListAllArticles(cment.Pagination) ([]entity.Article, int, error)

	Delete(cmprimitive.Username, cmprimitive.Id) error

	AddArticle(*entity.ArticleInfo) (articleId cmprimitive.Id, err error)

	Update(articleId cmprimitive.Id, articleInfo *entity.ArticleInfo) error
}
