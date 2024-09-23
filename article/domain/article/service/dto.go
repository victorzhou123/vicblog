package service

import (
	"errors"
	"sort"

	"github.com/victorzhou123/vicblog/article/domain/article/entity"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmdmerr "github.com/victorzhou123/vicblog/common/domain/error"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

// list article
type ArticleListCmd struct {
	cment.Pagination

	User cmprimitive.Username
}

func (cmd *ArticleListCmd) Validate() error {

	if cmd.User == nil {
		return errors.New("user cannot be empty")
	}

	return nil
}

type ArticleListDto struct {
	cment.PaginationStatus

	Articles []entity.Article
}

// get articles classify by month
type ArticleListClassifyByMonthDto struct {
	cment.PaginationStatus

	ArticleArchives []ArticleArchiveDto
}

type ArticleArchiveDto struct {
	Time         cmprimitive.Timex
	ArticleCards []entity.ArticleCard
}

func toArticleListClassifyByMonthDto(articleCards []entity.ArticleCard, cmd *cment.Pagination, total int) ArticleListClassifyByMonthDto {

	// ensure len of articles not 0
	if len(articleCards) == 0 {
		return ArticleListClassifyByMonthDto{PaginationStatus: cmd.ToPaginationStatus(total)}
	}

	// order by time desc
	sort.Slice(articleCards, func(i, j int) bool {
		return articleCards[i].CreatedAt.TimeUnix() > articleCards[j].CreatedAt.TimeUnix()
	})

	// convert
	d := ArticleListClassifyByMonthDto{PaginationStatus: cmd.ToPaginationStatus(total)}
	for i, j := 0, 0; i < len(articleCards) && j < len(articleCards); i++ {

		if i+1 >= len(articleCards) || !articleCards[i].IsSameMonthCreated(articleCards[i+1]) {
			d.ArticleArchives = append(d.ArticleArchives, ArticleArchiveDto{
				Time:         articleCards[i].CreatedAt,
				ArticleCards: articleCards[j : i+1],
			})

			j = i + 1
		}
	}

	return d

}

// get article card
type ArticleCardsCmd struct {
	cment.Pagination

	ArticleIds []cmprimitive.Id
}

func (cmd *ArticleCardsCmd) validate() error {

	if len(cmd.ArticleIds) == 0 {
		return cmdmerr.NewInvalidParam("at least one article id required")
	}

	return nil
}

type ArticleCardsDto struct {
	cment.PaginationStatus

	ArticleCards []entity.ArticleCard
}

// get article
type GetArticleCmd struct {
	User      cmprimitive.Username
	ArticleId cmprimitive.Id
}

// get prev and next article
type ArticlePrevAndNextDto struct {
	Prev *entity.ArticleIdTitle
	Next *entity.ArticleIdTitle
}

// article search
type ArticleCardWithSummaryDto struct {
	cment.PaginationStatus

	ArticleCardsWithSummary []entity.ArticleCardWithSummary
}
