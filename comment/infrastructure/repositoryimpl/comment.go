package repositoryimpl

import (
	"github.com/victorzhou123/vicblog/comment/domain/comment/entity"
	"github.com/victorzhou123/vicblog/comment/domain/comment/repository"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
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
		NickName:        comment.NickName.CommentNickname(),
		Email:           comment.Email.Email(),
		Content:         comment.Content.Text(),
		Website:         comment.Website.Urlx(),
		ArticleId:       comment.ArticleId.IdNum(),
		Status:          comment.Status.CommentStatus(),
		ReplyCommentId:  comment.ReplyCommentId.IdNum(),
		ParentCommentId: comment.ParentCommentId.IdNum(),
	}

	return impl.db.Add(&CommentDO{}, &commentDo)
}

func (impl *commentRepoImpl) GetCommentsByArticleId(articleId cmprimitive.Id) ([]entity.Comment, error) {

	commentDos := []CommentDO{}

	if err := impl.db.GetRecords(&CommentDO{}, impl.db.EqualQuery(fieldNameArticleId), &commentDos, articleId.Id()); err != nil {
		return nil, err
	}

	var err error
	comments := make([]entity.Comment, len(commentDos))
	for i := range comments {
		if comments[i], err = commentDos[i].toComment(); err != nil {
			return nil, err
		}
	}

	return comments, nil
}
