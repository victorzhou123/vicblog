package service

import (
	"reflect"
	"testing"

	"github.com/victorzhou123/vicblog/article/domain/article/entity"
	cmentt "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

var (
	opt                                                                                           *cmentt.Pagination
	paginationStatus                                                                              cmentt.PaginationStatus
	total                                                                                         int
	cardsEmpty, cardsOne, cardsTwoInSameMonth, cardsTwoInDiffMonth, cardsFour                     []entity.ArticleCard
	rightSortEmpty, rightSortOne, rightSortTwoInSameMonth, rightSortTwoInDiffMonth, rightSortFour ArticleListClassifyByMonthDto
)

func init() {

	curPage, _ := cmprimitive.NewCurPage(1)
	pageSize, _ := cmprimitive.NewPageSize(1)
	opt = &cmentt.Pagination{CurPage: curPage, PageSize: pageSize}
	total = 1

	paginationStatus = opt.ToPaginationStatus(total)

	articleCard20230111 := entity.ArticleCard{CreatedAt: cmprimitive.NewTimeXWithUnix(1673366400)}
	articleCard20240720 := entity.ArticleCard{CreatedAt: cmprimitive.NewTimeXWithUnix(1721404800)}
	articleCard20240820 := entity.ArticleCard{CreatedAt: cmprimitive.NewTimeXWithUnix(1724083200)}
	articleCard20240823 := entity.ArticleCard{CreatedAt: cmprimitive.NewTimeXWithUnix(1724342400)}

	cardsEmpty = []entity.ArticleCard{}
	cardsOne = []entity.ArticleCard{articleCard20240820}
	cardsTwoInSameMonth = []entity.ArticleCard{articleCard20240820, articleCard20240823}
	cardsTwoInDiffMonth = []entity.ArticleCard{articleCard20240820, articleCard20230111}
	cardsFour = []entity.ArticleCard{articleCard20230111, articleCard20240823, articleCard20240720, articleCard20240820}

	rightSortEmpty = ArticleListClassifyByMonthDto{PaginationStatus: paginationStatus}
	rightSortOne = ArticleListClassifyByMonthDto{PaginationStatus: paginationStatus, ArticleArchives: []ArticleArchiveDto{{Time: articleCard20240820.CreatedAt, ArticleCards: []entity.ArticleCard{articleCard20240820}}}}
	rightSortTwoInSameMonth = ArticleListClassifyByMonthDto{PaginationStatus: paginationStatus, ArticleArchives: []ArticleArchiveDto{{Time: articleCard20240820.CreatedAt, ArticleCards: []entity.ArticleCard{articleCard20240823, articleCard20240820}}}}
	rightSortTwoInDiffMonth = ArticleListClassifyByMonthDto{PaginationStatus: paginationStatus, ArticleArchives: []ArticleArchiveDto{{Time: articleCard20240820.CreatedAt, ArticleCards: []entity.ArticleCard{articleCard20240820}}, {Time: articleCard20230111.CreatedAt, ArticleCards: []entity.ArticleCard{articleCard20230111}}}}
	rightSortFour = ArticleListClassifyByMonthDto{PaginationStatus: paginationStatus, ArticleArchives: []ArticleArchiveDto{{Time: articleCard20240820.CreatedAt, ArticleCards: []entity.ArticleCard{articleCard20240823, articleCard20240820}}, {Time: articleCard20240720.CreatedAt, ArticleCards: []entity.ArticleCard{articleCard20240720}}, {Time: articleCard20230111.CreatedAt, ArticleCards: []entity.ArticleCard{articleCard20230111}}}}
}

func Test_toArticleListClassifyByMonthDto(t *testing.T) {
	type args struct {
		articleCards []entity.ArticleCard
		cmd          *cmentt.Pagination
		total        int
	}
	tests := []struct {
		name string
		args args
		want ArticleListClassifyByMonthDto
	}{
		{
			"test1. articleCards empty",
			args{
				cardsEmpty,
				opt,
				total,
			},
			rightSortEmpty,
		},
		{
			"test2. two articleCards in the same month",
			args{
				cardsTwoInSameMonth,
				opt,
				total,
			},
			rightSortTwoInSameMonth,
		},
		{
			"test3. two articleCards in the different month",
			args{
				cardsTwoInDiffMonth,
				opt,
				total,
			},
			rightSortTwoInDiffMonth,
		},
		{
			"test4. four articleCards",
			args{
				cardsFour,
				opt,
				total,
			},
			rightSortFour,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toArticleListClassifyByMonthDto(tt.args.articleCards, tt.args.cmd, tt.args.total); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toArticleListClassifyByMonthDto() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
