package service

import (
	"victorzhou123/vicblog/article/domain/article/entity"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type ArticleListDto struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	IsPublish bool   `json:"isPublish"`
	IsTop     bool   `json:"isTop"`
	CreatedAt string `json:"createTime"`
}

func toArticleListDto(v entity.Article) ArticleListDto {
	return ArticleListDto{
		Id:        v.Id.IdNum(),
		Title:     v.Title.Text(),
		Cover:     v.Cover.Urlx(),
		IsPublish: v.IsPublish,
		IsTop:     v.IsTop,
		CreatedAt: v.CreatedAt.TimeYearToSecond(),
	}
}

type ArticleCmd struct {
	Owner   cmprimitive.Username
	Title   cmprimitive.Text
	Summary entity.ArticleSummary
	Content cmprimitive.Text
	Cover   cmprimitive.Urlx
}
