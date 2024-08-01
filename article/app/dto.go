package app

import "victorzhou123/vicblog/article/domain"

type ArticleListDto struct {
	Id        string
	Title     string
	Cover     string
	IsPublish bool   `json:"isPublish"`
	IsTop     bool   `json:"isTop"`
	CreatedAt string `json:"createTime"`
}

func toArticleListDto(v domain.Article) ArticleListDto {
	return ArticleListDto{
		Id:        v.Id.Id(),
		Title:     v.Title.Text(),
		Cover:     v.Cover.Urlx(),
		IsPublish: v.IsPublish,
		IsTop:     v.IsTop,
		CreatedAt: v.CreatedAt.TimeYearToSecond(),
	}
}
