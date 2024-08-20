package entity

import cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"

type Article struct {
	Id        cmprimitive.Id
	Owner     cmprimitive.Username
	Title     cmprimitive.Text
	Summary   ArticleSummary
	Content   cmprimitive.Text
	Cover     cmprimitive.Urlx
	ReadTimes int
	IsPublish bool
	IsTop     bool
	UpdatedAt cmprimitive.Timex
	CreatedAt cmprimitive.Timex
}

type ArticleCard struct {
	Id        cmprimitive.Id
	Title     cmprimitive.Text
	Cover     cmprimitive.Urlx
	ReadTimes int
	UpdatedAt cmprimitive.Timex
	CreatedAt cmprimitive.Timex
}

func (r *ArticleCard) IsSameMonthCreated(article ArticleCard) bool {
	return r.CreatedAt.TimeYearMonthOnly() == article.CreatedAt.TimeYearMonthOnly()
}

type ArticleInfo struct {
	Owner   cmprimitive.Username
	Title   cmprimitive.Text
	Summary ArticleSummary
	Content cmprimitive.Text
	Cover   cmprimitive.Urlx
}

type ArticleIdTitle struct {
	Id    cmprimitive.Id
	Title cmprimitive.Text
}
