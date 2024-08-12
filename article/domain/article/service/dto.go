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

func toArticleListDto(articles []entity.Article, cmd *dmservice.PaginationCmd, total int) ArticleListDto {

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
	Summary   string `json:"summary"`
	Cover     string `json:"cover"`
	IsPublish bool   `json:"isPublish"`
	IsTop     bool   `json:"isTop"`
	CreatedAt string `json:"createTime"`
}

func toArticleDto(v entity.Article) ArticleDto {
	return ArticleDto{
		Id:        v.Id.IdNum(),
		Title:     v.Title.Text(),
		Summary:   v.Summary.ArticleSummary(),
		Cover:     v.Cover.Urlx(),
		IsPublish: v.IsPublish,
		IsTop:     v.IsTop,
		CreatedAt: v.CreatedAt.TimeYearToSecond(),
	}
}

// list all articles
type ListAllArticleCmd struct {
	dmservice.PaginationCmd
}

func (cmd *ListAllArticleCmd) toPageListOpt() repository.PageListOpt {
	return cmd.PaginationCmd.ToPageListOpt()
}

// type ArticleDetailListDto struct {
// 	dmservice.PaginationDto

// 	Articles []ArticleDetailDto
// }

// func toArticleDetailListDto(articles []entity.Article, cmd *ListAllArticleCmd, total int) ArticleDetailListDto {

// 	articleDtos := make([]ArticleDetailDto, len(articles))
// 	for i := range articles {
// 		articleDtos[i] = toArticleDetailDto(articles[i])
// 	}

// 	return ArticleDetailListDto{cmd.ToPaginationDto(total), articleDtos}
// }

type ArticleDetailDto struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Cover     string `json:"cover"`
	IsPublish bool   `json:"isPublish"`
	IsTop     bool   `json:"isTop"`
	CreatedAt string `json:"createTime"`
}

func toArticleDetailDto(v entity.Article) ArticleDetailDto {
	return ArticleDetailDto{
		Id:        v.Id.IdNum(),
		Title:     v.Title.Text(),
		Summary:   v.Summary.ArticleSummary(),
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

// get article
type GetArticleCmd struct {
	User      cmprimitive.Username
	ArticleId cmprimitive.Id
}

// update article
type UpdateArticleCmd struct {
	Id      cmprimitive.Id
	User    cmprimitive.Username
	Title   cmprimitive.Text
	Content cmprimitive.Text
	Summary entity.ArticleSummary
	Cover   cmprimitive.Urlx
}
