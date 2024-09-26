package repositoryimpl

import (
	"github.com/victorzhou123/vicblog/comment/domain/comment/entity"
	"github.com/victorzhou123/vicblog/comment/domain/comment/repository"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewCommentRepo(db mysql.Impl) repository.Comment {

	if err := mysql.AutoMigrate(&CommentDO{}); err != nil {
		return nil
	}

	return &commentRepoImpl{db}
}

type commentRepoImpl struct {
	db mysql.Impl
}

func (impl *commentRepoImpl) Add(comment entity.Comment) error {
	commentDo := CommentDO{
		Avatar:          comment.Avatar.Urlx(),
		NickName:        comment.NickName,
		Email:           comment.Email.Email(),
		Content:         comment.Content.Text(),
		Website:         comment.Website.Urlx(),
		RouterUrl:       comment.RouterUrl.Urlx(),
		Status:          comment.Status.CommentStatus(),
		ReplyCommentId:  comment.ReplyCommentId.IdNum(),
		ParentCommentId: comment.ParentCommentId.IdNum(),
	}

	return impl.db.Add(&CommentDO{}, &commentDo)
}
