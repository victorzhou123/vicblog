package dto

import (
	"strconv"

	"github.com/victorzhou123/vicblog/common/controller/rpc"
	"github.com/victorzhou123/vicblog/common/domain/entity"
	"github.com/victorzhou123/vicblog/common/domain/primitive"
)

func ToPaginationStatus(paginationStatus *rpc.PaginationStatus) (entity.PaginationStatus, error) {

	pagination, err := toPagination(paginationStatus.Pagination)
	if err != nil {
		return entity.PaginationStatus{}, err
	}

	return entity.PaginationStatus{
		Pagination: pagination,
		Total:      int(paginationStatus.GetTotal()),
		PageCount:  int(paginationStatus.GetPageCount()),
	}, nil
}

func toPagination(page *rpc.Pagination) (entity.Pagination, error) {

	curPage, err := primitive.NewCurPageWithString(page.GetCurPage())
	if err != nil {
		return entity.Pagination{}, err
	}

	pageSize, err := primitive.NewPageSizeWithString(page.GetPageSize())
	if err != nil {
		return entity.Pagination{}, err
	}

	return entity.Pagination{
		CurPage:  curPage,
		PageSize: pageSize,
	}, nil
}

func ToProtoPagination(pagination entity.Pagination) *rpc.Pagination {
	return &rpc.Pagination{
		CurPage:  strconv.Itoa(pagination.CurPage.CurPage()),
		PageSize: strconv.Itoa(pagination.PageSize.PageSize()),
	}
}

func ToProtoPaginationStatus(paginationStatus entity.PaginationStatus) *rpc.PaginationStatus {
	return &rpc.PaginationStatus{
		Pagination: toProtoPagination(paginationStatus.Pagination),
		Total:      int64(paginationStatus.Total),
		PageCount:  int64(paginationStatus.PageCount),
	}
}

func toProtoPagination(pagination entity.Pagination) *rpc.Pagination {
	return &rpc.Pagination{
		CurPage:  strconv.Itoa(pagination.CurPage.CurPage()),
		PageSize: strconv.Itoa(pagination.PageSize.PageSize()),
	}
}
