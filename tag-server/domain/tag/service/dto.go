package service

import (
	"github.com/victorzhou123/vicblog/common/controller/rpc"
	cmdto "github.com/victorzhou123/vicblog/common/domain/dto"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/tag-server/domain/tag/entity"
)

type TagListDto struct {
	cment.PaginationStatus

	Tags []entity.TagWithRelatedArticleAmount
}

func (dto *TagListDto) toProto() *rpc.TagList {

	tags := make([]*rpc.TagWithRelatedArticleAmount, len(dto.Tags))
	for i := range tags {
		tags[i] = toProtoTagWithRelatedArticleAmount(dto.Tags[i])
	}

	return &rpc.TagList{
		PaginationStatus: cmdto.ToProtoPaginationStatus(dto.PaginationStatus),
		Tags:             tags,
	}
}

func toTagListDto(tagList *rpc.TagList) (TagListDto, error) {

	paginationStatus, err := cmdto.ToPaginationStatus(tagList.PaginationStatus)
	if err != nil {
		return TagListDto{}, err
	}

	tas := make([]entity.TagWithRelatedArticleAmount, len(tagList.GetTags()))
	for i := range tas {
		tas[i], err = toTagWithRelatedArticleAmount(tagList.Tags[i])
		if err != nil {
			return TagListDto{}, err
		}
	}

	return TagListDto{
		PaginationStatus: paginationStatus,
		Tags:             tas,
	}, nil
}

func toTagWithRelatedArticleAmount(ta *rpc.TagWithRelatedArticleAmount) (entity.TagWithRelatedArticleAmount, error) {

	tag, err := toTag(ta.GetTag())
	if err != nil {
		return entity.TagWithRelatedArticleAmount{}, err
	}

	return entity.TagWithRelatedArticleAmount{
		Tag: tag,
	}, nil
}

func toProtoTagWithRelatedArticleAmount(ta entity.TagWithRelatedArticleAmount) *rpc.TagWithRelatedArticleAmount {
	return &rpc.TagWithRelatedArticleAmount{
		Tag:                  toProtoTag(ta.Tag),
		RelatedArticleAmount: int64(ta.RelatedArticleAmount.Amount()),
	}
}

func toProtoTag(tag entity.Tag) *rpc.Tag {
	return &rpc.Tag{
		Id:        tag.Id.Id(),
		Name:      tag.Name.TagName(),
		CreatedAt: tag.CreatedAt.TimeUnix(),
	}
}

func toProtoTagNames(names *rpc.TagNames) ([]entity.TagName, error) {
	var err error
	tagNames := make([]entity.TagName, len(names.Names))
	for i := range tagNames {
		if tagNames[i], err = entity.NewTagName(names.Names[i].GetName()); err != nil {
			return nil, err
		}
	}

	return tagNames, nil
}

func toTag(tag *rpc.Tag) (entity.Tag, error) {

	tagName, err := entity.NewTagName(tag.GetName())
	if err != nil {
		return entity.Tag{}, err
	}

	return entity.Tag{
		Id:        cmprimitive.NewId(tag.GetId()),
		Name:      tagName,
		CreatedAt: cmprimitive.NewTimeXWithUnix(tag.GetCreatedAt()),
	}, nil
}
