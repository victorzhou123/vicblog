package controller

import (
	"github.com/victorzhou123/vicblog/comment/app/dto"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type reqCommentInfo struct {
	Avatar          string `json:"avatar"`
	NickName        string `json:"nickname"`
	Email           string `json:"email"`
	Content         string `json:"content"`
	Website         string `json:"website"`
	RouterUrl       string `json:"routerUrl"`
	ReplyCommentId  string `json:"replyCommentId"`
	ParentCommentId string `json:"ParentCommentId"`
}

func (req *reqCommentInfo) toCmd() (cmd dto.CommentInfoCmd, err error) {

	if cmd.Avatar, err = cmprimitive.NewUrlx(req.Avatar); err != nil {
		return
	}

	if cmd.Email, err = cmprimitive.NewEmail(req.Email); err != nil {
		return
	}

	if cmd.Content, err = cmprimitive.NewCommentContent(req.Content); err != nil {
		return
	}

	if cmd.Website, err = cmprimitive.NewUrlx(req.Website); err != nil {
		return
	}

	if cmd.RouterUrl, err = cmprimitive.NewUrlx(req.RouterUrl); err != nil {
		return
	}

	cmd.ParentCommentId = cmprimitive.NewId(req.ParentCommentId)
	cmd.ReplyCommentId = cmprimitive.NewId(req.ReplyCommentId)
	cmd.NickName = req.NickName

	return
}
