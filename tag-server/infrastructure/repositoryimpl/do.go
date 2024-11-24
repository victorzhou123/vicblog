package repositoryimpl

import (
	"gorm.io/gorm"

	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	tagent "github.com/victorzhou123/vicblog/tag-server/domain/tag/entity"
)

const (
	fieldNamePrimaryKeyId     = "id"
	fieldNameArticleReadTimes = "read_times"
	fieldNameCategoryId       = "category_id"
	fieldNameTagId            = "tag_id"
	fieldCreatedAt            = "created_at"
	fieldTitle                = "title"
)

var (
	tableNameTag             = "tag"
	tableNameTagArticle      = "tag_article"
	tableNameCategoryArticle = "category_article"
)

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
