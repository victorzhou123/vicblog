package dto

import (
	"sort"

	"github.com/victorzhou123/vicblog/comment/domain/comment/entity"
)

type CommentInfoCmd struct {
	entity.CommentInfo
}

type CommentTreeDto struct {
	Total    int          `json:"total"`
	Comments []CommentDto `json:"comments"`
}

func ToCommentTreeDto(comments []entity.Comment) CommentTreeDto {

	// sort by create time ascend
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreatedAt.TimeUnix() < comments[j].CreatedAt.TimeUnix()
	})

	var cs []CommentDto
	for i := range comments {
		if comments[i].ParentCommentId.IdNum() == 0 {
			// find all sub comments
			var childrenComments []CommentDto
			for j := i + 1; j < len(comments); j++ {

				if comments[j].IsSubComment(comments[i]) {

					// if the comments[j] reply parent comment
					if comments[j].IsReplyParentComment() {
						childrenComments = append(childrenComments, toCommentDto(comments[j], ""))
						continue
					}

					// find reply nick name of comments[j]
					for k := i; k < j; k++ {

						if comments[j].IsReply(comments[k]) {

							childrenComments = append(childrenComments,
								toCommentDto(comments[j], comments[k].NickName.CommentNickname()))

							break
						}
					}
				}
			}

			// set commentDto into cs
			if len(childrenComments) != 0 {
				dto := toCommentDto(comments[i], "")
				dto.SubComments = childrenComments
				cs = append(cs, dto)

				continue
			}

			cs = append(cs, toCommentDto(comments[i], ""))
		}
	}

	return CommentTreeDto{
		Total:    len(comments),
		Comments: cs,
	}
}

type CommentDto struct {
	Id            uint         `json:"id"`
	Avatar        string       `json:"avatar"`
	Nickname      string       `json:"nickname"`
	WebSite       string       `json:"website"`
	Content       string       `json:"content"`
	CreatedAt     string       `json:"createdAt"`
	SubComments   []CommentDto `json:"subComments,omitempty"`
	ReplyNickname string       `json:"replyNickname,omitempty"`
}

func toCommentDto(comment entity.Comment, replyNickname string) CommentDto {
	return CommentDto{
		Id:            comment.Id.IdNum(),
		Avatar:        comment.Avatar.Urlx(),
		Nickname:      comment.NickName.CommentNickname(),
		WebSite:       comment.Website.Urlx(),
		Content:       comment.Content.Text(),
		ReplyNickname: replyNickname,
		CreatedAt:     comment.CreatedAt.TimeYearToSecond(),
	}
}
