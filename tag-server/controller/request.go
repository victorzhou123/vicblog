package controller

import (
	cmctl "github.com/victorzhou123/vicblog/common/controller"
	"github.com/victorzhou123/vicblog/tag-server/app/dto"
	tagent "github.com/victorzhou123/vicblog/tag-server/domain/tag/entity"
)

type reqTag struct {
	Names []string `json:"names"`
}

func (req *reqTag) toTagNames() ([]tagent.TagName, error) {
	tagNames := make([]tagent.TagName, len(req.Names))

	var err error
	for i := range req.Names {
		tagNames[i], err = tagent.NewTagName(req.Names[i])
		if err != nil {
			return nil, err
		}
	}

	return tagNames, nil
}

type reqTagList struct {
	cmctl.ReqList
}

func (req *reqTagList) emptyValue() bool {
	return req.ReqList.EmptyValue()
}

func (req *reqTagList) toCmd() (cmd dto.ListTagCmd, err error) {

	listCmd, err := req.ReqList.ToCmd()
	if err != nil {
		return
	}

	cmd = dto.ListTagCmd{
		PaginationCmd: listCmd,
	}

	return
}
