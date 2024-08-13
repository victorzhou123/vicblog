package dto

import (
	"victorzhou123/vicblog/article/domain/tag/entity"
	cmappdto "victorzhou123/vicblog/common/app/dto"
	cment "victorzhou123/vicblog/common/domain/entity"
)

// list tags
type ListTagCmd struct {
	cmappdto.PaginationCmd
}

type TagDto struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

func ToTagDto(tag entity.Tag) TagDto {
	return TagDto{
		Id:        tag.Id.IdNum(),
		Name:      tag.Name.TagName(),
		CreatedAt: tag.CreatedAt.TimeYearToSecond(),
	}
}

type TagListDto struct {
	cmappdto.PaginationDto

	Tag []TagDto `json:"tag"`
}

func ToTagListDto(
	ps cment.PaginationStatus, tags []entity.Tag,
) TagListDto {

	cateDtos := make([]TagDto, len(tags))
	for i := range tags {
		cateDtos[i] = ToTagDto(tags[i])
	}

	return TagListDto{
		PaginationDto: cmappdto.ToPaginationDto(ps),
		Tag:           cateDtos,
	}
}
