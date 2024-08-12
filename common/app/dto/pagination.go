package dto

import (
	"victorzhou123/vicblog/common/domain/entity"
	"victorzhou123/vicblog/common/domain/primitive"
)

type PaginationCmd struct {
	CurPage  primitive.CurPage
	PageSize primitive.PageSize
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
