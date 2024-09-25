package service

import cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"

type CommentCmd struct {
	Avatar          cmprimitive.Urlx
	NickName        cmprimitive.Text
	Email           cmprimitive.Email
	Content         cmprimitive.Text
	Website         cmprimitive.Urlx
	RouterUrl       cmprimitive.Urlx
	ReplyCommentId  cmprimitive.Id
	ParentCommentId cmprimitive.Id
}
