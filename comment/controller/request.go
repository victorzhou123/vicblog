package controller

import (
	"github.com/victorzhou123/vicblog/comment/app/dto"
	commentent "github.com/victorzhou123/vicblog/comment/domain/comment/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type reqCommentInfo struct {
	Avatar          string `json:"avatar"`
	NickName        string `json:"nickname" binding:"required"`
	Email           string `json:"email"`
	Content         string `json:"content" binding:"required"`
	Website         string `json:"website"`
	ArticleId       uint   `json:"articleId" binding:"required"`
	ReplyCommentId  uint   `json:"replyCommentId"`
	ParentCommentId uint   `json:"ParentCommentId"`
}

func (req *reqCommentInfo) toCmd() (cmd dto.CommentInfoCmd, err error) {

	if cmd.Avatar, err = cmprimitive.NewUrlx(req.Avatar); err != nil {
		return
	}

	if cmd.NickName, err = commentent.NewCommentNickname(req.NickName); err != nil {
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

	cmd.ArticleId = cmprimitive.NewIdByUint(req.ArticleId)
	cmd.ParentCommentId = cmprimitive.NewIdByUint(req.ParentCommentId)
	cmd.ReplyCommentId = cmprimitive.NewIdByUint(req.ReplyCommentId)

	return
}
