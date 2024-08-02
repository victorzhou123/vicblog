package repositoryimpl

import (
	"strconv"

	"gorm.io/gorm"

	"victorzhou123/vicblog/article/domain"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

var tableNameArticle string

type ArticleDO struct {
	gorm.Model

	Owner     string `gorm:"column:owner;uniqueIndex:owner_index;size:255"`
	Title     string `gorm:"column:title;size:255"`
	Content   string `gorm:"column:content;type:text;size:40000"`
	Cover     string `gorm:"column:cover;size:255"`
	IsPublish bool   `gorm:"column:is_publish"`
	IsTop     bool   `gorm:"column:is_top"`
}

func (do *ArticleDO) toArticle() (article domain.Article, err error) {

	if article.Owner, err = cmprimitive.NewUsername(do.Owner); err != nil {
		return
	}

	if article.Title, err = cmprimitive.NewTitle(do.Title); err != nil {
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

	article.IsPublish = do.IsPublish

	article.IsTop = do.IsTop

	return
}

func (do *ArticleDO) TableName() string {
	return tableNameArticle
}
