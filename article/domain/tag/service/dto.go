package service

import (
	"victorzhou123/vicblog/article/domain/tag/entity"
	cmappdto "victorzhou123/vicblog/common/app/dto"
)

type TagListCmd struct {
	cmappdto.PaginationCmd
}

type TagListDto struct {
	cmappdto.PaginationDto

	Tag []TagDto `json:"tag"`
}

func toTagListDto(tags []entity.Tag, cmd *TagListCmd, total int) TagListDto {

	dtos := make([]TagDto, len(tags))
	for i := range tags {
		dtos[i] = toTagDto(tags[i])
	}

	return TagListDto{
		PaginationDto: cmd.ToPaginationDto(total),
		Tag:           dtos,
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
