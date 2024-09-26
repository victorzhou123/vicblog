package service

import (
	"github.com/victorzhou123/vicblog/comment/app/dto"
	"github.com/victorzhou123/vicblog/comment/domain/comment/service"
)

type CommentAppService interface {
	AddComment(dto.CommentInfoCmd) error
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
