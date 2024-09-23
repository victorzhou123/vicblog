package dto

import (
	"errors"

	"github.com/victorzhou123/vicblog/article/domain/article/entity"
	articledmsvc "github.com/victorzhou123/vicblog/article/domain/article/service"
	cateent "github.com/victorzhou123/vicblog/article/domain/category/entity"
	tagent "github.com/victorzhou123/vicblog/article/domain/tag/entity"
	cmappdto "github.com/victorzhou123/vicblog/common/app/dto"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/common/util"
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
		UpdatedAt: article.UpdatedAt.TimeYearToSecond(),
	}
}

// update article
type UpdateArticleCmd struct {
	AddArticleCmd

	Id cmprimitive.Id
}

type ArticleDto struct {
	ArticleSummaryDto

	Content      string `json:"content"`
	ReadTimes    int    `json:"readTimes"`
	TotalChars   int    `json:"totalChars"`
	ReadDuration string `json:"readDuration"`
}

func toArticleDto(article entity.Article) ArticleDto {
	return ArticleDto{
		ArticleSummaryDto: toArticleSummaryDto(article),
		Content:           article.Content.Text(),
		ReadTimes:         article.ReadTimes,
		TotalChars:        util.CharacterLen(article.Content.Text()),
		ReadDuration:      util.ReadDurationAnalyze(article.Content.Text()),
	}
}

type ArticleWithTagCateDto struct {
	ArticleDto

	Category    CategoryDto        `json:"category"`
	Tags        []TagDto           `json:"tags"`
	PrevArticle *ArticleIdTitleDto `json:"prevArticle"`
	NextArticle *ArticleIdTitleDto `json:"nextArticle"`
}

func ToArticleWithTagCateDto(
	article entity.Article, tags []tagent.Tag, cate cateent.Category,
	prevArticle, nextArticle *entity.ArticleIdTitle,
) ArticleWithTagCateDto {

	tagDtos := make([]TagDto, len(tags))
	for i := range tags {
		tagDtos[i] = ToTagDto(tags[i])
	}

	return ArticleWithTagCateDto{
		ArticleDto:  toArticleDto(article),
		Category:    ToCategoryDto(cate),
		Tags:        tagDtos,
		PrevArticle: toArticleIdTitleDto(prevArticle),
		NextArticle: toArticleIdTitleDto(nextArticle),
	}
}

type ArticleIdTitleDto struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

func toArticleIdTitleDto(article *entity.ArticleIdTitle) *ArticleIdTitleDto {
	if article == nil {
		return nil
	}

	return &ArticleIdTitleDto{
		Id:    article.Id.IdNum(),
		Title: article.Title.Text(),
	}
}

// article cards
type GetArticleCardListByCateIdCmd struct {
	cmappdto.PaginationCmd

	CategoryId cmprimitive.Id
}

func (cmd *GetArticleCardListByCateIdCmd) Validate() error {
	if cmd.CategoryId == nil {
		return errors.New("category id must exist")
	}

	return cmd.PaginationCmd.Validate()
}

type GetArticleCardListByTagIdCmd struct {
	cmappdto.PaginationCmd

	TagId cmprimitive.Id
}

func (cmd *GetArticleCardListByTagIdCmd) Validate() error {
	if cmd.TagId == nil {
		return errors.New("tag id must exist")
	}

	return cmd.PaginationCmd.Validate()
}

type ArticleCardListDto struct {
	cmappdto.PaginationDto

	ArticleCards []ArticleCardDto `json:"articleCards"`
}

func ToArticleCardListDto(ps cment.PaginationStatus,
	articleCards []entity.ArticleCard,
) ArticleCardListDto {

	articleCadDtos := make([]ArticleCardDto, len(articleCards))
	for i := range articleCards {
		articleCadDtos[i] = ToArticleCardDto(articleCards[i])
	}

	return ArticleCardListDto{
		PaginationDto: cmappdto.ToPaginationDto(ps),
		ArticleCards:  articleCadDtos,
	}
}

type ArticleCardDto struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	ReadTimes int    `json:"readTimes"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}

func ToArticleCardDto(articleCard entity.ArticleCard) ArticleCardDto {
	return ArticleCardDto{
		Id:        articleCard.Id.IdNum(),
		Title:     articleCard.Title.Text(),
		Cover:     articleCard.Cover.Urlx(),
		ReadTimes: articleCard.ReadTimes,
		UpdatedAt: articleCard.UpdatedAt.TimeYearToSecond(),
		CreatedAt: articleCard.CreatedAt.TimeYearToSecond(),
	}
}

// list articles which is classified by month
type ArticlesClassifiedByMonthDto struct {
	cmappdto.PaginationDto

	ArticlesInSameMonth []ArticleCreatedInSameMonth `json:"archives"`
}

type ArticleCreatedInSameMonth struct {
	Date     string           `json:"date"` // yy-mm
	Articles []ArticleCardDto `json:"articles"`
}

// search article
type SearchArticlesCmd struct {
	cmappdto.PaginationCmd

	Word cmprimitive.Text
}

type ArticleCardWithSummaryDto struct {
	ArticleCardDto

	Summary string
}

type ArticleCardsWithSummaryDto struct {
	cmappdto.PaginationDto

	ArticleCardsWithSummary []ArticleCardWithSummaryDto `json:"searchResults"`
}

func ToArticleCardsWithSummaryDto(sa articledmsvc.ArticleCardWithSummaryDto) ArticleCardsWithSummaryDto {

	articleCardsWithSummary := make([]ArticleCardWithSummaryDto, len(sa.ArticleCardsWithSummary))
	for i := range articleCardsWithSummary {
		articleCardsWithSummary[i].ArticleCardDto = ToArticleCardDto(sa.ArticleCardsWithSummary[i].ArticleCard)
		articleCardsWithSummary[i].Summary = sa.ArticleCardsWithSummary[i].Summary.ArticleSummary()
	}

	return ArticleCardsWithSummaryDto{
		PaginationDto:           cmappdto.ToPaginationDto(sa.PaginationStatus),
		ArticleCardsWithSummary: articleCardsWithSummary,
	}
}
