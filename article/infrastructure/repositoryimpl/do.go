package repositoryimpl

import (
	"strconv"

	"gorm.io/gorm"

	"github.com/victorzhou123/vicblog/article/domain/article/entity"
	categoryent "github.com/victorzhou123/vicblog/article/domain/category/entity"
	tagent "github.com/victorzhou123/vicblog/article/domain/tag/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

const (
	fieldNamePrimaryKeyId     = "id"
	fieldNameArticleReadTimes = "read_times"
	fieldNameCategoryId       = "category_id"
)

var (
	tableNameArticle         = "article"
	tableNameCategory        = "category"
	tableNameTag             = "tag"
	tableNameTagArticle      = "tag_article"
	tableNameCategoryArticle = "category_article"
)

// article
type ArticleDO struct {
	gorm.Model

	Owner     string `gorm:"column:owner;index:owner_index;size:255"`
	Summary   string `gorm:"column:summary;size:255"`
	Title     string `gorm:"column:title;size:255"`
	Content   string `gorm:"column:content;type:text;size:40000"`
	Cover     string `gorm:"column:cover;size:255"`
	ReadTimes int    `gorm:"column:read_times;size:11"`
	IsPublish bool   `gorm:"column:is_publish"`
	IsTop     bool   `gorm:"column:is_top"`
}

func (do *ArticleDO) toArticle() (article entity.Article, err error) {

	if article.Owner, err = cmprimitive.NewUsername(do.Owner); err != nil {
		return
	}

	if article.Title, err = cmprimitive.NewTitle(do.Title); err != nil {
		return
	}

	if article.Summary, err = entity.NewArticleSummary(do.Summary); err != nil {
		return
	}

	if article.Content, err = cmprimitive.NewArticleContent(do.Content); err != nil {
		return
	}

	if article.Cover, err = cmprimitive.NewUrlx(do.Cover); err != nil {
		return
	}

	article.Id = cmprimitive.NewId(strconv.FormatUint(uint64(do.ID), 10))

	article.CreatedAt = cmprimitive.NewTimeXWithUnix(do.CreatedAt.Unix())

	article.UpdatedAt = cmprimitive.NewTimeXWithUnix(do.UpdatedAt.Unix())

	article.ReadTimes = do.ReadTimes

	article.IsPublish = do.IsPublish

	article.IsTop = do.IsTop

	return
}

func (do *ArticleDO) TableName() string {
	return tableNameArticle
}

type ArticleIdTitleDO struct {
	ID    uint
	Title string `gorm:"column:title"`
}

func (do *ArticleIdTitleDO) toArticleIdTitle() (*entity.ArticleIdTitle, error) {

	title, err := cmprimitive.NewTitle(do.Title)
	if err != nil {
		return &entity.ArticleIdTitle{}, err
	}

	return &entity.ArticleIdTitle{
		Id:    cmprimitive.NewIdByUint(do.ID),
		Title: title,
	}, nil
}

// category
type CategoryDO struct {
	gorm.Model

	Name string `gorm:"column:name;unique;size:60"`
}

func (do *CategoryDO) toCategory() (category categoryent.Category, err error) {

	if category.Name, err = categoryent.NewCategoryName(do.Name); err != nil {
		return
	}

	category.Id = cmprimitive.NewIdByUint(do.ID)

	category.CreatedAt = cmprimitive.NewTimeXWithUnix(do.CreatedAt.Unix())

	return
}

func (do *CategoryDO) TableName() string {
	return tableNameCategory
}

// tag
type TagDO struct {
	gorm.Model

	Name string `gorm:"column:name;unique;size:60"`
}

func (do *TagDO) toTag() (tag tagent.Tag, err error) {
	if tag.Name, err = tagent.NewTagName(do.Name); err != nil {
		return
	}

	tag.Id = cmprimitive.NewIdByUint(do.ID)

	tag.CreatedAt = cmprimitive.NewTimeXWithUnix(do.CreatedAt.Unix())

	return
}

func (do *TagDO) TableName() string {
	return tableNameTag
}

// tag-article
type TagArticleDO struct {
	gorm.Model

	TagId     uint `gorm:"column:tag_id;not null;uniqueIndex:idx_tag_id_article_id"`
	ArticleId uint `gorm:"column:article_id;not null;uniqueIndex:idx_tag_id_article_id"`
}

func (do *TagArticleDO) TableName() string {
	return tableNameTagArticle
}

// category-article
type CategoryArticleDO struct {
	gorm.Model

	CategoryId uint `gorm:"column:category_id;not null;uniqueIndex:idx_category_id_article_id"`
	ArticleId  uint `gorm:"column:article_id;not null;uniqueIndex:idx_category_id_article_id"`
}

func (do *CategoryArticleDO) TableName() string {
	return tableNameCategoryArticle
}
