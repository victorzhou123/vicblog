package repositoryimpl

import "gorm.io/gorm"

const (
	tableNameComment = "comment"
)

type CommentDO struct {
	gorm.Model

	Avatar          string `gorm:"column:avatar;size:255"`
	NickName        string `gorm:"column:nickname;size:255"`
	Email           string `gorm:"column:email;size:255"`
	Content         string `gorm:"column:content;type:text;size:10000"`
	Website         string `gorm:"column:website;size:255"`
	RouterUrl       string `gorm:"column:router_url;index:router_url_index;size:255"`
	Status          int    `gorm:"column:status;size:255"`
	ReplyCommentId  uint   `gorm:"column:reply_comment_id;index:reply_comment_id_index;size:255"`
	ParentCommentId uint   `gorm:"column:parent_comment_id;index:parent_comment_id_index;size:255"`
}

func (do *CommentDO) TableName() string {
	return tableNameComment
}
