package controller

import (
	"github.com/victorzhou123/vicblog/category-server/app/dto"
	"github.com/victorzhou123/vicblog/category-server/domain/category/entity"
	cmctl "github.com/victorzhou123/vicblog/common/controller"
)

type reqCategory struct {
	Name string `json:"name"`
}

func (req *reqCategory) toCategoryName() (entity.CategoryName, error) {
	return entity.NewCategoryName(req.Name)
}

type reqCategoryList struct {
	cmctl.ReqList
}

func (req *reqCategoryList) toCmd() (cmd dto.ListCategoryCmd, err error) {

	listCmd, err := req.ReqList.ToCmd()
	if err != nil {
		return
	}

	cmd = dto.ListCategoryCmd{
		PaginationCmd: listCmd,
	}

	return
}
