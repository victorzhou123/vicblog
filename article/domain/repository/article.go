package repository

import (
	"victorzhou123/vicblog/article/domain"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type Article interface {
	GetArticles(cmprimitive.Username) ([]domain.Article, error)
	Delete(cmprimitive.Username, cmprimitive.Id) error
}
