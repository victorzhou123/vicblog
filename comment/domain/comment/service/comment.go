package service

import (
	"github.com/victorzhou123/vicblog/comment/domain/comment/entity"
	"github.com/victorzhou123/vicblog/comment/domain/comment/repository"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type CommentService interface {
	AddComment(*entity.CommentInfo) error
	GetCommentsByArticleId(articleId *cmprimitive.Id) ([]entity.Comment, error)
}

type commentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) CommentService {
	return &commentService{
		repo: repo,
	}
}

func (s *commentService) AddComment(c *entity.CommentInfo) error {

	comment := entity.Comment{
		CommentInfo: *c,
	}

	comment.SetDefaultForCreateAction()

	return s.repo.Add(comment)
}

func (s *commentService) GetCommentsByArticleId(articleId *cmprimitive.Id) ([]entity.Comment, error) {
	return s.repo.GetCommentsByArticleId(*articleId)
}
