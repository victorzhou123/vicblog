package controller

import (
	"strconv"
	"victorzhou123/vicblog/common/domain/service"
)

type ReqList struct {
	CurPage  string `json:"current"`
	PageSize string `json:"size"`
}

func (req *ReqList) EmptyValue() bool {
	return req.CurPage == "" && req.PageSize == ""
}

func (req *ReqList) ToCmd() (cmd service.PaginationCmd, err error) {

	if cmd.CurPage, err = strconv.Atoi(req.CurPage); err != nil {
		return
	}

	if cmd.PageSize, err = strconv.Atoi(req.PageSize); err != nil {
		return
	}

	return
}
