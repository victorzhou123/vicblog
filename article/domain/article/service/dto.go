package service

import (
	"errors"

	"github.com/victorzhou123/vicblog/article/domain/article/entity"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

// list article
type ArticleListCmd struct {
	cment.Pagination

	User cmprimitive.Username
}

func (cmd *ArticleListCmd) Validate() error {

	if cmd.User == nil {
		return errors.New("user cannot be empty")
	}

	return nil
}

type ArticleListDto struct {
	cment.PaginationStatus

	Articles []entity.Article
}

// get article
type GetArticleCmd struct {
	User      cmprimitive.Username
	ArticleId cmprimitive.Id
}

// get prev and next article
type ArticlePrevAndNextDto struct {
	Prev *entity.ArticleIdTitle
	Next *entity.ArticleIdTitle
}
