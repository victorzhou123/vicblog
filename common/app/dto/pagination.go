package dto

import (
	"errors"

	"github.com/victorzhou123/vicblog/common/domain/entity"
	"github.com/victorzhou123/vicblog/common/domain/primitive"
)

type PaginationCmd struct {
	CurPage  primitive.CurPage
	PageSize primitive.PageSize
}

func (cmd *PaginationCmd) Validate() error {
	if cmd.CurPage == nil {
		return errors.New("current page must exist")
	}

	if cmd.PageSize == nil {
		return errors.New("page size must exist")
	}

	return nil
}

func (cmd *PaginationCmd) ToPaginationDto(total int) PaginationDto {
	return PaginationDto{
		Total:     total,
		PageCount: total/cmd.PageSize.PageSize() + 1,
		PageSize:  cmd.PageSize.PageSize(),
		CurPage:   cmd.CurPage.CurPage(),
	}
}

func (cmd *PaginationCmd) ToPagination() *entity.Pagination {
	return &entity.Pagination{
		CurPage:  cmd.CurPage,
		PageSize: cmd.PageSize,
	}
}

type PaginationDto struct {
	Total     int `json:"total"`
	PageCount int `json:"pages"`
	PageSize  int `json:"size"`
	CurPage   int `json:"current"`
}

func ToPaginationDto(ps entity.PaginationStatus) PaginationDto {
	return PaginationDto{
		Total:     ps.Total,
		PageCount: ps.PageCount,
		PageSize:  ps.PageSize.PageSize(),
		CurPage:   ps.CurPage.CurPage(),
	}
}
