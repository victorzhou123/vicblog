package service

import (
	"victorzhou123/vicblog/article/domain/tag/entity"
	cmapp "victorzhou123/vicblog/common/app"
)

type TagListCmd struct {
	cmapp.ListCmd
}

type TagListDto struct {
	Total     int      `json:"total"`
	PageCount int      `json:"pages"`
	PageSize  int      `json:"size"`
	CurPage   int      `json:"current"`
	Tag       []TagDto `json:"tag"`
}

func toTagListDto(tags []entity.Tag, cmd *TagListCmd, total int) TagListDto {

	pageCount := total/cmd.PageSize + 1

	dtos := make([]TagDto, len(tags))
	for i := range tags {
		dtos[i] = toTagDto(tags[i])
	}

	return TagListDto{
		Total:     total,
		PageCount: pageCount,
		PageSize:  cmd.PageSize,
		CurPage:   cmd.CurPage,
		Tag:       dtos,
	}
}

type TagDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createTime"`
}

func toTagDto(tag entity.Tag) TagDto {
	return TagDto{
		Id:        tag.Id.Id(),
		Name:      tag.Name.TagName(),
		CreatedAt: tag.CreatedAt.TimeYearToSecond(),
	}
}
