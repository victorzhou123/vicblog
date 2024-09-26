package repositoryimpl

import (
	"gorm.io/gorm"

	"github.com/victorzhou123/vicblog/comment/domain/comment/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

const (
	fieldNameArticleId = "article_id"

	tableNameComment = "comment"
)

type CommentDO struct {
	gorm.Model

	Avatar          string `gorm:"column:avatar;size:255"`
	NickName        string `gorm:"column:nickname;size:255"`
	Email           string `gorm:"column:email;size:255"`
	Content         string `gorm:"column:content;type:text;size:10000"`
	Website         string `gorm:"column:website;size:255"`
	Status          int    `gorm:"column:status;size:255"`
	ArticleId       uint   `gorm:"column:article_id;index:article_id_index;size:255"`
	ReplyCommentId  uint   `gorm:"column:reply_comment_id;index:reply_comment_id_index;size:255"`
	ParentCommentId uint   `gorm:"column:parent_comment_id;index:parent_comment_id_index;size:255"`
}

func (do *CommentDO) TableName() string {
	return tableNameComment
}

func (do *CommentDO) toComment() (comment entity.Comment, err error) {

	if comment.Avatar, err = cmprimitive.NewUrlx(do.Avatar); err != nil {
		return
	}

	if comment.NickName, err = entity.NewCommentNickname(do.NickName); err != nil {
		return
	}

	if comment.Email, err = cmprimitive.NewEmail(do.Email); err != nil {
		return
	}

	if comment.Content, err = cmprimitive.NewCommentContent(do.Content); err != nil {
		return
	}

	if comment.Website, err = cmprimitive.NewUrlx(do.Website); err != nil {
		return
	}

	if comment.Status, err = entity.NewCommentStatus(do.Status); err != nil {
		return
	}

	comment.Id = cmprimitive.NewIdByUint(do.ID)
	comment.ArticleId = cmprimitive.NewIdByUint(do.ArticleId)
	comment.ReplyCommentId = cmprimitive.NewIdByUint(do.ReplyCommentId)
	comment.ParentCommentId = cmprimitive.NewIdByUint(do.ParentCommentId)
	comment.CreatedAt = cmprimitive.NewTimeXWithUnix(do.CreatedAt.Unix())
	comment.UpdatedAt = cmprimitive.NewTimeXWithUnix(do.UpdatedAt.Unix())

	return
}
