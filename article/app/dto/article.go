package dto

import (
	"victorzhou123/vicblog/article/domain/article/entity"
	articledmsvc "victorzhou123/vicblog/article/domain/article/service"
	cateent "victorzhou123/vicblog/article/domain/category/entity"
	tagent "victorzhou123/vicblog/article/domain/tag/entity"
	cmappdto "victorzhou123/vicblog/common/app/dto"
	cment "victorzhou123/vicblog/common/domain/entity"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

// add article
type AddArticleCmd struct {
	Owner    cmprimitive.Username
	Title    cmprimitive.Text
	Summary  entity.ArticleSummary
	Content  cmprimitive.Text
	Cover    cmprimitive.Urlx
	Category cmprimitive.Id
	Tags     []cmprimitive.Id
}

func (cmd *AddArticleCmd) ToArticleInfo() *entity.ArticleInfo {
	return &entity.ArticleInfo{
		Owner:   cmd.Owner,
		Title:   cmd.Title,
		Summary: cmd.Summary,
		Content: cmd.Content,
		Cover:   cmd.Cover,
	}
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
	article entity.Article,
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

// get article list
type GetArticleListCmd struct {
	cmappdto.PaginationCmd

	User cmprimitive.Username
}

type ArticleListDto struct {
	cmappdto.PaginationDto

	Articles []ArticleSummaryDto `json:"articles"`
}

func ToArticleListDto(
	ps cment.PaginationStatus,
	articles []entity.Article,
) ArticleListDto {

	articleSummaryDtos := make([]ArticleSummaryDto, len(articles))
	for i := range articles {
		articleSummaryDtos[i] = toArticleSummaryDto(articles[i])
	}

	return ArticleListDto{
		PaginationDto: cmappdto.ToPaginationDto(ps),
		Articles:      articleSummaryDtos,
	}
}

// list all articles
type ListAllArticlesCmd struct {
	cmappdto.PaginationCmd
}

type ArticleDetailsListDto struct {
	cmappdto.PaginationDto

	Articles []ArticleDetailListDto `json:"articles"`
}

type ArticleDetailListDto struct {
	ArticleSummaryDto

	Category CategoryDto `json:"category"`
	Tags     []TagDto    `json:"tags"`
}

func ToArticleDetailListDto(
	article entity.Article, category cateent.Category, tags []tagent.Tag,
) ArticleDetailListDto {

	tagDtos := make([]TagDto, len(tags))
	for i := range tags {
		tagDtos[i] = ToTagDto(tags[i])
	}

	return ArticleDetailListDto{
		ArticleSummaryDto: toArticleSummaryDto(article),
		Category:          ToCategoryDto(category),
		Tags:              tagDtos,
	}
}

type ArticleSummaryDto struct {
	Id        uint   `json:"id"`
	Owner     string `json:"owner"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Cover     string `json:"cover"`
	IsPublish bool   `json:"isPublish"`
	IsTop     bool   `json:"isTop"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func toArticleSummaryDto(article entity.Article) ArticleSummaryDto {
	return ArticleSummaryDto{
		Id:        article.Id.IdNum(),
		Owner:     article.Owner.Username(),
		Title:     article.Title.Text(),
		Summary:   article.Summary.ArticleSummary(),
		Cover:     article.Cover.Urlx(),
		IsPublish: article.IsPublish,
		IsTop:     article.IsTop,
		CreatedAt: article.CreatedAt.TimeYearToSecond(),
	}
}

// update article
type UpdateArticleCmd struct {
	AddArticleCmd

	Id cmprimitive.Id
}
