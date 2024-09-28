package entity

import (
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type Comment struct {
	CommentInfo

	Id        cmprimitive.Id
	Status    CommentStatus
	IsDeleted bool
	CreatedAt cmprimitive.Timex
	UpdatedAt cmprimitive.Timex
}

type CommentInfo struct {
	CommentUserInfo

	Content         cmprimitive.Text
	Website         cmprimitive.Urlx
	ArticleId       cmprimitive.Id
	ReplyCommentId  cmprimitive.Id
	ParentCommentId cmprimitive.Id
}

type CommentUserInfo struct {
	Avatar   cmprimitive.Urlx
	NickName CommentNickname
	Email    cmprimitive.Email
}

func (r *Comment) IsShow() bool {
	return !r.IsDeleted && r.Status.IsAuditPassed()
}

func (r *Comment) SetDefaultForCreateAction() {
	if r.Status == nil {
		r.Status = NewCommentStatusWaiting()
	}

	r.IsDeleted = false
}

// IsReplyParentComment to judge if the comment reply parent comment
func (r *Comment) IsReplyParentComment() bool {
	return r.ParentCommentId.IdNum() == r.ReplyCommentId.IdNum()
}

// IsSubComment to judge if c is parent comment of r
func (r *Comment) IsSubComment(c Comment) bool {
	return r.ParentCommentId.IdNum() == c.Id.IdNum()
}

// IsReply to judge if r is the reply comment of c
func (r *Comment) IsReply(c Comment) bool {
	return r.ReplyCommentId.IdNum() == c.Id.IdNum()
}
