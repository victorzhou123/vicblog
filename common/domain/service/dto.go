package service

import (
	"fmt"
	"victorzhou123/vicblog/common/domain/repository"
)

type PaginationCmd struct {
	CurPage  int
	PageSize int
}

func (cmd *PaginationCmd) Validate() error {
	if cmd.CurPage <= 0 {
		return fmt.Errorf("current page must > 0")
	}

	if cmd.PageSize <= 0 {
		return fmt.Errorf("page size must > 0")
	}

	return nil
}

func (cmd *PaginationCmd) ToPageListOpt() repository.PageListOpt {
	return repository.PageListOpt{
		CurPage:  cmd.CurPage,
		PageSize: cmd.PageSize,
	}
}

func (cmd *PaginationCmd) ToPaginationDto(total int) PaginationDto {
	return PaginationDto{
		Total:     total,
		PageCount: total/cmd.PageSize + 1,
		PageSize:  cmd.PageSize,
		CurPage:   cmd.CurPage,
	}
}

type PaginationDto struct {
	Total     int `json:"total"`
	PageCount int `json:"pages"`
	PageSize  int `json:"size"`
	CurPage   int `json:"current"`
}
