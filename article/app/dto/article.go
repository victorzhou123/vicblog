package dto

import (
	articleent "victorzhou123/vicblog/article/domain/article/entity"
	articledmsvc "victorzhou123/vicblog/article/domain/article/service"
	catesvc "victorzhou123/vicblog/article/domain/category/service"
	tagsvc "victorzhou123/vicblog/article/domain/tag/service"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	dmsvc "victorzhou123/vicblog/common/domain/service"
)

// add article
type AddArticleCmd struct {
	Owner    cmprimitive.Username
	Title    cmprimitive.Text
	Summary  articleent.ArticleSummary
	Content  cmprimitive.Text
	Cover    cmprimitive.Urlx
	Category cmprimitive.Id
	Tags     []cmprimitive.Id
}

// get article
type GetArticleCmd struct {
	articledmsvc.GetArticleCmd
}

type ArticleDetailDto struct {
	Id         uint   `json:"id"`
	Owner      string `json:"owner"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
	Cover      string `json:"cover"`
	ReadTimes  int    `json:"readTimes"`
	IsPublish  bool   `json:"isPublish"`
	IsTop      bool   `json:"isTop"`
	CategoryId uint   `json:"categoryId"`
	TagIds     []uint `json:"tagIds"`
	UpdatedAt  string `json:"updatedAt"`
	CreatedAt  string `json:"createdAt"`
}

func ToArticleDetailDto(
	article articleent.Article,
	tagIds []cmprimitive.Id,
	cateId cmprimitive.Id,
) ArticleDetailDto {

	tags := make([]uint, len(tagIds))
	for i := range tagIds {
		tags[i] = tagIds[i].IdNum()
	}

	return ArticleDetailDto{
		Id:         article.Id.IdNum(),
		Owner:      article.Owner.Username(),
		Title:      article.Title.Text(),
		Summary:    article.Summary.ArticleSummary(),
		Content:    article.Content.Text(),
		Cover:      article.Cover.Urlx(),
		ReadTimes:  article.ReadTimes,
		IsPublish:  article.IsPublish,
		IsTop:      article.IsTop,
		CategoryId: cateId.IdNum(),
		TagIds:     tags,
		UpdatedAt:  article.UpdatedAt.TimeYearToSecond(),
		CreatedAt:  article.CreatedAt.TimeYearToSecond(),
	}
}

// list all articles
type ListAllArticlesCmd struct {
	dmsvc.PaginationCmd
}

type ArticleDetailsListDto struct {
	dmsvc.PaginationDto

	Articles []ArticleDetailListDto `json:"articles"`
}

type ArticleDetailListDto struct {
	articledmsvc.ArticleDto

	Category catesvc.CategoryDto `json:"category"`
	Tags     []tagsvc.TagDto     `json:"tags"`
}

// update article
type UpdateArticleCmd struct {
	articledmsvc.UpdateArticleCmd

	CategoryId cmprimitive.Id
	TagIds     []cmprimitive.Id
}
