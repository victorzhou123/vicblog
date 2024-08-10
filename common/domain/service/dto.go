package service

import (
	"fmt"
	"victorzhou123/vicblog/common/domain/repository"
)

type ListCmd struct {
	CurPage  int
	PageSize int
}

func (cmd *ListCmd) Validate() error {
	if cmd.CurPage <= 0 {
		return fmt.Errorf("current page must > 0")
	}

	if cmd.PageSize <= 0 {
		return fmt.Errorf("page size must > 0")
	}

	return nil
}

func (cmd *ListCmd) ToPageListOpt() repository.PageListOpt {
	return repository.PageListOpt{
		CurPage:  cmd.CurPage,
		PageSize: cmd.PageSize,
	}
}
