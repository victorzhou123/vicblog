package controller

import (
	"victorzhou123/vicblog/common/app/dto"
	"victorzhou123/vicblog/common/domain/primitive"
)

type ReqList struct {
	CurPage  string `json:"current"`
	PageSize string `json:"size"`
}

func (req *ReqList) EmptyValue() bool {
	return req.CurPage == "" && req.PageSize == ""
}

func (req *ReqList) ToCmd() (cmd dto.PaginationCmd, err error) {

	if cmd.CurPage, err = primitive.NewCurPageWithString(req.CurPage); err != nil {
		return
	}

	if cmd.PageSize, err = primitive.NewPageSizeWithString(req.PageSize); err != nil {
		return
	}

	return
}
