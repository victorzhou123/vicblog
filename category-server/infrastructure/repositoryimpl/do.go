package repositoryimpl

import (
	"gorm.io/gorm"

	"github.com/victorzhou123/vicblog/category-server/domain/category/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
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
	tableNameCategory        = "category"
	tableNameCategoryArticle = "category_article"
)

// category
type CategoryDO struct {
	gorm.Model

	Name string `gorm:"column:name;unique;size:60"`
}

func (do *CategoryDO) toCategory() (category entity.Category, err error) {

	if category.Name, err = entity.NewCategoryName(do.Name); err != nil {
		return
	}

	category.Id = cmprimitive.NewIdByUint(do.ID)

	category.CreatedAt = cmprimitive.NewTimeXWithUnix(do.CreatedAt.Unix())

	return
}

func (do *CategoryDO) TableName() string {
	return tableNameCategory
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
