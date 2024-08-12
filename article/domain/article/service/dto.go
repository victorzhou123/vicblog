package service

import (
	"errors"

	"victorzhou123/vicblog/article/domain/article/entity"
	cment "victorzhou123/vicblog/common/domain/entity"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
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
