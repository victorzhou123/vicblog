package repository

import (
	"github.com/victorzhou123/vicblog/article/domain/article/entity"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type Article interface {
	GetArticleById(articleId cmprimitive.Id) (entity.Article, error)
	GetArticle(user cmprimitive.Username, articleId cmprimitive.Id) (entity.Article, error)
	ListArticles(cmprimitive.Username, cment.Pagination) ([]entity.Article, int, error)
	ListArticleCards(articleIds []cmprimitive.Id, pagination cment.Pagination) ([]entity.ArticleCard, int, error)
	ListAllArticles(cment.Pagination) ([]entity.Article, int, error)
	GetPreAndNextArticle(articleId cmprimitive.Id) (articleArr [2]*entity.ArticleIdTitle, err error)
	GetTotalNumberOfArticle() (cmprimitive.Amount, error)

	Delete(cmprimitive.Username, cmprimitive.Id) error

	AddArticle(*entity.ArticleInfo) (articleId cmprimitive.Id, err error)
	AddArticleReadTimes(cmprimitive.Id, cmprimitive.Amount) error

	Update(articleId cmprimitive.Id, articleInfo *entity.ArticleInfo) error
}
