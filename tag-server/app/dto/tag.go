package dto

import (
	cmappdto "github.com/victorzhou123/vicblog/common/app/dto"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	"github.com/victorzhou123/vicblog/tag-server/domain/tag/entity"
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

type TagWithRelatedArticleAmountDto struct {
	TagDto

	RelatedArticleAmount int `json:"relatedArticleAmount"`
}

func ToTagWithRelatedArticleAmountDto(
	c entity.TagWithRelatedArticleAmount,
) TagWithRelatedArticleAmountDto {
	return TagWithRelatedArticleAmountDto{
		TagDto:               ToTagDto(c.Tag),
		RelatedArticleAmount: c.RelatedArticleAmount.Amount(),
	}
}

type TagListDto struct {
	cmappdto.PaginationDto

	Tag []TagWithRelatedArticleAmountDto `json:"tag"`
}

func ToTagListDto(
	ps cment.PaginationStatus, tags []entity.TagWithRelatedArticleAmount,
) TagListDto {

	tagDtos := make([]TagWithRelatedArticleAmountDto, len(tags))
	for i := range tags {
		tagDtos[i] = ToTagWithRelatedArticleAmountDto(tags[i])
	}

	return TagListDto{
		PaginationDto: cmappdto.ToPaginationDto(ps),
		Tag:           tagDtos,
	}
}
