package controller

type ReqList struct {
	CurPage  string `json:"current"`
	PageSize string `json:"size"`
}

func (req *ReqList) EmptyValue() bool {
	return req.CurPage == "" && req.PageSize == ""
}
