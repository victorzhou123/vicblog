package service

import (
	"github.com/victorzhou123/vicblog/comment/app/dto"
	"github.com/victorzhou123/vicblog/comment/domain/comment/service"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type CommentAppService interface {
	AddComment(dto.CommentInfoCmd) error
	GetCommentsTreeByArticleId(articleId cmprimitive.Id) (dto.CommentTreeDto, error)
}

type commentAppService struct {
	comment service.CommentService
}

func NewCommentAppService(comment service.CommentService) CommentAppService {
	return &commentAppService{
		comment: comment,
	}
}

func (s *commentAppService) AddComment(cmd dto.CommentInfoCmd) error {
	return s.comment.AddComment(&cmd.CommentInfo)
}

func (s *commentAppService) GetCommentsTreeByArticleId(articleId cmprimitive.Id) (dto.CommentTreeDto, error) {

	comments, err := s.comment.GetCommentsByArticleId(&articleId)
	if err != nil {
		return dto.CommentTreeDto{}, err
	}

	return dto.ToCommentTreeDto(comments), nil
}
