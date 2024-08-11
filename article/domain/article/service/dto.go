package service

import (
	"errors"

	"victorzhou123/vicblog/article/domain/article/entity"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/domain/repository"
	dmservice "victorzhou123/vicblog/common/domain/service"
)

// list article
type ArticleListCmd struct {
	dmservice.PaginationCmd

	User cmprimitive.Username
}

func (cmd *ArticleListCmd) Validate() error {

	if cmd.User == nil {
		return errors.New("user cannot be empty")
	}

	return cmd.PaginationCmd.Validate()
}

func (cmd *ArticleListCmd) toPageListOpt() repository.PageListOpt {
	return cmd.PaginationCmd.ToPageListOpt()
}

type ArticleListDto struct {
	dmservice.PaginationDto

	Articles []ArticleDto `json:"articles"`
}

func toArticleListDto(articles []entity.Article, cmd *ArticleListCmd, total int) ArticleListDto {

	dtos := make([]ArticleDto, len(articles))
	for i := range articles {
		dtos[i] = toArticleDto(articles[i])
	}

	return ArticleListDto{
		PaginationDto: cmd.ToPaginationDto(total),
		Articles:      dtos,
	}
}

type ArticleDto struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	IsPublish bool   `json:"isPublish"`
	IsTop     bool   `json:"isTop"`
	CreatedAt string `json:"createTime"`
}

func toArticleDto(v entity.Article) ArticleDto {
	return ArticleDto{
		Id:        v.Id.IdNum(),
		Title:     v.Title.Text(),
		Cover:     v.Cover.Urlx(),
		IsPublish: v.IsPublish,
		IsTop:     v.IsTop,
		CreatedAt: v.CreatedAt.TimeYearToSecond(),
	}
}

// add article
type ArticleCmd struct {
	Owner   cmprimitive.Username
	Title   cmprimitive.Text
	Summary entity.ArticleSummary
	Content cmprimitive.Text
	Cover   cmprimitive.Urlx
}
