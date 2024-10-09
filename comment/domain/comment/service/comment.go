package service

import (
	"github.com/victorzhou123/vicblog/comment/domain/comment/entity"
	"github.com/victorzhou123/vicblog/comment/domain/comment/repository"
	"github.com/victorzhou123/vicblog/common/domain/audit"
	cmerr "github.com/victorzhou123/vicblog/common/domain/error"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type CommentService interface {
	AddComment(*entity.CommentInfo) error
	GetCommentsByArticleId(articleId *cmprimitive.Id) ([]entity.Comment, error)
}

type commentService struct {
	repo  repository.Comment
	audit audit.Audit
}

func NewCommentService(repo repository.Comment, audit audit.Audit) CommentService {
	return &commentService{
		repo:  repo,
		audit: audit,
	}
}

func (s *commentService) AddComment(c *entity.CommentInfo) error {

	if s.audit.Check(c.Content.Text()) {
		return cmerr.NewInvalidParam("input comment not valid")
	}

	comment := entity.Comment{
		CommentInfo: *c,
	}

	comment.SetDefaultForCreateAction()

	return s.repo.Add(comment)
}

func (s *commentService) GetCommentsByArticleId(articleId *cmprimitive.Id) ([]entity.Comment, error) {
	return s.repo.GetCommentsByArticleId(*articleId)
}
