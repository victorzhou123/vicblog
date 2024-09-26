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

// func (c *Comment) IsCommentInput() bool {
// 	return c.Id == nil && c.Content != nil && c.Avatar != nil &&
// 		c.NickName != nil && c.Email != nil && c.ArticleId != nil &&
// 		c.Status == nil && c.CreatedAt == nil && c.UpdatedAt == nil
// }
